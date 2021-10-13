package cmd

import (
	"fmt"
	ganime "malctl/cmd/get/anime"
	guser "malctl/cmd/get/user"

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
			guser.GetUserInfo(args[1:])
		} else if args[0] == "anime" {
			ganime.GetAnimeInfo(args[1:])
		} else {
			fmt.Println("Incorrect resource type")
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
