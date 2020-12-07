package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version    string = "0.1"
	author     string = "hau.tran"
	versionCmd        = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of Harbor CLI",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Harbor CLI v%s - Author: %s.", version, author)
		},
	}
)

func init() {
	rootCmd.AddCommand(versionCmd)
}
