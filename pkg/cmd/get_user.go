package cmd

import (
	"fmt"

	usr "harborctl/pkg/users"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	getUserCmd = &cobra.Command{
		Use:   "user",
		Short: "List user's information of registry",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("[!] Please input username to get information.")
			} else if len(args) == 1 {
				usr.GetUser(viper.GetString("URL"), args[0])
			} else {
				fmt.Println("[!] Do not support.")
			}
		},
	}
)

func init() {
	getCmd.AddCommand(getUserCmd)
}
