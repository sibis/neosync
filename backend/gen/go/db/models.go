// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package db_queries

import (
	"github.com/jackc/pgx/v5/pgtype"
	jsonmodels "github.com/nucleuscloud/neosync/backend/internal/nucleusdb/json-models"
)

type NeosyncApiAccount struct {
	ID          pgtype.UUID
	CreatedAt   pgtype.Timestamp
	UpdatedAt   pgtype.Timestamp
	AccountType int16
	AccountSlug string
}

type NeosyncApiAccountUserAssociation struct {
	ID        pgtype.UUID
	AccountID pgtype.UUID
	UserID    pgtype.UUID
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

type NeosyncApiConnection struct {
	ID               pgtype.UUID
	CreatedAt        pgtype.Timestamp
	UpdatedAt        pgtype.Timestamp
	Name             string
	AccountID        pgtype.UUID
	ConnectionConfig *jsonmodels.ConnectionConfig
	CreatedByID      pgtype.UUID
	UpdatedByID      pgtype.UUID
}

type NeosyncApiJob struct {
	ID                 pgtype.UUID
	CreatedAt          pgtype.Timestamp
	UpdatedAt          pgtype.Timestamp
	Name               string
	AccountID          pgtype.UUID
	Status             int16
	ConnectionSourceID pgtype.UUID
	ConnectionOptions  *jsonmodels.JobSourceOptions
	Mappings           []*jsonmodels.JobMapping
	CronSchedule       pgtype.Text
	CreatedByID        pgtype.UUID
	UpdatedByID        pgtype.UUID
}

type NeosyncApiJobDestinationConnectionAssociation struct {
	ID           pgtype.UUID
	CreatedAt    pgtype.Timestamp
	UpdatedAt    pgtype.Timestamp
	JobID        pgtype.UUID
	ConnectionID pgtype.UUID
	Options      *jsonmodels.JobDestinationOptions
}

type NeosyncApiUser struct {
	ID        pgtype.UUID
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

type NeosyncApiUserIdentityProviderAssociation struct {
	ID              pgtype.UUID
	UserID          pgtype.UUID
	Auth0ProviderID string
	CreatedAt       pgtype.Timestamp
	UpdatedAt       pgtype.Timestamp
}
