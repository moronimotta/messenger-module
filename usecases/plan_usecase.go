package usecases

import (
	"context"
	"errors"
	"strings"

	"messenger-module/entities"
)

type PlanUsecase struct { repo PlanRepo }

func NewPlanUsecase(repo PlanRepo) *PlanUsecase { return &PlanUsecase{repo: repo} }

func (u *PlanUsecase) Create(ctx context.Context, in entities.Plan) (entities.Plan, error) {
	in.Name = strings.TrimSpace(in.Name)
	if in.Name == "" {
		return entities.Plan{}, errors.New("name is required")
	}
	if in.PriceCents <= 0 {
		return entities.Plan{}, errors.New("price_cents must be > 0")
	}
	return u.repo.CreatePlan(ctx, in)
}

func (u *PlanUsecase) Get(ctx context.Context, id string) (entities.Plan, error) { return u.repo.GetPlan(ctx, id) }
func (u *PlanUsecase) List(ctx context.Context) ([]entities.Plan, error) { return u.repo.ListPlans(ctx) }
func (u *PlanUsecase) Update(ctx context.Context, id string, in entities.Plan) (entities.Plan, error) {
	return u.repo.UpdatePlan(ctx, id, in)
}
func (u *PlanUsecase) Delete(ctx context.Context, id string) error { return u.repo.DeletePlan(ctx, id) }
