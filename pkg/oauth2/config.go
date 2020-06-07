package oauth2

import (
	"github.com/apex/log"
	"gopkg.in/yaml.v2"
)

type dashYaml struct {
	Oauth2 []struct {
		Provider        string   `yaml:"provider"`
		ClientID        string   `yaml:"clientId"`
		ClientSecret    string   `yaml:"clientSecret"`
		AuthURL         string   `yaml:"authURL"`
		TokenURL        string   `yaml:"tokenURL"`
		URLLoginSuccess string   `yaml:"urlLoginSuccess"`
		Scopes          []string `yaml:"scopes"`
	} `yaml:"oauth2"`
}

func loadConfig(file []byte) dashYaml {
	dc := dashYaml{}

	err := yaml.Unmarshal(file, &dc)
	if err != nil {
		log.WithError(err).Fatal("parse yaml config")
	}

	return dc
}