package usecases

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"messenger-module/entities"
	"messenger-module/handlers"
)

// MessageUsecaseRepo combines all repositories needed by MessageUsecase
type MessageUsecaseRepo interface {
	MessageRepo
	IntegrationRepo
	PlanRepo
	MessageStatusRepo
	UserRepo
	UserPlanRepo
}

type MessageUsecase struct {
	repo           MessageUsecaseRepo
	handlerFactory *handlers.MessageHandlerFactory
}

func NewMessageUsecase(repo MessageUsecaseRepo) *MessageUsecase {
	return &MessageUsecase{
		repo:           repo,
		handlerFactory: handlers.NewMessageHandlerFactory(),
	}
}

func (u *MessageUsecase) Create(ctx context.Context, in entities.Message) (entities.Message, error) {
	// Basic validation - type will be set automatically by the handler
	if in.Content == "" || in.Destination == "" {
		return entities.Message{}, errors.New("content and destination are required")
	}

	// Validate user_id is provided and user exists
	if in.UserID == "" {
		return entities.Message{}, errors.New("user_id is required")
	}
	_, err := u.repo.GetUser(ctx, in.UserID)
	if err != nil {
		return entities.Message{}, fmt.Errorf("user not found: %w", err)
	}

	// Validate integration exists
	if in.IntegrationID == "" {
		return entities.Message{}, errors.New("integration_id is required")
	}
	integration, err := u.repo.GetIntegration(ctx, in.IntegrationID)
	if err != nil {
		return entities.Message{}, fmt.Errorf("integration not found: %w", err)
	}

	// Validate plan permits this integration
	if integration.PlanID == "" {
		return entities.Message{}, errors.New("integration has no associated plan")
	}
	plan, err := u.repo.GetPlan(ctx, integration.PlanID)
	if err != nil {
		return entities.Message{}, fmt.Errorf("plan not found: %w", err)
	}

	// Validate user has access to this plan
	err = u.validateUserPlanAccess(ctx, in.UserID, plan)
	if err != nil {
		return entities.Message{}, err
	}

	// Send via factory (includes validation and handler selection)
	updatedMessage, externalID, err := u.handlerFactory.SendMessage(integration, plan, in)
	if err != nil {
		return entities.Message{}, fmt.Errorf("failed to send message: %w", err)
	}

	// Use the updated message with correct type and external ID
	updatedMessage.ExternalID = externalID

	// Store message in DB
	createdMessage, err := u.repo.CreateMessage(ctx, updatedMessage)
	if err != nil {
		return entities.Message{}, fmt.Errorf("failed to store message: %w", err)
	}

	// For Ntfy messages, automatically create a "sent" status since Ntfy doesn't provide webhooks
	if updatedMessage.Type == "ntfy" {
		messageStatus := entities.MessageStatus{
			MessageID: createdMessage.ID,
			Status:    "sent",
		}
		_, err := u.repo.CreateMessageStatus(ctx, messageStatus)
		if err != nil {
			// Log the error but don't fail the message creation
			// In a real app, you might want to use a proper logger here
			fmt.Printf("Warning: failed to create message status for Ntfy message %s: %v\n", createdMessage.ID, err)
		}
	}

	return createdMessage, nil
}

// validateUserPlanAccess checks if user has access to the required plan
// Pro users have access to both Free and Pro features
// Free users only have access to Free features
func (u *MessageUsecase) validateUserPlanAccess(ctx context.Context, userID string, requiredPlan entities.Plan) error {
	// Get all user plans (active ones)
	userPlans, err := u.repo.ListUserPlans(ctx)
	if err != nil {
		return fmt.Errorf("failed to get user plans: %w", err)
	}

	// Filter for this user's active plans
	var userActivePlans []entities.UserPlan
	for _, up := range userPlans {
		if up.UserID == userID && up.Active {
			userActivePlans = append(userActivePlans, up)
		}
	}

	if len(userActivePlans) == 0 {
		return errors.New("user has no active plans")
	}

	// Get the user's highest plan level
	hasProPlan := false
	hasFreePlan := false

	for _, userPlan := range userActivePlans {
		plan, err := u.repo.GetPlan(ctx, userPlan.PlanID)
		if err != nil {
			continue // Skip invalid plans
		}

		if strings.EqualFold(plan.Name, "pro") {
			hasProPlan = true
		}
		if strings.EqualFold(plan.Name, "free") {
			hasFreePlan = true
		}
	}

	// Check access permissions
	requiredPlanName := requiredPlan.Name

	// Pro users can access everything (Pro and Free features)
	if hasProPlan {
		return nil
	}

	// Free users can only access Free features
	if hasFreePlan && strings.EqualFold(requiredPlanName, "free") {
		return nil
	}

	// User doesn't have required access
	return fmt.Errorf("user does not have access to %s plan features", requiredPlanName)
}

func (u *MessageUsecase) Get(ctx context.Context, id string) (entities.Message, error) {
	return u.repo.GetMessage(ctx, id)
}
func (u *MessageUsecase) List(ctx context.Context) ([]entities.Message, error) {
	return u.repo.ListMessages(ctx)
}
func (u *MessageUsecase) Update(ctx context.Context, id string, in entities.Message) (entities.Message, error) {
	return u.repo.UpdateMessage(ctx, id, in)
}
func (u *MessageUsecase) Delete(ctx context.Context, id string) error {
	return u.repo.DeleteMessage(ctx, id)
}
