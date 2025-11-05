package repositories

import (
	"context"
	"errors"

	"messenger-module/db"
	"messenger-module/entities"
)

// Plans CRUD methods
func (r *DBRepository) CreatePlan(ctx context.Context, in entities.Plan) (entities.Plan, error) {
	m := toDBPlan(in)
	if err := r.database.GetDB().WithContext(ctx).Create(&m).Error; err != nil {
		return entities.Plan{}, err
	}
	return toDomainPlan(m), nil
}

func (r *DBRepository) GetPlan(ctx context.Context, id string) (entities.Plan, error) {
	var m db.PlanModel
	if err := r.database.GetDB().WithContext(ctx).First(&m, "id = ?", id).Error; err != nil {
		return entities.Plan{}, err
	}
	return toDomainPlan(m), nil
}

func (r *DBRepository) ListPlans(ctx context.Context) ([]entities.Plan, error) {
	var rows []db.PlanModel
	if err := r.database.GetDB().WithContext(ctx).Find(&rows).Error; err != nil {
		return nil, err
	}
	out := make([]entities.Plan, 0, len(rows))
	for _, m := range rows {
		out = append(out, toDomainPlan(m))
	}
	return out, nil
}

func (r *DBRepository) UpdatePlan(ctx context.Context, id string, in entities.Plan) (entities.Plan, error) {
	var m db.PlanModel
	if err := r.database.GetDB().WithContext(ctx).First(&m, "id = ?", id).Error; err != nil {
		return entities.Plan{}, err
	}
	if in.Name != "" {
		m.Name = in.Name
	}
	if in.PriceCents != 0 {
		m.PriceCents = in.PriceCents
	}
	if in.ExternalID != "" {
		m.ExternalID = in.ExternalID
	}
	if err := r.database.GetDB().WithContext(ctx).Save(&m).Error; err != nil {
		return entities.Plan{}, err
	}
	return toDomainPlan(m), nil
}

func (r *DBRepository) DeletePlan(ctx context.Context, id string) error {
	res := r.database.GetDB().WithContext(ctx).Delete(&db.PlanModel{}, "id = ?", id)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("not found")
	}
	return nil
}
