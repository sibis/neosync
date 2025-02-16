package tablesync_workflow_register

import (
	"github.com/nucleuscloud/neosync/backend/gen/go/protos/mgmt/v1alpha1/mgmtv1alpha1connect"
	benthosstream "github.com/nucleuscloud/neosync/internal/benthos-stream"
	connectionmanager "github.com/nucleuscloud/neosync/internal/connection-manager"
	neosync_benthos_mongodb "github.com/nucleuscloud/neosync/worker/pkg/benthos/mongodb"
	neosync_benthos_sql "github.com/nucleuscloud/neosync/worker/pkg/benthos/sql"
	sync_activity "github.com/nucleuscloud/neosync/worker/pkg/workflows/tablesync/activities/sync"
	tablesync_workflow "github.com/nucleuscloud/neosync/worker/pkg/workflows/tablesync/workflow"

	"go.opentelemetry.io/otel/metric"
)

type Worker interface {
	RegisterWorkflow(workflow any)
	RegisterActivity(activity any)
}

func Register(
	w Worker,
	connclient mgmtv1alpha1connect.ConnectionServiceClient,
	jobclient mgmtv1alpha1connect.JobServiceClient,
	sqlconnmanager connectionmanager.Interface[neosync_benthos_sql.SqlDbtx],
	mongoconnmanager connectionmanager.Interface[neosync_benthos_mongodb.MongoClient],
	meter metric.Meter, // optional
	benthosStreamManager benthosstream.BenthosStreamManagerClient,
) {
	tsWf := tablesync_workflow.New()
	w.RegisterWorkflow(tsWf.TableSync)

	syncActivity := sync_activity.New(
		connclient,
		jobclient,
		sqlconnmanager,
		mongoconnmanager,
		meter,
		benthosStreamManager,
	)

	w.RegisterActivity(syncActivity.Sync)
	w.RegisterActivity(syncActivity.SyncTable)
}
