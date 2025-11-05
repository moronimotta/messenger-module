package handlers

import (
	"errors"
	"fmt"
	"strings"

	"messenger-module/entities"
)

type MessageHandlerFactory struct {
	sendgridHandler *SendGridHandler
	twillioHandler  *TwillioHandler
	ntfyHandler     *NtfyHandler
}

func NewMessageHandlerFactory() *MessageHandlerFactory {
	factory := &MessageHandlerFactory{}

	if sg, err := NewSendGridHandler(); err == nil {
		factory.sendgridHandler = sg
	}

	factory.twillioHandler = NewTwillioHandler()

	factory.ntfyHandler = NewNtfyHandler()

	return factory
}

func (f *MessageHandlerFactory) GetHandler(integration entities.Integration) (MessageHandler, error) {
	switch strings.ToLower(integration.Name) {
	case "sendgrid":
		if f.sendgridHandler == nil {
			return nil, errors.New("sendgrid handler not configured - missing API key")
		}
		return f.sendgridHandler, nil
	case "twilio":
		if f.twillioHandler == nil {
			return nil, errors.New("twilio handler not configured")
		}
		return f.twillioHandler, nil
	case "ntfy":
		if f.ntfyHandler == nil {
			return nil, errors.New("ntfy handler not configured")
		}
		return f.ntfyHandler, nil
	default:
		return nil, fmt.Errorf("unknown integration name: %s", integration.Name)
	}
}

func (f *MessageHandlerFactory) SendMessage(integration entities.Integration, plan entities.Plan, message entities.Message) (entities.Message, string, error) {
	switch strings.ToLower(integration.Name) {
	case "sendgrid":
		message.Type = "email"
	case "twilio":
		message.Type = "sms"
	case "ntfy":
		message.Type = "ntfy"
	}

	if strings.EqualFold(plan.Name, "free") && message.Type != "ntfy" {
		return message, "", errors.New("free plan only supports ntfy messages")
	}

	handler, err := f.GetHandler(integration)
	if err != nil {
		return message, "", err
	}

	externalID, err := handler.SendMessage(message)
	return message, externalID, err
}

func (f *MessageHandlerFactory) IsHandlerAvailable(integrationName string) bool {
	switch strings.ToLower(integrationName) {
	case "sendgrid":
		return f.sendgridHandler != nil
	case "twilio":
		return f.twillioHandler != nil
	case "ntfy":
		return f.ntfyHandler != nil
	default:
		return false
	}
}

func (f *MessageHandlerFactory) ListAvailableHandlers() []string {
	var available []string

	if f.sendgridHandler != nil {
		available = append(available, "sendgrid")
	}
	if f.twillioHandler != nil {
		available = append(available, "twilio")
	}
	if f.ntfyHandler != nil {
		available = append(available, "ntfy")
	}

	return available
}

func (f *MessageHandlerFactory) ListAvailableHandlersForPlan(plan entities.Plan) []string {
	var available []string

	if strings.EqualFold(plan.Name, "free") {
		if f.ntfyHandler != nil {
			available = append(available, "ntfy")
		}
		return available
	}

	return f.ListAvailableHandlers()
}
