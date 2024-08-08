package routes

import (
	"database/sql"
	"fmt"
	"sync"

	"go.uber.org/zap"

	"github.com/Improwised/GPAT/config"
	"github.com/Improwised/GPAT/constants"
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

	err := setupFiltersController(v1, db, logger, middlewares)
	if err != nil {
		return err
	}

	err = setupMatrixController(v1, db, logger, middlewares)
	if err != nil {
		return err
	}

	err = setupContributionsController(v1, db, logger, middlewares)
	if err != nil {
		return err
	}

	err = setupGithubDataController(v1, db, config, logger, middlewares)
	if err != nil{
		return err
	}

	err = setupSyncDataController(v1, db, config, logger, middlewares)
	if err != nil{
		return err
	}

	mu.Unlock()
	return nil
}

func setupFiltersController(v1 fiber.Router, db *sql.DB, logger *zap.Logger, middlewares middlewares.Middleware) error {
	filtersController, err := controller.NewFiltersController(db, logger)
	if err != nil {
		return err
	}
	filtersRouter := v1.Group("/filters")
	filtersRouter.Get("/organization", filtersController.GetOrganizationFilterOptions)
	filtersRouter.Get("/member", filtersController.GetMemberFilterOptions)
	filtersRouter.Get("/repository", filtersController.GetRepositoryFilterOptions)
	return nil
}

func setupMatrixController(v1 fiber.Router, db *sql.DB, logger *zap.Logger, middlewares middlewares.Middleware) error {
	matrixController, err := controller.NewMatrixController(db, logger)
	if err != nil {
		return err
	}
	matrixRouter := v1.Group("/matrics")
	matrixRouter.Get("/", matrixController.GetMatrics)
	return nil
}

func setupContributionsController(v1 fiber.Router, db *sql.DB, logger *zap.Logger, middlewares middlewares.Middleware) error {
	contributionController, err := controller.NewContributionController(db, logger)
	if err != nil {
		return err
	}
	contributionRouter := v1.Group("/contributions")
	contributionRouter.Get("/organization", contributionController.GetOrganizationContributions)
	contributionRouter.Get(fmt.Sprintf("/pullrequest/status/:%s", constants.ParamStatus), contributionController.GetPullRequestContributions)
	contributionRouter.Get(fmt.Sprintf("/issue/status/:%s", constants.ParamStatus), contributionController.GetIssueContributions)
	contributionRouter.Get("/commit", contributionController.GetCommitContributions)
	contributionRouter.Get("/pullrequest/details", contributionController.GetPullRequestContributionInDetailsByFilters)
	contributionRouter.Get("/issue/details", contributionController.GetIssueContributionInDetailsByFilters)
	contributionRouter.Get("/commit/details", contributionController.GetCommitContributionsDetailsByFilters)
	contributionRouter.Get(fmt.Sprintf("/organizations/:%s/repository/:%s/member/:%s", constants.ParamOrg, constants.ParamRepo, constants.ParamMember), contributionController.GetDefultBranchCommitsByFilters)
	return nil
}

func setupGithubDataController(v1 fiber.Router, db *sql.DB, config config.AppConfig, logger *zap.Logger, middlewares middlewares.Middleware) error {
	githubDataController, err := controller.NewGithubDataController(db, logger, config)
	if err != nil {
		return err
	}
	githubDataRouter := v1.Group("/github")
	githubDataRouter.Post("/data", githubDataController.GetGithubData)
	return nil
}

func setupSyncDataController(v1 fiber.Router, db *sql.DB, config config.AppConfig, logger *zap.Logger, middlewares middlewares.Middleware) error {
	syncDataController, err := controller.NewSyncController(db, logger)
	if err != nil {
		return err
	}
	syncDataRouter := v1.Group("/sync")
	syncDataRouter.Get("/", syncDataController.GetSyncedDateWiseData)
	return nil
}
