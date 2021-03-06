package vars

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/Sirupsen/logrus"
	"github.com/catalyzeio/cli/commands/services"
	"github.com/catalyzeio/cli/lib/httpclient"
	"gopkg.in/yaml.v2"
)

type Formatter interface {
	Output(envVars map[string]string) error
}

type JSONFormatter struct{}

func (j *JSONFormatter) Output(envVars map[string]string) error {
	jsonVars := map[string]string{}
	for k, v := range envVars {
		jsonVars[k] = v
	}
	b, err := json.MarshalIndent(jsonVars, "", "    ")
	if err != nil {
		return err
	}
	logrus.Println(string(b))
	return nil
}

type YAMLFormatter struct{}

func (y *YAMLFormatter) Output(envVars map[string]string) error {
	b, err := yaml.Marshal(envVars)
	if err != nil {
		return err
	}
	logrus.Println(string(b))
	return nil
}

type PlainFormatter struct{}

func (p *PlainFormatter) Output(envVars map[string]string) error {
	var keys []string
	for k := range envVars {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, key := range keys {
		logrus.Printf("%s=%s", key, envVars[key])
	}
	return nil
}

func CmdList(svcName, defaultSvcID string, formatter Formatter, iv IVars, is services.IServices) error {
	if svcName != "" {
		service, err := is.RetrieveByLabel(svcName)
		if err != nil {
			return err
		}
		if service == nil {
			return fmt.Errorf("Could not find a service with the label \"%s\". You can list services with the \"catalyze services\" command.", svcName)
		}
		defaultSvcID = service.ID
	}
	envVars, err := iv.List(defaultSvcID)
	if err != nil {
		return err
	}
	if len(envVars) == 0 {
		logrus.Println("No environment variables found")
		return nil
	}
	return formatter.Output(envVars)
}

// List lists all environment variables.
func (v *SVars) List(svcID string) (map[string]string, error) {
	headers := httpclient.GetHeaders(v.Settings.SessionToken, v.Settings.Version, v.Settings.Pod, v.Settings.UsersID)
	resp, statusCode, err := httpclient.Get(nil, fmt.Sprintf("%s%s/environments/%s/services/%s/env", v.Settings.PaasHost, v.Settings.PaasHostVersion, v.Settings.EnvironmentID, svcID), headers)
	if err != nil {
		return nil, err
	}
	var envVars map[string]string
	err = httpclient.ConvertResp(resp, statusCode, &envVars)
	if err != nil {
		return nil, err
	}
	return envVars, nil
}
