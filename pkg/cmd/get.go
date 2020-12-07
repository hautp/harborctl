package cmd

import (
	"fmt"
	pj "harborctl/pkg/project"
	rp "harborctl/pkg/repository"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	getCmd = &cobra.Command{
		Use:   "get",
		Short: "Display one or many resources",
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

`, rootCmd.Use)
				os.Exit(0)
			}

			// Project section
			if strings.Contains(args[0], "project") {
				if len(args) == 1 {
					pj.ListProjects(viper.GetString("URL"))
				} else if len(args) == 2 {
					pj.GetProject(viper.GetString("URL"), args[1])
				} else {
					fmt.Println("[!] Do not support to get multiple projects.")
				}

				// Repository section
			} else if strings.Contains(args[0], "repo") {
				if len(args) == 1 {
					fmt.Println("[!] Please input project name to get repository.")
				} else if len(args) == 2 {
					rp.ListRepositories(viper.GetString("URL"), args[1])
				} else {
					fmt.Println("[!] Do not support to get multiple repositories.")
				}

				// Artifact section
			} else if strings.Contains(args[0], "tags") {
				if len(args) == 1 {
					fmt.Println("[!] Please input repository name to get artifacts.")
				} else if len(args) == 2 {
					rp.ListArtifactsToTable(viper.GetString("URL"), args[1])
				} else {
					fmt.Println("[!] Do not support to get artifacts on multiple repositories.")
				}

				// Exception
			} else {
				fmt.Printf("[!] Resource \"%s\" not found.\n", args[0])
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(getCmd)
}
