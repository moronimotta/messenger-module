package server

import (
	"messenger-module/db"
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
	httphdl.NewUserHandler(userUC).Register(users)

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

	return s.app.Run(":" + port)
}