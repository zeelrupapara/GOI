package github

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Improwised/GPAT/models"
	"github.com/Improwised/GPAT/utils"
	"go.uber.org/zap"
	"time"

	"github.com/Improwised/GPAT/constants"
	"github.com/shurcooL/githubv4"
)

type GithubCommitRepoQ struct {
	ID          string
	Name        string
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
	DefaultBranchRef struct {
		ID     string
		Name   string
		Target struct {
			Commit struct {
				History struct {
					Nodes    []GithubCommitQ
					PageInfo PageInfo
				} `graphql:"history(first: $commitLimit, after: $commitCursor, since: $sinceTime, until: $untilTime, author: { id: $memberID })"`
			} `graphql:"... on Commit"`
		}
	}
}

func (github *GithubService) LoadRepoByCommits(orgMember GithubOrgMemberArgs, start, end time.Time) error {
	var noPages []string
	var commitsCursor *githubv4.String
	var commitsLimit githubv4.Int = githubv4.Int(constants.DefaultLimit)
	var commitQ struct {
		User struct {
			ContributionsCollection struct {
				CommitContributionsByRepository []struct {
					Repository GithubCommitRepoQ
				}
			} `graphql:"contributionsCollection(organizationID: $orgID, from: $startTime, to: $endTime)"`
		} `graphql:"user(login: $memberLogin)"`
	}

	for {
		// Set the cursor for pagination
		variables := map[string]interface{}{
			"commitLimit":  commitsLimit,
			"commitCursor": commitsCursor,
			"sinceTime":    githubv4.NewGitTimestamp(githubv4.GitTimestamp{start}),
			"untilTime":    githubv4.NewGitTimestamp(githubv4.GitTimestamp{end}),
			"startTime":    *githubv4.NewDateTime(githubv4.DateTime{start}),
			"endTime":      *githubv4.NewDateTime(githubv4.DateTime{end}),
			"orgID":        orgMember.ID,
			"memberLogin":  githubv4.String(orgMember.Member.Login),
			"memberID":     orgMember.Member.ID,
		}
		err := github.client.Query(context.Background(), &commitQ, variables)
		if err != nil {
			github.CommitLog(ERROR, err)
			return nil
		}

		if len(commitQ.User.ContributionsCollection.CommitContributionsByRepository) > 0 {
			for _, repo := range commitQ.User.ContributionsCollection.CommitContributionsByRepository {
				github.CommitLog(DEBUG, fmt.Sprintf("📦️ Repo: %s", repo.Repository.Name))
				_, err := github.model.GetRepoByID(github.ctx, repo.Repository.ID)
				if err != nil {
					if err == sql.ErrNoRows {
						_, err = github.model.InsertRepo(github.ctx, models.InsertRepoParams{
							ID:              repo.Repository.ID,
							Name:            sql.NullString{String: repo.Repository.Name, Valid: true},
							IsPrivate:       sql.NullBool{Bool: repo.Repository.IsPrivate, Valid: true},
							DefaultBranch:   sql.NullString{String: repo.Repository.DefaultBranchRef.Name, Valid: true},
							Url:             sql.NullString{String: repo.Repository.URL, Valid: true},
							HomepageUrl:     sql.NullString{String: repo.Repository.HomepageUrl, Valid: true},
							OpenIssues:      sql.NullInt32{Int32: int32(repo.Repository.OpenIssues.TotalCount), Valid: true},
							ClosedIssues:    sql.NullInt32{Int32: int32(repo.Repository.ClosedIssues.TotalCount), Valid: true},
							OpenPrs:         sql.NullInt32{Int32: int32(repo.Repository.OpenPRs.TotalCount), Valid: true},
							ClosedPrs:       sql.NullInt32{Int32: int32(repo.Repository.ClosedPRs.TotalCount), Valid: true},
							MergedPrs:       sql.NullInt32{Int32: int32(repo.Repository.MergedPRs.TotalCount), Valid: true},
							GithubCreatedAt: sql.NullTime{Time: repo.Repository.CreatedAt, Valid: true},
							GithubUpdatedAt: sql.NullTime{Time: repo.Repository.UpdatedAt, Valid: true},
						})
						if err != nil {
							github.CommitLog(ERROR, err)
							return err
						}
					} else {
						github.CommitLog(ERROR, err)
						return err
					}
				}

				// Get the branch name and set default branch to true
				github.CommitLog(DEBUG, fmt.Sprintf("🌳 Branch: %s", repo.Repository.DefaultBranchRef.Name))
				branchID, err := github.model.GetBranchByID(github.ctx, models.GetBranchByIDParams{
					Name:         repo.Repository.DefaultBranchRef.Name,
					RepositoryID: repo.Repository.ID,
				})
				if err != nil {
					if err == sql.ErrNoRows {
						branchID, err = github.model.InsertBranch(github.ctx, models.InsertBranchParams{
							ID:           repo.Repository.DefaultBranchRef.ID,
							Name:         repo.Repository.DefaultBranchRef.Name,
							RepositoryID: repo.Repository.ID,
							IsDefault:    true,
						})
						if err != nil {
							github.CommitLog(ERROR, err)
							return err
						}
					} else {
						github.CommitLog(ERROR, err)
						return err
					}
				}

				// Get the contributors commits
				if len(repo.Repository.DefaultBranchRef.Target.Commit.History.Nodes) > 0 {
					for _, repoCommit := range repo.Repository.DefaultBranchRef.Target.Commit.History.Nodes {
						github.CommitLog(DEBUG, fmt.Sprintf("💬 Commit: %s", repoCommit.Message))
						github.CommitLog(DEBUG, fmt.Sprintf("💬👤 Committer: %s", repoCommit.Author.User.Login))
						committerID, err := github.model.GetMemberByLogin(github.ctx, repoCommit.Author.User.Login)
						if err != nil {
							github.CommitLog(ERROR, err)
							return err
						}
						_, err = github.model.GetCommitByID(github.ctx, models.GetCommitByIDParams{
							ID:       repoCommit.ID,
							BranchID: branchID,
						})
						if err != nil {
							if err == sql.ErrNoRows {
								_, err := github.model.InsertCommit(github.ctx, models.InsertCommitParams{
									ID:                  repoCommit.ID,
									Message:             sql.NullString{String: repoCommit.Message, Valid: true},
									BranchID:            branchID,
									AuthorID:            committerID,
									Url:                 sql.NullString{String: repoCommit.URL},
									CommitUrl:           sql.NullString{String: repoCommit.CommitUrl},
									GithubCommittedTime: sql.NullTime{Time: repoCommit.CommittedDate},
								})
								if err != nil {
									github.CommitLog(ERROR, err)
									return err
								}
							} else {
								github.CommitLog(ERROR, err)
								return err
							}
						}
					}
				} else {
					if !utils.Contains("Commit", noPages) {
						noPages = append(noPages, "Commit")
						commitsLimit = githubv4.Int(0)
					}
					break
				}
				// Commit page break
				if !repo.Repository.DefaultBranchRef.Target.Commit.History.PageInfo.HasNextPage {
					if !utils.Contains("Commit", noPages) {
						noPages = append(noPages, "Commit")
						commitsLimit = githubv4.Int(0)
					}
				}
				commitsCursor = &repo.Repository.DefaultBranchRef.Target.Commit.History.PageInfo.EndCursor
			}
		} else {
			break
		}
		if len(noPages) == 1 {
			break
		}
	}
	return nil
}

func (github *GithubService) CommitLog(level string, message interface{}) {
	const path = "commit -> LoadRepoByCommits -"
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
