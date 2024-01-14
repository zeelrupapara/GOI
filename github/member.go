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

func (github *GithubService) LoadMembers(org GithubOrganizationQ) error {
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
			fmt.Println("Error executing query:", err)
			return nil
		}

		for _, member := range memberQ.Organization.MembersWithRole.Nodes {
			fmt.Println(">>>>>>Member:", member.Login)
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
						return err
					}
				} else {
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
						fmt.Println("err", err)
						return err
					}
				} else {
					return err
				}
			}

			err = github.LoadRepo(GithubOrgMemberArgs{
				ID:       org.ID,
				Login:    org.Login,
				Member:   member,
				OrgMemID: orgMemID,
			})
			if err != nil {
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
