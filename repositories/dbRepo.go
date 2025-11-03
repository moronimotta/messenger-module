package repositories

import (
	"context"
	"errors"
	"time"

	"messenger-module/db"
	"messenger-module/entities"
)

type DBRepository struct {
	database db.Database
}

func NewDBRepository(database db.Database) *DBRepository {
	return &DBRepository{database: database}
}

// Helper mappers between domain and DB models
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
	// CreatedAt/UpdatedAt handled by DB defaults
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
		APIKey:    m.APIKey,
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
		APIKey:    e.APIKey,
	}
}

func toDomainMessage(m db.MessageModel) entities.Message {
	var del string
	if m.DeletedAt != nil {
		del = m.DeletedAt.Format(time.RFC3339)
	}
	return entities.Message{
		ID:          m.ID,
		CreatedAt:   m.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   m.UpdatedAt.Format(time.RFC3339),
		DeletedAt:   del,
		Type:        m.Type,
		Subject:     m.Subject,
		Content:     m.Content,
		Destination: m.Destination,
		ExternalID:  m.ExternalID,
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
		ID:          e.ID,
		DeletedAt:   del,
		Type:        e.Type,
		Subject:     e.Subject,
		Content:     e.Content,
		Destination: e.Destination,
		ExternalID:  e.ExternalID,
	}
}

func toDomainMessageStatus(m db.MessageStatusModel) entities.MessageStatus {
	var del string
	if m.DeletedAt != nil {
		del = m.DeletedAt.Format(time.RFC3339)
	}
	ms := entities.MessageStatus{
		ID:              m.ID,
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

// CRUD methods
// Users
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
	// update fields
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

// Plans
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

// UserPlans
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

// Integrations
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
	if in.APIKey != "" {
		m.APIKey = in.APIKey
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

// Messages
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

// Message Status
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
	// dates
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