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

type GithubMemberQ struct {
	ID              string
	Login           string
	Name            string
	Email           string
	URL             string
	AvatarURL       string    `graphql:"avatarUrl"`
	WebsiteURL      string    `graphql:"websiteUrl"`
	GithubUpdatedAt time.Time `graphql:"updatedAt"`
	GithubCreatedAt time.Time `graphql:"createdAt"`
}

func (github *GithubService) LoadMembers(org GithubOrganizationQ, start, end time.Time) error {
	var memberQ struct {
		Organization struct {
			MembersWithRole struct {
				Nodes    []GithubMemberQ
				PageInfo PageInfo
			} `graphql:"membersWithRole(first: $limit, after: $cursor)"`
		} `graphql:"organization(login: $login)"`
	}

	var limit githubv4.Int = githubv4.Int(constants.DefaultLimit)
	var cursor *githubv4.String
	var login githubv4.String = githubv4.String(org.Login)

	for {
		// Set the cursor for pagination
		variables := map[string]interface{}{
			"cursor": cursor,
			"limit":  limit,
			"login":  login,
		}

		// Execute the graphQL query
		err := github.client.Query(context.Background(), &memberQ, variables)
		if err != nil {
			github.LoadMembersLog(ERROR, err)
			return nil
		}

		for _, member := range memberQ.Organization.MembersWithRole.Nodes {
			github.LoadMembersLog(DEBUG, fmt.Sprintf("👤 Member: %s", member.Login))
			_, err := github.model.GetMemberByLogin(github.ctx, member.Login)
			if err != nil {
				if err == sql.ErrNoRows {
					_, err := github.model.InsertMember(github.ctx, models.InsertMemberParams{
						ID:              member.ID,
						Login:           member.Login,
						Name:            sql.NullString{String: member.Name, Valid: true},
						Email:           sql.NullString{String: member.Email, Valid: true},
						Url:             sql.NullString{String: member.URL, Valid: true},
						AvatarUrl:       sql.NullString{String: member.AvatarURL, Valid: true},
						WebsiteUrl:      sql.NullString{String: member.WebsiteURL, Valid: true},
						GithubCreatedAt: sql.NullTime{Time: member.GithubCreatedAt, Valid: true},
						GithubUpdatedAt: sql.NullTime{Time: member.GithubUpdatedAt, Valid: true},
					})
					if err != nil {
						github.LoadMembersLog(ERROR, err)
						return err
					}
				} else {
					github.LoadMembersLog(ERROR, err)
					return err
				}
			}

			// Check OrgMember Exist
			orgMemID, err := github.model.GetOrgMemberByID(github.ctx, models.GetOrgMemberByIDParams{
				OrganizationID: org.ID,
				CollaboratorID: member.ID,
			})
			if err != nil {
				if err == sql.ErrNoRows {
					orgMemID, err = github.model.InsertOrgMember(github.ctx, models.InsertOrgMemberParams{
						ID:             utils.GenerateUUID(),
						OrganizationID: org.ID,
						CollaboratorID: member.ID,
					})
					if err != nil {
						github.LoadMembersLog(ERROR, err)
						return err
					}
				} else {
					github.LoadMembersLog(ERROR, err)
					return err
				}
			}

			err = github.LoadRepo(GithubOrgMemberArgs{
				ID:       org.ID,
				Login:    org.Login,
				Member:   member,
				OrgMemID: orgMemID,
			}, start, end)
			if err != nil {
				github.LoadMembersLog(ERROR, err)
				return err
			}
		}

		// Check for pagination
		if !memberQ.Organization.MembersWithRole.PageInfo.HasNextPage {
			break
		}

		cursor = &memberQ.Organization.MembersWithRole.PageInfo.EndCursor

	}
	return nil
}

func (github *GithubService) LoadMember(username string) (string, error) {
	var memID string
	var memberQ struct {
		User GithubMemberQ `graphql:"user(login: $username)"`
	}

	variables := map[string]interface{}{
		"username": githubv4.String(username),
	}

	err := github.client.Query(context.Background(), &memberQ, variables)
	if err != nil {
		github.LoadMemberLog(ERROR, err)
		return memID, err
	}

	memID, err = github.model.GetMemberByLogin(github.ctx, username)
	if err != nil {
		if err == sql.ErrNoRows {
			memID, err = github.model.InsertMember(github.ctx, models.InsertMemberParams{
				ID:              memberQ.User.ID,
				Login:           memberQ.User.Login,
				Name:            sql.NullString{String: memberQ.User.Name, Valid: true},
				Email:           sql.NullString{String: memberQ.User.Email, Valid: true},
				Url:             sql.NullString{String: memberQ.User.URL, Valid: true},
				AvatarUrl:       sql.NullString{String: memberQ.User.AvatarURL, Valid: true},
				WebsiteUrl:      sql.NullString{String: memberQ.User.WebsiteURL, Valid: true},
				GithubCreatedAt: sql.NullTime{Time: memberQ.User.GithubCreatedAt, Valid: true},
				GithubUpdatedAt: sql.NullTime{Time: memberQ.User.GithubUpdatedAt, Valid: true},
			})
			if err != nil {
				github.LoadMemberLog(ERROR, err)
				return memID, err
			}
		} else {
			github.LoadMemberLog(ERROR, err)
			return memID, err
		}
	}

	return memID, nil
}

func (github *GithubService) LoadMembersLog(level string, message interface{}) {
	const path = "member -> LoadMembers -"
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

func (github *GithubService) LoadMemberLog(level string, message interface{}) {
	const path = "member -> LoadMember -"
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
