package v1

import (
	"database/sql"
	"encoding/json"
	"strings"
	"time"

	"github.com/Improwised/GPAT/constants"
	"github.com/Improwised/GPAT/models"
	"github.com/Improwised/GPAT/utils"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type ContributionControllers struct {
	model *models.Queries
}

func NewContributionController(db *sql.DB, logger *zap.Logger) (*ContributionControllers, error) {
	contributionModel := models.New(db)
	return &ContributionControllers{
		model: contributionModel,
	}, nil
}

// Get matrics for the dashboard
func (ctrl *ContributionControllers) GetOrganizationContributions(c *fiber.Ctx) error {
	var orgs []string
	var repos []string
	var members []string
	var from time.Time
	var to time.Time

	// get orgs
	orgsQP := c.Query(constants.ORG_QP)
	if orgsQP == "" {
		orgs, err = ctrl.model.GetOrganizationIDs(c.Context())
		if err != nil {
			return utils.JSONError(c, 400, constants.ErrGetOrganizationContributions)
		}
	} else {
		err = json.Unmarshal([]byte(orgsQP), &orgs)
		if err != nil {
			return utils.JSONError(c, 400, constants.ErrGetOrganizationContributions)
		}
	}
	orgStrings := strings.Join(orgs, ",")

	// get repos
	reposQP := c.Query(constants.REPO_QP)
	if reposQP == "" {
		repos, err = ctrl.model.GetRepoIDs(c.Context())
		if err != nil {
			return utils.JSONError(c, 400, constants.ErrGetOrganizationContributions)
		}
	} else {
		err = json.Unmarshal([]byte(reposQP), &repos)
		if err != nil {
			return utils.JSONError(c, 400, constants.ErrGetOrganizationContributions)
		}
	}
	reposStrings := strings.Join(repos, ",")

	// get membs
	membsQP := c.Query(constants.MEMBER_QP)
	if membsQP == "" {
		members, err = ctrl.model.GetMemberIDs(c.Context())
		if err != nil {
			return utils.JSONError(c, 400, constants.ErrGetOrganizationContributions)
		}
	} else {
		err = json.Unmarshal([]byte(membsQP), &members)
		if err != nil {
			return utils.JSONError(c, 400, constants.ErrGetOrganizationContributions)
		}
	}
	membersStrings := strings.Join(members, ",")

	// get the from and to
	fromQP := c.Query(constants.FROM)
	toQP := c.Query(constants.TO)
	if fromQP == "" || toQP == "" {
		// get the 1 week data from the utils
		to, from = utils.GetWeekTimestamps()
	} else {
		from, err = utils.ConvertEpochToTime(fromQP)
		if err != nil {
			return utils.JSONError(c, 400, constants.ErrGetOrganizationContributions)
		}
		to, err = utils.ConvertEpochToTime(toQP)
		if err != nil {
			return utils.JSONError(c, 400, constants.ErrGetOrganizationContributions)
		}
	}

	// Get Organization Contribution
	orgContributions, err := ctrl.model.GetOrganizationContributionsByFilters(c.Context(), models.GetOrganizationContributionsByFiltersParams{
		GithubUpdatedAt:   sql.NullTime{Time: from, Valid: true},
		GithubUpdatedAt_2: sql.NullTime{Time: to, Valid: true},
		StringToArray:     membersStrings,
		StringToArray_2:   orgStrings,
		StringToArray_3:   reposStrings,
	})
	if err != nil {
		return utils.JSONError(c, 400, constants.ErrGetOrganizationContributions)
	}
	return utils.JSONSuccess(c, 200, orgContributions)
}

func (ctrl *ContributionControllers) GetPullRequestContributions(c *fiber.Ctx) error {
	var orgs []string
	var repos []string
	var members []string
	var from time.Time
	var to time.Time

	// get orgs
	orgsQP := c.Query(constants.ORG_QP)
	if orgsQP == "" {
		orgs, err = ctrl.model.GetOrganizationIDs(c.Context())
		if err != nil {
			return utils.JSONError(c, 400, constants.ErrGetPullRequestContributions)
		}
	} else {
		err = json.Unmarshal([]byte(orgsQP), &orgs)
		if err != nil {
			return utils.JSONError(c, 400, constants.ErrGetPullRequestContributions)
		}
	}
	orgStrings := strings.Join(orgs, ",")

	// get repos
	reposQP := c.Query(constants.REPO_QP)
	if reposQP == "" {
		repos, err = ctrl.model.GetRepoIDs(c.Context())
		if err != nil {
			return utils.JSONError(c, 400, constants.ErrGetPullRequestContributions)
		}
	} else {
		err = json.Unmarshal([]byte(reposQP), &repos)
		if err != nil {
			return utils.JSONError(c, 400, constants.ErrGetPullRequestContributions)
		}
	}
	reposStrings := strings.Join(repos, ",")

	// get membs
	membsQP := c.Query(constants.MEMBER_QP)
	if membsQP == "" {
		members, err = ctrl.model.GetMemberIDs(c.Context())
		if err != nil {
			return utils.JSONError(c, 400, constants.ErrGetPullRequestContributions)
		}
	} else {
		err = json.Unmarshal([]byte(membsQP), &members)
		if err != nil {
			return utils.JSONError(c, 400, constants.ErrGetPullRequestContributions)
		}
	}
	membersStrings := strings.Join(members, ",")

	// get the from and to
	fromQP := c.Query(constants.FROM)
	toQP := c.Query(constants.TO)
	if fromQP == "" || toQP == "" {
		// get the 1 week data from the utils
		to, from = utils.GetWeekTimestamps()
	} else {
		from, err = utils.ConvertEpochToTime(fromQP)
		if err != nil {
			return utils.JSONError(c, 400, constants.ErrGetPullRequestContributions)
		}
		to, err = utils.ConvertEpochToTime(toQP)
		if err != nil {
			return utils.JSONError(c, 400, constants.ErrGetPullRequestContributions)
		}
	}

	// Get Pull Request Contribution
	pullRequestContributions, err := ctrl.model.GetPullRequestContributionByFilters(c.Context(), models.GetPullRequestContributionByFiltersParams{
		GithubUpdatedAt:   sql.NullTime{Time: from, Valid: true},
		GithubUpdatedAt_2: sql.NullTime{Time: to, Valid: true},
		StringToArray:     membersStrings,
		StringToArray_2:   orgStrings,
		StringToArray_3:   reposStrings,
	})
	if err != nil {
	    return utils.JSONError(c, 400, constants.ErrGetPullRequestContributions)
	}
	return utils.JSONSuccess(c, 200, pullRequestContributions)
}

func (ctrl *ContributionControllers) GetIssueContributions(c *fiber.Ctx) error {
	var orgs []string
	var repos []string
	var members []string
	var from time.Time
	var to time.Time

	// get orgs
	orgsQP := c.Query(constants.ORG_QP)
	if orgsQP == "" {
		orgs, err = ctrl.model.GetOrganizationIDs(c.Context())
		if err != nil {
			return utils.JSONError(c, 400, constants.ErrGetIssueContributions)
		}
	} else {
		err = json.Unmarshal([]byte(orgsQP), &orgs)
		if err != nil {
			return utils.JSONError(c, 400, constants.ErrGetIssueContributions)
		}
	}
	orgStrings := strings.Join(orgs, ",")

	// get repos
	reposQP := c.Query(constants.REPO_QP)
	if reposQP == "" {
		repos, err = ctrl.model.GetRepoIDs(c.Context())
		if err != nil {
			return utils.JSONError(c, 400, constants.ErrGetIssueContributions)
		}
	} else {
		err = json.Unmarshal([]byte(reposQP), &repos)
		if err != nil {
			return utils.JSONError(c, 400, constants.ErrGetIssueContributions)
		}
	}
	reposStrings := strings.Join(repos, ",")

	// get membs
	membsQP := c.Query(constants.MEMBER_QP)
	if membsQP == "" {
		members, err = ctrl.model.GetMemberIDs(c.Context())
		if err != nil {
			return utils.JSONError(c, 400, constants.ErrGetIssueContributions)
		}
	} else {
		err = json.Unmarshal([]byte(membsQP), &members)
		if err != nil {
			return utils.JSONError(c, 400, constants.ErrGetIssueContributions)
		}
	}
	membersStrings := strings.Join(members, ",")

	// get the from and to
	fromQP := c.Query(constants.FROM)
	toQP := c.Query(constants.TO)
	if fromQP == "" || toQP == "" {
		// get the 1 week data from the utils
		to, from = utils.GetWeekTimestamps()
	} else {
		from, err = utils.ConvertEpochToTime(fromQP)
		if err != nil {
			return utils.JSONError(c, 400, constants.ErrGetIssueContributions)
		}
		to, err = utils.ConvertEpochToTime(toQP)
		if err != nil {
			return utils.JSONError(c, 400, constants.ErrGetIssueContributions)
		}
	}

	// Get Issue Contribution
	issueContributions, err := ctrl.model.GetIssueContributionByFilters(c.Context(), models.GetIssueContributionByFiltersParams{
		GithubUpdatedAt:   sql.NullTime{Time: from, Valid: true},
		GithubUpdatedAt_2: sql.NullTime{Time: to, Valid: true},
		StringToArray:     membersStrings,
		StringToArray_2:   orgStrings,
		StringToArray_3:   reposStrings,
	})
	if err != nil {
	    return utils.JSONError(c, 400, constants.ErrGetIssueContributions)
	}
	return utils.JSONSuccess(c, 200, issueContributions)
}
