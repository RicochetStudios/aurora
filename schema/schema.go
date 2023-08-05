package schema

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Resources struct {
	CPU    string `yaml:"cpu"`
	Memory string `yaml:"memory"`
}

type Sizes struct {
	XS Resources `yaml:"xs"`
	S  Resources `yaml:"s"`
	M  Resources `yaml:"m"`
	L  Resources `yaml:"l"`
	XL Resources `yaml:"xl"`
}

type Network struct {
	Name     string `yaml:"name"`
	Port     int    `yaml:"port"`
	Protocol string `yaml:"protocol"`
}

type Setting struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

type Volume struct {
	Name  string `yaml:"name"`
	Path  string `yaml:"path"`
	Class string `yaml:"class"`
	Size  string `yaml:"size"`
}

type Probes struct {
	Command      []string `yaml:"command"`
	StartupProbe struct {
		FailureThreshold int `yaml:"failureThreshold"`
		PeriodSeconds    int `yaml:"periodSeconds"`
	} `yaml:"startupProbe"`
	ReadyLiveProbe struct {
		InitialDelaySeconds int `yaml:"initialDelaySeconds"`
		PeriodSeconds       int `yaml:"periodSeconds"`
		FailureThreshold    int `yaml:"failureThreshold"`
		SuccessThreshold    int `yaml:"successThreshold"`
		TimeoutSeconds      int `yaml:"timeoutSeconds"`
	} `yaml:"readyLiveProbe"`
}

type Schema struct {
	Name     string    `yaml:"name"`
	Image    string    `yaml:"image"`
	URL      string    `yaml:"url"`
	Ratio    string    `yaml:"ratio"`
	Sizes    Sizes     `yaml:"sizes"`
	Network  []Network `yaml:"network"`
	Settings []Setting `yaml:"settings"`
	Volumes  []Volume  `yaml:"volumes"`
	Probes   Probes    `yaml:"probes"`
}

func GetSchema(game string) (Schema, error) {
	// Load the file; returns []byte.
	f, err := os.ReadFile(game + "/schema.yaml")
	if err != nil {
		return Schema{}, err
	}

	// Create an empty Schema to be are target of unmarshalling.
	var schema Schema

	// Unmarshal the YAML file into empty Schema.
	if err := yaml.Unmarshal(f, &schema); err != nil {
		return Schema{}, err
	}

	return schema, nil
}
