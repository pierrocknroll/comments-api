package parameters

import (
	"comments-api/internal"
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print version number",
	Long:  `print the Comments API service version number`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Comments API version %s\n", internal.PROJECT_VERSION)
	},
}
