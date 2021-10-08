package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get anime <name",
	Short: `Get anime by name`,
	Long:  `WIP`,
	Run: func(cmd *cobra.Command, args []string) {
		///
		fmt.Println("get called")
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
