package httphdl

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"messenger-module/entities"
	"messenger-module/usecases"
)

type UserPlanHandler struct{ uc *usecases.UserPlanUsecase }

func NewUserPlanHandler(uc *usecases.UserPlanUsecase) *UserPlanHandler { return &UserPlanHandler{uc: uc} }

func (h *UserPlanHandler) Register(rg *gin.RouterGroup) {
	rg.POST("/", h.create)
	rg.GET("/", h.list)
	rg.GET(":id", h.get)
	rg.PUT(":id", h.update)
	rg.DELETE(":id", h.delete)
}

func (h *UserPlanHandler) create(c *gin.Context) {
	var in entities.UserPlan
	if err := c.ShouldBindJSON(&in); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return }
	out, err := h.uc.Create(c.Request.Context(), in)
	if err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return }
	c.JSON(http.StatusCreated, out)
}
func (h *UserPlanHandler) list(c *gin.Context) {
	out, err := h.uc.List(c.Request.Context())
	if err != nil { c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}); return }
	c.JSON(http.StatusOK, out)
}
func (h *UserPlanHandler) get(c *gin.Context) {
	id := c.Param("id")
	out, err := h.uc.Get(c.Request.Context(), id)
	if err != nil { c.JSON(http.StatusNotFound, gin.H{"error": err.Error()}); return }
	c.JSON(http.StatusOK, out)
}
func (h *UserPlanHandler) update(c *gin.Context) {
	id := c.Param("id")
	var in entities.UserPlan
	if err := c.ShouldBindJSON(&in); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return }
	out, err := h.uc.Update(c.Request.Context(), id, in)
	if err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return }
	c.JSON(http.StatusOK, out)
}
func (h *UserPlanHandler) delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.uc.Delete(c.Request.Context(), id); err != nil { c.JSON(http.StatusNotFound, gin.H{"error": err.Error()}); return }
	c.Status(http.StatusNoContent)
}
