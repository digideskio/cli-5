package disassociate

import (
	"github.com/Sirupsen/logrus"
	"github.com/catalyzeio/cli/config"
	"github.com/catalyzeio/cli/models"
	"github.com/jault3/mow.cli"
)

// Cmd is the contract between the user and the CLI. This specifies the command
// name, arguments, and required/optional arguments and flags for the command.
var Cmd = models.Command{
	Name:      "disassociate",
	ShortHelp: "Remove the association with an environment",
	LongHelp: "`disassociate` removes the environment from your list of associated environments but **does not** remove the catalyze git remote on the git repo. " +
		"Disassociate does not have to be run from within a git repo. Here is a sample command\n\n" +
		"```\ncatalyze disassociate myprod\n```",
	CmdFunc: func(settings *models.Settings) func(cmd *cli.Cmd) {
		return func(cmd *cli.Cmd) {
			alias := cmd.StringArg("ENV_ALIAS", "", "The alias of an already associated environment to disassociate")
			cmd.Action = func() {
				if err := config.CheckRequiredAssociation(true, false, settings); err != nil {
					logrus.Fatal(err.Error())
				}
				err := CmdDisassociate(*alias, New(settings))
				if err != nil {
					logrus.Fatal(err.Error())
				}
			}
			cmd.Spec = "ENV_ALIAS"
		}
	},
}

// IDisassociate
type IDisassociate interface {
	Disassociate(alias string) error
}

// SDisassociate is a concrete implementation of IDisassociate
type SDisassociate struct {
	Settings *models.Settings
}

// New returns an instance of IDisassociate
func New(settings *models.Settings) IDisassociate {
	return &SDisassociate{
		Settings: settings,
	}
}
