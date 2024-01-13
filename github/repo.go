package github

import (
	"fmt"
	"time"
)

// type GithubMember struct {
// 	ID              string
// 	Login           string
// 	Name            string
// 	Email           string
// 	URL             string
// 	AvatarURL       string    `graphql:"avatarUrl"`
// 	WebsiteURL      string    `graphql:"websiteUrl"`
// 	GithubUpdatedAt time.Time `graphql:"updatedAt"`
// 	GithubCreatedAt time.Time `graphql:"createdAt"`
// }

type GithubOrgMemberArgs struct {
	ID     string
	Login  string
	Member GithubMemberQ
}

type GithubRepoQ struct {
	ID            int
	Name          string
	NameWithOwner string
	DefaultBranch struct {
		Name string
	}
	URL         string
	HomepageUrl string
	Description string
	IsPrivate   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	OpenIssues  struct {
		TotalCount int
	} `graphql:"issues(states: OPEN)"`
	ClosedIssues struct {
		TotalCount int
	} `graphql:"issues(states: CLOSED)"`
	OpenPRs struct {
		TotalCount int
	} `graphql:"pullRequests(states: OPEN)"`
	ClosedPRs struct {
		TotalCount int
	} `graphql:"pullRequests(states: CLOSED)"`
	MergedPRs struct {
		TotalCount int
	} `graphql:"pullRequests(states: MERGED)"`
}

type GithubPullRequestQ struct {
	ID        string
	Title     string
	State     string
	Number    int
	IsDraft   bool
	URL       string
	Reviews   GithubReviewsQ   `graphql:"reviews(first: $reviewsLimit, after: $reviewsCursor)"`
	Labels    GithubLabelsQ    `graphql:"labels(first: $labelsLimit, after: $labelsCursor)"`
	Commits   GithubCommitsQ   `graphql:"commits(first: $commitsLimit, after: $commitsCursor)"`
	Assignees GithubAssigneesQ `graphql:"assignees(first: $assigneesLimit, after: $assigneesCursor)"`
	MergedAt  time.Time
	ClosedAt  time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

type GithubReviewQ struct {
	ID          string
	State       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	SubmittedAt time.Time
	Author      struct {
		Login time.Time
	}
}

type GithubLabelQ struct {
	ID   string
	Name string
}

type GithubCommitQ struct {
	ID            string
	Message       string
	CommittedDate string
	URL           string
	CommitUrl     string
	Author        struct {
		Login time.Time
	}
}

type GithubAssigneeQ struct {
	ID    string
	Login string
}

// nodes with pagination
type GithubAssigneesQ struct {
	Nodes    []GithubAssigneeQ
	PageInfo PageInfo
}

type GithubReviewsQ struct {
	Nodes    []GithubReviewQ
	PageInfo PageInfo
}

type GithubLabelsQ struct {
	Nodes    []GithubLabelQ
	PageInfo PageInfo
}

type GithubCommitsQ struct {
	Nodes    []GithubCommitQ
	PageInfo PageInfo
}

func (github *GithubService) LoadRepo(orgMember GithubOrgMemberArgs) error {
	var repoQ struct {
		User struct {
			ContributionsCollection struct {
				PullRequestContributionsByRepository struct {
					Contributions struct {
						Nodes    []GithubPullRequestQ
						PageInfo PageInfo
					} `graphql:"contributions(first: $contributionsLimit, after: $contributionsCursor)"`
				}
			} `graphql:"contributionsCollection(from: $startTime, to: $endTime, organizationID: $orgID)"`
		} `graphql:"user(login: $memberLogin)"`
	}
	fmt.Println(repoQ)
	return nil
}
