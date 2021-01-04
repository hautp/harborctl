package cmd

import (
	"fmt"
	rp "harborctl/pkg/repository"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	getTagsCmd = &cobra.Command{
		Use:   "tags",
		Short: "",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("[!] Please input repository name to get artifacts.")
			} else if len(args) == 1 {
				rp.ListArtifactsToTable(viper.GetString("URL"), args[0])
			} else {
				fmt.Println("[!] Do not support to get artifacts on multiple repositories.")
			}
		},
	}
)

func init() {
	getCmd.AddCommand(getTagsCmd)
}
