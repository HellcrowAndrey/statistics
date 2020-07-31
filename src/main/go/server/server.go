package server

import (
	"../config"
	"../handlers"
	"../migration"
	"../services"
	"log"
	"net/http"
)

type Server struct {
	config            *config.Config
	handler           *handlers.RestHandler
	eurekaService     *services.EurekaService
	dataBaseContainer *migration.DataBaseMigration
}

func NewServer(config *config.Config, handler *handlers.RestHandler, eurekaService *services.EurekaService, dataBaseContainer *migration.DataBaseMigration) *Server {
	return &Server{config: config, handler: handler, eurekaService: eurekaService, dataBaseContainer: dataBaseContainer}
}

func (server *Server) Run() {
	httpServer := &http.Server{
		Addr:    ":" + server.config.ServerPort,
		Handler: server.handler.Handler(),
	}
	err := server.dataBaseContainer.RunBuildDataBase()
	if server.config.EurekaConfig.EnableEureka {
		go server.eurekaService.Run()
	}
	err = httpServer.ListenAndServe()
	if err != nil {
		log.Fatal("Server crashed!")
	}
}
