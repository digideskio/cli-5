package files

import (
	"fmt"
	"os"

	"github.com/catalyzeio/cli/models"
	"github.com/jawher/mow.cli"
)

// Cmd is the contract between the user and the CLI. This specifies the command
// name, arguments, and required/optional arguments and flags for the command.
var Cmd = models.Command{
	Name:      "files",
	ShortHelp: "Tasks for managing service files",
	LongHelp:  "Tasks for managing service files",
	CmdFunc: func(settings *models.Settings) func(cmd *cli.Cmd) {
		return func(cmd *cli.Cmd) {
			cmd.Command(DownloadSubCmd.Name, DownloadSubCmd.ShortHelp, DownloadSubCmd.CmdFunc(settings))
			cmd.Command(ListSubCmd.Name, ListSubCmd.ShortHelp, ListSubCmd.CmdFunc(settings))
		}
	},
}

var DownloadSubCmd = models.Command{
	Name:      "download",
	ShortHelp: "Download a file to your localhost with the same file permissions as on the remote host or print it to stdout",
	LongHelp:  "Download a file to your localhost with the same file permissions as on the remote host or print it to stdout",
	CmdFunc: func(settings *models.Settings) func(cmd *cli.Cmd) {
		return func(subCmd *cli.Cmd) {
			serviceName := subCmd.StringArg("SERVICE_NAME", "", "The name of the service to download a file from")
			fileName := subCmd.StringArg("FILE_NAME", "", "The name of the service file from running \"catalyze files list\"")
			output := subCmd.StringOpt("o output", "", "The downloaded file will be saved to the given location with the same file permissions as it has on the remote host. If those file permissions cannot be applied, a warning will be printed and default 0644 permissions applied. If no output is specified, stdout is used.")
			force := subCmd.BoolOpt("f force", false, "If the specified output file already exists, automatically overwrite it")
			subCmd.Action = func() {
				ifiles := New(settings, *serviceName, *fileName, *output, *force)
				err := CmdDownload(ifiles)
				if err != nil {
					fmt.Println(err.Error())
					os.Exit(1)
				}
			}
			subCmd.Spec = "SERVICE_NAME FILE_NAME [-o] [-f]"
		}
	},
}

var ListSubCmd = models.Command{
	Name:      "list",
	ShortHelp: "List all files available for a given service",
	LongHelp:  "List all files available for a given service",
	CmdFunc: func(settings *models.Settings) func(cmd *cli.Cmd) {
		return func(subCmd *cli.Cmd) {
			serviceName := subCmd.StringArg("SERVICE_NAME", "", "The name of the service to list files for")
			subCmd.Action = func() {
				ifiles := New(settings, *serviceName, *fileName, *output, *force)
				err := CmdList(ifiles)
				if err != nil {
					fmt.Println(err.Error())
					os.Exit(1)
				}
			}
			subCmd.Spec = "SERVICE_NAME"
		}
	},
}

// IFiles
type IFiles interface {
	List() (*[]models.ServiceFile, error)
	Retrieve() (*models.ServiceFile, error)
}

// SFiles is a concrete implementation of IFiles
type SFiles struct {
	Settings *models.Settings
	Services services.IService

	SvcName  string
	FileName string
	Output   string
	Force    bool
}

// New generates a new instance of IFiles
func New(settings *models.Settings, services services.IService, svcName, fileName, output string, force bool) IFiles {
	return &SFiles{
		Settings: settings,
		Services: services,

		SvcName:  svcName,
		FileName: fileName,
		Output:   output,
		Force:    force,
	}
}
