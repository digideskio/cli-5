package files

import (
	"fmt"
	"os"
	"strconv"

	"github.com/catalyzeio/cli/helpers"
	"github.com/catalyzeio/cli/models"
)

// CmdList lists all service files that are able to be downloaded
// by a member of the environment. Typically service files of interest
// will be on the service_proxy.
func CmdList(ifiles IFiles) error {
	service := helpers.RetrieveServiceByLabel(serviceName, settings)
	if service == nil {
		fmt.Printf("Could not find a service with the name \"%s\"\n", serviceName)
		os.Exit(1)
	}
	files := helpers.ListServiceFiles(service.ID, settings)
	if len(*files) == 0 {
		fmt.Println("No service files found")
		return
	}
	fmt.Println("NAME")
	for _, sf := range *files {
		fmt.Println(sf.Name)
	}
}

func fileModeToRWXString(perms uint64) string {
	permissionString := ""
	binaryString := strconv.FormatUint(perms, 2)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if string(binaryString[len(binaryString)-1-i*3-j]) == "1" {
				switch j {
				case 0:
					permissionString = "x" + permissionString
				case 1:
					permissionString = "w" + permissionString
				case 2:
					permissionString = "r" + permissionString
				}
			} else {
				permissionString = "-" + permissionString
			}
		}
	}
	permissionString = "-" + permissionString // we don't store folders
	return permissionString
}

func (f *SFiles) List() (*[]models.ServiceFile, error) {
	return nil, nil
}
