package cmd

import (
	hc "harborctl/pkg/healthcheck"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var healthcheckCmd = &cobra.Command{
	Use:   "health",
	Short: "Check status of Harbor's component services",
	Run: func(cmd *cobra.Command, args []string) {
		hc.HealthCheck(viper.GetString("URL"))
	},
}

func init() {
	rootCmd.AddCommand(healthcheckCmd)
}
