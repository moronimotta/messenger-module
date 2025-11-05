package httphdl

import (
	"context"
	"messenger-module/entities"
	"messenger-module/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MessageHandler struct {
	uc *usecases.MessageUsecase
}

func NewMessageHandler(uc *usecases.MessageUsecase) *MessageHandler { return &MessageHandler{uc: uc} }

func (h *MessageHandler) Register(rg *gin.RouterGroup) {
	rg.POST("/", h.create)
	rg.GET("/", h.list)
	rg.GET(":id", h.get)
	rg.PUT(":id", h.update)
	rg.DELETE(":id", h.delete)
}

func (h *MessageHandler) create(c *gin.Context) {
	var input entities.Message
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// If this is a redirect, preserve the original request body
	if c.Request.Method == http.MethodPost && c.Request.Header.Get("X-Forwarded-For") != "" {
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid redirected request"})
			return
		}
	}

	message, err := h.uc.Create(c, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, message)
}

func (h *MessageHandler) list(c *gin.Context) {
	out, err := h.uc.List(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, out)
}

func (h *MessageHandler) get(c *gin.Context) {
	id := c.Param("id")
	out, err := h.uc.Get(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, out)
}

func (h *MessageHandler) update(c *gin.Context) {
	id := c.Param("id")
	var in entities.Message
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	out, err := h.uc.Update(c.Request.Context(), id, in)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, out)
}

func (h *MessageHandler) delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.uc.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
