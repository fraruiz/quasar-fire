package server

import (
	"fmt"
	"log"

	"github.com/franciscoruizar/quasar-fire/internal/infrastructure/server/handler/health"
	"github.com/franciscoruizar/quasar-fire/internal/infrastructure/server/handler/topsecret"
	"github.com/franciscoruizar/quasar-fire/internal/usecases"
	"github.com/gin-gonic/gin"
)

type Server struct {
	httpAddr         string
	engine           *gin.Engine
	sateliteFinder   usecases.SateliteFinder
	topSecretCreator usecases.TopSecretCreator
}

func New(host string, port uint, topSecretCreator usecases.TopSecretCreator, sateliteFinder usecases.SateliteFinder) Server {
	srv := Server{
		engine:           gin.New(),
		httpAddr:         fmt.Sprintf("%s:%d", host, port),
		topSecretCreator: topSecretCreator,
		sateliteFinder:   sateliteFinder,
	}

	srv.registerRoutes()
	return srv
}

func (s *Server) Run() error {
	log.Println("Server running on", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}

func (s *Server) registerRoutes() {
	s.engine.GET("/health", health.CheckHandler())
	s.engine.POST("/topsecret", topsecret.TopSecretCreateHandler(s.topSecretCreator))
}
