package schema

import (
	"os"
	"regexp"

	"gopkg.in/yaml.v3"
)

type Sizes struct {
	XS Size `yaml:"xs"`
	S  Size `yaml:"s"`
	M  Size `yaml:"m"`
	L  Size `yaml:"l"`
	XL Size `yaml:"xl"`
}

type Size struct {
	Resources Resources `yaml:"resources"`
	Players   int       `yaml:"players"`
}

type Resources struct {
	CPU    string `yaml:"cpu"`
	Memory string `yaml:"memory"`
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
	Command        []string `yaml:"command"`
	StartupProbe   Probe    `yaml:"startupProbe"`
	ReadynessProbe Probe    `yaml:"readynessProbe"`
	LivenessProbe  Probe    `yaml:"livenessProbe"`
}

type Probe struct {
	InitialDelaySeconds int `yaml:"initialDelaySeconds"`
	PeriodSeconds       int `yaml:"periodSeconds"`
	FailureThreshold    int `yaml:"failureThreshold"`
	SuccessThreshold    int `yaml:"successThreshold"`
	TimeoutSeconds      int `yaml:"timeoutSeconds"`
}

type Schema struct {
	Name     string          `yaml:"name"`
	Image    string          `yaml:"image"`
	URL      string          `yaml:"url"`
	Ratio    string          `yaml:"ratio"`
	Sizes    map[string]Size `yaml:"sizes"`
	Network  []Network       `yaml:"network"`
	Settings []Setting       `yaml:"settings"`
	Volumes  []Volume        `yaml:"volumes"`
	Probes   Probes          `yaml:"probes"`
}

// GetSchema gets a game schema from a yaml file and stores it as a Schema.
func GetSchema(game string) (Schema, error) {
	// We need to correct the directory path when testing.
	var dir string = "/"
	wd, err := os.Getwd()
	if err != nil {
		return Schema{}, err
	}
	matched, err := regexp.MatchString(`/schema$`, wd)
	if err != nil {
		return Schema{}, err
	}
	if !matched {
		dir = "/schema/"
	}

	// Load the file; returns []byte.
	f, err := os.ReadFile(wd + dir + game + "/schema.yaml")
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
