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

type GithubCommitRepoQ struct {
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
	Refs struct {
		Nodes []struct {
			Name   string
			Target struct {
				History struct {
					Nodes    []GithubCommitQ
					PageInfo struct {
						HasNextPage bool
						EndCursor   string
					}
				} `graphql:"history(first: $commitLimit, after: $commitCursor, author: { id: $memberID })"`
			} `graphql:"... on Commit"`
		}
		PageInfo struct {
			HasNextPage bool
			EndCursor   string
		}
	} `graphql:"refs(refPrefix: \"refs/heads/\", first: $branchLimit, after: $branchCursor)"`
}

func (github *GithubService) LoadRepoByCommits(orgMember GithubOrgMemberArgs) error {
	var noPages []string
	end, start := utils.GetWeekTimestamps()
	var commitsCursor *githubv4.String
	var commitsLimit githubv4.Int = githubv4.Int(constants.DefaultLimit)
	var branchCursor *githubv4.String
	var branchLimit githubv4.Int = githubv4.Int(constants.DefaultLimit)
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
			"branchLimit":  branchLimit,
			"branchCursor": branchCursor,
			"startTime":    *githubv4.NewDateTime(githubv4.DateTime{start}),
			"endTime":      *githubv4.NewDateTime(githubv4.DateTime{end}),
			"orgID":        orgMember.ID,
			"memberLogin":  githubv4.String(orgMember.Member.Login),
			"memberID":     orgMember.Member.ID,
		}
		err := github.client.Query(context.Background(), &commitQ, variables)
		if err != nil {
			fmt.Println("Error executing query:", err)
			return nil
		}

		for _, repo := range commitQ.User.ContributionsCollection.CommitContributionsByRepository {
			_, err := github.model.GetRepoByID(github.ctx, repo.Repository.ID)
			if err != nil {
				if err == sql.ErrNoRows {
					_, err = github.model.InsertRepo(github.ctx, models.InsertRepoParams{
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

			for _, repoBranch := range repo.Repository.Refs.Nodes {
				branchID, err := github.model.GetBranchByID(github.ctx, models.GetBranchByIDParams{
					Name:         repoBranch.Name,
					RepositoryID: repo.Repository.ID,
				})
				if err != nil {
					if err == sql.ErrNoRows {
						branchID, err = github.model.InsertBranch(github.ctx, models.InsertBranchParams{
							ID:           utils.GenerateUUID(),
							Name:         repoBranch.Name,
							RepositoryID: repo.Repository.ID,
						})
						if err != nil {
							return err
						}
					} else {
						return err
					}
				}

				for _, repoCommit := range repoBranch.Target.History.Nodes {
					fmt.Println(repoCommit.ID)
					committerID, err := github.model.GetMemberByLogin(github.ctx, repoCommit.Author.User.Login)
					if err != nil {
						return err
					}
					_, err = github.model.GetCommitByID(github.ctx, repoCommit.ID)
					if err != nil {
						if err == sql.ErrNoRows {
							github.model.InsertCommit(github.ctx, models.InsertCommitParams{
								ID:                  repoCommit.ID,
								Message:             sql.NullString{String: repoCommit.Message, Valid: true},
								BranchID:            branchID,
								AuthorID:            committerID,
								Url:                 sql.NullString{String: repoCommit.URL},
								CommitUrl:           sql.NullString{String: repoCommit.CommitUrl},
								GithubCommittedTime: sql.NullTime{Time: repoCommit.CommittedDate},
							})
						} else {
							return err
						}
					}

				}
				if !repoBranch.Target.History.PageInfo.HasNextPage {
					if !utils.Contains("Commit", noPages) {
						noPages = append(noPages, "Commit")
						commitsLimit = githubv4.Int(0)
					}
				}
				commitsCursor = (*githubv4.String)(&repoBranch.Target.History.PageInfo.EndCursor)
			}
			if !repo.Repository.Refs.PageInfo.HasNextPage {
				if !utils.Contains("Branch", noPages) {
					noPages = append(noPages, "Branch")
					branchLimit = githubv4.Int(0)
				}
			}
			branchCursor = (*githubv4.String)(&repo.Repository.Refs.PageInfo.EndCursor)
		}

		if len(commitQ.User.ContributionsCollection.CommitContributionsByRepository) == 0 {
			break
		}
		if (len(noPages)) == 2 {
			break
		}
	}
	return nil
}
