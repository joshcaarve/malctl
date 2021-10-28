package cmd

import (
	"fmt"
	"malctl/cmd/seed"

	"github.com/spf13/cobra"
)

var seedCmd = &cobra.Command{
	Use:   "seed <resource> <name>",
	Short: `Get resource`,
	Long:  `WIP`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			seed.SeedAnimeHash("")
		} else {
			fmt.Println("too many args :)")
		}
	},
}

func init() {
	rootCmd.AddCommand(seedCmd)
}
