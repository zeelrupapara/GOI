package github

import (
	"fmt"
	"time"

	"go.uber.org/zap"
)

type GithubOrgMemberArgs struct {
	ID       string
	Login    string
	Member   GithubMemberQ
	OrgMemID string
}

type GithubRepoQ struct {
	ID            string
	Name          string
	DefaultBranch struct {
		Name string
	} `graphql:"defaultBranchRef"`
	URL         string
	HomepageUrl string
	Description string
	IsPrivate   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	OpenIssues  struct {
		TotalCount int
	} `graphql:"openIssues: issues(states: OPEN)"`
	ClosedIssues struct {
		TotalCount int
	} `graphql:"closeIssue: issues(states: CLOSED)"`
	OpenPRs struct {
		TotalCount int
	} `graphql:"openPRs: pullRequests(states: OPEN)"`
	ClosedPRs struct {
		TotalCount int
	} `graphql:"closedPRs: pullRequests(states: CLOSED)"`
	MergedPRs struct {
		TotalCount int
	} `graphql:"mergedPRs: pullRequests(states: MERGED)"`
}

func (github *GithubService) LoadRepo(orgMember GithubOrgMemberArgs) error {
	err := github.LoadRepoByPullRequests(orgMember)
	if err != nil {
		github.LoadRepoLog(ERROR, err)
		return err
	}
	err = github.LoadRepoByIssues(orgMember)
	if err != nil {
		github.LoadRepoLog(ERROR, err)
		return err
	}
	// err = github.LoadRepoByCommits(orgMember)
	// if err != nil {
	// 	github.LoadRepoLog(ERROR, err)
	// 	return err
	// }
	return nil
}

func (github *GithubService) LoadRepoLog(level string, message interface{}) {
	const path = "repo -> LoadRepo -"
	switch level {
	case DEBUG:
		github.logger.Debug(fmt.Sprintf("%s, %s", path, message))
	case INFO:
		github.logger.Info(fmt.Sprintf("%s, %s", path, message))
	case ERROR:
		github.logger.Error(path, zap.Error(fmt.Errorf("%s", message)))
	case WARNING:
		github.logger.Warn(fmt.Sprintf("%s, %s", path, message))
	}
}
