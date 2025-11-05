package httphdl

import (
	"net/http"
	"strings"

	"messenger-module/entities"
	"messenger-module/usecases"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUC     *usecases.UserUsecase
	planUC     *usecases.PlanUsecase
	userPlanUC *usecases.UserPlanUsecase
}

func NewUserHandler(userUC *usecases.UserUsecase, planUC *usecases.PlanUsecase, userPlanUC *usecases.UserPlanUsecase) *UserHandler {
	return &UserHandler{userUC: userUC, planUC: planUC, userPlanUC: userPlanUC}
}

func (h *UserHandler) Register(rg *gin.RouterGroup) {
	rg.POST("/", h.create)
	rg.GET("/", h.list)
	rg.GET(":id", h.get)
	rg.PUT(":id", h.update)
	rg.DELETE(":id", h.delete)
}

func (h *UserHandler) create(c *gin.Context) {
	var in entities.User
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 1) Create user
	user, err := h.userUC.Create(c.Request.Context(), in)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 2) Ensure we have a Free plan; create if missing (price 0)
	plans, err := h.planUC.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var freePlan entities.Plan
	for _, p := range plans {
		if strings.EqualFold(p.Name, "free") {
			freePlan = p
			break
		}
	}
	if freePlan.ID == "" {
		// create Free plan with zero price if not exists
		p, err := h.planUC.Create(c.Request.Context(), entities.Plan{Name: "Free", PriceCents: 0})
		if err != nil {
			// Possible race: plan might have been created just now by another request. Re-list once.
			plans2, err2 := h.planUC.List(c.Request.Context())
			if err2 != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to ensure Free plan: " + err.Error()})
				return
			}
			for _, p2 := range plans2 {
				if strings.EqualFold(p2.Name, "free") {
					freePlan = p2
					break
				}
			}
			if freePlan.ID == "" {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to ensure Free plan: " + err.Error()})
				return
			}
		} else {
			freePlan = p
		}
	}
	// 3) Create UserPlan linking to Free plan
	_, err = h.userPlanUC.Create(c.Request.Context(), entities.UserPlan{UserID: user.ID, PlanID: freePlan.ID, Active: true})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user plan: " + err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}
func (h *UserHandler) list(c *gin.Context) {
	out, err := h.userUC.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, out)
}
func (h *UserHandler) get(c *gin.Context) {
	id := c.Param("id")
	out, err := h.userUC.Get(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, out)
}
func (h *UserHandler) update(c *gin.Context) {
	id := c.Param("id")
	var in entities.User
	if err := c.ShouldBindJSON(&in); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	out, err := h.userUC.Update(c.Request.Context(), id, in)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, out)
}
func (h *UserHandler) delete(c *gin.Context) {
	id := c.Param("id")
	if err := h.userUC.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
