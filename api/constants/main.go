package constants

// variables
const (
	DateTimeFormat string = "2006-01-02T15:04:05Z"
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
)
