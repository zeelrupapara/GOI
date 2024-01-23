package routes

import (
	"database/sql"
	"sync"

	"go.uber.org/zap"

	"github.com/Improwised/GPAT/config"
	controller "github.com/Improwised/GPAT/controllers/api/v1"

	"github.com/Improwised/GPAT/middlewares"
	"github.com/gofiber/fiber/v2"
)

var mu sync.Mutex

// Setup func
func Setup(app *fiber.App, db *sql.DB, logger *zap.Logger, config config.AppConfig) error {
	mu.Lock()

	app.Use(middlewares.LogHandler(logger))

	app.Static("/assets/", "./assets")

	app.Get("/docs", func(c *fiber.Ctx) error {
		return c.Render("./assets/index.html", fiber.Map{})
	})

	router := app.Group("/api")
	v1 := router.Group("/v1")

	middlewares := middlewares.NewMiddleware(config, logger)

	err := setupOrganizationController(v1, db, logger, middlewares)
	if err != nil {
		return err
	}

	mu.Unlock()
	return nil
}

func setupOrganizationController(v1 fiber.Router, db *sql.DB, logger *zap.Logger, middlewares middlewares.Middleware) error {
	organizationController, err := controller.NewOrganizationController(db, logger)
	if err != nil {
		return err
	}

	organizationRouter := v1.Group("/organizations")
	organizationRouter.Get("/", organizationController.GetOrganizations)
	return nil
}
