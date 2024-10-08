package v1

import (
	"database/sql"
	"fmt"

	"github.com/Improwised/GPAT/constants"
	"github.com/Improwised/GPAT/models"
	"github.com/Improwised/GPAT/utils"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type FiltersControllers struct {
	model *models.Queries
}

type Organization struct {
	ID          string `json:"key"`
	Login       string `json:"name"`
	Name        string `json:"full_name,omitempty"`
	Email       string `json:"email,omitempty"`
	Location    string `json:"location,omitempty"`
	Description string `json:"description,omitempty"`
	Url         string `json:"url,omitempty"`
	AvatarUrl   string `json:"avatar_url,omitempty"`
	WebsiteUrl  string `json:"website_url,omitempty"`
}

type Member struct {
	ID         string `json:"key"`
	Login      string `json:"name"`
	Name       string `json:"full_name,omitempty"`
	Email      string `json:"email,omitempty"`
	Url        string `json:"url,omitempty"`
	AvatarUrl  string `json:"avatar_url,omitempty"`
	WebsiteUrl string `json:"website_url,omitempty"`
}

type Repository struct {
	ID            string `json:"key"`
	Name          string `json:"name"`
	IsPrivate     bool   `json:"is_private,omitempty"`
	DefaultBranch string `json:"default_branch,omitempty"`
	Url           string `json:"url,omitempty"`
	HomepageUrl   string `json:"homepage_url,omitempty"`
	OpenIssues    int32  `json:"open_issues,omitempty"`
	ClosedIssues  int32  `json:"closed_issues,omitempty"`
	OpenPrs       int32  `json:"open_prs,omitempty"`
	ClosedPrs     int32  `json:"closed_prs,omitempty"`
	MergedPrs     int32  `json:"merged_prs,omitempty"`
}

func NewFiltersController(db *sql.DB, logger *zap.Logger) (*FiltersControllers, error) {
	filtersModel := models.New(db)
	return &FiltersControllers{
		model: filtersModel,
	}, nil
}

// Get organization filter option
func (ctrl *FiltersControllers) GetOrganizationFilterOptions(c *fiber.Ctx) error {
	var organizations []Organization
	orgs, err := ctrl.model.GetOrganizations(c.Context())
	for _, org := range orgs {
		organizations = append(organizations, Organization{
			ID:    org.ID,
			Login: org.Login,
		})
	}
	if err != nil {
		return utils.JSONError(c, 400, constants.ErrGetOrganizationFilter)
	}
	return utils.JSONSuccess(c, 200, organizations)
}

// Get member filter option
func (ctrl *FiltersControllers) GetMemberFilterOptions(c *fiber.Ctx) error {
	var members []Member
	membs, err := ctrl.model.GetMembers(c.Context())
	for _, m := range membs {
		members = append(members, Member{
			ID:    m.ID,
			Login: m.Login,
		})
	}
	if err != nil {
		return utils.JSONError(c, 400, constants.ErrGetMemberFilter)
	}
	return utils.JSONSuccess(c, 200, members)
}

// Get repository filter option
func (ctrl *FiltersControllers) GetRepositoryFilterOptions(c *fiber.Ctx) error {
	var repositoies []Repository
	repos, err := ctrl.model.GetRepositories(c.Context())
	for _, repo := range repos {
		repository_name := fmt.Sprintf("%s(%s)", utils.SqlNullString(repo.RepoName), repo.OrgLogin)
		repositoies = append(repositoies, Repository{
			ID:   repo.RepoID,
			Name: repository_name,
		})
	}
	if err != nil {
		return utils.JSONError(c, 400, constants.ErrGetRepositoryFilter)
	}
	return utils.JSONSuccess(c, 200, repositoies)
}
