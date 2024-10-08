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
	"go.uber.org/zap"
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
	URL                      string
	ReviewRequests           GithubReviewRequestsQ           `graphql:"reviewRequests(first: $reviewRequestsLimit, after: $reviewRequestsCursor)"`
	LatestOpinionatedReviews GithubLatestOpinionatedReviewsQ `graphql:"latestOpinionatedReviews(first: $latestOpinionatedReviewsLimit, after: $latestOpinionatedReviewsCursor)"`
	Labels                   GithubLabelsQ                   `graphql:"labels(first: $labelsLimit, after: $labelsCursor)"`
	Commits                  GithubCommitsQ                  `graphql:"commits(first: $commitsLimit, after: $commitsCursor)"`
	Assignees                GithubAssigneesQ                `graphql:"assignees(first: $assigneesLimit, after: $assigneesCursor)"`
	MergedAt                 time.Time
	ClosedAt                 time.Time
	CreatedAt                time.Time
	UpdatedAt                time.Time
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

type GithubLatestOpinionatedReviewQ struct {
	ID     string
	Author struct {
		Login string
	}
	State       string
	SubmittedAt time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type GithubReviewRequestQ struct {
	ID                string
	RequestedReviewer struct {
		User struct {
			Login string
		} `graphql:"... on User"`
	}
}

type GithubLabelQ struct {
	ID   string
	Name string
}

type GithubCommits struct {
	Commit GithubCommitQ
}

type GithubCommitQ struct {
	ID            string
	Message       string
	CommittedDate time.Time
	URL           string
	CommitUrl     string
	Author        struct {
		User struct {
			Login string
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

type GithubLatestOpinionatedReviewsQ struct {
	Nodes    []GithubLatestOpinionatedReviewQ
	PageInfo PageInfo
}

type GithubReviewRequestsQ struct {
	Nodes    []GithubReviewRequestQ
	PageInfo PageInfo
}

type GithubLabelsQ struct {
	Nodes    []GithubLabelQ
	PageInfo PageInfo
}

type GithubCommitsQ struct {
	Nodes    []GithubCommits
	PageInfo PageInfo
}

func (github *GithubService) LoadRepoByPullRequests(orgMember GithubOrgMemberArgs, start, end time.Time) error {
	var noPages []string
	var contributionsLimit githubv4.Int = githubv4.Int(constants.DefaultLimit)
	var contributionsCursor *githubv4.String
	var memberName githubv4.String = githubv4.String(orgMember.Member.Login)
	var reviewRequestsLimit githubv4.Int = githubv4.Int(constants.DefaultLimit)
	var labelsLimit githubv4.Int = githubv4.Int(constants.DefaultLimit)
	var commitsLimit githubv4.Int = githubv4.Int(constants.DefaultLimit)
	var assigneesLimit githubv4.Int = githubv4.Int(constants.DefaultLimit)
	var latestOpinionatedReviewsLimit githubv4.Int = githubv4.Int(constants.DefaultLimit)
	var assigneesCursor *githubv4.String
	var commitsCursor *githubv4.String
	var labelsCursor *githubv4.String
	var reviewRequestsCursor *githubv4.String
	var latestOpinionatedReviewsCursor *githubv4.String

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
			"labelsLimit":                    labelsLimit,
			"commitsLimit":                   commitsLimit,
			"assigneesLimit":                 assigneesLimit,
			"labelsCursor":                   labelsCursor,
			"commitsCursor":                  commitsCursor,
			"assigneesCursor":                assigneesCursor,
			"reviewRequestsLimit":            reviewRequestsLimit,
			"reviewRequestsCursor":           reviewRequestsCursor,
			"latestOpinionatedReviewsLimit":  latestOpinionatedReviewsLimit,
			"latestOpinionatedReviewsCursor": latestOpinionatedReviewsCursor,
			"contributionsLimit":             contributionsLimit,
			"contributionsCursor":            contributionsCursor,
			"startTime":                      *githubv4.NewDateTime(githubv4.DateTime{start}),
			"endTime":                        *githubv4.NewDateTime(githubv4.DateTime{end}),
			"orgID":                          orgMember.ID,
			"memberLogin":                    memberName,
		}

		// Execute the graphQL query
		err := github.client.Query(context.Background(), &pullRequestsQ, variables)
		if err != nil {
			github.PRLog(ERROR, err)
			return nil
		}
		if len(pullRequestsQ.User.ContributionsCollection.PullRequestContributionsByRepository) == 0 {
			break
		}

		for _, repo := range pullRequestsQ.User.ContributionsCollection.PullRequestContributionsByRepository {
			github.PRLog(DEBUG, fmt.Sprintf("📦️ Repo: %s", repo.Repository.Name))
			var repoMemberID string
			repoID, err := github.model.GetRepoByID(github.ctx, repo.Repository.ID)
			if err != nil {
				if err == sql.ErrNoRows {
					repoID, err = github.model.InsertRepo(github.ctx, models.InsertRepoParams{
						ID:              repo.Repository.ID,
						Name:            sql.NullString{String: repo.Repository.Name, Valid: true},
						IsPrivate:       sql.NullBool{Bool: repo.Repository.IsPrivate, Valid: true},
						DefaultBranch:   sql.NullString{String: repo.Repository.DefaultBranch.Name},
						Url:             sql.NullString{String: repo.Repository.URL},
						HomepageUrl:     sql.NullString{String: repo.Repository.HomepageUrl},
						OpenIssues:      sql.NullInt32{Int32: int32(repo.Repository.OpenIssues.TotalCount), Valid: true},
						ClosedIssues:    sql.NullInt32{Int32: int32(repo.Repository.ClosedIssues.TotalCount), Valid: true},
						OpenPrs:         sql.NullInt32{Int32: int32(repo.Repository.OpenPRs.TotalCount), Valid: true},
						ClosedPrs:       sql.NullInt32{Int32: int32(repo.Repository.ClosedPRs.TotalCount), Valid: true},
						MergedPrs:       sql.NullInt32{Int32: int32(repo.Repository.MergedPRs.TotalCount), Valid: true},
						GithubCreatedAt: sql.NullTime{Time: repo.Repository.CreatedAt, Valid: true},
						GithubUpdatedAt: sql.NullTime{Time: repo.Repository.UpdatedAt, Valid: true},
					})
					if err != nil {
						github.PRLog(ERROR, err)
						return err
					}
				} else {
					github.PRLog(ERROR, err)
					return err
				}
			}
			repoMemberID, err = github.model.GetRepoMemberByOrgRepoID(github.ctx, models.GetRepoMemberByOrgRepoIDParams{
				RepoID:                     repoID,
				OrganizationCollaboratorID: orgMember.OrgMemID,
			})
			if err != nil {
				if err == sql.ErrNoRows {
					repoMemberID, err = github.model.InsertOrgRepoMember(github.ctx, models.InsertOrgRepoMemberParams{
						ID:                         utils.GenerateUUID(),
						RepoID:                     repoID,
						OrganizationCollaboratorID: orgMember.OrgMemID,
					})
					if err != nil {
						github.PRLog(ERROR, err)
						return err
					}
				} else {
					github.PRLog(ERROR, err)
					return err
				}
			}
			if len(repo.Contributions.Nodes) > 0 {
				for _, prContribution := range repo.Contributions.Nodes {
					github.PRLog(DEBUG, fmt.Sprintf("📥️ PR: %s", prContribution.PullRequest.Title))
					prID, err := github.model.GetPRByID(github.ctx, prContribution.PullRequest.ID)
					if err != nil {
						if err == sql.ErrNoRows {
							prAuthorID, err := github.model.GetMemberByLogin(github.ctx, prContribution.PullRequest.Author.Login)
							if err != nil {
								if err == sql.ErrNoRows {
									prAuthorID, err = github.LoadMember(prContribution.PullRequest.Author.Login)
									if err != nil {
										github.PRLog(ERROR, err)
										return err
									}
								} else {
									github.PRLog(ERROR, err)
									return err
								}
							}
							prID, err = github.model.InsertPR(github.ctx, models.InsertPRParams{
								ID:                        prContribution.PullRequest.ID,
								Title:                     sql.NullString{String: prContribution.PullRequest.Title, Valid: true},
								Status:                    sql.NullString{String: prContribution.PullRequest.State, Valid: true},
								Url:                       sql.NullString{String: prContribution.PullRequest.URL, Valid: true},
								IsDraft:                   sql.NullBool{Bool: prContribution.PullRequest.IsDraft, Valid: true},
								Number:                    sql.NullInt32{Int32: int32(prContribution.PullRequest.Number), Valid: true},
								Branch:                    sql.NullString{String: prContribution.PullRequest.Branch, Valid: true},
								AuthorID:                  prAuthorID,
								RepositoryCollaboratorsID: repoMemberID,
								GithubClosedAt:            sql.NullTime{Time: prContribution.PullRequest.ClosedAt, Valid: true},
								GithubMergedAt:            sql.NullTime{Time: prContribution.PullRequest.MergedAt, Valid: true},
								GithubCreatedAt:           sql.NullTime{Time: prContribution.PullRequest.ClosedAt, Valid: true},
								GithubUpdatedAt:           sql.NullTime{Time: prContribution.PullRequest.UpdatedAt, Valid: true},
							})
							if err != nil {
								github.PRLog(ERROR, err)
								return err
							}
						} else {
							github.PRLog(ERROR, err)
							return err
						}
					} else {
						// Update PR
						err = github.model.UpdatePR(github.ctx, models.UpdatePRParams{
							ID:              prID,
							Title:           sql.NullString{String: prContribution.PullRequest.Title, Valid: true},
							Status:          sql.NullString{String: prContribution.PullRequest.State, Valid: true},
							IsDraft:         sql.NullBool{Bool: prContribution.PullRequest.IsDraft, Valid: true},
							GithubClosedAt:  sql.NullTime{Time: prContribution.PullRequest.ClosedAt, Valid: true},
							GithubMergedAt:  sql.NullTime{Time: prContribution.PullRequest.MergedAt, Valid: true},
							GithubUpdatedAt: sql.NullTime{Time: prContribution.PullRequest.UpdatedAt, Valid: true},
							UpdatedAt:       time.Now(),
						})
						if err != nil {
							github.PRLog(ERROR, err)
						}
					}

					// Review Request
					if len(prContribution.PullRequest.ReviewRequests.Nodes) > 0 {
						for _, review := range prContribution.PullRequest.ReviewRequests.Nodes {
							github.PRLog(DEBUG, fmt.Sprintf("👀 Reviewer: %s", review.RequestedReviewer.User.Login))
							reviwerID, err := github.model.GetMemberByLogin(github.ctx, review.RequestedReviewer.User.Login)
							if err != nil {
								if err == sql.ErrNoRows {
									reviwerID, err = github.LoadMember(review.RequestedReviewer.User.Login)
									if err != nil {
										github.PRLog(ERROR, err)
										return err
									}
								} else {
									github.PRLog(ERROR, err)
									return err
								}
							}
							_, err = github.model.GetReviewByID(github.ctx, review.ID)
							if err != nil {
								if err == sql.ErrNoRows {
									github.model.InsertReview(github.ctx, models.InsertReviewParams{
										ID:                review.ID,
										ReviewerID:        reviwerID,
										PrID:              prContribution.PullRequest.ID,
										Status:            constants.ReviewRequestInitState,
										GithubCreatedAt:   sql.NullTime{Time: prContribution.PullRequest.CreatedAt, Valid: true},
										GithubUpdatedAt:   sql.NullTime{Time: prContribution.PullRequest.CreatedAt, Valid: true},
										GithubSubmittedAt: sql.NullTime{Time: prContribution.PullRequest.CreatedAt, Valid: true},
									})
								} else {
									github.PRLog(ERROR, err)
									return err
								}
							}
						}
					}

					// Latest Opinionated Review
					if len(prContribution.PullRequest.LatestOpinionatedReviews.Nodes) > 0 {
						for _, latestOpinionatedReview := range prContribution.PullRequest.LatestOpinionatedReviews.Nodes {
							github.PRLog(DEBUG, fmt.Sprintf("✍️ State: %s, 👀 Reviewer: %s", latestOpinionatedReview.State, latestOpinionatedReview.Author.Login))
							reviewerID, err := github.model.GetMemberByLogin(github.ctx, latestOpinionatedReview.Author.Login)
							if err != nil {
								if err == sql.ErrNoRows {
									reviewerID, err = github.LoadMember(latestOpinionatedReview.Author.Login)
									if err != nil {
										github.PRLog(ERROR, err)
										return err
									}
								} else {
									github.PRLog(ERROR, err)
									return err
								}
							}

							reviewID, err := github.model.GetReviewByPRAndReviewerID(github.ctx, models.GetReviewByPRAndReviewerIDParams{
								PrID:       prContribution.PullRequest.ID,
								ReviewerID: reviewerID,
							})
							if err != nil {
								if err == sql.ErrNoRows {
									_, err = github.model.InsertReview(github.ctx, models.InsertReviewParams{
										ID:                latestOpinionatedReview.ID,
										ReviewerID:        reviewerID,
										PrID:              prContribution.PullRequest.ID,
										Status:            latestOpinionatedReview.State,
										GithubCreatedAt:   sql.NullTime{Time: latestOpinionatedReview.CreatedAt, Valid: true},
										GithubUpdatedAt:   sql.NullTime{Time: latestOpinionatedReview.UpdatedAt, Valid: true},
										GithubSubmittedAt: sql.NullTime{Time: latestOpinionatedReview.SubmittedAt, Valid: true},
									})
									if err != nil {
										github.PRLog(ERROR, err)
										return err
									}
									continue
								} else {
									github.PRLog(ERROR, err)
									return err
								}
							}
							_, err = github.model.UpdateReview(github.ctx, models.UpdateReviewParams{
								ID:                reviewID,
								ReviewerID:        reviewerID,
								PrID:              prContribution.PullRequest.ID,
								Status:            latestOpinionatedReview.State,
								GithubCreatedAt:   sql.NullTime{Time: latestOpinionatedReview.CreatedAt, Valid: true},
								GithubUpdatedAt:   sql.NullTime{Time: latestOpinionatedReview.UpdatedAt, Valid: true},
								GithubSubmittedAt: sql.NullTime{Time: latestOpinionatedReview.SubmittedAt, Valid: true},
							})
							if err != nil {
								github.PRLog(ERROR, err)
								return err
							}
						}
					}

					// Labal
					if len(prContribution.PullRequest.Labels.Nodes) > 0 {
						for _, labal := range prContribution.PullRequest.Labels.Nodes {
							github.PRLog(DEBUG, fmt.Sprintf("🪧 Labal: %s", labal.Name))
							labalID, err := github.model.GetLabalByID(github.ctx, labal.ID)
							if err != nil {
								if err == sql.ErrNoRows {
									labalID, err = github.model.InsertLabal(github.ctx, models.InsertLabalParams{
										ID:   labal.ID,
										Name: sql.NullString{String: labal.Name, Valid: true},
									})
									if err != nil {
										github.PRLog(ERROR, err)
										return err
									}
								} else {
									github.PRLog(ERROR, err)
									return err
								}
							}

							// Assignned labal
							_, err = github.model.GetAssignedLabalByPR(github.ctx, models.GetAssignedLabalByPRParams{
								LabalID: labalID,
								PrID:    sql.NullString{String: prContribution.PullRequest.ID, Valid: true},
							})
							if err != nil {
								if err == sql.ErrNoRows {
									_, err = github.model.InsertAssignedLabal(github.ctx, models.InsertAssignedLabalParams{
										ID:           utils.GenerateUUID(),
										LabalID:      labalID,
										PrID:         sql.NullString{String: prContribution.PullRequest.ID, Valid: true},
										ActivityType: constants.ActivityPR,
									})
									if err != nil {
										github.PRLog(ERROR, err)
										return err
									}
								} else {
									github.PRLog(ERROR, err)
									return err
								}
							}
						}
					}

					// Assignee
					if len(prContribution.PullRequest.Assignees.Nodes) > 0 {
						for _, assignee := range prContribution.PullRequest.Assignees.Nodes {
							github.PRLog(DEBUG, fmt.Sprintf("🧑‍💻 Assginee: %s", assignee.Login))
							memID, err := github.model.GetMemberByLogin(github.ctx, assignee.Login)
							if err != nil {
								if err == sql.ErrNoRows {
									memID, err = github.LoadMember(assignee.Login)
									if err != nil {
										github.PRLog(ERROR, err)
										return err
									}
								} else {
									github.PRLog(ERROR, err)
									return err
								}
							}
							_, err = github.model.GetAssigneeByPR(github.ctx, models.GetAssigneeByPRParams{
								CollaboratorID: memID,
								PrID:           sql.NullString{String: prContribution.PullRequest.ID},
							})
							if err != nil {
								if err == sql.ErrNoRows {
									_, err = github.model.InsertAssignee(github.ctx, models.InsertAssigneeParams{
										ID:             utils.GenerateUUID(),
										CollaboratorID: memID,
										PrID:           sql.NullString{String: prID, Valid: true},
										ActivityType:   constants.ActivityPR,
									})
									if err != nil {
										github.PRLog(ERROR, err)
										return err
									}
								} else {
									github.PRLog(ERROR, err)
									return err
								}
							}
						}
					}

					if len(prContribution.PullRequest.Commits.Nodes) > 0 {
						github.PRLog(DEBUG, fmt.Sprintf("🌳 Branch: %s", prContribution.PullRequest.Branch))

						// Branch
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
									github.PRLog(ERROR, err)
									return err
								}
							} else {
								github.PRLog(ERROR, err)
								return err
							}
						}

						// Commits
						for _, commit := range prContribution.PullRequest.Commits.Nodes {
							github.PRLog(DEBUG, fmt.Sprintf("💬 Commit: %s", commit.Commit.Message))
							github.PRLog(DEBUG, fmt.Sprintf("💬👤 Committer: %s", commit.Commit.Author.User.Login))
							if commit.Commit.Author.User.Login == "" {
								continue
							}
							committerID, err := github.model.GetMemberByLogin(github.ctx, commit.Commit.Author.User.Login)
							if err != nil {
								if err == sql.ErrNoRows {
									committerID, err = github.LoadMember(commit.Commit.Author.User.Login)
									if err != nil {
										github.PRLog(ERROR, err)
										return err
									}
								} else {
									github.PRLog(ERROR, err)
									return err
								}
							}
							_, err = github.model.GetCommitByID(github.ctx, models.GetCommitByIDParams{
								HashID:   commit.Commit.ID,
								BranchID: branchID,
							})
							if err != nil {
								if err == sql.ErrNoRows {
									github.model.InsertCommit(github.ctx, models.InsertCommitParams{
										ID:                  utils.GenerateUUID(),
										HashID:              commit.Commit.ID,
										Message:             sql.NullString{String: commit.Commit.Message, Valid: true},
										BranchID:            branchID,
										AuthorID:            committerID,
										PrID:                sql.NullString{String: prContribution.PullRequest.ID, Valid: true},
										Url:                 sql.NullString{String: commit.Commit.URL, Valid: true},
										CommitUrl:           sql.NullString{String: commit.Commit.CommitUrl, Valid: true},
										GithubCommittedTime: sql.NullTime{Time: commit.Commit.CommittedDate, Valid: true},
									})
								} else {
									github.PRLog(ERROR, err)
									return err
								}
							}
						}
					}

					// reviewRequested page break
					if !prContribution.PullRequest.ReviewRequests.PageInfo.HasNextPage {
						if !utils.Contains("ReviewRequests", noPages) {
							noPages = append(noPages, "ReviewRequests")
							reviewRequestsLimit = githubv4.Int(0)
						}
					}
					reviewRequestsCursor = &prContribution.PullRequest.ReviewRequests.PageInfo.EndCursor

					// latest opinionated review page break
					if !prContribution.PullRequest.ReviewRequests.PageInfo.HasNextPage {
						if !utils.Contains("LatestOpinionatedReviews", noPages) {
							noPages = append(noPages, "LatestOpinionatedReviews")
							latestOpinionatedReviewsLimit = githubv4.Int(0)
						}
					}
					latestOpinionatedReviewsCursor = &prContribution.PullRequest.ReviewRequests.PageInfo.EndCursor

					// assignees page break
					if !prContribution.PullRequest.Assignees.PageInfo.HasNextPage {
						if !utils.Contains("Assignee", noPages) {
							noPages = append(noPages, "Assignee")
							assigneesLimit = githubv4.Int(0)
						}
					}
					assigneesCursor = &prContribution.PullRequest.Assignees.PageInfo.EndCursor

					// commit page break
					if !prContribution.PullRequest.Commits.PageInfo.HasNextPage {
						if !utils.Contains("Commit", noPages) {
							noPages = append(noPages, "Commit")
							commitsLimit = githubv4.Int(0)
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
				if (!repo.Contributions.PageInfo.HasNextPage) && len(noPages) == 5 {
					if !utils.Contains("PullRequest", noPages) {
						noPages = append(noPages, "PullRequest")
						contributionsLimit = githubv4.Int(0)
					}
				}
				if repo.Contributions.PageInfo.HasNextPage && len(noPages) == 5 {
					contributionsCursor = &repo.Contributions.PageInfo.EndCursor
				}
				if len(noPages) == 5 {
					// ReviewRequests Reset
					reviewRequestsLimit = githubv4.Int(constants.DefaultLimit)
					reviewRequestsCursor = nil

					// LatestOpinionatedReviews Reset
					latestOpinionatedReviewsLimit = githubv4.Int(constants.DefaultLimit)
					latestOpinionatedReviewsCursor = nil

					// Commit Reset
					commitsLimit = githubv4.Int(constants.DefaultLimit)
					commitsCursor = nil

					// Assaignee Reset
					assigneesCursor = nil
					assigneesLimit = githubv4.Int(constants.DefaultLimit)

					// Label Reset
					labelsCursor = nil
					labelsLimit = githubv4.Int(constants.DefaultLimit)
				}
			}
		}
		if (len(noPages)) == 6 {
			break
		}
	}
	return nil
}

func (github *GithubService) PRLog(level string, message interface{}) {
	const path = "pull_request -> LoadRepoByPullRequests -"
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
