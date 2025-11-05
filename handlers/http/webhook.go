package httphdl

import (
	"fmt"
	"net/http"
	"strings"
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

// SendGrid event structure
type sendgridEvent struct {
	Event       string `json:"event"`         // processed, delivered, bounce, etc.
	Email       string `json:"email"`         // recipient email
	Timestamp   int64  `json:"timestamp"`     // unix timestamp
	SMTPId      string `json:"smtp-id"`       // message ID
	SGMessageID string `json:"sg_message_id"` // alternative message ID
	Category    string `json:"category"`
	Reason      string `json:"reason"` // for bounces/drops
	Status      string `json:"status"` // for bounces
	Response    string `json:"response"`
}

func (h *WebhookHandler) sendgrid(c *gin.Context) {
	// SendGrid sends an array of events
	var events []sendgridEvent
	if err := c.ShouldBindJSON(&events); err != nil {
		fmt.Printf("SendGrid webhook error: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid event format"})
		return
	}

	fmt.Printf("SendGrid webhook received %d events\n", len(events))

	// Process each event
	processedCount := 0
	for idx, event := range events {
		fmt.Printf("\n=== Processing SendGrid Event %d/%d ===\n", idx+1, len(events))
		fmt.Printf("Event Type: %s\n", event.Event)
		fmt.Printf("Email: %s\n", event.Email)
		fmt.Printf("SMTP-ID: '%s'\n", event.SMTPId)
		fmt.Printf("SG-Message-ID: '%s'\n", event.SGMessageID)
		fmt.Printf("Timestamp: %d\n", event.Timestamp)

				// Get external ID from event
		externalID := event.SMTPId
		if externalID == "" {
			externalID = event.SGMessageID
		}

		if externalID == "" {
			fmt.Printf("⚠️  No external ID in SendGrid event\n")
			continue
		}

		// Try to extract just the message ID part (SendGrid may prefix with extra data)
		// Format can be like "<TBSQO9UpQ2Cl_s8PtRwPmA@smtp.sendgrid.net>"
		// or "<TBSQO9UpQ2Cl_s8PtRwPmA.filter0001p1mdw1.12345.67890ABCDEF@sendgrid.net>"
		// We need the base64-like part before any @ or .
		extractedID := externalID
		if idx := strings.Index(externalID, "@"); idx != -1 {
			extractedID = externalID[0:idx]
		}
		if idx := strings.Index(extractedID, "."); idx != -1 {
			extractedID = extractedID[0:idx]
		}
		// Remove angle brackets if present
		extractedID = strings.Trim(extractedID, "<>")
		
		fmt.Printf("🔍 Looking for message:\n")
		fmt.Printf("   Raw external_id: '%s'\n", externalID)
		fmt.Printf("   Extracted ID: '%s'\n", extractedID)

		// Map SendGrid event to our status
		var status string
		switch event.Event {
		case "processed":
			status = "sent"
		case "delivered":
			status = "delivered"
		case "open":
			status = "read"
		case "bounce", "dropped", "spamreport":
			status = "error"
		case "deferred":
			status = "deferred"
		default:
			fmt.Printf("Unknown SendGrid event type: %s\n", event.Event)
			continue
		}

		// Find message by external id
		msgs, err := h.msgUC.List(c.Request.Context())
		if err != nil {
			fmt.Printf("Error listing messages: %v\n", err)
			continue
		}

		var msg entities.Message
		var foundWithID string
		for _, m := range msgs {
			// Try exact match with extracted ID first
			if m.ExternalID == extractedID {
				msg = m
				foundWithID = "extracted"
				break
			}
			// Try exact match with raw ID
			if m.ExternalID == externalID {
				msg = m
				foundWithID = "raw"
				break
			}
		}

		if msg.ID == "" {
			fmt.Printf("❌ Message not found in %d messages:\n", len(msgs))
			fmt.Printf("   Extracted ID: '%s'\n", extractedID)
			fmt.Printf("   Raw ID: '%s'\n", externalID)
			// Show first few messages to debug
			fmt.Printf("   Sample message external_ids from database:\n")
			for i, m := range msgs {
				if i >= 3 {
					break
				}
				fmt.Printf("      [%d] '%s'\n", i+1, m.ExternalID)
			}
			continue
		}

		fmt.Printf("✅ Found message %s using %s ID\n", msg.ID, foundWithID)

		// Build gateway response
		gatewayResponse := fmt.Sprintf(
			"Event=%s, Email=%s, SMTPId=%s, SGMessageID=%s, Reason=%s, Response=%s, Status=%s",
			event.Event, event.Email, event.SMTPId, event.SGMessageID,
			event.Reason, event.Response, event.Status,
		)

		var sentAt *string
		if event.Timestamp > 0 {
			t := time.Unix(event.Timestamp, 0).UTC().Format(time.RFC3339)
			sentAt = &t
		}

		statusIn := entities.MessageStatus{
			ExternalID:      externalID,
			MessageID:       msg.ID,
			Status:          status,
			GatewayResponse: gatewayResponse,
			DateSent:        sentAt,
		}

		_, err = h.statUC.Create(c.Request.Context(), statusIn)
		if err != nil {
			fmt.Printf("Error creating status: %v\n", err)
			continue
		}

		processedCount++
	}

	c.JSON(http.StatusOK, gin.H{"ok": true, "processed": processedCount, "total": len(events)})
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
