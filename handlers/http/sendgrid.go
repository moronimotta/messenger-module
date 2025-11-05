package httphdl

import (
	"net/http"

	"messenger-module/entities"
	"messenger-module/handlers"

	"github.com/gin-gonic/gin"
)

// SendGridEmailRequest represents the JSON request body for sending an email
type SendGridEmailRequest struct {
	ToName      string `json:"to_name"`
	ToEmail     string `json:"to_email" binding:"required,email"`
	Subject     string `json:"subject" binding:"required"`
	PlainText   string `json:"plain_text" binding:"required"`
	HTMLContent string `json:"html_content"`
}

type SendGridHTTPHandler struct {
	sender *handlers.SendGridHandler
}

func NewSendGridHTTPHandler(sender *handlers.SendGridHandler) *SendGridHTTPHandler {
	return &SendGridHTTPHandler{sender: sender}
}

func (h *SendGridHTTPHandler) Register(rg *gin.RouterGroup) {
	rg.POST("/send", h.sendEmail)
}

func (h *SendGridHTTPHandler) sendEmail(c *gin.Context) {
	var req SendGridEmailRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Build entities.Message to send via MessageHandler interface
	msg := entities.Message{
		Type:        "email",
		Subject:     req.Subject,
		Content:     req.PlainText,
		Destination: req.ToEmail,
	}

	_, err := h.sender.SendMessage(msg)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Email sent successfully"})
}
