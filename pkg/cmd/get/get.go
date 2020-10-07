package get

import (
	"github.com/jenkins-x-plugins/jx-health/pkg/cmd/get/status"
	"github.com/jenkins-x/jx-helpers/v3/pkg/cobras"
	"github.com/jenkins-x/jx-logging/v3/pkg/log"
	"github.com/spf13/cobra"
)

// NewCmdGet the get command
func NewCmdGet() *cobra.Command {
	command := &cobra.Command{
		Use:   "get",
		Short: "used for getting resources",
		Run: func(command *cobra.Command, args []string) {
			err := command.Help()
			if err != nil {
				log.Logger().Errorf(err.Error())
			}
		},
	}
	command.AddCommand(cobras.SplitCommand(status.NewCmdStatus()))
	return command
}
