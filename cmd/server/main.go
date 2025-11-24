package main

import (
	"fmt"
	"log"

	"go-starter/config"
	"go-starter/internal/handlers"
	"go-starter/internal/repository"
	"go-starter/internal/services"
	"go-starter/pkg/database"
	"go-starter/pkg/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, err := database.NewPostgresDB(cfg.Database.DSN())
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	fmt.Println("Connected to database successfully")

	if err := database.RunMigrations(db, "migrations"); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	jwtManager := utils.NewJWTManager(cfg.JWT.Secret, cfg.JWT.ExpireHour)
	userHandler := handlers.NewUserHandler(userService, jwtManager)

	gin.SetMode(cfg.Server.Mode)
	router := handlers.SetupRouter(userHandler, jwtManager)

	addr := fmt.Sprintf(":%s", cfg.Server.Port)
	fmt.Printf("Server starting on %s\n", addr)

	if err := router.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
