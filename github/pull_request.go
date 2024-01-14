package github

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Improwised/GPAT/constants"
	"github.com/Improwised/GPAT/models"
	"github.com/Improwised/GPAT/utils"
	"github.com/shurcooL/githubv4"
)

type GithubPRContribution struct {
	PullRequest GithubPullRequestQ
}

type GithubPullRequestQ struct {
	ID      string
	Title   string
	State   string
	Number  int
	IsDraft bool
	Branch  string `graphql:"headRefName"`
	Author  struct {
		Login string
	}
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
		CommittedDate time.Time
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
	var ActivityType string = "PR"
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

			// Check repo exist or not?
			repoID, err := github.model.GetRepoByID(github.ctx, repo.Repository.ID)
			if err != nil {
				if err == sql.ErrNoRows {
					repoID, err = github.model.InsertRepo(github.ctx, models.InsertRepoParams{
						ID:                         repo.Repository.ID,
						Name:                       sql.NullString{String: repo.Repository.Name, Valid: true},
						IsPrivate:                  sql.NullBool{Bool: repo.Repository.IsPrivate, Valid: true},
						DefaultBranch:              sql.NullString{String: repo.Repository.DefaultBranch.Name},
						Url:                        sql.NullString{String: repo.Repository.URL},
						HomepageUrl:                sql.NullString{String: repo.Repository.HomepageUrl},
						OpenIssues:                 sql.NullInt32{Int32: int32(repo.Repository.OpenIssues.TotalCount), Valid: true},
						ClosedIssues:               sql.NullInt32{Int32: int32(repo.Repository.ClosedIssues.TotalCount), Valid: true},
						OpenPrs:                    sql.NullInt32{Int32: int32(repo.Repository.OpenPRs.TotalCount), Valid: true},
						ClosedPrs:                  sql.NullInt32{Int32: int32(repo.Repository.ClosedPRs.TotalCount), Valid: true},
						MergedPrs:                  sql.NullInt32{Int32: int32(repo.Repository.MergedPRs.TotalCount), Valid: true},
						OrganizationCollaboratorID: orgMember.OrgMemID,
						GithubCreatedAt:            sql.NullTime{Time: repo.Repository.CreatedAt, Valid: true},
						GithubUpdatedAt:            sql.NullTime{Time: repo.Repository.UpdatedAt, Valid: true},
					})
					if err != nil {
						return err
					}
				} else {
					return err
				}
			}

			for _, prContribution := range repo.Contributions.Nodes {
				fmt.Println(prContribution.PullRequest.Title)
				_, err := github.model.GetPRByID(github.ctx, prContribution.PullRequest.ID)
				if err != nil {
					if err == sql.ErrNoRows {
						prAuthorID, err := github.model.GetMemberByLogin(github.ctx, prContribution.PullRequest.Author.Login)
						if err != nil {
							return err
						}
						_, err = github.model.InsertPR(github.ctx, models.InsertPRParams{
							ID:              prContribution.PullRequest.ID,
							Title:           sql.NullString{String: prContribution.PullRequest.Title, Valid: true},
							Status:          sql.NullString{String: prContribution.PullRequest.State, Valid: true},
							Url:             sql.NullString{String: prContribution.PullRequest.URL, Valid: true},
							IsDraft:         sql.NullBool{Bool: prContribution.PullRequest.IsDraft, Valid: true},
							Branch:          sql.NullString{String: prContribution.PullRequest.Branch, Valid: true},
							AuthorID:        prAuthorID,
							RepositoryID:    repoID,
							GithubClosedAt:  sql.NullTime{Time: prContribution.PullRequest.ClosedAt, Valid: true},
							GithubMergedAt:  sql.NullTime{Time: prContribution.PullRequest.MergedAt, Valid: true},
							GithubCreatedAt: sql.NullTime{Time: prContribution.PullRequest.ClosedAt, Valid: true},
							GithubUpdatedAt: sql.NullTime{Time: prContribution.PullRequest.UpdatedAt, Valid: true},
						})
						if err != nil {
							return err
						}
					} else {
						return err
					}
				}

				for _, review := range prContribution.PullRequest.Reviews.Nodes {
					fmt.Println(review.Author.Login)
					reviwerID, err := github.model.GetMemberByLogin(github.ctx, review.Author.Login)
					if err != nil {
						return err
					}
					_, err = github.model.GetReviewByID(github.ctx, review.ID)
					if err != nil {
						if err == sql.ErrNoRows {
							github.model.InsertReview(github.ctx, models.InsertReviewParams{
								ID:                review.ID,
								ReviewerID:        reviwerID,
								PrID:              prContribution.PullRequest.ID,
								Status:            review.State,
								GithubCreatedAt:   sql.NullTime{Time: review.CreatedAt, Valid: true},
								GithubUpdatedAt:   sql.NullTime{Time: review.UpdatedAt, Valid: true},
								GithubSubmittedAt: sql.NullTime{Time: review.SubmittedAt, Valid: true},
							})
						} else {
							return err
						}
					}
				}
				for _, labal := range prContribution.PullRequest.Labels.Nodes {
					fmt.Println(labal.Name)
					labalID, err := github.model.GetLabalByID(github.ctx, labal.ID)
					if err != nil {
						if err == sql.ErrNoRows {
							labalID, err = github.model.InsertLabal(github.ctx, models.InsertLabalParams{
								ID:   labal.ID,
								Name: sql.NullString{String: labal.Name, Valid: true},
							})
							if err != nil {
								return err
							}
						} else {
							return err
						}
					}

					// assign labal
					_, err = github.model.GetAssignedLabal(github.ctx, models.GetAssignedLabalParams{
						LabalID:    labalID,
						ActivityID: prContribution.PullRequest.ID,
					})
					if err != nil {
						if err == sql.ErrNoRows {
							_, err = github.model.InsertAssignedLabal(github.ctx, models.InsertAssignedLabalParams{
								ID:           utils.GenerateUUID(),
								LabalID:      labalID,
								ActivityID:   prContribution.PullRequest.ID,
								ActivityType: ActivityType,
							})
							if err != nil {
								return err
							}
						} else {
							return err
						}
					}
				}
				for _, assignee := range prContribution.PullRequest.Assignees.Nodes {
					fmt.Println(assignee.Login)
					memID, err := github.model.GetMemberByLogin(github.ctx, assignee.Login)
					if err != nil {
						return err
					}
					_, err = github.model.GetAssigneeByID(github.ctx, models.GetAssigneeByIDParams{
						CollaboratorID: memID,
						ActivityID:     prContribution.PullRequest.ID,
					})
					if err != nil {
						if err == sql.ErrNoRows {
							_, err = github.model.InsertAssignee(github.ctx, models.InsertAssigneeParams{
								ID:             utils.GenerateUUID(),
								CollaboratorID: memID,
								ActivityID:     prContribution.PullRequest.ID,
								ActivityType:   ActivityType,
							})
							if err != nil {
								return err
							}
						} else {
							return err
						}
					}
				}
				branchID, err := github.model.GetBranchByID(github.ctx, models.GetBranchByIDParams{
					Name:         prContribution.PullRequest.Branch,
					RepositoryID: repo.Repository.ID,
				})
				if err != nil {
					if err == sql.ErrNoRows {
						branchID, err = github.model.InsertBranch(github.ctx, models.InsertBranchParams{
							ID:           utils.GenerateUUID(),
							Name:         prContribution.PullRequest.Branch,
							RepositoryID: repo.Repository.ID,
						})
						if err != nil {
							return err
						}
					} else {
						return err
					}
				}
				for _, commit := range prContribution.PullRequest.Commits.Nodes {
					fmt.Println(commit.Commit.Committer.Name)
					committerID, err := github.model.GetMemberByLogin(github.ctx, commit.Commit.Committer.Name)
					if err != nil {
						return err
					}
					_, err = github.model.GetCommitByID(github.ctx, commit.Commit.ID)
					if err != nil {
						if err == sql.ErrNoRows {
							github.model.InsertCommit(github.ctx, models.InsertCommitParams{
								ID:                  commit.Commit.ID,
								Message:             sql.NullString{String: commit.Commit.Message, Valid: true},
								BranchID:            branchID,
								AuthorID:            committerID,
								PrID:                sql.NullString{String: prContribution.PullRequest.ID, Valid: true},
								Url:                 sql.NullString{String: commit.Commit.URL},
								CommitUrl:           sql.NullString{String: commit.Commit.CommitUrl},
								GithubCommittedTime: sql.NullTime{Time: commit.Commit.CommittedDate},
							})
						} else {
							return err
						}
					}
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
