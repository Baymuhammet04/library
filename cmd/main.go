package main

import (
	"context"
	"log"

	"github.com/2004942/library/internal/config"
	controller "github.com/2004942/library/internal/controller/http/v1"
	"github.com/2004942/library/internal/repository/postgres"
	"github.com/2004942/library/internal/usecase"
	"github.com/2004942/library/pkg/connection"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	group := app.Group("/api/v0")

	cfg := config.LoadConfig()

	psqlDB, err := connection.NewDBConnection(context.Background(), cfg.Postgres)
	if err!=nil{
		log.Printf("failed to connection db: %v", err )
	}

	subjectRepo := postgres.NewSubjectRepository(psqlDB)

	subjectUC := usecase.NewSubjectUC(subjectRepo)

	subjectController := controller.NewSubjectUC(subjectUC)

	controller.MapRoutes(group, subjectController)

	app.Listen(":8080")
}