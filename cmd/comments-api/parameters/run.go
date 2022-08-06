package parameters

import (
	"comments-api/internal"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(runCmd)
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run service",
	Long:  `run Comments API service`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		internal.Run()
	},
}
