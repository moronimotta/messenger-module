package repositories

import (
	"context"
	"errors"

	"messenger-module/db"
	"messenger-module/entities"
)

// Integrations CRUD methods
func (r *DBRepository) CreateIntegration(ctx context.Context, in entities.Integration) (entities.Integration, error) {
	m := toDBIntegration(in)
	if err := r.database.GetDB().WithContext(ctx).Create(&m).Error; err != nil {
		return entities.Integration{}, err
	}
	return toDomainIntegration(m), nil
}

func (r *DBRepository) GetIntegration(ctx context.Context, id string) (entities.Integration, error) {
	var m db.IntegrationModel
	if err := r.database.GetDB().WithContext(ctx).First(&m, "id = ?", id).Error; err != nil {
		return entities.Integration{}, err
	}
	return toDomainIntegration(m), nil
}

func (r *DBRepository) ListIntegrations(ctx context.Context) ([]entities.Integration, error) {
	var rows []db.IntegrationModel
	if err := r.database.GetDB().WithContext(ctx).Find(&rows).Error; err != nil {
		return nil, err
	}
	out := make([]entities.Integration, 0, len(rows))
	for _, m := range rows {
		out = append(out, toDomainIntegration(m))
	}
	return out, nil
}

func (r *DBRepository) UpdateIntegration(ctx context.Context, id string, in entities.Integration) (entities.Integration, error) {
	var m db.IntegrationModel
	if err := r.database.GetDB().WithContext(ctx).First(&m, "id = ?", id).Error; err != nil {
		return entities.Integration{}, err
	}
	if in.Name != "" {
		m.Name = in.Name
	}
	if in.Type != "" {
		m.Type = in.Type
	}
	if in.PlanID != "" {
		m.PlanID = in.PlanID
	}
	if err := r.database.GetDB().WithContext(ctx).Save(&m).Error; err != nil {
		return entities.Integration{}, err
	}
	return toDomainIntegration(m), nil
}

func (r *DBRepository) DeleteIntegration(ctx context.Context, id string) error {
	res := r.database.GetDB().WithContext(ctx).Delete(&db.IntegrationModel{}, "id = ?", id)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("not found")
	}
	return nil
}
