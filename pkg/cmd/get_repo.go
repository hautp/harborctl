package cmd

import (
	"fmt"

	rp "harborctl/pkg/repository"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	getRepoCmd = &cobra.Command{
		Use:   "repo",
		Short: "List all repository of project",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("[!] Please input project name to get repository.")
			} else if len(args) == 1 {
				rp.ListRepositories(viper.GetString("URL"), args[0])
			} else {
				fmt.Println("[!] Do not support to get multiple repositories.")
			}
		},
	}
)

func init() {
	getCmd.AddCommand(getRepoCmd)
}
