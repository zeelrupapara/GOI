package v1

import (
	"database/sql"
	"time"

	"github.com/Improwised/GPAT/constants"
	"github.com/Improwised/GPAT/models"
	"github.com/Improwised/GPAT/utils"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type OrganizationControllers struct {
	model *models.Queries
}

type Organization struct {
	ID              string    `json:"id"`
	Login           string    `json:"login"`
	Name            string    `json:"name"`
	Email           string    `json:"email"`
	Location        string    `json:"location"`
	Description     string    `json:"description"`
	Url             string    `json:"url"`
	AvatarUrl       string    `json:"avatar_url"`
	WebsiteUrl      string    `json:"website_url"`
	GithubUpdatedAt time.Time `json:"github_updated_at"`
	GithubCreatedAt time.Time `json:"github_created_at"`
}

func NewOrganizationController(db *sql.DB, logger *zap.Logger) (*OrganizationControllers, error) {
	orgModel := models.New(db)
	return &OrganizationControllers{
		model: orgModel,
	}, nil
}

func (ctrl *OrganizationControllers) GetOrganizations(c *fiber.Ctx) error {
	var orgs []Organization
	organizations, err := ctrl.model.GetOrganizations(c.Context())
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
		return utils.JSONError(c, 400, constants.ErrGetOrganizations)
	}
	return utils.JSONSuccess(c, 200, orgs)
}
