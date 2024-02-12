package cli

import (
	"fmt"
	"time"

	"github.com/Improwised/GPAT/config"
	"github.com/Improwised/GPAT/constants"
	gh "github.com/Improwised/GPAT/github"
	"github.com/Improwised/GPAT/utils"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var start, end string
var startTime, endTime time.Time
var err error

// var repos, orgs, users []string

func GetGithubCommandDef(cfg config.AppConfig, logger *zap.Logger) cobra.Command {
	githubCmd := cobra.Command{
		Use:   "github",
		Short: "To fetch data from the github",
		Run: func(cmd *cobra.Command, args []string) {

			// get time from the flag
			if start != "" && end != "" {
				startTime, err = utils.ParseTimeFromString(start)
				if err != nil {
					logger.Error("github -> GetGithubCommandDef() - Error while parsing time", zap.Error(err))
					return
				}
				endTime, err = utils.ParseTimeFromString(end)
				if err != nil {
					logger.Error("github -> GetGithubCommandDef() - Error while parsing time", zap.Error(err))
					return
				}
			} else if start != "" {
				endTime = time.Now().UTC()
				startTime, err = utils.ParseTimeFromString(start)
				if err != nil {
					logger.Error("github -> GetGithubCommandDef() - Error while parsing time", zap.Error(err))
					return
				}
			} else {
				endTime, startTime = utils.GetWeekTimestamps()
			}

			// setup github service
			githubService, err := gh.NewGithubService(cfg, logger)
			if err != nil {
				logger.Error(err.Error())
			}

			// Execute command week wise
			weekWiseTime := utils.SplitTimeRange(startTime, endTime, constants.CommandIntervalTime)
			for _, weekTime := range weekWiseTime {
				err = githubService.LoadOrganizations(weekTime[0], weekTime[1])
				if err != nil {
					fmt.Println(err)
				}
				if err != nil {
					panic(err)
				}
			}
		},
	}
	githubCmd.Flags().StringVarP(&start, "start", "s", "", "Start time in ISO 8601 format ex: '2006-01-02T15:04:05Z'")
	githubCmd.Flags().StringVarP(&end, "end", "e", "", "End time in ISO 8601 format ex: '2006-01-02T15:04:05Z'")
	// githubCmd.Flags().StringSliceVarP(&repos, "repos", "r", []string{}, "Repositories (comma-separated list)")
	// githubCmd.Flags().StringSliceVarP(&orgs, "orgs", "o", []string{}, "Organizations (comma-separated list)")
	// githubCmd.Flags().StringSliceVarP(&users, "users", "u", []string{}, "Users (comma-separated list)")
	// repoCmd := github.GetGithubRepoCommand(cfg, logger)
	// githubCmd.AddCommand(&repoCmd)
	return githubCmd
}
