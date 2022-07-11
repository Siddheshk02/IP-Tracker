package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// traceCmd represents the trace command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Get the version",
	Long:  `Get the version.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v1.0.1")

	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

}
