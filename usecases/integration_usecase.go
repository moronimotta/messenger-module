package usecases

import (
	"context"
	"errors"
	"strings"

	"messenger-module/entities"
)

type IntegrationUsecase struct{ repo IntegrationRepo }

func NewIntegrationUsecase(repo IntegrationRepo) *IntegrationUsecase {
	return &IntegrationUsecase{repo: repo}
}

func (u *IntegrationUsecase) Create(ctx context.Context, in entities.Integration) (entities.Integration, error) {
	in.Name = strings.TrimSpace(in.Name)
	if in.Name == "" {
		return entities.Integration{}, errors.New("name is required")
	}
	in.Type = strings.TrimSpace(in.Type)
	if in.Type == "" {
		return entities.Integration{}, errors.New("type is required")
	}

	return u.repo.CreateIntegration(ctx, in)
}
func (u *IntegrationUsecase) Get(ctx context.Context, id string) (entities.Integration, error) {
	return u.repo.GetIntegration(ctx, id)
}
func (u *IntegrationUsecase) List(ctx context.Context) ([]entities.Integration, error) {
	return u.repo.ListIntegrations(ctx)
}
func (u *IntegrationUsecase) Update(ctx context.Context, id string, in entities.Integration) (entities.Integration, error) {
	return u.repo.UpdateIntegration(ctx, id, in)
}
func (u *IntegrationUsecase) Delete(ctx context.Context, id string) error {
	return u.repo.DeleteIntegration(ctx, id)
}
