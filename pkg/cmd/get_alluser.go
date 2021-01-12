package cmd

import (
	"fmt"

	usr "harborctl/pkg/users"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	getAllUsersCmd = &cobra.Command{
		Use:   "users",
		Short: "List all users of registry",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
			    fmt.Println("[!] Do not support multi arguments.")
			} else {
                            usr.GetAllUsers(viper.GetString("URL"))
                        }        
		},
	}
)

func init() {
	getCmd.AddCommand(getAllUsersCmd)
}
