package server

import (
	"log"
	"messenger-module/db"
	"messenger-module/handlers"
	httphdl "messenger-module/handlers/http"
	"messenger-module/repositories"
	"messenger-module/usecases"

	"github.com/gin-gonic/gin"
)

type Server struct {
	app *gin.Engine
}

func NewServer() *Server {
	return &Server{app: gin.Default()}
}

func (s *Server) Start(database db.Database, port string) error {

	s.app.RedirectFixedPath = true
	s.app.RedirectTrailingSlash = true
	s.app.HandleMethodNotAllowed = true

	s.app.Static("/examples", "./examples")
	// repositories and usecases
	repo := repositories.NewDBRepository(database)
	userUC := usecases.NewUserUsecase(repo)
	planUC := usecases.NewPlanUsecase(repo)
	userPlanUC := usecases.NewUserPlanUsecase(repo)
	integrationUC := usecases.NewIntegrationUsecase(repo)
	messageUC := usecases.NewMessageUsecase(repo)
	messageStatusUC := usecases.NewMessageStatusUsecase(repo)

	// routes
	api := s.app.Group("/api/v1")

	users := api.Group("/users/")
	httphdl.NewUserHandler(userUC, planUC, userPlanUC).Register(users)

	plans := api.Group("/plans/")
	httphdl.NewPlanHandler(planUC).Register(plans)

	userplans := api.Group("/user-plans/")
	httphdl.NewUserPlanHandler(userPlanUC).Register(userplans)

	integrations := api.Group("/integrations/")
	httphdl.NewIntegrationHandler(integrationUC).Register(integrations)

	messages := api.Group("/messages/")
	httphdl.NewMessageHandler(messageUC).Register(messages)

	statuses := api.Group("/message-statuses/")
	httphdl.NewMessageStatusHandler(messageStatusUC).Register(statuses)

	// webhooks
	webhooks := api.Group("/webhooks/")
	httphdl.NewWebhookHandler(messageUC, messageStatusUC).Register(webhooks)

	// emails via SendGrid
	if sg, err := handlers.NewSendGridHandler(); err != nil {
		log.Printf("sendgrid disabled: %v", err)
	} else {
		emails := api.Group("/emails/")
		httphdl.NewSendGridHTTPHandler(sg).Register(emails)
	}

	return s.app.Run(":" + port)
}
