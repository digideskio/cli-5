package worker

import (
	"github.com/Sirupsen/logrus"
	"github.com/catalyzeio/cli/commands/services"
	"github.com/catalyzeio/cli/config"
	"github.com/catalyzeio/cli/lib/auth"
	"github.com/catalyzeio/cli/lib/prompts"
	"github.com/catalyzeio/cli/models"
	"github.com/jawher/mow.cli"
)

// Cmd is the contract between the user and the CLI. This specifies the command
// name, arguments, and required/optional arguments and flags for the command.
var Cmd = models.Command{
	Name:      "worker",
	ShortHelp: "Start a background worker",
	LongHelp:  "Start a background worker",
	CmdFunc: func(settings *models.Settings) func(cmd *cli.Cmd) {
		return func(cmd *cli.Cmd) {
			serviceName := cmd.StringArg("SERVICE_NAME", "", "The name of the service to use to start a worker. Defaults to the associated service.")
			target := cmd.StringArg("TARGET", "", "The name of the Procfile target to invoke as a worker")
			cmd.Action = func() {
				if _, err := auth.New(settings, prompts.New()).Signin(); err != nil {
					logrus.Fatal(err.Error())
				}
				if err := config.CheckRequiredAssociation(true, true, settings); err != nil {
					logrus.Fatal(err.Error())
				}
				err := CmdWorker(*serviceName, settings.ServiceID, *target, New(settings), services.New(settings))
				if err != nil {
					logrus.Fatal(err.Error())
				}
			}
			cmd.Spec = "[SERVICE_NAME] TARGET"
		}
	},
}

// IWorker
type IWorker interface {
	Start(svcID, target string) error
}

// SWorker is a concrete implementation of IWorker
type SWorker struct {
	Settings *models.Settings
}

// New returns an instance of IWorker
func New(settings *models.Settings) IWorker {
	return &SWorker{
		Settings: settings,
	}
}
