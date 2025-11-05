package httphdl

import (
	"fmt"
	"net/http"
	"time"

	"messenger-module/entities"
	"messenger-module/usecases"

	"github.com/gin-gonic/gin"
)

type WebhookHandler struct {
	msgUC  *usecases.MessageUsecase
	statUC *usecases.MessageStatusUsecase
}

func NewWebhookHandler(msgUC *usecases.MessageUsecase, statUC *usecases.MessageStatusUsecase) *WebhookHandler {
	return &WebhookHandler{msgUC: msgUC, statUC: statUC}
}

func (h *WebhookHandler) Register(rg *gin.RouterGroup) {
	rg.POST("/sendgrid", h.sendgrid)
	rg.POST("/twilio", h.twilio)
}

type genericWebhook struct {
	ExternalID      string `json:"external_id"` // gateway message id
	Status          string `json:"status"`
	GatewayResponse string `json:"gateway_response"`
	Timestamp       int64  `json:"timestamp"`
}

func (h *WebhookHandler) handleGeneric(c *gin.Context, body genericWebhook) {
	if body.ExternalID == "" || body.Status == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "external_id and status are required"})
		return
	}
	// Find message by external id
	msgs, err := h.msgUC.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var msg entities.Message
	for _, m := range msgs {
		if m.ExternalID == body.ExternalID {
			msg = m
			break
		}
	}
	if msg.ID == "" {
		c.JSON(http.StatusNotFound, gin.H{"error": "message not found for external_id"})
		return
	}
	var sentAt *string
	if body.Timestamp > 0 {
		t := time.Unix(body.Timestamp, 0).UTC().Format(time.RFC3339)
		sentAt = &t
	}
	statusIn := entities.MessageStatus{
		ExternalID:      body.ExternalID,
		MessageID:       msg.ID,
		Status:          body.Status,
		GatewayResponse: body.GatewayResponse,
		DateSent:        sentAt,
	}
	_, err = h.statUC.Create(c.Request.Context(), statusIn)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ok": true})
}

func (h *WebhookHandler) sendgrid(c *gin.Context) {
	var body genericWebhook
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.handleGeneric(c, body)
}

func (h *WebhookHandler) twilio(c *gin.Context) {
	// Twilio posts form-encoded by default
	if c.ContentType() == "application/x-www-form-urlencoded" || c.Request.Method == "POST" {
		// Parse form data first
		if err := c.Request.ParseForm(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "failed to parse form"})
			return
		}

		messageSid := c.PostForm("MessageSid")
		messageStatus := c.PostForm("MessageStatus")

		if messageSid == "" || messageStatus == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "MessageSid and MessageStatus are required"})
			return
		}

		// Build comprehensive gateway response from all Twilio fields
		gatewayResponse := fmt.Sprintf(
			"MessageSid=%s, AccountSid=%s, MessagingServiceSid=%s, From=%s, To=%s, Body=%s, Status=%s, ErrorCode=%s, ErrorMessage=%s, DateCreated=%s, DateSent=%s, DateUpdated=%s",
			c.PostForm("MessageSid"),
			c.PostForm("AccountSid"),
			c.PostForm("MessagingServiceSid"),
			c.PostForm("From"),
			c.PostForm("To"),
			c.PostForm("Body"),
			c.PostForm("MessageStatus"),
			c.PostForm("ErrorCode"),
			c.PostForm("ErrorMessage"),
			c.PostForm("DateCreated"),
			c.PostForm("DateSent"),
			c.PostForm("DateUpdated"),
		)

		fmt.Printf("Twilio webhook received: %s\n", gatewayResponse)

		body := genericWebhook{
			ExternalID:      messageSid,
			Status:          messageStatus,
			GatewayResponse: gatewayResponse,
		}
		h.handleGeneric(c, body)
		return
	}

	// Fallback to JSON if not form-encoded
	var body genericWebhook
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.handleGeneric(c, body)
}
