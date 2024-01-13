package github

import (
	"context"
	"fmt"
	"time"

	"github.com/Improwised/GPAT/constants"
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
			fmt.Println("Member:", member.Login)
		}

		// Check for pagination
		if !memberQ.Organization.MembersWithRole.PageInfo.HasNextPage {
			break
		}

		cursor = &memberQ.Organization.MembersWithRole.PageInfo.EndCursor

	}
	return nil
}
