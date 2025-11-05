package repositories

import (
	"context"
	"errors"

	"messenger-module/db"
	"messenger-module/entities"
)

// UserPlans CRUD methods
func (r *DBRepository) CreateUserPlan(ctx context.Context, in entities.UserPlan) (entities.UserPlan, error) {
	m := toDBUserPlan(in)
	if err := r.database.GetDB().WithContext(ctx).Create(&m).Error; err != nil {
		return entities.UserPlan{}, err
	}
	return toDomainUserPlan(m), nil
}

func (r *DBRepository) GetUserPlan(ctx context.Context, id string) (entities.UserPlan, error) {
	var m db.UserPlanModel
	if err := r.database.GetDB().WithContext(ctx).First(&m, "id = ?", id).Error; err != nil {
		return entities.UserPlan{}, err
	}
	return toDomainUserPlan(m), nil
}

func (r *DBRepository) ListUserPlans(ctx context.Context) ([]entities.UserPlan, error) {
	var rows []db.UserPlanModel
	if err := r.database.GetDB().WithContext(ctx).Find(&rows).Error; err != nil {
		return nil, err
	}
	out := make([]entities.UserPlan, 0, len(rows))
	for _, m := range rows {
		out = append(out, toDomainUserPlan(m))
	}
	return out, nil
}

func (r *DBRepository) UpdateUserPlan(ctx context.Context, id string, in entities.UserPlan) (entities.UserPlan, error) {
	var m db.UserPlanModel
	if err := r.database.GetDB().WithContext(ctx).First(&m, "id = ?", id).Error; err != nil {
		return entities.UserPlan{}, err
	}
	if in.UserID != "" {
		m.UserID = in.UserID
	}
	if in.PlanID != "" {
		m.PlanID = in.PlanID
	}
	m.Active = in.Active
	if err := r.database.GetDB().WithContext(ctx).Save(&m).Error; err != nil {
		return entities.UserPlan{}, err
	}
	return toDomainUserPlan(m), nil
}

func (r *DBRepository) DeleteUserPlan(ctx context.Context, id string) error {
	res := r.database.GetDB().WithContext(ctx).Delete(&db.UserPlanModel{}, "id = ?", id)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("not found")
	}
	return nil
}
