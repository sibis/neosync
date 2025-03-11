package piidetect_workflow_register

import (
	"github.com/nucleuscloud/neosync/backend/gen/go/protos/mgmt/v1alpha1/mgmtv1alpha1connect"
	"github.com/nucleuscloud/neosync/internal/connectiondata"
	piidetect_job_workflow "github.com/nucleuscloud/neosync/worker/pkg/workflows/ee/piidetect/workflows/job"
	piidetect_job_activities "github.com/nucleuscloud/neosync/worker/pkg/workflows/ee/piidetect/workflows/job/activities"
	piidetect_table_workflow "github.com/nucleuscloud/neosync/worker/pkg/workflows/ee/piidetect/workflows/table"
	piidetect_table_activities "github.com/nucleuscloud/neosync/worker/pkg/workflows/ee/piidetect/workflows/table/activities"
	"github.com/openai/openai-go"
)

type Worker interface {
	RegisterWorkflow(workflow any)
	RegisterActivity(activity any)
}

func Register(
	w Worker,
	connclient mgmtv1alpha1connect.ConnectionServiceClient,
	jobclient mgmtv1alpha1connect.JobServiceClient,
	openaiclient *openai.Client,
	connectiondatabuilder connectiondata.ConnectionDataBuilder,
) {
	tablePiiDetectWorkflow := piidetect_table_workflow.New()
	jobPiiDetectWorkflow := piidetect_job_workflow.New()

	w.RegisterWorkflow(tablePiiDetectWorkflow.TablePiiDetect)
	w.RegisterWorkflow(jobPiiDetectWorkflow.JobPiiDetect)

	tablePiiDetectActivitites := piidetect_table_activities.New(connclient, openaiclient.Chat.Completions, connectiondatabuilder, jobclient)
	w.RegisterActivity(tablePiiDetectActivitites.GetColumnData)
	w.RegisterActivity(tablePiiDetectActivitites.DetectPiiRegex)
	w.RegisterActivity(tablePiiDetectActivitites.DetectPiiLLM)
	w.RegisterActivity(tablePiiDetectActivitites.SaveTablePiiDetectReport)

	jobPiiDetectActivitites := piidetect_job_activities.New(jobclient, connclient, connectiondatabuilder)
	w.RegisterActivity(jobPiiDetectActivitites.GetPiiDetectJobDetails)
	w.RegisterActivity(jobPiiDetectActivitites.GetTablesToPiiScan)
	w.RegisterActivity(jobPiiDetectActivitites.SaveJobPiiDetectReport)
}
