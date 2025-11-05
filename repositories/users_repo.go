package repositories

import (
	"context"
	"errors"

	"messenger-module/db"
	"messenger-module/entities"
)

// Users CRUD methods
func (r *DBRepository) CreateUser(ctx context.Context, in entities.User) (entities.User, error) {
	m := toDBUser(in)
	if err := r.database.GetDB().WithContext(ctx).Create(&m).Error; err != nil {
		return entities.User{}, err
	}
	return toDomainUser(m), nil
}

func (r *DBRepository) GetUser(ctx context.Context, id string) (entities.User, error) {
	var m db.UserModel
	if err := r.database.GetDB().WithContext(ctx).First(&m, "id = ?", id).Error; err != nil {
		return entities.User{}, err
	}
	return toDomainUser(m), nil
}

func (r *DBRepository) ListUsers(ctx context.Context) ([]entities.User, error) {
	var rows []db.UserModel
	if err := r.database.GetDB().WithContext(ctx).Find(&rows).Error; err != nil {
		return nil, err
	}
	out := make([]entities.User, 0, len(rows))
	for _, m := range rows {
		out = append(out, toDomainUser(m))
	}
	return out, nil
}

func (r *DBRepository) UpdateUser(ctx context.Context, id string, in entities.User) (entities.User, error) {
	var m db.UserModel
	if err := r.database.GetDB().WithContext(ctx).First(&m, "id = ?", id).Error; err != nil {
		return entities.User{}, err
	}
	// Update fields
	if in.Name != "" {
		m.Name = in.Name
	}
	if in.APIKey != "" {
		m.APIKey = in.APIKey
	}
	m.Active = in.Active
	if err := r.database.GetDB().WithContext(ctx).Save(&m).Error; err != nil {
		return entities.User{}, err
	}
	return toDomainUser(m), nil
}

func (r *DBRepository) DeleteUser(ctx context.Context, id string) error {
	res := r.database.GetDB().WithContext(ctx).Delete(&db.UserModel{}, "id = ?", id)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("not found")
	}
	return nil
}
