package handlers

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"messenger-module/entities"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type SendGridHandler struct {
	client    *sendgrid.Client
	fromName  string
	fromEmail string
}

func NewSendGridHandler() (*SendGridHandler, error) {
	apiKey := os.Getenv("SENDGRID_API_KEY")
	if apiKey == "" {
		return nil, errors.New("SENDGRID_API_KEY is required")
	}

	fromName := os.Getenv("SENDGRID_FROM_NAME")
	fromEmail := os.Getenv("SENDGRID_FROM_EMAIL")
	if fromEmail == "" {
		return nil, errors.New("SENDGRID_FROM_EMAIL is required")
	}

	client := sendgrid.NewSendClient(apiKey)

	return &SendGridHandler{
		client:    client,
		fromName:  fromName,
		fromEmail: fromEmail,
	}, nil
}

func (h *SendGridHandler) SendMessage(input entities.Message) (string, error) {
	input.Type = "email"

	if err := h.ValidateMessage(input); err != nil {
		return "", err
	}

	from := mail.NewEmail(h.fromName, h.fromEmail)
	to := mail.NewEmail("", input.Destination)
	message := mail.NewSingleEmail(from, "Message from API", to, input.Content, input.Content)

	// Set unique message ID for tracking
	message.SetHeader("X-Message-ID", fmt.Sprintf("msg_%s", strings.Replace(input.ID, "-", "", -1)))

	response, err := h.client.Send(message)
	if err != nil {
		return "", fmt.Errorf("failed to send email: %w", err)
	}

	if response.StatusCode >= 400 {
		return "", fmt.Errorf("sendgrid error: status=%d body=%s", response.StatusCode, response.Body)
	}

	// Extract message ID from headers
	msgID := response.Headers["X-Message-Id"]
	if len(msgID) == 0 {
		return "", errors.New("no message ID in response")
	}

	return msgID[0], nil
}

func (h *SendGridHandler) ValidateMessage(input entities.Message) error {
	if input.Destination == "" {
		return errors.New("destination (email) is required")
	}

	if !strings.Contains(input.Destination, "@") {
		return errors.New("invalid email format")
	}

	if input.Content == "" {
		return errors.New("content is required")
	}

	return nil
}
