package cmd

import (
	"fmt"
	ganime "malctl/cmd/get/anime"
	guser "malctl/cmd/get/user"
	"malctl/db"

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
			if len(args) == 1 {
				fmt.Println("get user <>")
			} else {
				guser.GetUserInfo(args[1:])
			}

		} else if args[0] == "anime" {
			if len(args) == 1 {
				fmt.Println("get anime <>")
			} else {
				if args[1] == "id" {
					ganime.GetAnimeInfo(args[2])
				} else if args[1] == "name" {
					id := db.GetAnimeIDFromName(args[2])
					if id == "" {
						fmt.Println("Unknown Name")
					} else {
						ganime.GetAnimeInfo(id)
					}
				}
			}
		} else {
			fmt.Println("Incorrect resource type")
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
