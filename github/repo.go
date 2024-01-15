package github

import (
	"time"
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
		return err
	}
	// err = github.LoadRepoByIssues(orgMember)
	// if err != nil {
	// 	return err
	// }
	// err = github.LoadRepoByCommits(orgMember)
	// if err != nil {
	// 	return err
	// }

	return nil
}
