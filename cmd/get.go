package cmd

import (
	"fmt"
	"malctl/cmd/get/user"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get <resource> <name>",
	Short: `Get resource`,
	Long:  `WIP`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Get What?!")
		} else if args[0] == "user" {
			user.GetUserInfo(args[1:])
		} else if args[0] == "anime" {
			fmt.Println("WIP")
		} else {
			fmt.Println("Incorrect resource type")
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
