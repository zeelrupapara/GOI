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

var err error

type MatrixControllers struct {
	model *models.Queries
}

type Matrix struct {
	Title string
	Count int64
}

func NewMatrixController(db *sql.DB, logger *zap.Logger) (*MatrixControllers, error) {
	matrixModel := models.New(db)
	return &MatrixControllers{
		model: matrixModel,
	}, nil
}

// Get matrics for the dashboard
func (ctrl *MatrixControllers) GetMatrics(c *fiber.Ctx) error {
	var orgs []string
	var repos []string
	var members []string
	var metrics []Matrix
	var from time.Time
	var to time.Time

	// get orgs
	orgsQP := c.Query(constants.ORG_QP)
	if orgsQP == "" {
		orgs, err = ctrl.model.GetOrganizationIDs(c.Context())
		if err != nil {
			return utils.JSONError(c, 400, constants.ErrGetMatrics)
		}
	} else {
		err = json.Unmarshal([]byte(orgsQP), &orgs)
		if err != nil {
			return utils.JSONError(c, 400, constants.ErrGetMatrics)
		}
	}
	orgStrings := strings.Join(orgs, ",")

	// get repos
	reposQP := c.Query(constants.REPO_QP)
	if reposQP == "" {
		repos, err = ctrl.model.GetRepoIDs(c.Context())
		if err != nil {
			return utils.JSONError(c, 400, constants.ErrGetMatrics)
		}
	} else {
		err = json.Unmarshal([]byte(reposQP), &repos)
		if err != nil {
			return utils.JSONError(c, 400, constants.ErrGetMatrics)
		}
	}
	reposStrings := strings.Join(repos, ",")

	// get membs
	membsQP := c.Query(constants.MEMBER_QP)
	if membsQP == "" {
		members, err = ctrl.model.GetMemberIDs(c.Context())
		if err != nil {
			return utils.JSONError(c, 400, constants.ErrGetMatrics)
		}
	} else {
		err = json.Unmarshal([]byte(membsQP), &members)
		if err != nil {
			return utils.JSONError(c, 400, constants.ErrGetMatrics)
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
			return utils.JSONError(c, 400, constants.ErrGetMatrics)
		}
		to, err = utils.ConvertEpochToTime(toQP)
		if err != nil {
			return utils.JSONError(c, 400, constants.ErrGetMatrics)
		}
	}

	// Get PR Matrix
	prCount, err := ctrl.model.GetPRCountByFilters(c.Context(), models.GetPRCountByFiltersParams{
		GithubUpdatedAt:   sql.NullTime{Time: from, Valid: true},
		GithubUpdatedAt_2: sql.NullTime{Time: to, Valid: true},
		StringToArray:     membersStrings,
		StringToArray_2:   orgStrings,
		StringToArray_3:   reposStrings,
	})
	if err != nil {
		return utils.JSONError(c, 400, constants.ErrGetMatrics)
	}
	metrics = append(metrics, Matrix{
		Title: "Pull Requests",
		Count: prCount,
	})

	// Get Issue Matrix
	issueCount, err := ctrl.model.GetIssueCountByFilters(c.Context(), models.GetIssueCountByFiltersParams{
		GithubUpdatedAt:   sql.NullTime{Time: from, Valid: true},
		GithubUpdatedAt_2: sql.NullTime{Time: to, Valid: true},
		StringToArray:     membersStrings,
		StringToArray_2:   orgStrings,
		StringToArray_3:   reposStrings,
	})
	if err != nil {
		return utils.JSONError(c, 400, constants.ErrGetMatrics)
	}
	metrics = append(metrics, Matrix{
		Title: "Issues",
		Count: issueCount,
	})

	// Get Repo Matrix
	repoCount, err := ctrl.model.GetRepoCountByFilters(c.Context(), models.GetRepoCountByFiltersParams{
		GithubUpdatedAt:   sql.NullTime{Time: from, Valid: true},
		GithubUpdatedAt_2: sql.NullTime{Time: to, Valid: true},
		StringToArray:     membersStrings,
		StringToArray_2:   orgStrings,
		StringToArray_3:   reposStrings,
	})
	if err != nil {
		return utils.JSONError(c, 400, constants.ErrGetMatrics)
	}
	metrics = append(metrics, Matrix{
		Title: "Repositories",
		Count: repoCount,
	})

	// Get Org Matrix
	orgCount, err := ctrl.model.GetOrganizationByFilter(c.Context(), models.GetOrganizationByFilterParams{
		GithubUpdatedAt:   sql.NullTime{Time: from, Valid: true},
		GithubUpdatedAt_2: sql.NullTime{Time: to, Valid: true},
		StringToArray:     membersStrings,
		StringToArray_2:   orgStrings,
		StringToArray_3:   reposStrings,
	})
	if err != nil {
		return utils.JSONError(c, 400, constants.ErrGetMatrics)
	}
	metrics = append(metrics, Matrix{
		Title: "Organizations",
		Count: orgCount,
	})
	return utils.JSONSuccess(c, 200, metrics)
}
