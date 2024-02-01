package v1

import (
	"database/sql"
	"encoding/json"
	"sort"
	"strconv"
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

type ContributionsDetails struct {
	ID           string    `json:"id"`
	Url          string    `json:"url"`
	Title        string    `json:"title"`
	Status       string    `json:"status"`
	Assignee     string    `json:"assignee"`
	Repository   string    `json:"repository"`
	Organization string    `json:"organization"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type PageInfo struct {
	Previuos bool `json:"previous"`
	Next     bool `json:"next"`
}

type ContributionsDetailsRes struct {
	Details  []ContributionsDetails `json:"details"`
	PageInfo PageInfo               `json:"page_info"`
}

type UserPrCount struct {
	User  string `json:"user"`
	Count int64  `json:"count"`
}

type DateWisePrContributision struct {
	Date time.Time     `json:"date"`
	Data []UserPrCount `json:"data"`
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

// Get Count of User Wise PR by Status
func (ctrl *ContributionControllers) GetPullRequestContributions(c *fiber.Ctx) error {
	var orgs []string
	var repos []string
	var members []string
	var from time.Time
	var to time.Time
	var status string

	// get status from the params
	statusQP := c.Params(constants.ParamStatus)
	if statusQP == "" {
		return utils.JSONError(c, 400, constants.ErrGetIssueContributions)
	}

	status = strings.ToUpper(statusQP)

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
	pullRequestContributions, err := ctrl.model.GetUserWisePullRequestContributionByFilters(c.Context(), models.GetUserWisePullRequestContributionByFiltersParams{
		GithubUpdatedAt:   sql.NullTime{Time: from, Valid: true},
		GithubUpdatedAt_2: sql.NullTime{Time: to, Valid: true},
		StringToArray:     membersStrings,
		StringToArray_2:   orgStrings,
		StringToArray_3:   reposStrings,
		Status:            sql.NullString{String: status, Valid: true},
	})
	if err != nil {
		return utils.JSONError(c, 400, constants.ErrGetPullRequestContributions)
	}

	dateWisePrContributionOutput := make(map[time.Time]*DateWisePrContributision)

	for _, pullRequestContribution := range pullRequestContributions {
		if _, ok := dateWisePrContributionOutput[pullRequestContribution.UserDate]; !ok {
			dateWisePrContributionOutput[pullRequestContribution.UserDate] = &DateWisePrContributision{
				Date: pullRequestContribution.UserDate,
				Data: make([]UserPrCount, 0),
			}
		}
		dateWisePrContributionOutput[pullRequestContribution.UserDate].Data = append(dateWisePrContributionOutput[pullRequestContribution.UserDate].Data, UserPrCount{
			User:  utils.SqlNullString(pullRequestContribution.Login),
			Count: pullRequestContribution.PrCount,
		})
	}

	dateWiseUserPrContributionRes := make([]DateWisePrContributision, 0, len(dateWisePrContributionOutput))
	for _, value := range dateWisePrContributionOutput {
		dateWiseUserPrContributionRes = append(dateWiseUserPrContributionRes, *value)
	}

	sort.Slice(dateWiseUserPrContributionRes, func(i, j int) bool {
		return dateWiseUserPrContributionRes[i].Date.Before(dateWiseUserPrContributionRes[j].Date)
	})

	return utils.JSONSuccess(c, 200, dateWiseUserPrContributionRes)
}

// Get Count of User Wise Issue by Status
func (ctrl *ContributionControllers) GetIssueContributions(c *fiber.Ctx) error {
	var orgs []string
	var repos []string
	var members []string
	var from time.Time
	var to time.Time
	var status string

	// get status from the params
	statusQP := c.Params(constants.ParamStatus)
	if statusQP == "" {
		return utils.JSONError(c, 400, constants.ErrGetIssueContributions)
	}

	status = strings.ToUpper(statusQP)

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
	issueRequestContributions, err := ctrl.model.GetUserWiseIssueContributionByFilters(c.Context(), models.GetUserWiseIssueContributionByFiltersParams{
		GithubUpdatedAt:   sql.NullTime{Time: from, Valid: true},
		GithubUpdatedAt_2: sql.NullTime{Time: to, Valid: true},
		StringToArray:     membersStrings,
		StringToArray_2:   orgStrings,
		StringToArray_3:   reposStrings,
		Status:            status,
	})
	if err != nil {
		return utils.JSONError(c, 400, constants.ErrGetIssueContributions)
	}

	dateWiseIssueContributionOutput := make(map[time.Time]*DateWisePrContributision)

	for _, issueRequestContribution := range issueRequestContributions {
		if _, ok := dateWiseIssueContributionOutput[issueRequestContribution.UserDate]; !ok {
			dateWiseIssueContributionOutput[issueRequestContribution.UserDate] = &DateWisePrContributision{
				Date: issueRequestContribution.UserDate,
				Data: make([]UserPrCount, 0),
			}
		}
		dateWiseIssueContributionOutput[issueRequestContribution.UserDate].Data = append(dateWiseIssueContributionOutput[issueRequestContribution.UserDate].Data, UserPrCount{
			User:  utils.SqlNullString(issueRequestContribution.Login),
			Count: issueRequestContribution.IssueCount,
		})
	}

	dateWiseUserIssueContributionRes := make([]DateWisePrContributision, 0, len(dateWiseIssueContributionOutput))
	for _, value := range dateWiseIssueContributionOutput {
		dateWiseUserIssueContributionRes = append(dateWiseUserIssueContributionRes, *value)
	}

	sort.Slice(dateWiseUserIssueContributionRes, func(i, j int) bool {
		return dateWiseUserIssueContributionRes[i].Date.Before(dateWiseUserIssueContributionRes[j].Date)
	})

	return utils.JSONSuccess(c, 200, dateWiseUserIssueContributionRes)
}

func (ctrl *ContributionControllers) GetPullRequestContributionInDetailsByFilters(c *fiber.Ctx) error {
	var orgs []string
	var repos []string
	var members []string
	var from time.Time
	var to time.Time
	var page int32
	var status string
	var hasPreviousPage bool = true
	var hasNextPage bool = true

	// get status
	statusQP := c.Query(constants.PR_STATUS)
	if statusQP == "" {
		status = "OPEN, CLOSED, MERGED"
	} else {
		status = strings.ToUpper(statusQP)
	}

	// get orgs
	orgsQP := c.Query(constants.ORG_QP)
	if orgsQP == "" {
		orgs, err = ctrl.model.GetOrganizationIDs(c.Context())
		if err != nil {
			return utils.JSONError(c, 400, constants.ErrGetPullRequestContributionInDetailsByFilters)
		}
	} else {
		err = json.Unmarshal([]byte(orgsQP), &orgs)
		if err != nil {
			return utils.JSONError(c, 400, constants.ErrGetPullRequestContributionInDetailsByFilters)
		}
	}
	orgStrings := strings.Join(orgs, ",")

	// get repos
	reposQP := c.Query(constants.REPO_QP)
	if reposQP == "" {
		repos, err = ctrl.model.GetRepoIDs(c.Context())
		if err != nil {
			return utils.JSONError(c, 400, constants.ErrGetPullRequestContributionInDetailsByFilters)
		}
	} else {
		err = json.Unmarshal([]byte(reposQP), &repos)
		if err != nil {
			return utils.JSONError(c, 400, constants.ErrGetPullRequestContributionInDetailsByFilters)
		}
	}
	reposStrings := strings.Join(repos, ",")

	// get membs
	membsQP := c.Query(constants.MEMBER_QP)
	if membsQP == "" {
		members, err = ctrl.model.GetMemberIDs(c.Context())
		if err != nil {
			return utils.JSONError(c, 400, constants.ErrGetPullRequestContributionInDetailsByFilters)
		}
	} else {
		err = json.Unmarshal([]byte(membsQP), &members)
		if err != nil {
			return utils.JSONError(c, 400, constants.ErrGetPullRequestContributionInDetailsByFilters)
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
			return utils.JSONError(c, 400, constants.ErrGetPullRequestContributionInDetailsByFilters)
		}
		to, err = utils.ConvertEpochToTime(toQP)
		if err != nil {
			return utils.JSONError(c, 400, constants.ErrGetPullRequestContributionInDetailsByFilters)
		}
	}

	// Get Page Number
	pageQP := c.Query(constants.PR_PAGE_NUMBER)
	if pageQP == "" {
		page = 1
	} else {
		pageInt, err := strconv.ParseInt(pageQP, 10, 32)
		if err != nil {
			return utils.JSONError(c, 400, constants.ErrGetPullRequestContributionInDetailsByFilters)
		}
		page = int32(pageInt)
	}

	// Get PullRequest Contribution
	prContributionsDetails, err := ctrl.model.GetPullRequestContributionDetailsByFilters(c.Context(), models.GetPullRequestContributionDetailsByFiltersParams{
		GithubUpdatedAt:   sql.NullTime{Time: from, Valid: true},
		GithubUpdatedAt_2: sql.NullTime{Time: to, Valid: true},
		StringToArray:     membersStrings,
		StringToArray_2:   orgStrings,
		StringToArray_3:   reposStrings,
		Limit:             constants.PAGINATION_LIMIT,
		Offset:            constants.PAGINATION_LIMIT * (page - 1),
		StringToArray_4:   status,
	})
	if err != nil {
		return utils.JSONError(c, 400, constants.ErrGetPullRequestContributionInDetailsByFilters)
	}
	prContributionsDetailsStructure := []ContributionsDetails{}
	for _, prContributionDetails := range prContributionsDetails {
		prContributionsDetailsStructure = append(prContributionsDetailsStructure, ContributionsDetails{
			ID:           utils.SqlNullString(prContributionDetails.ID),
			Url:          utils.SqlNullString(prContributionDetails.Url),
			Title:        utils.SqlNullString(prContributionDetails.Title),
			Status:       utils.SqlNullString(prContributionDetails.Status),
			Assignee:     utils.SqlNullString(prContributionDetails.AssigneeName),
			Repository:   utils.SqlNullString(prContributionDetails.RepositoryName),
			Organization: prContributionDetails.OrganizationName,
			UpdatedAt:    utils.SqlNullTime(prContributionDetails.UpdatedAt),
		})
	}

	if page <= 1 {
		hasPreviousPage = false
	}
	if len(prContributionsDetails) < int(constants.PAGINATION_LIMIT) {
		hasNextPage = false
	}

	prContributionsDetailsRes := ContributionsDetailsRes{
		Details:  prContributionsDetailsStructure,
		PageInfo: PageInfo{Previuos: hasPreviousPage, Next: hasNextPage},
	}
	return utils.JSONSuccess(c, 200, prContributionsDetailsRes)
}

func (ctrl *ContributionControllers) GetIssueContributionInDetailsByFilters(c *fiber.Ctx) error {
	var orgs []string
	var repos []string
	var members []string
	var from time.Time
	var to time.Time
	var page int32
	var status string
	var hasPreviousPage bool = true
	var hasNextPage bool = true

	// get status
	statusQP := c.Query(constants.ISSUE_STATUS)
	if statusQP == "" {
		status = "OPEN, CLOSED"
	} else {
		status = strings.ToUpper(statusQP)
	}

	// get orgs
	orgsQP := c.Query(constants.ORG_QP)
	if orgsQP == "" {
		orgs, err = ctrl.model.GetOrganizationIDs(c.Context())
		if err != nil {
			return utils.JSONError(c, 400, constants.ErrGetIssueContributionsInDetailsByFilters)
		}
	} else {
		err = json.Unmarshal([]byte(orgsQP), &orgs)
		if err != nil {
			return utils.JSONError(c, 400, constants.ErrGetIssueContributionsInDetailsByFilters)
		}
	}
	orgStrings := strings.Join(orgs, ",")

	// get repos
	reposQP := c.Query(constants.REPO_QP)
	if reposQP == "" {
		repos, err = ctrl.model.GetRepoIDs(c.Context())
		if err != nil {
			return utils.JSONError(c, 400, constants.ErrGetIssueContributionsInDetailsByFilters)
		}
	} else {
		err = json.Unmarshal([]byte(reposQP), &repos)
		if err != nil {
			return utils.JSONError(c, 400, constants.ErrGetIssueContributionsInDetailsByFilters)
		}
	}
	reposStrings := strings.Join(repos, ",")

	// get membs
	membsQP := c.Query(constants.MEMBER_QP)
	if membsQP == "" {
		members, err = ctrl.model.GetMemberIDs(c.Context())
		if err != nil {
			return utils.JSONError(c, 400, constants.ErrGetIssueContributionsInDetailsByFilters)
		}
	} else {
		err = json.Unmarshal([]byte(membsQP), &members)
		if err != nil {
			return utils.JSONError(c, 400, constants.ErrGetIssueContributionsInDetailsByFilters)
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
			return utils.JSONError(c, 400, constants.ErrGetIssueContributionsInDetailsByFilters)
		}
		to, err = utils.ConvertEpochToTime(toQP)
		if err != nil {
			return utils.JSONError(c, 400, constants.ErrGetIssueContributionsInDetailsByFilters)
		}
	}

	// Get Page Number
	pageQP := c.Query(constants.ISSUE_PAGE_NUMBER)
	if pageQP == "" {
		page = 1
	} else {
		pageInt, err := strconv.ParseInt(pageQP, 10, 32)
		if err != nil {
			return utils.JSONError(c, 400, constants.ErrGetIssueContributionsInDetailsByFilters)
		}
		page = int32(pageInt)
	}

	// Get issue Contribution
	issueContributionsDetails, err := ctrl.model.GetIssueContributionDetailsByFilters(c.Context(), models.GetIssueContributionDetailsByFiltersParams{
		GithubUpdatedAt:   sql.NullTime{Time: from, Valid: true},
		GithubUpdatedAt_2: sql.NullTime{Time: to, Valid: true},
		StringToArray:     membersStrings,
		StringToArray_2:   orgStrings,
		StringToArray_3:   reposStrings,
		Limit:             constants.PAGINATION_LIMIT,
		Offset:            constants.PAGINATION_LIMIT * (page - 1),
		StringToArray_4:   status,
	})
	if err != nil {
		return utils.JSONError(c, 400, constants.ErrGetIssueContributionsInDetailsByFilters)
	}
	issueContributionsDetailsStructure := []ContributionsDetails{}
	for _, issueContributionDetails := range issueContributionsDetails {
		issueContributionsDetailsStructure = append(issueContributionsDetailsStructure, ContributionsDetails{
			ID:           utils.SqlNullString(issueContributionDetails.ID),
			Url:          utils.SqlNullString(issueContributionDetails.Url),
			Title:        utils.SqlNullString(issueContributionDetails.Title),
			Status:       utils.SqlNullString(issueContributionDetails.Status),
			Assignee:     utils.SqlNullString(issueContributionDetails.AssigneeName),
			Repository:   utils.SqlNullString(issueContributionDetails.RepositoryName),
			Organization: issueContributionDetails.OrganizationName,
			UpdatedAt:    utils.SqlNullTime(issueContributionDetails.UpdatedAt),
		})
	}

	if page <= 1 {
		hasPreviousPage = false
	}
	if len(issueContributionsDetails) < int(constants.PAGINATION_LIMIT) {
		hasNextPage = false
	}

	issueContributionsDetailsRes := ContributionsDetailsRes{
		Details:  issueContributionsDetailsStructure,
		PageInfo: PageInfo{Previuos: hasPreviousPage, Next: hasNextPage},
	}
	return utils.JSONSuccess(c, 200, issueContributionsDetailsRes)
}
