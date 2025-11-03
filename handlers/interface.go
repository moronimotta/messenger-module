package handlers

import "messenger-module/entities"

type MessageHandler interface {
	SendMessage(input entities.Message) (string, error)
	ValidateMessage(input entities.Message) error
}
