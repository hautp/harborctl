package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	getCmd = &cobra.Command{
		Use:       "get",
		Short:     "Display one or many resources",
		Args:      cobra.OnlyValidArgs,
		ValidArgs: []string{"project", "repo", "tags"},
		Run: func(cmd *cobra.Command, args []string) {

			// Check len(args)
			if len(args) < 1 {
				fmt.Printf(`[!] You must specify the type of resource to get.

# Project
$ %[1]s get project # Get details all projects 
$ %[1]s get project <PROJECT_NAME> # Get details project <PROJECT_NAME>

# Repository
$ %[1]s get repo <PROJECT_NAME> # Get all repositories of <PROJECT_NAME>

# Artifact
$ %[1]s get tags <REPOSITORY_NAME> # Get artifacts of <REPOSITORY_NAME>

# All users
$ %[1]s get users # Get all user of registry

# Get a user
$ %[1]s get user <USERNAME> # Get information of <USERNAME>

`, rootCmd.Use)
				os.Exit(0)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(getCmd)
}
