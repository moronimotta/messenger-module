package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        string `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at,omitempty"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	APIKey    string `json:"api_key"`
	Active    bool   `json:"active"`
}

type Plan struct {
	ID        string `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at,omitempty"`
	Name      string `json:"name"` // Free, Pro
	// If Free, then only Ntfy Features, IF Pro, then Twilio and SendGrid
	PriceCents int    `json:"price_cents"`
	ExternalID string `json:"external_id,omitempty"` // ProductID for PaymentAPI
}

type UserPlan struct {
	ID        string `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at,omitempty"`
	UserID    string `json:"user_id"`
	PlanID    string `json:"plan_id"`
	Active    bool   `json:"active"` // If he do not pay, then false
}

type Integration struct {
	ID        string `json:"id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at,omitempty"`
	Name      string `json:"name"` // Twilio, SendGrid, Ntfy
	Type      string `json:"type"` // email, phone, ntfy
	PlanID    string `json:"plan_id"`
}

type Message struct {
	ID            string        `json:"id"`
	CreatedAt     string        `json:"created_at"`
	UpdatedAt     string        `json:"updated_at"`
	DeletedAt     string        `json:"deleted_at,omitempty"`
	IntegrationID string        `json:"integration_id"`
	UserID        string        `json:"user_id"`
	Type          string        `json:"type"` // email, sms
	Subject       string        `json:"subject"`
	Content       string        `json:"content"`
	Destination   string        `json:"destination"`
	ExternalID    string        `json:"external_id,omitempty"`
	Status        MessageStatus `json:"status"`
}

type MessageStatus struct {
	ID              string  `json:"id"`
	ExternalID      string  `json:"external_id"`
	CreatedAt       string  `json:"created_at"`
	UpdatedAt       string  `json:"updated_at"`
	DeletedAt       string  `json:"deleted_at,omitempty"`
	MessageID       string  `json:"message_id"`
	Status          string  `json:"status"` // sent, delivered, read
	GatewayResponse string  `json:"gateway_response,omitempty"`
	DateSent        *string `json:"date_sent,omitempty"`
	DateOpened      *string `json:"date_opened,omitempty"`
	DateError       *string `json:"date_error,omitempty"`
	DateCanceled    *string `json:"date_canceled,omitempty"`
	DateDeferred    *string `json:"date_deferred,omitempty"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.ID == "" {
		u.ID = uuid.New().String()
	}

	u.APIKey = uuid.New().String()

	u.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	u.UpdatedAt = u.CreatedAt
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) error {
	u.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	return nil
}

func (p *Plan) BeforeCreate(tx *gorm.DB) error {
	if p.ID == "" {
		p.ID = uuid.New().String()
	}
	p.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	p.UpdatedAt = p.CreatedAt
	return nil
}

func (p *Plan) BeforeUpdate(tx *gorm.DB) error {
	p.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	return nil
}

func (up *UserPlan) BeforeCreate(tx *gorm.DB) error {
	if up.ID == "" {
		up.ID = uuid.New().String()
	}
	up.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	up.UpdatedAt = up.CreatedAt
	return nil
}

func (up *UserPlan) BeforeUpdate(tx *gorm.DB) error {
	up.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	return nil
}

func (i *Integration) BeforeCreate(tx *gorm.DB) error {
	if i.ID == "" {
		i.ID = uuid.New().String()
	}

	i.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	i.UpdatedAt = i.CreatedAt
	return nil
}

func (i *Integration) BeforeUpdate(tx *gorm.DB) error {
	i.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	return nil
}

func (m *Message) BeforeCreate(tx *gorm.DB) error {
	if m.ID == "" {
		m.ID = uuid.New().String()
	}
	m.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	m.UpdatedAt = m.CreatedAt
	return nil
}

func (m *Message) BeforeUpdate(tx *gorm.DB) error {
	m.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	return nil
}

func (ms *MessageStatus) BeforeCreate(tx *gorm.DB) error {
	if ms.ID == "" {
		ms.ID = uuid.New().String()
	}
	ms.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	ms.UpdatedAt = ms.CreatedAt
	return nil
}

func (ms *MessageStatus) BeforeUpdate(tx *gorm.DB) error {
	ms.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
	return nil
}
