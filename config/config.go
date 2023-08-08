package config

import (
	"encoding/json"
	"os"
	"ricochet/aurora/types"
)

// configPath is the path to the config file.
var configPath string = "/config.json"

// Config is a struct of the local, persistent configuration of this instance.
type Config struct {
	ClusterID string       // The cluster this instance belongs to.
	ID        string       // The identifier of the server.
	Server    types.Server // The server configuration.
}

// Init runs a setup of the configuration on first run of the application.
func Init() error {
	return nil
}

// Update creates or modifies config properties.
func Update(config Config) (Config, error) {
	// Get the working directory.
	wd, err := os.Getwd()
	if err != nil {
		return Config{}, err
	}

	as := []A{
		{Name: "John", Surname: "Black"},
		{Name: "Mary", Surname: "Brown"},
	}
	as_json, _ := json.MarshalIndent(as, "", "\t")
	f.Write(as_json)
}

// Read returns the current configuration.
func Read() (Config, error) {
	// Get the working directory.
	wd, err := os.Getwd()
	if err != nil {
		return Config{}, err
	}

	// Load the file; returns []byte.
	f, err := os.ReadFile(wd + configPath)
	if err != nil {
		return Config{}, err
	}

	// Create an empty Config to be are target of unmarshalling.
	var config Config

	// Unmarshal the JSON file into empty Config.
	if err := json.Unmarshal(f, &config); err != nil {
		return Config{}, err
	}

	return config, nil
}

func createFile() (*os.File, error) {
	// Get the working directory.
	wd, err := os.Getwd()
	if err != nil {
		return &os.File{}, err
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		// Create the file if it doesn't exist.
		f, err := os.Create(wd + configPath)
		if err != nil {
			return &os.File{}, err
		}
		defer f.Close()
		return f, nil
	} else if os.IsExist(err) {
		// Otherwise return the existing file.
		f, err := os.Open(wd + configPath)
		if err != nil {
			return &os.File{}, err
		}
		defer f.Close()
		return f, nil
	} else {
		if err != nil {
			return &os.File{}, err
		}
	}
}
