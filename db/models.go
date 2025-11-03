package db

import (
	"time"
)

type UserModel struct {
	ID        string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	CreatedAt time.Time `gorm:"not null;default:now()"`
	UpdatedAt time.Time `gorm:"not null;default:now()"`
	DeletedAt *time.Time
	Name      string `gorm:"not null"`
	APIKey    string `gorm:"not null;unique"`
	Active    bool   `gorm:"not null;default:true"`
}

type PlanModel struct {
	ID         string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	CreatedAt  time.Time `gorm:"not null;default:now()"`
	UpdatedAt  time.Time `gorm:"not null;default:now()"`
	DeletedAt  *time.Time
	Name       string `gorm:"not null;unique"`
	PriceCents int    `gorm:"not null"`
	ExternalID string
}

type UserPlanModel struct {
	ID        string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	CreatedAt time.Time `gorm:"not null;default:now()"`
	UpdatedAt time.Time `gorm:"not null;default:now()"`
	DeletedAt *time.Time
	UserID    string `gorm:"not null;index"`
	PlanID    string `gorm:"not null;index"`
	Active    bool   `gorm:"not null;default:true"`
}

type IntegrationModel struct {
	ID        string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	CreatedAt time.Time `gorm:"not null;default:now()"`
	UpdatedAt time.Time `gorm:"not null;default:now()"`
	DeletedAt *time.Time
	Name      string `gorm:"not null"`
	APIKey    string `gorm:"not null"`
}

type MessageModel struct {
	ID          string    `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	CreatedAt   time.Time `gorm:"not null;default:now()"`
	UpdatedAt   time.Time `gorm:"not null;default:now()"`
	DeletedAt   *time.Time
	Type        string `gorm:"not null"`
	Subject     string
	Content     string `gorm:"not null"`
	Destination string `gorm:"not null"`
	ExternalID  string
}

type MessageStatusModel struct {
	ID               string     `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	CreatedAt        time.Time  `gorm:"not null;default:now()"`
	UpdatedAt        time.Time  `gorm:"not null;default:now()"`
	DeletedAt        *time.Time
	MessageID        string     `gorm:"not null;index"`
	Status           string     `gorm:"not null"`
	GatewayResponse  string
	DateSent         *time.Time
	DateOpened       *time.Time
	DateError        *time.Time
	DateCanceled     *time.Time
	DateDeferred     *time.Time
}
