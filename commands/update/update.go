package update

import (
	"fmt"
	"runtime"

	"github.com/Sirupsen/logrus"

	"github.com/catalyzeio/cli/lib/pods"
	"github.com/catalyzeio/cli/lib/updater"
	"github.com/catalyzeio/cli/models"
)

func CmdUpdate(iu IUpdate) error {
	logrus.Println("Checking for available updates...")
	needsUpdate, err := iu.Check()
	if err != nil {
		return err
	}
	// check if we can overwrite exe
	if needsUpdate && (runtime.GOOS == "linux" || runtime.GOOS == "darwin") {
		err = verifyExeDirWriteable()
		if err != nil {
			return err
		}
	}
	if !needsUpdate {
		logrus.Println("You are already running the latest version of the Catalyze CLI")
		return nil
	}
	logrus.Printf("Version %s is available. Updating your CLI...", updater.AutoUpdater.Info.Version)
	err = iu.Update()
	if err != nil {
		return err
	}
	logrus.Printf("Your CLI has been updated to version %s", updater.AutoUpdater.Info.Version)
	return nil
}

func exeGenericError() error {
	return fmt.Errorf("There was an error trying to find where your CLI is on your system. You may need to manually update your CLI")
}

func (u *SUpdate) Check() (bool, error) {
	updater.AutoUpdater.FetchInfo()
	if updater.AutoUpdater.CurrentVersion == updater.AutoUpdater.Info.Version {
		return false, nil
	}
	return true, nil
}

// Update updates the  CLI if a new update is available.
func (u *SUpdate) Update() error {
	updater.AutoUpdater.ForcedUpgrade()
	return nil
}

func updatePods(settings *models.Settings) {
	logrus.Debugf("Updating cached pods...")
	p := pods.New(settings)
	pods, err := p.List()
	if err == nil {
		settings.Pods = pods
	} else {
		settings.Pods = &[]models.Pod{}
		logrus.Debugf("Error listing pods: %s", err.Error())
	}
}
