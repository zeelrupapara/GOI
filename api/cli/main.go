package cli

import (
	"github.com/Improwised/GPAT/config"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// Init app initialization
func Init(cfg config.AppConfig, logger *zap.Logger) error {
	migrationCmd := GetMigrationCommandDef(cfg)
	APICmd := GetAPICommandDef(cfg, logger)
	githubCmd := GetGithubCommandDef(cfg, logger)
	// use is stands for a binary after build a golang app
	rootCmd := &cobra.Command{Use: "GPAT"}
	rootCmd.AddCommand(&migrationCmd, &APICmd, &githubCmd)
	return rootCmd.Execute()
}
