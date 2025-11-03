package httphdl

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"messenger-module/entities"
	"messenger-module/usecases"
)

type PlanHandler struct{ uc *usecases.PlanUsecase }

func NewPlanHandler(uc *usecases.PlanUsecase) *PlanHandler { return &PlanHandler{uc: uc} }

func (h *PlanHandler) Register(rg *gin.RouterGroup) {
	rg.POST("/", h.create)
	rg.GET("/", h.list)
	rg.GET(":id", h.get)
	rg.PUT(":id", h.update)
	rg.DELETE(":id", h.delete)
}

func (h *PlanHandler) create(c *gin.Context) {
	var in entities.Plan
	if err := c.ShouldBindJSON(&in); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return }
	out, err := h.uc.Create(c.Request.Context(), in)
	if err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return }
	c.JSON(http.StatusCreated, out)
}
func (h *PlanHandler) list(c *gin.Context) {
	out, err := h.uc.List(c.Request.Context())
	if err != nil { c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}); return }
	c.JSON(http.StatusOK, out)
}
func (h *PlanHandler) get(c *gin.Context) {
	id := c.Param("id")
	out, err := h.uc.Get(c.Request.Context(), id)
	if err != nil { c.JSON(http.StatusNotFound, gin.H{"error": err.Error()}); return }
	c.JSON(http.StatusOK, out)
}
func (h *PlanHandler) update(c *gin.Context) {
	id := c.Param("id")
	var in entities.Plan
	if err := c.ShouldBindJSON(&in); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return }
	out, err := h.uc.Update(c.Request.Context(), id, in)
	if err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return }
	c.JSON(http.StatusOK, out)
}
func (h *PlanHandler) delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.uc.Delete(c.Request.Context(), id); err != nil { c.JSON(http.StatusNotFound, gin.H{"error": err.Error()}); return }
	c.Status(http.StatusNoContent)
}
