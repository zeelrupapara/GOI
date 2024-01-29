package constants

// variables
const (
	DateTimeFormat string = "2006-01-02T15:04:05Z"
	PAGINATION_LIMIT int32 = 5
)

// fiber contexts
const ()

// params
const (
	ParamOid string = "OrgID"
)

// query params
const (
	ORG_QP string = "orgs"
	MEMBER_QP string = "membs"
	REPO_QP string = "repos"
	FROM string = "from"
	TO string = "to"
	PR_PAGE_NUMBER string = "pr_page"
	ISSUE_PAGE_NUMBER string = "issue_page"
)

// Success messages
// ...
const ()

// Fail messages
// ...
const ()

// Error messages
const (
	// filters
	ErrGetOrganizationFilter = "Unable to get organization filter options"
	ErrGetMemberFilter = "Unable to get member filter options"
	ErrGetRepositoryFilter = "Unable to get repository filter options"

	// matrix
	ErrGetMatrics = "Unable to get matrics"

	// contribution
	ErrGetOrganizationContributions = "Unable to get organization contributions"
	ErrGetPullRequestContributions = "Unable to get pull request contributions"
	ErrGetIssueContributions = "Unable to get issue contributions"
	ErrGetPullRequestContributionInDetailsByFilters = "Unable to get pull request contribution in details by filters"
	ErrGetIssueContributionsInDetailsByFilters = "Unable to get issue contribution in details by filters"
)
