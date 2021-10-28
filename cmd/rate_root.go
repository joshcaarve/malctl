package cmd

import (
	"fmt"
	"malctl/cmd/rate"
	"malctl/db"
	"malctl/guess"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var rateCmd = &cobra.Command{
	Use:   "rate",
	Short: `rate`,
	Long:  `WIP`,
	Run: func(cmd *cobra.Command, args []string) {
		if args[0] == "anime" {
			if len(args) != 4 {
				os.Exit(1)
			}
			var id int
			if args[1] == "id" {
				id, _ = strconv.Atoi(args[2])
			} else if args[1] == "name" {
				idStr := db.GetAnimeIDFromName(args[2])
				if idStr == "" {
					id, _ = strconv.Atoi(guess.GetIDFromGuess(args[2]))
				} else {
					id, _ = strconv.Atoi(idStr)
				}
			} else if args[1] == "guess" {
				id, _ = strconv.Atoi(guess.GetIDFromGuess(args[2]))
			}
			rating, _ := strconv.Atoi(args[3])
			err := rate.RateAnime(id, rating)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("rate what?")
		}
	},
}

func init() {
	rootCmd.AddCommand(rateCmd)
}
