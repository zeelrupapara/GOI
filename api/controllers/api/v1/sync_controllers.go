package v1

import (
	"database/sql"

	"github.com/Improwised/GPAT/constants"
	"github.com/Improwised/GPAT/models"
	"github.com/Improwised/GPAT/utils"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type SyncControllers struct {
	model *models.Queries
}

func NewSyncController(db *sql.DB, logger *zap.Logger) (*SyncControllers, error) {
	syncModel := models.New(db)
	return &SyncControllers{
		model: syncModel,
	}, nil
}

// Get organization filter option
func (ctrl *SyncControllers) GetSyncedDateWiseData(c *fiber.Ctx) error {
	data, err := ctrl.model.GetSyncDates(c.Context())
	if err != nil {
		return utils.JSONError(c, 400, constants.ErrGetSyncData)
	}
	return utils.JSONSuccess(c, 200, data)
}
