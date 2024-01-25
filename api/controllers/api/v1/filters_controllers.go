package v1

import (
	"database/sql"

	"github.com/Improwised/GPAT/constants"
	"github.com/Improwised/GPAT/models"
	"github.com/Improwised/GPAT/utils"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type FiltersControllers struct {
	model *models.Queries
}

func NewFiltersController(db *sql.DB, logger *zap.Logger) (*FiltersControllers, error) {
	filtersModel := models.New(db)
	return &FiltersControllers{
		model: filtersModel,
	}, nil
}

func (ctrl *FiltersControllers) GetOrganizationFilterOptions(c *fiber.Ctx) error {
	var orgs []Organization
	organizations, err := ctrl.model.GetOrganizationList(c.Context())
	for _, organization := range organizations {
		orgs = append(orgs, Organization{
			ID:              organization.ID,
			Login:           organization.Login,
			Name:            utils.SqlNullString(organization.Name),
			Email:           utils.SqlNullString(organization.Email),
			Location:        utils.SqlNullString(organization.Location),
			Description:     utils.SqlNullString(organization.Description),
			Url:             utils.SqlNullString(organization.Url),
			AvatarUrl:       utils.SqlNullString(organization.AvatarUrl),
			WebsiteUrl:      utils.SqlNullString(organization.WebsiteUrl),
			GithubCreatedAt: utils.SqlNullTime(organization.GithubCreatedAt),
			GithubUpdatedAt: utils.SqlNullTime(organization.GithubUpdatedAt),
		})
	}
	if err != nil {
		return utils.JSONError(c, 400, constants.ErrGetFilterOrganization)
	}
	return utils.JSONSuccess(c, 200, orgs)
}

