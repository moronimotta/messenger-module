package usecases

import (
	"context"
	"errors"
	"strings"

	"messenger-module/entities"
)

type MessageUsecase struct { repo MessageRepo }

func NewMessageUsecase(repo MessageRepo) *MessageUsecase { return &MessageUsecase{repo: repo} }

func (u *MessageUsecase) Create(ctx context.Context, in entities.Message) (entities.Message, error) {
	in.Type = strings.TrimSpace(in.Type)
	if in.Type == "" {
		return entities.Message{}, errors.New("type is required")
	}
	if in.Content == "" || in.Destination == "" {
		return entities.Message{}, errors.New("content and destination are required")
	}
	return u.repo.CreateMessage(ctx, in)
}
func (u *MessageUsecase) Get(ctx context.Context, id string) (entities.Message, error) { return u.repo.GetMessage(ctx, id) }
func (u *MessageUsecase) List(ctx context.Context) ([]entities.Message, error) { return u.repo.ListMessages(ctx) }
func (u *MessageUsecase) Update(ctx context.Context, id string, in entities.Message) (entities.Message, error) {
	return u.repo.UpdateMessage(ctx, id, in)
}
func (u *MessageUsecase) Delete(ctx context.Context, id string) error { return u.repo.DeleteMessage(ctx, id) }
