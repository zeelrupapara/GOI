package constants

import "time"

// variables
const (
	DefaultLimit           uint8 = 100
	ReviewRequestInitState       = "PENDING"
	ActivityPR                   = "PR"
	ActivityIssue                = "Issue"
)

// errors messages
const (
	QueryError          string        = "Error while doing query: "
	CommandIntervalTime time.Duration = time.Hour * 24 * 7
)
