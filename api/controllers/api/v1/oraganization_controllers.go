package v1

import (
	"database/sql"

	"github.com/Improwised/GPAT/constants"
	"github.com/Improwised/GPAT/models"
	"github.com/Improwised/GPAT/utils"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type OrganizationControllers struct {
	model *models.Queries
}

func NewOrganizationController(db *sql.DB, logger *zap.Logger) (*OrganizationControllers, error) {
	authorModel := models.New(db)
	return &OrganizationControllers{
		model: authorModel,
	}, nil
}

func (ctrl *OrganizationControllers) GetOrganizations(c *fiber.Ctx) error {
	organizations, err := ctrl.model.GetOrganizationList(c.Context())
	if err != nil {
		return utils.JSONError(c, 400, constants.ErrGetOrganizations)
	}

	return utils.JSONSuccess(c, 200, organizations)
}
