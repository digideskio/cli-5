package catalyze

import (
	"flag"
	"fmt"
	"os"
	"reflect"
	"testing"
	"unsafe"

	"github.com/catalyzeio/cli/config"
	"github.com/catalyzeio/cli/models"
	"github.com/jawher/mow.cli"
)

var app *cli.Cli

func setup() {
	app = cli.App("catalyze", fmt.Sprintf("Catalyze CLI. Version %s", config.VERSION))
	settings := &models.Settings{}
	InitGlobalOpts(app, settings)
	InitCLI(app, settings)
}

func TestMain(m *testing.M) {
	flag.Parse()
	setup()
	os.Exit(m.Run())
}

var globalArgsData = map[string]string{
	"-U": "--username",
	"-P": "--password",
	"-E": "--env",
	"-v": "--version",
}

func TestGlobalArgs(t *testing.T) {
	s := reflect.ValueOf(*app)
	optsFields := s.FieldByName("options")
	for i := 0; i < optsFields.Len(); i++ {
		names := optsFields.Index(i).Elem().FieldByName("names")
		shortName := names.Index(0).String()
		longName := ""
		if names.Len() > 1 {
			longName = names.Index(1).String()
		}
		expectedLongName, ok := globalArgsData[shortName]
		if !ok {
			t.Errorf("Global arg not found: %s. Was %s removed or renamed?", shortName, shortName)
			continue
		}
		if longName != expectedLongName {
			t.Errorf("Global arg long name incorrect for arg %s: expected %s but got %s", shortName, expectedLongName, longName)
			continue
		}
	}
}

func TestCommandArgs(t *testing.T) {
	s := reflect.ValueOf(*app)
	cmdField := s.FieldByName("commands")
	ptr := unsafe.Pointer(cmdField.UnsafeAddr())
	cmds := (*[]*cli.Cmd)(ptr)
	for _, c := range *cmds {
		t.Logf("cmd: %+v\n", c)
	}
	//t.Fail()
}
