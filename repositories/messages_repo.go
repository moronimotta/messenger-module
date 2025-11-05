package repositories

import (
	"context"
	"errors"

	"messenger-module/db"
	"messenger-module/entities"
)

// Messages CRUD methods
func (r *DBRepository) CreateMessage(ctx context.Context, in entities.Message) (entities.Message, error) {
	m := toDBMessage(in)
	if err := r.database.GetDB().WithContext(ctx).Create(&m).Error; err != nil {
		return entities.Message{}, err
	}
	return toDomainMessage(m), nil
}

func (r *DBRepository) GetMessage(ctx context.Context, id string) (entities.Message, error) {
	var m db.MessageModel
	if err := r.database.GetDB().WithContext(ctx).First(&m, "id = ?", id).Error; err != nil {
		return entities.Message{}, err
	}
	return toDomainMessage(m), nil
}

func (r *DBRepository) ListMessages(ctx context.Context) ([]entities.Message, error) {
	var rows []db.MessageModel
	if err := r.database.GetDB().WithContext(ctx).Find(&rows).Error; err != nil {
		return nil, err
	}
	out := make([]entities.Message, 0, len(rows))
	for _, m := range rows {
		out = append(out, toDomainMessage(m))
	}
	return out, nil
}

func (r *DBRepository) UpdateMessage(ctx context.Context, id string, in entities.Message) (entities.Message, error) {
	var m db.MessageModel
	if err := r.database.GetDB().WithContext(ctx).First(&m, "id = ?", id).Error; err != nil {
		return entities.Message{}, err
	}
	if in.Type != "" {
		m.Type = in.Type
	}
	if in.Subject != "" {
		m.Subject = in.Subject
	}
	if in.Content != "" {
		m.Content = in.Content
	}
	if in.Destination != "" {
		m.Destination = in.Destination
	}
	if in.ExternalID != "" {
		m.ExternalID = in.ExternalID
	}
	if err := r.database.GetDB().WithContext(ctx).Save(&m).Error; err != nil {
		return entities.Message{}, err
	}
	return toDomainMessage(m), nil
}

func (r *DBRepository) DeleteMessage(ctx context.Context, id string) error {
	res := r.database.GetDB().WithContext(ctx).Delete(&db.MessageModel{}, "id = ?", id)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("not found")
	}
	return nil
}
