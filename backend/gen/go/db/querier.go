// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db_queries

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	pg_models "github.com/nucleuscloud/neosync/backend/sql/postgresql/models"
)

type Querier interface {
	AreConnectionsInAccount(ctx context.Context, db DBTX, arg AreConnectionsInAccountParams) (int64, error)
	ConvertPersonalAccountToTeam(ctx context.Context, db DBTX, arg ConvertPersonalAccountToTeamParams) (NeosyncApiAccount, error)
	CreateAccountApiKey(ctx context.Context, db DBTX, arg CreateAccountApiKeyParams) (NeosyncApiAccountApiKey, error)
	CreateAccountHook(ctx context.Context, db DBTX, arg CreateAccountHookParams) (NeosyncApiAccountHook, error)
	CreateAccountInvite(ctx context.Context, db DBTX, arg CreateAccountInviteParams) (NeosyncApiAccountInvite, error)
	CreateAccountUserAssociation(ctx context.Context, db DBTX, arg CreateAccountUserAssociationParams) error
	CreateConnection(ctx context.Context, db DBTX, arg CreateConnectionParams) (NeosyncApiConnection, error)
	CreateIdentityProviderAssociation(ctx context.Context, db DBTX, arg CreateIdentityProviderAssociationParams) (NeosyncApiUserIdentityProviderAssociation, error)
	CreateJob(ctx context.Context, db DBTX, arg CreateJobParams) (NeosyncApiJob, error)
	CreateJobConnectionDestination(ctx context.Context, db DBTX, arg CreateJobConnectionDestinationParams) (NeosyncApiJobDestinationConnectionAssociation, error)
	CreateJobConnectionDestinations(ctx context.Context, db DBTX, arg []CreateJobConnectionDestinationsParams) (int64, error)
	CreateJobHook(ctx context.Context, db DBTX, arg CreateJobHookParams) (NeosyncApiJobHook, error)
	CreateMachineUser(ctx context.Context, db DBTX) (NeosyncApiUser, error)
	CreateNonMachineUser(ctx context.Context, db DBTX) (NeosyncApiUser, error)
	CreatePersonalAccount(ctx context.Context, db DBTX, arg CreatePersonalAccountParams) (NeosyncApiAccount, error)
	CreateSlackOAuthConnection(ctx context.Context, db DBTX, arg CreateSlackOAuthConnectionParams) (NeosyncApiSlackOauthConnection, error)
	CreateTeamAccount(ctx context.Context, db DBTX, accountSlug string) (NeosyncApiAccount, error)
	CreateUserDefinedTransformer(ctx context.Context, db DBTX, arg CreateUserDefinedTransformerParams) (NeosyncApiTransformer, error)
	DeleteJob(ctx context.Context, db DBTX, id pgtype.UUID) error
	DeleteSlackOAuthConnection(ctx context.Context, db DBTX, accountID pgtype.UUID) error
	DeleteUserDefinedTransformerById(ctx context.Context, db DBTX, id pgtype.UUID) error
	DoesJobHaveConnectionId(ctx context.Context, db DBTX, arg DoesJobHaveConnectionIdParams) (bool, error)
	GetAccount(ctx context.Context, db DBTX, id pgtype.UUID) (NeosyncApiAccount, error)
	GetAccountApiKeyById(ctx context.Context, db DBTX, id pgtype.UUID) (NeosyncApiAccountApiKey, error)
	GetAccountApiKeyByKeyValue(ctx context.Context, db DBTX, keyValue string) (NeosyncApiAccountApiKey, error)
	GetAccountApiKeys(ctx context.Context, db DBTX, accountid pgtype.UUID) ([]NeosyncApiAccountApiKey, error)
	GetAccountHookById(ctx context.Context, db DBTX, id pgtype.UUID) (NeosyncApiAccountHook, error)
	GetAccountHooksByAccount(ctx context.Context, db DBTX, accountID pgtype.UUID) ([]NeosyncApiAccountHook, error)
	GetAccountIdFromJobId(ctx context.Context, db DBTX, id pgtype.UUID) (pgtype.UUID, error)
	GetAccountIds(ctx context.Context, db DBTX) ([]pgtype.UUID, error)
	GetAccountInvite(ctx context.Context, db DBTX, id pgtype.UUID) (NeosyncApiAccountInvite, error)
	GetAccountInviteByToken(ctx context.Context, db DBTX, token string) (NeosyncApiAccountInvite, error)
	GetAccountOnboardingConfig(ctx context.Context, db DBTX, id pgtype.UUID) (*pg_models.AccountOnboardingConfig, error)
	GetAccountUserAssociation(ctx context.Context, db DBTX, arg GetAccountUserAssociationParams) (NeosyncApiAccountUserAssociation, error)
	GetAccountUsers(ctx context.Context, db DBTX, accountid pgtype.UUID) ([]pgtype.UUID, error)
	GetAccountsByUser(ctx context.Context, db DBTX, id pgtype.UUID) ([]NeosyncApiAccount, error)
	GetActiveAccountHooksByEvent(ctx context.Context, db DBTX, arg GetActiveAccountHooksByEventParams) ([]NeosyncApiAccountHook, error)
	GetActiveAccountInvites(ctx context.Context, db DBTX, accountid pgtype.UUID) ([]NeosyncApiAccountInvite, error)
	GetActiveJobHooks(ctx context.Context, db DBTX, jobID pgtype.UUID) ([]NeosyncApiJobHook, error)
	GetActivePostSyncJobHooks(ctx context.Context, db DBTX, jobID pgtype.UUID) ([]NeosyncApiJobHook, error)
	GetActivePreSyncJobHooks(ctx context.Context, db DBTX, jobID pgtype.UUID) ([]NeosyncApiJobHook, error)
	GetAnonymousUser(ctx context.Context, db DBTX) (NeosyncApiUser, error)
	GetBilledAccounts(ctx context.Context, db DBTX, accountids []pgtype.UUID) ([]NeosyncApiAccount, error)
	GetConnectionById(ctx context.Context, db DBTX, id pgtype.UUID) (NeosyncApiConnection, error)
	GetConnectionByNameAndAccount(ctx context.Context, db DBTX, arg GetConnectionByNameAndAccountParams) (NeosyncApiConnection, error)
	GetConnectionsByAccount(ctx context.Context, db DBTX, accountid pgtype.UUID) ([]NeosyncApiConnection, error)
	GetConnectionsByIds(ctx context.Context, db DBTX, dollar_1 []pgtype.UUID) ([]NeosyncApiConnection, error)
	GetJobById(ctx context.Context, db DBTX, id pgtype.UUID) (NeosyncApiJob, error)
	GetJobByNameAndAccount(ctx context.Context, db DBTX, arg GetJobByNameAndAccountParams) (NeosyncApiJob, error)
	GetJobConnectionDestination(ctx context.Context, db DBTX, id pgtype.UUID) (NeosyncApiJobDestinationConnectionAssociation, error)
	GetJobConnectionDestinations(ctx context.Context, db DBTX, id pgtype.UUID) ([]NeosyncApiJobDestinationConnectionAssociation, error)
	GetJobConnectionDestinationsByJobIds(ctx context.Context, db DBTX, jobids []pgtype.UUID) ([]NeosyncApiJobDestinationConnectionAssociation, error)
	GetJobHookById(ctx context.Context, db DBTX, id pgtype.UUID) (NeosyncApiJobHook, error)
	GetJobHooksByJob(ctx context.Context, db DBTX, jobID pgtype.UUID) ([]NeosyncApiJobHook, error)
	GetJobsByAccount(ctx context.Context, db DBTX, accountid pgtype.UUID) ([]NeosyncApiJob, error)
	GetPersonalAccountByUserId(ctx context.Context, db DBTX, userid pgtype.UUID) (NeosyncApiAccount, error)
	GetRunContextByKey(ctx context.Context, db DBTX, arg GetRunContextByKeyParams) (NeosyncApiRuncontext, error)
	GetSlackAccessToken(ctx context.Context, db DBTX, accountID pgtype.UUID) (string, error)
	GetTeamAccountsByUserId(ctx context.Context, db DBTX, userid pgtype.UUID) ([]NeosyncApiAccount, error)
	GetTemporalConfigByAccount(ctx context.Context, db DBTX, id pgtype.UUID) (*pg_models.TemporalConfig, error)
	GetTemporalConfigByUserAccount(ctx context.Context, db DBTX, arg GetTemporalConfigByUserAccountParams) (*pg_models.TemporalConfig, error)
	GetUser(ctx context.Context, db DBTX, id pgtype.UUID) (NeosyncApiUser, error)
	GetUserAssociationByProviderSub(ctx context.Context, db DBTX, providerSub string) (NeosyncApiUserIdentityProviderAssociation, error)
	GetUserByProviderSub(ctx context.Context, db DBTX, providerSub string) (NeosyncApiUser, error)
	GetUserDefinedTransformerById(ctx context.Context, db DBTX, id pgtype.UUID) (NeosyncApiTransformer, error)
	GetUserDefinedTransformersByAccount(ctx context.Context, db DBTX, accountid pgtype.UUID) ([]NeosyncApiTransformer, error)
	GetUserIdentitiesByTeamAccount(ctx context.Context, db DBTX, accountid pgtype.UUID) ([]NeosyncApiUserIdentityProviderAssociation, error)
	GetUserIdentityAssociationsByUserIds(ctx context.Context, db DBTX, dollar_1 []pgtype.UUID) ([]NeosyncApiUserIdentityProviderAssociation, error)
	GetUserIdentityByUserId(ctx context.Context, db DBTX, userID pgtype.UUID) (NeosyncApiUserIdentityProviderAssociation, error)
	IsAccountHookNameAvailable(ctx context.Context, db DBTX, arg IsAccountHookNameAvailableParams) (bool, error)
	IsConnectionInAccount(ctx context.Context, db DBTX, arg IsConnectionInAccountParams) (int64, error)
	IsConnectionNameAvailable(ctx context.Context, db DBTX, arg IsConnectionNameAvailableParams) (int64, error)
	IsJobHookNameAvailable(ctx context.Context, db DBTX, arg IsJobHookNameAvailableParams) (bool, error)
	IsJobNameAvailable(ctx context.Context, db DBTX, arg IsJobNameAvailableParams) (int64, error)
	IsTransformerNameAvailable(ctx context.Context, db DBTX, arg IsTransformerNameAvailableParams) (int64, error)
	IsUserInAccount(ctx context.Context, db DBTX, arg IsUserInAccountParams) (int64, error)
	IsUserInAccountApiKey(ctx context.Context, db DBTX, arg IsUserInAccountApiKeyParams) (int64, error)
	RemoveAccountApiKey(ctx context.Context, db DBTX, id pgtype.UUID) error
	RemoveAccountHookById(ctx context.Context, db DBTX, id pgtype.UUID) error
	RemoveAccountInvite(ctx context.Context, db DBTX, id pgtype.UUID) error
	RemoveAccountUser(ctx context.Context, db DBTX, arg RemoveAccountUserParams) error
	RemoveConnectionById(ctx context.Context, db DBTX, id pgtype.UUID) error
	RemoveConnectionByNameAndAccount(ctx context.Context, db DBTX, arg RemoveConnectionByNameAndAccountParams) error
	RemoveJobById(ctx context.Context, db DBTX, id pgtype.UUID) error
	RemoveJobConnectionDestination(ctx context.Context, db DBTX, id pgtype.UUID) error
	RemoveJobConnectionDestinations(ctx context.Context, db DBTX, jobids []pgtype.UUID) error
	RemoveJobHookById(ctx context.Context, db DBTX, id pgtype.UUID) error
	SetAccountCreatedAt(ctx context.Context, db DBTX, arg SetAccountCreatedAtParams) (NeosyncApiAccount, error)
	SetAccountHookEnabled(ctx context.Context, db DBTX, arg SetAccountHookEnabledParams) (NeosyncApiAccountHook, error)
	SetAnonymousUser(ctx context.Context, db DBTX) (NeosyncApiUser, error)
	SetJobHookEnabled(ctx context.Context, db DBTX, arg SetJobHookEnabledParams) (NeosyncApiJobHook, error)
	SetJobSyncOptions(ctx context.Context, db DBTX, arg SetJobSyncOptionsParams) (NeosyncApiJob, error)
	SetJobWorkflowOptions(ctx context.Context, db DBTX, arg SetJobWorkflowOptionsParams) (NeosyncApiJob, error)
	SetNewAccountStripeCustomerId(ctx context.Context, db DBTX, arg SetNewAccountStripeCustomerIdParams) (NeosyncApiAccount, error)
	SetRunContext(ctx context.Context, db DBTX, arg SetRunContextParams) error
	UpdateAccountApiKeyValue(ctx context.Context, db DBTX, arg UpdateAccountApiKeyValueParams) (NeosyncApiAccountApiKey, error)
	UpdateAccountHook(ctx context.Context, db DBTX, arg UpdateAccountHookParams) (NeosyncApiAccountHook, error)
	UpdateAccountInviteToAccepted(ctx context.Context, db DBTX, id pgtype.UUID) (NeosyncApiAccountInvite, error)
	UpdateAccountOnboardingConfig(ctx context.Context, db DBTX, arg UpdateAccountOnboardingConfigParams) (NeosyncApiAccount, error)
	UpdateActiveAccountInvitesToExpired(ctx context.Context, db DBTX, arg UpdateActiveAccountInvitesToExpiredParams) (NeosyncApiAccountInvite, error)
	UpdateConnection(ctx context.Context, db DBTX, arg UpdateConnectionParams) (NeosyncApiConnection, error)
	UpdateJobConnectionDestination(ctx context.Context, db DBTX, arg UpdateJobConnectionDestinationParams) (NeosyncApiJobDestinationConnectionAssociation, error)
	UpdateJobHook(ctx context.Context, db DBTX, arg UpdateJobHookParams) (NeosyncApiJobHook, error)
	UpdateJobMappings(ctx context.Context, db DBTX, arg UpdateJobMappingsParams) (NeosyncApiJob, error)
	UpdateJobSchedule(ctx context.Context, db DBTX, arg UpdateJobScheduleParams) (NeosyncApiJob, error)
	UpdateJobSource(ctx context.Context, db DBTX, arg UpdateJobSourceParams) (NeosyncApiJob, error)
	UpdateJobVirtualForeignKeys(ctx context.Context, db DBTX, arg UpdateJobVirtualForeignKeysParams) (NeosyncApiJob, error)
	UpdateTemporalConfigByAccount(ctx context.Context, db DBTX, arg UpdateTemporalConfigByAccountParams) (NeosyncApiAccount, error)
	UpdateUserDefinedTransformer(ctx context.Context, db DBTX, arg UpdateUserDefinedTransformerParams) (NeosyncApiTransformer, error)
}

var _ Querier = (*Queries)(nil)
