package github

import (
	"context"
	"fmt"
	"time"

	"github.com/Improwised/GPAT/constants"
	"github.com/Improwised/GPAT/utils"
	"github.com/shurcooL/githubv4"
)

type GithubPRContribution struct {
	PullRequest GithubPullRequestQ
}

type GithubPullRequestQ struct {
	ID        string
	Title     string
	State     string
	Number    int
	IsDraft   bool
	Branch    string `graphql:"headRefName"`
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
		Login string
	}
}

type GithubLabelQ struct {
	ID   string
	Name string
}

type GithubCommitQ struct {
	Commit struct {
		ID            string
		Message       string
		CommittedDate string
		URL           string
		CommitUrl     string
		Committer     struct {
			Name string
		}
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

func (github *GithubService) LoadRepoByPullRequests(orgMember GithubOrgMemberArgs) error {
	var noPages []string
	end, start := utils.GetWeekTimestamps()
	var contributionsLimit githubv4.Int = githubv4.Int(constants.DefaultLimit)
	var contributionsCursor *githubv4.String
	var memberName githubv4.String = githubv4.String(orgMember.Member.Login)
	var reviewsLimit githubv4.Int = githubv4.Int(constants.DefaultLimit)
	var labelsLimit githubv4.Int = githubv4.Int(constants.DefaultLimit)
	var commitsLimit githubv4.Int = githubv4.Int(constants.DefaultLimit)
	var assigneesLimit githubv4.Int = githubv4.Int(constants.DefaultLimit)
	var assigneesCursor *githubv4.String
	var commitsCursor *githubv4.String
	var labelsCursor *githubv4.String
	var reviewsCursor *githubv4.String

	var pullRequestsQ struct {
		User struct {
			ContributionsCollection struct {
				PullRequestContributionsByRepository []struct {
					Repository    GithubRepoQ
					Contributions struct {
						Nodes    []GithubPRContribution
						PageInfo PageInfo
					} `graphql:"contributions(first: $contributionsLimit, after: $contributionsCursor)"`
				}
			} `graphql:"contributionsCollection(organizationID: $orgID, from: $startTime, to: $endTime)"`
		} `graphql:"user(login: $memberLogin)"`
	}

	for {
		// Set the cursor for pagination
		variables := map[string]interface{}{
			"labelsLimit":         labelsLimit,
			"commitsLimit":        commitsLimit,
			"assigneesLimit":      assigneesLimit,
			"labelsCursor":        labelsCursor,
			"commitsCursor":       commitsCursor,
			"assigneesCursor":     assigneesCursor,
			"reviewsLimit":        reviewsLimit,
			"reviewsCursor":       reviewsCursor,
			"contributionsLimit":  contributionsLimit,
			"contributionsCursor": contributionsCursor,
			"startTime":           *githubv4.NewDateTime(githubv4.DateTime{start}),
			"endTime":             *githubv4.NewDateTime(githubv4.DateTime{end}),
			"orgID":               orgMember.ID,
			"memberLogin":         memberName,
		}

		// Execute the graphQL query
		err := github.client.Query(context.Background(), &pullRequestsQ, variables)
		if err != nil {
			fmt.Println("Error executing query:", err)
			return nil
		}
		if len(pullRequestsQ.User.ContributionsCollection.PullRequestContributionsByRepository) == 0 {
			break
		}

		for _, repo := range pullRequestsQ.User.ContributionsCollection.PullRequestContributionsByRepository {

			for _, prContribution := range repo.Contributions.Nodes {
				// fmt.Println(prContribution.PullRequest.Title)

				for _, review := range prContribution.PullRequest.Reviews.Nodes {
					fmt.Println(review.Author.Login)
				}
				for _, labal := range prContribution.PullRequest.Labels.Nodes {
					fmt.Println(labal.Name)
				}
				for _, assignee := range prContribution.PullRequest.Assignees.Nodes {
					fmt.Println(assignee.Login)
				}
				for _, commit := range prContribution.PullRequest.Commits.Nodes {
					fmt.Println(commit.Commit.Committer)
				}

				// reviews page break
				if !prContribution.PullRequest.Reviews.PageInfo.HasNextPage {
					if !utils.Contains("Review", noPages) {
						noPages = append(noPages, "Review")
						contributionsLimit = githubv4.Int(0)
					}
				}
				reviewsCursor = &prContribution.PullRequest.Reviews.PageInfo.EndCursor

				// assignees page break
				if !prContribution.PullRequest.Assignees.PageInfo.HasNextPage {
					if !utils.Contains("Assignee", noPages) {
						noPages = append(noPages, "Assignee")
						reviewsLimit = githubv4.Int(0)
					}
				}
				assigneesCursor = &prContribution.PullRequest.Assignees.PageInfo.EndCursor

				// commit page break
				if !prContribution.PullRequest.Commits.PageInfo.HasNextPage {
					if !utils.Contains("Commit", noPages) {
						noPages = append(noPages, "Commit")
						assigneesLimit = githubv4.Int(0)
					}
				}
				commitsCursor = &prContribution.PullRequest.Commits.PageInfo.EndCursor

				// labal page break
				if !prContribution.PullRequest.Labels.PageInfo.HasNextPage {
					if !utils.Contains("Label", noPages) {
						noPages = append(noPages, "Label")
						labelsLimit = githubv4.Int(0)
					}
				}
				labelsCursor = &prContribution.PullRequest.Labels.PageInfo.EndCursor
			}

			// pullrequest contribution page break
			if !repo.Contributions.PageInfo.HasNextPage {
				if !utils.Contains("PullRequest", noPages) {
					noPages = append(noPages, "PullRequest")
					contributionsLimit = githubv4.Int(0)
				}
			}
			contributionsCursor = &repo.Contributions.PageInfo.EndCursor

		}
		if (len(noPages)) == 5 {
			break
		}
	}
	return nil
}
