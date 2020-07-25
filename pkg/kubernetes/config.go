package kubernetes

import (
	"log"

	"gopkg.in/yaml.v2"
)

type dashYaml struct {
	Kubernetes []kubernetesConfig `yaml:"kubernetes"`
}

type kubernetesConfig struct {
	Name       string        `yaml:"name"`
	Kubeconfig string        `yaml:"kubeconfig"`
	Context    string        `yaml:"context"`
	Permission k8sPermission `yaml:"permission"`
	Listen     string        `yaml:"-"`
}

type k8sPermission struct {
	Deployments struct {
		Namespaces []string `yaml:"namespaces" json:"namespaces"`
		Start      []string `yaml:"start" json:"start"`
		Stop       []string `yaml:"stop" json:"stop"`
	} `yaml:"deployments" json:"deployments"`
}

func loadConfig(file []byte) dashYaml {
	dc := dashYaml{}

	err := yaml.Unmarshal(file, &dc)
	if err != nil {
		log.Fatalln("parse yaml config", err)
	}

	return dc
}
