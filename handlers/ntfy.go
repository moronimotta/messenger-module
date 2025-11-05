package handlers

import (
	"errors"
	nethttp "net/http"
	"strings"

	"messenger-module/entities"
)

type NtfyHandler struct{}

func NewNtfyHandler() *NtfyHandler { return &NtfyHandler{} }

func (h *NtfyHandler) ValidateMessage(input entities.Message) error {
	if input.Destination == "" {
		return errors.New("destination is required")
	}
	if input.Content == "" {
		return errors.New("content is required")
	}
	return nil
}

func (h *NtfyHandler) SendMessage(input entities.Message) (string, error) {
	input.Type = "ntfy"

	if err := h.ValidateMessage(input); err != nil {
		return "", err
	}
	httpClient := &nethttp.Client{}
	req, err := nethttp.NewRequest("POST", "https://ntfy.sh/"+input.Destination, strings.NewReader(input.Content))
	if err != nil {
		return "", err
	}
	resp, err := httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	return "", nil
}
