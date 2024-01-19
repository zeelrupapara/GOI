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

type GithubIssueContribution struct {
	Issue GithubIssueQ
}

type GithubIssueQ struct {
	ID     string
	Title  string
	State  string
	Number int
	Author struct {
		Login string
	}
	URL       string
	Labels    GithubLabelsQ    `graphql:"labels(first: $labelsLimit, after: $labelsCursor)"`
	Assignees GithubAssigneesQ `graphql:"assignees(first: $assigneesLimit, after: $assigneesCursor)"`
	ClosedAt  time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (github *GithubService) LoadRepoByIssues(orgMember GithubOrgMemberArgs, start, end time.Time) error {
	var noPages []string
	var contributionsLimit githubv4.Int = githubv4.Int(constants.DefaultLimit)
	var contributionsCursor *githubv4.String
	var memberName githubv4.String = githubv4.String(orgMember.Member.Login)
	var labelsLimit githubv4.Int = githubv4.Int(constants.DefaultLimit)
	var assigneesLimit githubv4.Int = githubv4.Int(constants.DefaultLimit)
	var assigneesCursor *githubv4.String
	var labelsCursor *githubv4.String

	var issuesQ struct {
		User struct {
			ContributionsCollection struct {
				IssueContributionsByRepository []struct {
					Repository    GithubRepoQ
					Contributions struct {
						Nodes    []GithubIssueContribution
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
			"assigneesLimit":      assigneesLimit,
			"labelsCursor":        labelsCursor,
			"assigneesCursor":     assigneesCursor,
			"contributionsLimit":  contributionsLimit,
			"contributionsCursor": contributionsCursor,
			"startTime":           *githubv4.NewDateTime(githubv4.DateTime{start}),
			"endTime":             *githubv4.NewDateTime(githubv4.DateTime{end}),
			"orgID":               orgMember.ID,
			"memberLogin":         memberName,
		}

		// Execute the graphQL query
		err := github.client.Query(context.Background(), &issuesQ, variables)
		if err != nil {
			github.IssuesLog(ERROR, err)
			return nil
		}
		if len(issuesQ.User.ContributionsCollection.IssueContributionsByRepository) == 0 {
			break
		}

		for _, repo := range issuesQ.User.ContributionsCollection.IssueContributionsByRepository {
			github.IssuesLog(DEBUG, fmt.Sprintf("📦️ Repo: %s", repo.Repository.Name))
			// Check repo exist or not?
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
						github.IssuesLog(ERROR, err)
						return err
					}
				} else {
					github.IssuesLog(ERROR, err)
					return err
				}
			}
			repoMemberID, err := github.model.GetRepoMemberByOrgRepoID(github.ctx, models.GetRepoMemberByOrgRepoIDParams{
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
						github.IssuesLog(ERROR, err)
						return err
					}
				} else {
					github.IssuesLog(ERROR, err)
					return err
				}
			}
			if len(repo.Contributions.Nodes) > 0 {
				for _, issueContribution := range repo.Contributions.Nodes {
					github.IssuesLog(DEBUG, fmt.Sprintf("🎯 Issue: %s", issueContribution.Issue.Title))
					issueID, err := github.model.GetIssueByID(github.ctx, issueContribution.Issue.ID)
					if err != nil {
						if err == sql.ErrNoRows {
							issueAuthorID, err := github.model.GetMemberByLogin(github.ctx, issueContribution.Issue.Author.Login)
							if err != nil {
								if err == sql.ErrNoRows {
									issueAuthorID, err = github.LoadMember(issueContribution.Issue.Author.Login)
									if err != nil {
										github.IssuesLog(ERROR, err)
										return err
									}
								} else {
									github.IssuesLog(ERROR, err)
									return err
								}
							}
							if err != nil {
								github.IssuesLog(ERROR, err)
								return err
							}
							issueID, err = github.model.InsertIssue(github.ctx, models.InsertIssueParams{
								ID:                        issueContribution.Issue.ID,
								Title:                     issueContribution.Issue.Title,
								Status:                    issueContribution.Issue.State,
								Url:                       sql.NullString{String: issueContribution.Issue.URL, Valid: true},
								Number:                    sql.NullInt32{Int32: int32(issueContribution.Issue.Number), Valid: true},
								AuthorID:                  issueAuthorID,
								RepositoryCollaboratorsID: repoMemberID,
								GithubClosedAt:            sql.NullTime{Time: issueContribution.Issue.ClosedAt, Valid: true},
								GithubCreatedAt:           sql.NullTime{Time: issueContribution.Issue.ClosedAt, Valid: true},
								GithubUpdatedAt:           sql.NullTime{Time: issueContribution.Issue.UpdatedAt, Valid: true},
							})
							if err != nil {
								github.IssuesLog(ERROR, err)
								return err
							}
						} else {
							github.IssuesLog(ERROR, err)
							return err
						}
					}
					// labal
					if len(issueContribution.Issue.Labels.Nodes) > 0 {
						for _, labal := range issueContribution.Issue.Labels.Nodes {
							github.IssuesLog(DEBUG, fmt.Sprintf("🪧 Labal: %s", labal.Name))
							labalID, err := github.model.GetLabalByID(github.ctx, labal.ID)
							if err != nil {
								if err == sql.ErrNoRows {
									labalID, err = github.model.InsertLabal(github.ctx, models.InsertLabalParams{
										ID:   labal.ID,
										Name: sql.NullString{String: labal.Name, Valid: true},
									})
									if err != nil {
										github.IssuesLog(ERROR, err)
										return err
									}
								} else {
									github.IssuesLog(ERROR, err)
									return err
								}
							}

							// assigned labal
							_, err = github.model.GetAssignedLabalByIssue(github.ctx, models.GetAssignedLabalByIssueParams{
								LabalID: labalID,
								IssueID: sql.NullString{String: issueContribution.Issue.ID, Valid: true},
							})
							if err != nil {
								if err == sql.ErrNoRows {
									_, err = github.model.InsertAssignedLabal(github.ctx, models.InsertAssignedLabalParams{
										ID:           utils.GenerateUUID(),
										LabalID:      labalID,
										IssueID:      sql.NullString{String: issueID, Valid: true},
										ActivityType: constants.ActivityIssue,
									})
									if err != nil {
										github.IssuesLog(ERROR, err)
										return err
									}
								} else {
									github.IssuesLog(ERROR, err)
									return err
								}
							}
						}
					}

					if len(issueContribution.Issue.Assignees.Nodes) > 0 {
						for _, assignee := range issueContribution.Issue.Assignees.Nodes {
							github.IssuesLog(DEBUG, fmt.Sprintf("🧑‍💻 Assignee: %s", assignee.Login))
							memID, err := github.model.GetMemberByLogin(github.ctx, assignee.Login)
							if err != nil {
								if err == sql.ErrNoRows {
									memID, err = github.LoadMember(issueContribution.Issue.Author.Login)
									if err != nil {
										github.IssuesLog(ERROR, err)
										return err
									}
								} else {
									github.IssuesLog(ERROR, err)
									return err
								}
							}
							_, err = github.model.GetAssigneeByIssue(github.ctx, models.GetAssigneeByIssueParams{
								CollaboratorID: memID,
								IssueID:        sql.NullString{String: issueContribution.Issue.ID, Valid: true},
							})
							if err != nil {
								if err == sql.ErrNoRows {
									_, err = github.model.InsertAssignee(github.ctx, models.InsertAssigneeParams{
										ID:             utils.GenerateUUID(),
										CollaboratorID: memID,
										IssueID:        sql.NullString{String: issueID, Valid: true},
										ActivityType:   constants.ActivityIssue,
									})
									if err != nil {
										github.IssuesLog(ERROR, err)
										return err
									}
								} else {
									github.IssuesLog(ERROR, err)
									return err
								}
							}
						}
					}

					// assignees page break
					if !issueContribution.Issue.Assignees.PageInfo.HasNextPage {
						if !utils.Contains("Assignee", noPages) {
							noPages = append(noPages, "Assignee")
							assigneesLimit = githubv4.Int(0)
						}
					}
					assigneesCursor = &issueContribution.Issue.Assignees.PageInfo.EndCursor

					// labal page break
					if !issueContribution.Issue.Labels.PageInfo.HasNextPage {
						if !utils.Contains("Label", noPages) {
							noPages = append(noPages, "Label")
							labelsLimit = githubv4.Int(0)
						}
					}
					labelsCursor = &issueContribution.Issue.Labels.PageInfo.EndCursor
				}

				// Issue contribution page break
				if (!repo.Contributions.PageInfo.HasNextPage) && len(noPages) == 2 {
					if !utils.Contains("Issue", noPages) {
						noPages = append(noPages, "Issue")
					}
				}
				contributionsCursor = &repo.Contributions.PageInfo.EndCursor
				if len(noPages) == 2 {
					// Assaignee Reset
					assigneesCursor = nil
					assigneesLimit = githubv4.Int(constants.DefaultLimit)
					
					// Label Reset
					labelsCursor = nil
					labelsLimit = githubv4.Int(constants.DefaultLimit)
				}
			}
		}
		if len(noPages) == 3 {
			break
		}
	}
	return nil
}

func (github *GithubService) IssuesLog(level string, message interface{}) {
	const path = "issue -> LoadRepoByIssues -"
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
