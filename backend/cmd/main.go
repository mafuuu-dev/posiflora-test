package main

import (
	"backend/core/config"
	"backend/core/pkg/errorsx"
	"backend/core/pkg/repository"
	"backend/core/pkg/response"
	"backend/core/pkg/scope"
	"backend/core/pkg/storage"
	"backend/internal/api/routes"
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	cfg := config.Load()

	database := storage.New(ctx, cfg)
	defer database.PG.Close()

	container := scope.New(
		ctx,
		cfg,
		database,
		scope.Support{
			Factory: &scope.Factory{
				Repository: repository.New(ctx, database.PG),
			},
		},
	)

	server := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return response.Error(c, errorsx.Extract(err).Error(), http.StatusBadRequest)
		},
	})

	server.Use(logger.New())
	server.Use(recover.New())
	server.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowMethods:     "GET,POST,PUT,PATCH,OPTIONS",
		AllowHeaders:     "Authorization, Content-Type, X-Requested-With",
		ExposeHeaders:    "Content-Length",
		AllowCredentials: true,
	}))

	routes.Register(server, container)

	go func() {
		<-ctx.Done()

		shutdown, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		if err := server.ShutdownWithContext(shutdown); err != nil {
			log.Printf("error during shutdown: %v", err)
		}
	}()

	fmt.Println("Server started on port 8080")
	if err := server.Listen(":8080"); err != nil {
		log.Fatalf("listen error: %v", err)
	}
}
