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

type FiltersControllers struct {
	model *models.Queries
}

type Member struct {
	ID              string    `json:"id"`
	Login           string    `json:"login"`
	Name            string    `json:"name"`
	Email           string    `json:"email"`
	Url             string    `json:"url"`
	AvatarUrl       string    `json:"avatar_url"`
	WebsiteUrl      string    `json:"website_url"`
	GithubCreatedAt time.Time `json:"github_created_at"`
	GithubUpdatedAt time.Time `json:"github_updated_at"`
}

func NewFiltersController(db *sql.DB, logger *zap.Logger) (*FiltersControllers, error) {
	filtersModel := models.New(db)
	return &FiltersControllers{
		model: filtersModel,
	}, nil
}

func (ctrl *FiltersControllers) GetOrganizationFilterOptions(c *fiber.Ctx) error {
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
		return utils.JSONError(c, 400, constants.ErrGetFilterOrganization)
	}
	return utils.JSONSuccess(c, 200, orgs)
}

func (ctrl *FiltersControllers) GetMemberFilterOptions(c *fiber.Ctx) error {
	var members []Member
	membs, err := ctrl.model.GetMembers(c.Context())
	for _, m := range membs {
		members = append(members, Member{
			ID:              m.ID,
			Login:           m.Login,
			Name:            utils.SqlNullString(m.Name),
			Email:           utils.SqlNullString(m.Email),
			Url:             utils.SqlNullString(m.Url),
			AvatarUrl:       utils.SqlNullString(m.AvatarUrl),
			WebsiteUrl:      utils.SqlNullString(m.WebsiteUrl),
			GithubCreatedAt: utils.SqlNullTime(m.GithubCreatedAt),
			GithubUpdatedAt: utils.SqlNullTime(m.GithubUpdatedAt),
		})
	}
	if err != nil {
		return utils.JSONError(c, 400, constants.ErrGetFilterMembers)
	}
	return utils.JSONSuccess(c, 200, members)
}
