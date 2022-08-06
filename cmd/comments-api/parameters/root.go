package parameters

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "comments-api",
	Short: "comments-api",
	Long:  `comments-api`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Usage()
		if err != nil {
			log.Debugf("Error on usage [err: %v]", err)
		}
	},
}
