package v1

import (
	"database/sql"

	"github.com/Improwised/GPAT/models"
	"github.com/Improwised/GPAT/utils"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type ReportController struct {
	model *models.Queries
}

func NewReportController(db *sql.DB, logger *zap.Logger) (*ReportController, error) {
	reportModel := models.New(db)
	return &ReportController{
		model: reportModel,
	}, nil
}

func (ctrl *ReportController) GetUsersReport(c *fiber.Ctx) error {
	users, err := ctrl.model.GetCollaborators(c.Context())
	if err != nil {
		return utils.JSONError(c, 400, err.Error())
	}
	// for _, user := range users {
	// 	ctrl.model.Get
	// }

	return utils.JSONSuccess(c, 200, users)
}
