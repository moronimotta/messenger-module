package usecases

import (
	"context"
	"errors"

	"messenger-module/entities"
)

type MessageStatusUsecase struct { repo MessageStatusRepo }

func NewMessageStatusUsecase(repo MessageStatusRepo) *MessageStatusUsecase {
	return &MessageStatusUsecase{repo: repo}
}

func (u *MessageStatusUsecase) Create(ctx context.Context, in entities.MessageStatus) (entities.MessageStatus, error) {
	if in.MessageID == "" || in.Status == "" {
		return entities.MessageStatus{}, errors.New("message_id and status are required")
	}
	return u.repo.CreateMessageStatus(ctx, in)
}
func (u *MessageStatusUsecase) Get(ctx context.Context, id string) (entities.MessageStatus, error) {
	return u.repo.GetMessageStatus(ctx, id)
}
func (u *MessageStatusUsecase) List(ctx context.Context) ([]entities.MessageStatus, error) {
	return u.repo.ListMessageStatuses(ctx)
}
func (u *MessageStatusUsecase) Update(ctx context.Context, id string, in entities.MessageStatus) (entities.MessageStatus, error) {
	return u.repo.UpdateMessageStatus(ctx, id, in)
}
func (u *MessageStatusUsecase) Delete(ctx context.Context, id string) error { return u.repo.DeleteMessageStatus(ctx, id) }
