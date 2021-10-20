package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var updateCmd = &cobra.Command{
	Use:   "update <resource> <name>",
	Short: `Get resource`,
	Long:  `WIP`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("update What?!")
		} else {
			
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
