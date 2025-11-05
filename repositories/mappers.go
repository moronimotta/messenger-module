package repositories

import (
	"time"

	"messenger-module/db"
	"messenger-module/entities"
)

func toDomainUser(m db.UserModel) entities.User {
	var del string
	if m.DeletedAt != nil {
		del = m.DeletedAt.Format(time.RFC3339)
	}
	return entities.User{
		ID:        m.ID,
		CreatedAt: m.CreatedAt.Format(time.RFC3339),
		UpdatedAt: m.UpdatedAt.Format(time.RFC3339),
		DeletedAt: del,
		Name:      m.Name,
		APIKey:    m.APIKey,
		Active:    m.Active,
	}
}

func toDBUser(e entities.User) db.UserModel {
	var del *time.Time
	if e.DeletedAt != "" {
		if t, err := time.Parse(time.RFC3339, e.DeletedAt); err == nil {
			del = &t
		}
	}
	return db.UserModel{
		ID:        e.ID,
		DeletedAt: del,
		Name:      e.Name,
		APIKey:    e.APIKey,
		Active:    e.Active,
	}
}

func toDomainPlan(m db.PlanModel) entities.Plan {
	var del string
	if m.DeletedAt != nil {
		del = m.DeletedAt.Format(time.RFC3339)
	}
	return entities.Plan{
		ID:         m.ID,
		CreatedAt:  m.CreatedAt.Format(time.RFC3339),
		UpdatedAt:  m.UpdatedAt.Format(time.RFC3339),
		DeletedAt:  del,
		Name:       m.Name,
		PriceCents: m.PriceCents,
		ExternalID: m.ExternalID,
	}
}

func toDBPlan(e entities.Plan) db.PlanModel {
	var del *time.Time
	if e.DeletedAt != "" {
		if t, err := time.Parse(time.RFC3339, e.DeletedAt); err == nil {
			del = &t
		}
	}
	return db.PlanModel{
		ID:         e.ID,
		DeletedAt:  del,
		Name:       e.Name,
		PriceCents: e.PriceCents,
		ExternalID: e.ExternalID,
	}
}

func toDomainUserPlan(m db.UserPlanModel) entities.UserPlan {
	var del string
	if m.DeletedAt != nil {
		del = m.DeletedAt.Format(time.RFC3339)
	}
	return entities.UserPlan{
		ID:        m.ID,
		CreatedAt: m.CreatedAt.Format(time.RFC3339),
		UpdatedAt: m.UpdatedAt.Format(time.RFC3339),
		DeletedAt: del,
		UserID:    m.UserID,
		PlanID:    m.PlanID,
		Active:    m.Active,
	}
}

func toDBUserPlan(e entities.UserPlan) db.UserPlanModel {
	var del *time.Time
	if e.DeletedAt != "" {
		if t, err := time.Parse(time.RFC3339, e.DeletedAt); err == nil {
			del = &t
		}
	}
	return db.UserPlanModel{
		ID:        e.ID,
		DeletedAt: del,
		UserID:    e.UserID,
		PlanID:    e.PlanID,
		Active:    e.Active,
	}
}

func toDomainIntegration(m db.IntegrationModel) entities.Integration {
	var del string
	if m.DeletedAt != nil {
		del = m.DeletedAt.Format(time.RFC3339)
	}
	return entities.Integration{
		ID:        m.ID,
		CreatedAt: m.CreatedAt.Format(time.RFC3339),
		UpdatedAt: m.UpdatedAt.Format(time.RFC3339),
		DeletedAt: del,
		Name:      m.Name,
		Type:      m.Type,
		PlanID:    m.PlanID,
	}
}

func toDBIntegration(e entities.Integration) db.IntegrationModel {
	var del *time.Time
	if e.DeletedAt != "" {
		if t, err := time.Parse(time.RFC3339, e.DeletedAt); err == nil {
			del = &t
		}
	}
	return db.IntegrationModel{
		ID:        e.ID,
		DeletedAt: del,
		Name:      e.Name,
		Type:      e.Type,
		PlanID:    e.PlanID,
	}
}

func toDomainMessage(m db.MessageModel) entities.Message {
	var del string
	if m.DeletedAt != nil {
		del = m.DeletedAt.Format(time.RFC3339)
	}
	return entities.Message{
		ID:            m.ID,
		CreatedAt:     m.CreatedAt.Format(time.RFC3339),
		UpdatedAt:     m.UpdatedAt.Format(time.RFC3339),
		DeletedAt:     del,
		Type:          m.Type,
		Subject:       m.Subject,
		Content:       m.Content,
		Destination:   m.Destination,
		ExternalID:    m.ExternalID,
		UserID:        m.UserID,
		IntegrationID: m.IntegrationID,
	}
}

func toDBMessage(e entities.Message) db.MessageModel {
	var del *time.Time
	if e.DeletedAt != "" {
		if t, err := time.Parse(time.RFC3339, e.DeletedAt); err == nil {
			del = &t
		}
	}
	return db.MessageModel{
		ID:            e.ID,
		DeletedAt:     del,
		Type:          e.Type,
		Subject:       e.Subject,
		Content:       e.Content,
		Destination:   e.Destination,
		ExternalID:    e.ExternalID,
		UserID:        e.UserID,
		IntegrationID: e.IntegrationID,
	}
}

func toDomainMessageStatus(m db.MessageStatusModel) entities.MessageStatus {
	var del string
	if m.DeletedAt != nil {
		del = m.DeletedAt.Format(time.RFC3339)
	}
	ms := entities.MessageStatus{
		ID:              m.ID,
		ExternalID:      m.ExternalID,
		CreatedAt:       m.CreatedAt.Format(time.RFC3339),
		UpdatedAt:       m.UpdatedAt.Format(time.RFC3339),
		DeletedAt:       del,
		MessageID:       m.MessageID,
		Status:          m.Status,
		GatewayResponse: m.GatewayResponse,
	}
	if m.DateSent != nil {
		t := m.DateSent.Format(time.RFC3339)
		ms.DateSent = &t
	}
	if m.DateOpened != nil {
		t := m.DateOpened.Format(time.RFC3339)
		ms.DateOpened = &t
	}
	if m.DateError != nil {
		t := m.DateError.Format(time.RFC3339)
		ms.DateError = &t
	}
	if m.DateCanceled != nil {
		t := m.DateCanceled.Format(time.RFC3339)
		ms.DateCanceled = &t
	}
	if m.DateDeferred != nil {
		t := m.DateDeferred.Format(time.RFC3339)
		ms.DateDeferred = &t
	}
	return ms
}

func toDBMessageStatus(e entities.MessageStatus) db.MessageStatusModel {
	var del *time.Time
	if e.DeletedAt != "" {
		if t, err := time.Parse(time.RFC3339, e.DeletedAt); err == nil {
			del = &t
		}
	}
	var sent, opened, errt, canceled, deferred *time.Time
	if e.DateSent != nil {
		if t, err := time.Parse(time.RFC3339, *e.DateSent); err == nil {
			sent = &t
		}
	}
	if e.DateOpened != nil {
		if t, err := time.Parse(time.RFC3339, *e.DateOpened); err == nil {
			opened = &t
		}
	}
	if e.DateError != nil {
		if t, err := time.Parse(time.RFC3339, *e.DateError); err == nil {
			errt = &t
		}
	}
	if e.DateCanceled != nil {
		if t, err := time.Parse(time.RFC3339, *e.DateCanceled); err == nil {
			canceled = &t
		}
	}
	if e.DateDeferred != nil {
		if t, err := time.Parse(time.RFC3339, *e.DateDeferred); err == nil {
			deferred = &t
		}
	}
	return db.MessageStatusModel{
		ID:              e.ID,
		DeletedAt:       del,
		ExternalID:      e.ExternalID,
		MessageID:       e.MessageID,
		Status:          e.Status,
		GatewayResponse: e.GatewayResponse,
		DateSent:        sent,
		DateOpened:      opened,
		DateError:       errt,
		DateCanceled:    canceled,
		DateDeferred:    deferred,
	}
}
