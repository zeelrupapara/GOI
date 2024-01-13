package cli

import (
	"github.com/Improwised/GPAT/cli/github"
	"github.com/Improwised/GPAT/config"
	gh "github.com/Improwised/GPAT/github"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var start, end string
var repos, orgs, users []string

func GetGithubCommandDef(cfg config.AppConfig, logger *zap.Logger) cobra.Command {
	githubCmd := cobra.Command{
		Use:   "github",
		Short: "Interact with GitHub resources",
		Run: func(cmd *cobra.Command, args []string) {
			// setup github service
			githubService, err := gh.NewGithubService(cfg, logger)
			if err != nil {
				logger.Error(err.Error())
			}
			err = githubService.LoadOrganizations()
			// // for _, org := range orgs{
			// // err := githubService.LoadOrganization(org)
			if err != nil {
				panic(err)
			}
			// }
		},
	}
	githubCmd.Flags().StringVarP(&start, "start", "s", "", "Start time in ISO 8601 format")
	githubCmd.Flags().StringVarP(&end, "end", "e", "", "End time in ISO 8601 format")
	githubCmd.Flags().StringSliceVarP(&repos, "repos", "r", []string{}, "Repositories (comma-separated list)")
	githubCmd.Flags().StringSliceVarP(&orgs, "orgs", "o", []string{}, "Organizations (comma-separated list)")
	githubCmd.Flags().StringSliceVarP(&users, "users", "u", []string{}, "Users (comma-separated list)")
	repoCmd := github.GetGithubRepoCommand(cfg, logger)
	githubCmd.AddCommand(&repoCmd)
	return githubCmd
}
