package usecases

import (
	"context"
	"errors"

	"messenger-module/entities"
)

type UserPlanUsecase struct { repo UserPlanRepo }

func NewUserPlanUsecase(repo UserPlanRepo) *UserPlanUsecase { return &UserPlanUsecase{repo: repo} }

func (u *UserPlanUsecase) Create(ctx context.Context, in entities.UserPlan) (entities.UserPlan, error) {
	if in.UserID == "" || in.PlanID == "" {
		return entities.UserPlan{}, errors.New("user_id and plan_id are required")
	}
	return u.repo.CreateUserPlan(ctx, in)
}
func (u *UserPlanUsecase) Get(ctx context.Context, id string) (entities.UserPlan, error) { return u.repo.GetUserPlan(ctx, id) }
func (u *UserPlanUsecase) List(ctx context.Context) ([]entities.UserPlan, error) { return u.repo.ListUserPlans(ctx) }
func (u *UserPlanUsecase) Update(ctx context.Context, id string, in entities.UserPlan) (entities.UserPlan, error) {
	return u.repo.UpdateUserPlan(ctx, id, in)
}
func (u *UserPlanUsecase) Delete(ctx context.Context, id string) error { return u.repo.DeleteUserPlan(ctx, id) }
