package certs

/*var (
	settings     *models.Settings
	r            = config.FileSettingsRetriever{}
	validEnvName = os.Getenv("ENV_NAME")
)

const (
	invalidEnvName = "invalid-env-name"
	validSvcName   = "code-1"
	invalidSvcName = "invalid-svc-name"
	alias          = "e"
	remote         = "ctlyz"
)

func TestMain(m *testing.M) {
	flag.Parse()
	settings = r.GetSettings(os.Getenv(config.CatalyzeEnvironmentEnvVar), "", os.Getenv(config.AccountsHostEnvVar), os.Getenv(config.AuthHostEnvVar), "", os.Getenv(config.PaasHostEnvVar), "", os.Getenv(config.CatalyzeUsernameEnvVar), os.Getenv(config.CatalyzePasswordEnvVar))
	os.Exit(m.Run())
}

var createTests = []struct {
	name        string
	pubKeyPath  string
	privKeyPath string
	selfSigned  bool
	resolve     bool
	expectErr   bool
}{
	{validEnvName, validSvcName, alias, "", false, false},
	{invalidEnvName, validSvcName, alias, "", false, true},
	{validEnvName, invalidSvcName, alias, "", false, true},
	{validEnvName, validSvcName, "", "", false, false},
	{validEnvName, validSvcName, alias, remote, false, false},
	{validEnvName, validSvcName, alias, "", true, false},
	{validEnvName, validSvcName, alias, "", false, false},
}

func TestAssociate(t *testing.T) {
	mgit := mGit{}
	for _, data := range associateTests {
		t.Logf("%+v\n", data)
		initSettings()
		err := CmdAssociate(data.envLabel, data.svcLabel, data.alias, data.remote, data.defaultEnv, New(settings), &mgit, environments.New(settings), services.New(settings))
		if err != nil != data.expectErr {
			t.Errorf("Unexpected error: %s\n", err.Error())
			continue
		}
		if data.expectErr {
			continue
		}
		name := data.alias
		if name == "" {
			name = data.envLabel
		}

		envNames := []string{}
		found := false
		for key := range settings.Environments {
			envNames = append(envNames, key)
			if key == name {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Environment \"%s\" not added to the settings list of environments or the alias was not used. Found %+v instead", name, envNames)
			continue
		}
		if data.defaultEnv && settings.Default != name {
			t.Error("Default environment specified but was not stored in the settings")
			continue
		}
	}
}

// mock git implementation
type mGit models.Breadcrumb

func (m *mGit) Add(remote, gitURL string) error { return nil }
func (m *mGit) Exists() bool                    { return true }
func (m *mGit) List() ([]string, error)         { return []string{}, nil }
func (m *mGit) Rm(remote string) error          { return nil }
*/
