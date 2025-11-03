package usecases

import (
	"context"
	"messenger-module/entities"
)

// Repository abstracts data access for usecases
// It is implemented by repositories.DBRepository
// Split by entity for clarity

type UserRepo interface {
	CreateUser(ctx context.Context, in entities.User) (entities.User, error)
	GetUser(ctx context.Context, id string) (entities.User, error)
	ListUsers(ctx context.Context) ([]entities.User, error)
	UpdateUser(ctx context.Context, id string, in entities.User) (entities.User, error)
	DeleteUser(ctx context.Context, id string) error
}

type PlanRepo interface {
	CreatePlan(ctx context.Context, in entities.Plan) (entities.Plan, error)
	GetPlan(ctx context.Context, id string) (entities.Plan, error)
	ListPlans(ctx context.Context) ([]entities.Plan, error)
	UpdatePlan(ctx context.Context, id string, in entities.Plan) (entities.Plan, error)
	DeletePlan(ctx context.Context, id string) error
}

type UserPlanRepo interface {
	CreateUserPlan(ctx context.Context, in entities.UserPlan) (entities.UserPlan, error)
	GetUserPlan(ctx context.Context, id string) (entities.UserPlan, error)
	ListUserPlans(ctx context.Context) ([]entities.UserPlan, error)
	UpdateUserPlan(ctx context.Context, id string, in entities.UserPlan) (entities.UserPlan, error)
	DeleteUserPlan(ctx context.Context, id string) error
}

type IntegrationRepo interface {
	CreateIntegration(ctx context.Context, in entities.Integration) (entities.Integration, error)
	GetIntegration(ctx context.Context, id string) (entities.Integration, error)
	ListIntegrations(ctx context.Context) ([]entities.Integration, error)
	UpdateIntegration(ctx context.Context, id string, in entities.Integration) (entities.Integration, error)
	DeleteIntegration(ctx context.Context, id string) error
}

type MessageRepo interface {
	CreateMessage(ctx context.Context, in entities.Message) (entities.Message, error)
	GetMessage(ctx context.Context, id string) (entities.Message, error)
	ListMessages(ctx context.Context) ([]entities.Message, error)
	UpdateMessage(ctx context.Context, id string, in entities.Message) (entities.Message, error)
	DeleteMessage(ctx context.Context, id string) error
}

type MessageStatusRepo interface {
	CreateMessageStatus(ctx context.Context, in entities.MessageStatus) (entities.MessageStatus, error)
	GetMessageStatus(ctx context.Context, id string) (entities.MessageStatus, error)
	ListMessageStatuses(ctx context.Context) ([]entities.MessageStatus, error)
	UpdateMessageStatus(ctx context.Context, id string, in entities.MessageStatus) (entities.MessageStatus, error)
	DeleteMessageStatus(ctx context.Context, id string) error
}
