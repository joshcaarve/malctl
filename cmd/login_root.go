package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login <resource> <name>",
	Short: `Get resource`,
	Long:  `WIP`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("login What?!")
		} else {

		}
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
