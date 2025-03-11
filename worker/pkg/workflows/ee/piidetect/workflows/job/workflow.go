package piidetect_job_workflow

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	mgmtv1alpha1 "github.com/nucleuscloud/neosync/backend/gen/go/protos/mgmt/v1alpha1"
	piidetect_job_activities "github.com/nucleuscloud/neosync/worker/pkg/workflows/ee/piidetect/workflows/job/activities"
	piidetect_table_workflow "github.com/nucleuscloud/neosync/worker/pkg/workflows/ee/piidetect/workflows/table"
	"github.com/spf13/viper"
	"go.temporal.io/sdk/log"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

type Workflow struct{}

func New() *Workflow {
	return &Workflow{}
}

type PiiDetectRequest struct {
	JobId string
}

type PiiDetectResponse struct {
	ReportKey *mgmtv1alpha1.RunContextKey
}

func (w *Workflow) JobPiiDetect(ctx workflow.Context, req *PiiDetectRequest) (*PiiDetectResponse, error) {
	logger := log.With(
		workflow.GetLogger(ctx),
		"jobId", req.JobId,
	)

	logger.Info("starting PII detection")

	var activities *piidetect_job_activities.Activities

	var jobDetailsResp *piidetect_job_activities.GetPiiDetectJobDetailsResponse
	err := workflow.ExecuteActivity(
		workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
			StartToCloseTimeout: 1 * time.Minute,
			RetryPolicy: &temporal.RetryPolicy{
				MaximumAttempts: 3,
			},
		}),
		activities.GetPiiDetectJobDetails,
		&piidetect_job_activities.GetPiiDetectJobDetailsRequest{
			JobId: req.JobId,
		},
	).Get(ctx, &jobDetailsResp)
	if err != nil {
		return nil, err
	}

	var filter *mgmtv1alpha1.JobTypeConfig_JobTypePiiDetect_TableScanFilter
	if jobDetailsResp != nil && jobDetailsResp.PiiDetectConfig != nil && jobDetailsResp.PiiDetectConfig.TableScanFilter != nil {
		logger.Debug("using table scan filter", "filter", jobDetailsResp.PiiDetectConfig.TableScanFilter)
		filter = jobDetailsResp.PiiDetectConfig.TableScanFilter
	}

	var tablesToScanResp *piidetect_job_activities.GetTablesToPiiScanResponse
	err = workflow.ExecuteActivity(
		workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
			StartToCloseTimeout: 1 * time.Minute,
			RetryPolicy: &temporal.RetryPolicy{
				MaximumAttempts: 3,
			},
		}),
		activities.GetTablesToPiiScan,
		&piidetect_job_activities.GetTablesToPiiScanRequest{
			SourceConnectionId: jobDetailsResp.SourceConnectionId,
			Filter:             filter,
		},
	).Get(ctx, &tablesToScanResp)
	if err != nil {
		return nil, err
	}

	piiDetectConfig := jobDetailsResp.PiiDetectConfig
	if piiDetectConfig == nil {
		piiDetectConfig = &mgmtv1alpha1.JobTypeConfig_JobTypePiiDetect{}
	}

	report, err := w.orchestrateTables(
		ctx,
		jobDetailsResp.AccountId,
		tablesToScanResp.Tables,
		req.JobId,
		jobDetailsResp.SourceConnectionId,
		piiDetectConfig.GetDataSampling().GetIsEnabled(),
		piiDetectConfig.GetUserPrompt(),
		logger,
	)
	if err != nil {
		return nil, fmt.Errorf("unable to orchestrate tables: %w", err)
	}

	logger.Debug("saving job pii detect report")

	var saveResp *piidetect_job_activities.SaveJobPiiDetectReportResponse
	err = workflow.ExecuteActivity(
		workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
			StartToCloseTimeout: 1 * time.Minute,
			RetryPolicy: &temporal.RetryPolicy{
				MaximumAttempts: 3,
			},
		}),
		activities.SaveJobPiiDetectReport,
		&piidetect_job_activities.SaveJobPiiDetectReportRequest{
			AccountId: jobDetailsResp.AccountId,
			JobId:     req.JobId,
			Report:    report,
		},
	).Get(ctx, &saveResp)
	if err != nil {
		return nil, fmt.Errorf("unable to save job pii detect report: %w", err)
	}

	logger.Info("PII detection completed")
	return &PiiDetectResponse{
		ReportKey: saveResp.Key,
	}, nil
}

func (w *Workflow) orchestrateTables(
	ctx workflow.Context,
	accountId string,
	tables []piidetect_job_activities.TableIdentifier,
	jobId string,
	connectionId string,
	shouldSampleData bool,
	userPrompt string,
	logger log.Logger,
) (*piidetect_job_activities.JobPiiDetectReport, error) {
	// todo: retrieve previous report

	workselector := workflow.NewNamedSelector(ctx, "job_pii_detect")

	maxConcurrency := getTablePiiDetectMaxConcurrency()
	inFlightLimiter := workflow.NewSemaphore(ctx, int64(maxConcurrency))

	tableWf := piidetect_table_workflow.New()
	wfInfo := workflow.GetInfo(ctx)

	tableResultKeys := []*mgmtv1alpha1.RunContextKey{}
	mu := workflow.NewMutex(ctx)

	processTable := func(table piidetect_job_activities.TableIdentifier) error {
		if err := inFlightLimiter.Acquire(ctx, 1); err != nil {
			return fmt.Errorf("unable to acquire semaphore: %w", err)
		}
		workselector.AddFuture(
			workflow.ExecuteChildWorkflow(
				workflow.WithChildOptions(
					ctx,
					workflow.ChildWorkflowOptions{
						WorkflowID: getTablePiiDetectChildWorkflowId(
							wfInfo.WorkflowExecution.ID,
							fmt.Sprintf("%s.%s", table.Schema, table.Table),
							workflow.Now(ctx),
						),
						RetryPolicy: &temporal.RetryPolicy{
							MaximumAttempts: 1,
						},
						WorkflowRunTimeout: 5 * time.Minute,
					}),
				tableWf.TablePiiDetect,
				&piidetect_table_workflow.TablePiiDetectRequest{
					AccountId:          accountId,
					JobId:              jobId,
					ConnectionId:       connectionId,
					TableSchema:        table.Schema,
					TableName:          table.Table,
					ParentExecutionId:  &wfInfo.WorkflowExecution.ID,
					ShouldSampleData:   shouldSampleData,
					UserPrompt:         userPrompt,
					PreviousResultsKey: nil, // todo: retrieve previous report
				},
			),
			func(f workflow.Future) {
				var wfResult *piidetect_table_workflow.TablePiiDetectResponse
				err := f.Get(ctx, &wfResult)
				inFlightLimiter.Release(1)
				if err != nil {
					logger.Error("activity did not complete", "err", err)
					return
				}
				logger.Debug("table pii detect completed", "table", table.Table, "schema", table.Schema)
				err = mu.Lock(ctx)
				if err != nil {
					logger.Error("unable to lock mutex after table pii detect completed", "err", err)
					return
				}
				defer mu.Unlock()
				tableResultKeys = append(tableResultKeys, wfResult.ResultKey)
			},
		)
		return nil
	}

	for _, table := range tables {
		if err := processTable(table); err != nil {
			return nil, err
		}
	}

	for range tables {
		workselector.Select(ctx)
	}

	logger.Debug("all tables processed")
	return &piidetect_job_activities.JobPiiDetectReport{
		SuccessfulTableKeys: tableResultKeys,
	}, nil
}

func getTablePiiDetectChildWorkflowId(parentJobRunId, configName string, now time.Time) string {
	id := fmt.Sprintf("%s-%s-%d", parentJobRunId, sanitizeWorkflowID(strings.ToLower(configName)), now.UnixNano())
	if len(id) > 1000 {
		id = id[:1000]
	}
	return id
}

var invalidWorkflowIDChars = regexp.MustCompile(`[^a-zA-Z0-9_\-]`)

func sanitizeWorkflowID(id string) string {
	return invalidWorkflowIDChars.ReplaceAllString(id, "_")
}

func getTablePiiDetectMaxConcurrency() int {
	maxConcurrency := viper.GetInt("TABLE_PII_DETECT_MAX_CONCURRENCY")
	if maxConcurrency <= 0 {
		return 3 // default max concurrency
	}
	return maxConcurrency
}
