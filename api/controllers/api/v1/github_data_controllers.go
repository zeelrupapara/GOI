package v1

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Improwised/GPAT/config"
	"github.com/Improwised/GPAT/constants"
	"github.com/Improwised/GPAT/github"
	"github.com/Improwised/GPAT/models"
	"github.com/Improwised/GPAT/utils"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type GithubDataBody struct {
	StartTime int `json:"start_time"`
	EndTime   int `json:"end_time"`
}

type GithubDataController struct {
	model  *models.Queries
	logger *zap.Logger
	cfg    config.AppConfig
}

func NewGithubDataController(db *sql.DB, logger *zap.Logger, cfg config.AppConfig) (*GithubDataController, error) {
	matrixModel := models.New(db)
	return &GithubDataController{
		model:  matrixModel,
		logger: logger,
		cfg:    cfg,
	}, nil
}

// Fetch all github data by given daterange
func (ctrl *GithubDataController) GetGithubData(c *fiber.Ctx) error {
	githubDataBody := new(GithubDataBody)
	if err := c.BodyParser(githubDataBody); err != nil {
		return utils.JSONError(c, 400, constants.ErrGithubData)
	}
	githubService, err := github.NewGithubService(ctrl.cfg, ctrl.logger)
	if err != nil {
		return utils.JSONError(c, 400, constants.ErrGithubData)
	}
	startTime, err := utils.ConvertIntToTime(githubDataBody.StartTime)
	if err != nil {
		return utils.JSONError(c, 400, constants.ErrGithubData)
	}

	endTime, err := utils.ConvertIntToTime(githubDataBody.EndTime)
	if err != nil {
		return utils.JSONError(c, 400, constants.ErrGithubData)
	}

	// Execute the github command
	weekWiseTime := utils.SplitTimeRange(startTime, endTime, constants.CommandIntervalTime)

	for _, weekTime := range weekWiseTime {
		go func(start, end time.Time) {
			err := githubService.LoadOrganizations(start, end)
			if err != nil {
				fmt.Println(err)
			}
		}(weekTime[0], weekTime[1])
	}

	return utils.JSONSuccess(c, 200, constants.SuccessGetGithubData)
}
