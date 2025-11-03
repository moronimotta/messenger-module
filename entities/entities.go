package entities

type User struct {
	ID 	 string `json:"id"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
	DeletedAt string `json:"deleted_at,omitempty"`
	Name string `json:"name"`
	APIKey string `json:"api_key"`
	Active bool `json:"active"`
}

type Plan struct {
	ID 	 string `json:"id"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
	DeletedAt string `json:"deleted_at,omitempty"`
	Name string `json:"name"` // Free, Pro
	// If Free, then only Ntfy Features, IF Pro, then Twilio and SendGrid
	PriceCents int `json:"price_cents"`
	ExternalID string `json:"external_id,omitempty"` // ProductID for PaymentAPI
}

type UserPlan struct {
	ID 	 string `json:"id"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
	DeletedAt string `json:"deleted_at,omitempty"`
	UserID string `json:"user_id"`
	PlanID string `json:"plan_id"`
	Active bool `json:"active"` // If he do not pay, then false
}

type Integration struct {
	ID 	 string `json:"id"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
	DeletedAt string `json:"deleted_at,omitempty"`
	Name string `json:"name"` // Twilio, SendGrid, Ntfy
	APIKey string `json:"api_key"`
}


type Message struct {
	ID 	 string `json:"id"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
	DeletedAt string `json:"deleted_at,omitempty"`
	Type string `json:"type"` // email, sms
	Subject string `json:"subject"`
	Content string `json:"content"`
	Destination string `json:"destination"`
	ExternalID string `json:"external_id,omitempty"`
	Status MessageStatus `json:"status"`
}

type MessageStatus struct {
	ID 	 string `json:"id"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
	DeletedAt string `json:"deleted_at,omitempty"`
	MessageID string `json:"message_id"`
	Status string `json:"status"` // sent, delivered, read
	GatewayResponse string `json:"gateway_response,omitempty"`
	DateSent       *string `json:"date_sent,omitempty"`
	DateOpened     *string `json:"date_opened,omitempty"`
	DateError      *string `json:"date_error,omitempty"`
	DateCanceled   *string `json:"date_canceled,omitempty"`
	DateDeferred   *string `json:"date_deferred,omitempty"`
}

