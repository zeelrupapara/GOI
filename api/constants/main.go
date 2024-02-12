package constants

// variables
const (
	DateTimeFormat   string = "2006-01-02T15:04:05Z"
	PAGINATION_LIMIT int32  = 5
	OPEN_STATUS      string = "OPEN"
	CLOSED_STATUS    string = "CLOSED"
	MERGED_STATUS    string = "MERGED"
	PR_ALL_STATUS    string = "OPEN,CLOSED,MERGED"
	ISSUE_ALL_STATUS string = "OPEN,CLOSED"
)

// fiber contexts
const ()

// params
const (
	ParamOid    string = "OrgID"
	ParamStatus string = "Status"
	ParamOrg    string = "organization"
	ParamRepo   string = "repository"
	ParamMember string = "member"
)

// query params
const (
	ORG_QP             string = "orgs"
	MEMBER_QP          string = "membs"
	REPO_QP            string = "repos"
	FROM               string = "from"
	TO                 string = "to"
	PR_PAGE_NUMBER     string = "pr_page"
	ISSUE_PAGE_NUMBER  string = "issue_page"
	PR_STATUS          string = "pr_status"
	ISSUE_STATUS       string = "issue_status"
	COMMIT_PAGE_NUMBER string = "commit_page"
)

// Success messages
// ...
const (
	// github data
	SuccessGetGithubData = "Successfully executed github command"
)

// Fail messages
// ...
const ()

// Error messages
const (
	// filters
	ErrGetOrganizationFilter = "Unable to get organization filter options"
	ErrGetMemberFilter       = "Unable to get member filter options"
	ErrGetRepositoryFilter   = "Unable to get repository filter options"

	// matrix
	ErrGetMatrics = "Unable to get matrics"

	// contribution
	ErrGetOrganizationContributions                 = "Unable to get organization contributions"
	ErrGetPullRequestContributions                  = "Unable to get pull request contributions"
	ErrGetIssueContributions                        = "Unable to get issue contributions"
	ErrGetPullRequestContributionInDetailsByFilters = "Unable to get pull request contribution in details by filters"
	ErrGetIssueContributionsInDetailsByFilters      = "Unable to get issue contribution in details by filters"
	ErrNoStatusDefine                               = "No status defined"
	ErrGetUserWiseCommitContribution                = "Unable to get user wise commit contribution"
	ErrGetCommitContributionDetailsByFilters        = "Unable to get commit contribution details by filters"

	// commits
	ErrNotProvideOrganization = "Not provide organization"
	ErrNotProvideRepository   = "Not provide repository"
	ErrNotProvideMember       = "Not provide member"
	ErrGetCommits             = "Unable to get commits"

	// github data
	ErrGithubData = "Unable to get github data"
)
