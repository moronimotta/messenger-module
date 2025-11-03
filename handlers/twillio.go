package handlers

import (
	"errors"
	"os"
	"regexp"

	"messenger-module/entities"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

type TwillioHandler struct {
	client *twilio.RestClient
	from   string
}

func NewTwillioHandler() *TwillioHandler {
	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")
	fromNumber := os.Getenv("TWILIO_FROM_NUMBER")

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	return &TwillioHandler{
		client: client,
		from:   fromNumber,
	}
}

func (h *TwillioHandler) ValidateMessage(input entities.Message) error {
	if input.Destination == "" {
		return errors.New("recipient phone number is required")
	}

	// Basic phone number validation using regex
	phoneRegex := regexp.MustCompile(`^\+[1-9]\d{1,14}$`)
	if !phoneRegex.MatchString(input.Destination) {
		return errors.New("invalid phone number format. Must be E.164 format (e.g., +1234567890)")
	}

	if input.Destination == "" {
		return errors.New("message body is required")
	}

	if len(input.Destination) > 1600 {
		return errors.New("message body exceeds maximum length of 1600 characters")
	}

	return nil
}

func (h *TwillioHandler) SendMessage(input entities.Message) (string, error) {
	if err := h.ValidateMessage(input); err != nil {
		return "", err
	}

	params := &twilioApi.CreateMessageParams{}
	params.SetTo(input.Destination)
	params.SetFrom(h.from)
	params.SetBody(input.Content)

	resp, err := h.client.Api.CreateMessage(params)
	if err != nil {
		return "", err
	}

	return *resp.Sid, nil
}
