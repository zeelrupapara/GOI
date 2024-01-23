package github

import (
	"fmt"

	"github.com/Improwised/GPAT/config"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var prs, issues, branches []string
var name string

func GetGithubRepoCommand(cfg config.AppConfig, logger *zap.Logger) cobra.Command {
	repoCmd := cobra.Command{
		Use:   "repo",
		Short: "Interact with individual repositories",
		Run: func(cmd *cobra.Command, args []string) {
			// Handle repo command execution
			if (name == "") {
				fmt.Print("Please specify the name of repo. (ex: github repo --name 'repo_name' --help)")
			} else {
				fmt.Println("Repo command executed")
			}
		},
	}
	repoCmd.Flags().StringSliceVarP(&prs, "prs", "p", []string{}, "Pull Requestes (comma-separated list)")
	repoCmd.Flags().StringSliceVarP(&issues, "issues", "i", []string{}, "Issues (comma-separated list)")
	repoCmd.Flags().StringSliceVarP(&branches, "barnches", "b", []string{}, "Branches (comma-separated list)")
	repoCmd.Flags().StringVarP(&name, "name", "n", "", "Name of repository (required)")
	return repoCmd
}
