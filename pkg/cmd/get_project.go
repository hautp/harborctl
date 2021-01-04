package cmd

import (
	"fmt"
	pj "harborctl/pkg/project"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	getProjectCmd = &cobra.Command{
		Use:   "project",
		Short: "Get details all projects",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				pj.ListProjects(viper.GetString("URL"))
			} else if len(args) == 1 {
				pj.GetProject(viper.GetString("URL"), args[0])
			} else {
				fmt.Println("[!] Do not support to get multiple projects.")
			}
		},
	}
)

func init() {
	getCmd.AddCommand(getProjectCmd)
}
