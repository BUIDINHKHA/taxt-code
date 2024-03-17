package main

import (
	"fmt"
	"megabot/config"
	"megabot/routers"
	"megabot/usecase"
	"net/http"

	"megabot/pkg/logger"

	"github.com/Netflix/go-env"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func main() {
	id := uuid.New().String()
	log := logger.InitLog()
	cfg := loadEnvironment(log)

	config := usecase.NewConfig(
		log,
		cfg,
	)

	// implement api
	log.Info(id, "starting webserver")
	router := routers.InitRouter(
		log,
		cfg,
		config,
	)
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: router,
	}
	log.Info(id, fmt.Sprintf("Application started with port: %v", cfg.Port))
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(id, "Start HTTP Server Failed", err.Error())
	}
}

func loadEnvironment(log *logger.Logger) *config.Environment {
	id := uuid.New().String()

	_ = godotenv.Load()
	cfg := &config.Environment{}
	_, err := env.UnmarshalFromEnviron(cfg)
	if err != nil {
		log.Fatal(id, "error when load env, error", err.Error())
	}
	return cfg
}
