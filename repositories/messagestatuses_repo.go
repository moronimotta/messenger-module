package repositories

import (
	"context"
	"errors"
	"time"

	"messenger-module/db"
	"messenger-module/entities"
)

// MessageStatuses CRUD methods
func (r *DBRepository) CreateMessageStatus(ctx context.Context, in entities.MessageStatus) (entities.MessageStatus, error) {
	m := toDBMessageStatus(in)
	if err := r.database.GetDB().WithContext(ctx).Create(&m).Error; err != nil {
		return entities.MessageStatus{}, err
	}
	return toDomainMessageStatus(m), nil
}

func (r *DBRepository) GetMessageStatus(ctx context.Context, id string) (entities.MessageStatus, error) {
	var m db.MessageStatusModel
	if err := r.database.GetDB().WithContext(ctx).First(&m, "id = ?", id).Error; err != nil {
		return entities.MessageStatus{}, err
	}
	return toDomainMessageStatus(m), nil
}

func (r *DBRepository) ListMessageStatuses(ctx context.Context) ([]entities.MessageStatus, error) {
	var rows []db.MessageStatusModel
	if err := r.database.GetDB().WithContext(ctx).Find(&rows).Error; err != nil {
		return nil, err
	}
	out := make([]entities.MessageStatus, 0, len(rows))
	for _, m := range rows {
		out = append(out, toDomainMessageStatus(m))
	}
	return out, nil
}

func (r *DBRepository) UpdateMessageStatus(ctx context.Context, id string, in entities.MessageStatus) (entities.MessageStatus, error) {
	var m db.MessageStatusModel
	if err := r.database.GetDB().WithContext(ctx).First(&m, "id = ?", id).Error; err != nil {
		return entities.MessageStatus{}, err
	}
	if in.Status != "" {
		m.Status = in.Status
	}
	if in.GatewayResponse != "" {
		m.GatewayResponse = in.GatewayResponse
	}
	// Update date fields if provided
	if in.DateSent != nil {
		if t, err := time.Parse(time.RFC3339, *in.DateSent); err == nil {
			m.DateSent = &t
		}
	}
	if in.DateOpened != nil {
		if t, err := time.Parse(time.RFC3339, *in.DateOpened); err == nil {
			m.DateOpened = &t
		}
	}
	if in.DateError != nil {
		if t, err := time.Parse(time.RFC3339, *in.DateError); err == nil {
			m.DateError = &t
		}
	}
	if in.DateCanceled != nil {
		if t, err := time.Parse(time.RFC3339, *in.DateCanceled); err == nil {
			m.DateCanceled = &t
		}
	}
	if in.DateDeferred != nil {
		if t, err := time.Parse(time.RFC3339, *in.DateDeferred); err == nil {
			m.DateDeferred = &t
		}
	}
	if err := r.database.GetDB().WithContext(ctx).Save(&m).Error; err != nil {
		return entities.MessageStatus{}, err
	}
	return toDomainMessageStatus(m), nil
}

func (r *DBRepository) DeleteMessageStatus(ctx context.Context, id string) error {
	res := r.database.GetDB().WithContext(ctx).Delete(&db.MessageStatusModel{}, "id = ?", id)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("not found")
	}
	return nil
}
