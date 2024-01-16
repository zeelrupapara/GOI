package github

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Improwised/GPAT/constants"
	"github.com/Improwised/GPAT/models"
	"github.com/shurcooL/githubv4"
	"go.uber.org/zap"
)

type GithubOrganizationQ struct {
	ID              string
	Login           string
	Name            string
	Email           string
	Location        string
	Description     string
	URL             string
	AvatarURL       string    `graphql:"avatarUrl"`
	WebsiteURL      string    `graphql:"websiteUrl"`
	GithubUpdatedAt time.Time `graphql:"updatedAt"`
	GithubCreatedAt time.Time `graphql:"createdAt"`
}

type PageInfo struct {
	HasNextPage bool
	EndCursor   githubv4.String
}

func (github *GithubService) LoadOrganizations() error {
	var organizationQ struct {
		Viewer struct {
			Organizations struct {
				Nodes    []GithubOrganizationQ
				PageInfo PageInfo
			} `graphql:"organizations(first: $limit, after: $cursor)"`
		}
	}

	var limit githubv4.Int = githubv4.Int(constants.DefaultLimit)
	var cursor *githubv4.String

	for {
		// Set the cursor for pagination
		variables := map[string]interface{}{
			"cursor": cursor,
			"limit":  limit,
		}

		// Execute the graphQL query
		err := github.client.Query(context.Background(), &organizationQ, variables)
		if err != nil {
			github.LoadOrganizationsLog(ERROR, err)
			return nil
		}

		for _, org := range organizationQ.Viewer.Organizations.Nodes {
			github.LoadOrganizationsLog(DEBUG, fmt.Sprintf("🏦 Organization: %s", org.Login))

			// Check if org exists or not
			_, err := github.model.GetOrganizationByLogin(github.ctx, org.Login)
			if err != nil {
				if err == sql.ErrNoRows {
					_, err = github.model.InsertOrganization(github.ctx, models.InsertOrganizationParams{
						ID:              org.ID,
						Login:           org.Login,
						Name:            sql.NullString{String: org.Name, Valid: true},
						Email:           sql.NullString{String: org.Email, Valid: true},
						Location:        sql.NullString{String: org.Location, Valid: true},
						Description:     sql.NullString{String: org.Description, Valid: true},
						Url:             sql.NullString{String: org.URL, Valid: true},
						AvatarUrl:       sql.NullString{String: org.AvatarURL, Valid: true},
						WebsiteUrl:      sql.NullString{String: org.WebsiteURL, Valid: true},
						GithubUpdatedAt: sql.NullTime{Time: org.GithubUpdatedAt, Valid: true},
						GithubCreatedAt: sql.NullTime{Time: org.GithubCreatedAt, Valid: true},
					})
					if err != nil {
						github.LoadOrganizationsLog(ERROR, err)
						return nil
					}
				} else {
					github.LoadOrganizationsLog(ERROR, err)
					return err
				}
			}

			// Get the org members
			github.LoadMembers(org)
		}

		// Check for pagination
		if !organizationQ.Viewer.Organizations.PageInfo.HasNextPage {
			break
		}

		cursor = &organizationQ.Viewer.Organizations.PageInfo.EndCursor

	}
	return nil
}

func (github *GithubService) LoadOrganization(org string) error {
	var query struct {
		Organization GithubOrganizationQ `graphql:"organization(login: $login)"`
	}

	// Set the cursor for pagination
	variables := map[string]interface{}{
		"login": githubv4.String(org),
	}

	// Execute the graphQL query
	err := github.client.Query(context.Background(), &query, variables)
	if err != nil {
		github.LoadOrganizationsLog(ERROR, err)
		return err
	}
	return nil
}

func (github *GithubService) LoadOrganizationsLog(level string, message interface{}) {
	const path = "organization -> LoadOrganizations -"
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
