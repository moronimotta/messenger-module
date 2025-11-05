package handlers

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"

	"messenger-module/entities"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

type TwillioHandler struct {
	client *twilio.RestClient
	from   string
	env    string
}

func NewTwillioHandler() *TwillioHandler {
	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")
	fromNumber := os.Getenv("TWILIO_PHONE_NUMBER")
	env := os.Getenv("APP_ENV")

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	return &TwillioHandler{
		client: client,
		from:   fromNumber,
		env:    env,
	}
}

func (h *TwillioHandler) ValidateMessage(input entities.Message) error {
	if input.Destination == "" {
		return errors.New("recipient phone number is required")
	}

	if input.Content == "" {
		return errors.New("message content is required")
	}

	if len(input.Content) > 1600 {
		return errors.New("message content exceeds maximum length of 1600 characters")
	}

	phoneRegex := regexp.MustCompile(`^\+[1-9]\d{6,14}$`)
	if !phoneRegex.MatchString(input.Destination) {
		return fmt.Errorf(
			"invalid phone number format: %s. Must be E.164 (e.g., +14155552671 for US, +447911123456 for UK)",
			input.Destination,
		)
	}

	if h.from != "" && input.Destination == h.from {
		return errors.New("destination cannot be the same as the sender number")
	}

	return nil
}

func (h *TwillioHandler) SendMessage(input entities.Message) (string, error) {
	input.Type = "sms"

	dest := input.Destination
	if h.env == "development" {
		if v := os.Getenv("TWILIO_VIRTUAL_NUMBER"); strings.TrimSpace(v) != "" {
			dest = strings.TrimSpace(v)
		}
	}

	eff := input
	eff.Destination = dest
	if err := h.ValidateMessage(eff); err != nil {
		return "", err
	}

	params := &twilioApi.CreateMessageParams{}

	if base := strings.TrimRight(os.Getenv("WEBHOOK_BASE_URL"), "/"); base != "" {
		params.SetStatusCallback(fmt.Sprintf("%s/api/v1/webhooks/twilio", base))
	}

	params.SetTo(dest)
	params.SetFrom(h.from)
	params.SetBody(input.Content)

	resp, err := h.client.Api.CreateMessage(params)
	if err != nil {
		return "", err
	}
	return *resp.Sid, nil
}
