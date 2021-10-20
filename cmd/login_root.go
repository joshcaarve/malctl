package cmd

import (
	"fmt"
	"malctl/cmd/login"

	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: `login`,
	Long:  `WIP`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			login.LoginUser()
		} else {
			fmt.Println("too many args? :)")
		}
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
