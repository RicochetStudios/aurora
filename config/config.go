package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"dario.cat/mergo"
)

// configPath is the path to the config file.
const configPath string = "/aurora-config.json"

// Config is a struct of the local, persistent configuration of this instance.
type Config struct {
	ID        string `json:"id" yaml:"id" xml:"id" form:"id"`                             // The identifier of the instance.
	ClusterID string `json:"clusterId" yaml:"clusterId" xml:"clusterId" form:"clusterId"` // The cluster this instance belongs to.
}

// Update creates or modifies config properties.
func Update(newConfig Config) (Config, error) {
	// Get the file.
	file, err := getFile()
	if err != nil {
		return Config{}, fmt.Errorf("Update() error getting file: %v", err)
	}
	defer file.Close()

	// Read the existing config.
	config, err := Read()
	if err != nil {
		return Config{}, fmt.Errorf("Update() error reading config: %v", err)
	}

	// Merge the configs, taking precidence from the input config.
	// WithOverwriteWithEmptyValue allows us to overwrite populated fields even with empty fields.
	if err := mergo.Merge(&config, newConfig, mergo.WithOverwriteWithEmptyValue); err != nil {
		return Config{}, fmt.Errorf("Update() error merging configs: %v", err)
	}

	// Update the file.
	as_json, jsonErr := json.MarshalIndent(config, "", "\t")
	if jsonErr != nil {
		return Config{}, fmt.Errorf("Update() error converting config to json: %v", err)
	}
	if writeErr := os.WriteFile(file.Name(), as_json, 0666); writeErr != nil {
		return Config{}, fmt.Errorf("Update() error writing json to file: %v", err)
	}

	return config, nil
}

// Read returns the current configuration.
func Read() (Config, error) {
	// Get the file.
	file, err := getFile()
	if err != nil {
		return Config{}, fmt.Errorf("Read() error getting file: %v", err)
	}
	defer file.Close()

	// Load the file; returns []byte.
	content, err := os.ReadFile(file.Name())
	if err != nil {
		return Config{}, fmt.Errorf("Read() error reading json from file: %v", err)
	}

	// Create an empty Config to be are target of unmarshalling.
	var config Config

	// Unmarshal the JSON file into empty Config.
	if err := json.Unmarshal(content, &config); err != nil {
		return Config{}, fmt.Errorf("Read() error converting json from file to Config: %v", err)
	}

	return config, nil
}

// GetId gets the instance id from the config.
func GetId() (string, error) {
	// Read the existing config.
	config, err := Read()
	if err != nil {
		return "", fmt.Errorf("GetId() error reading config: %v", err)
	}

	return config.ID, nil
}

// UpdateId updates the instance id in the config, given an id.
func UpdateId(id string) (Config, error) {
	// Read the existing config.
	cfg, err := Read()
	if err != nil {
		return Config{}, fmt.Errorf("UpdateId() error reading config: %v", err)
	}

	// Replace the ID.
	cfg.ID = id

	// Update the config with changes.
	cfg, err = Update(cfg)
	if err != nil {
		return Config{}, fmt.Errorf("UpdateId() error updating config: %v", err)
	}

	return cfg, nil
}

// getFile returns the config file.
// Additionally it will create the file if it does not exist.
func getFile() (*os.File, error) {
	var file *os.File

	// Get the working directory.
	wd, wdErr := os.Getwd()
	if wdErr != nil {
		return file, fmt.Errorf("getFile() error getting working directory: %v", wdErr)
	}

	// Check path exists.
	_, pathErr := os.Stat(wd + configPath)

	// Scope variables correctly.
	var err error

	if errors.Is(pathErr, os.ErrNotExist) {
		// Create the file if it doesn't exist.
		file, err = os.Create(wd + configPath)
		if err != nil {
			return &os.File{}, fmt.Errorf("getFile() error creating file: %v", err)
		}

		// Create an empty json file.
		as_json, jsonErr := json.MarshalIndent(struct{}{}, "", "\t")
		if jsonErr != nil {
			return &os.File{}, fmt.Errorf("getFile() error marshalling struct{}{} to json: %v", jsonErr)
		}
		if writeErr := os.WriteFile(file.Name(), as_json, 0666); writeErr != nil {
			return &os.File{}, fmt.Errorf("getFile() error writing to file: %v", writeErr)
		}
	} else if !errors.Is(pathErr, os.ErrNotExist) {
		// Otherwise return the existing file.
		file, err = os.Open(wd + configPath)
		if err != nil {
			return &os.File{}, fmt.Errorf("getFile() error opening file: %v", err)
		}
	} else {
		return &os.File{}, fmt.Errorf("getFile() error checking path exists: %v", pathErr)
	}

	return file, nil
}
