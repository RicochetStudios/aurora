package config

import (
	"errors"
	"fmt"
	"os"
	"ricochet/aurora/types"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// TestRead calls Read,
// checking for a config in return.
func TestRead(t *testing.T) {
	// Cleanup at the end of the test.
	t.Cleanup(func() {
		if err := cleanup(); err != nil {
			t.Fatalf("TestUpdate() error cleaning up:\n%v", err)
		}
	})

	var want Config = Config{}

	got, err := Read()

	if err != nil {
		t.Fatalf("TestUpdate() returned an error: \n%v", err)
	}
	// Use cmp for more complex types.
	if diff := cmp.Diff(want, got); diff != "" {
		t.Fatalf("TestUpdate() mismatch (-want +got):\n%s", diff)
	}
}

// TestUpdateCreate calls Update with a valid config,
// with no existing config present, it will first create the config,
// then it will update the config.
// It will check the config file and config have been created and updated.
func TestUpdate(t *testing.T) {
	// Cleanup at the end of the test.
	t.Cleanup(func() {
		if err := cleanup(); err != nil {
			t.Fatalf("TestUpdate() (create) error cleaning up:\n%v", err)
		}
	})

	// Get the working directory.
	wd, wdErr := os.Getwd()
	if wdErr != nil {
		t.Fatalf("TestUpdate() (create) could not get the working directory, error: \n%v", wdErr)
	}

	// Test creation.
	var createWant Config = Config{
		ClusterID: "myclusterid",
		Server: types.Server{
			ID:   "00000001",
			Size: "xs",
			Game: types.Game{
				Name:      "minecraft_java",
				Modloader: "fabric",
			},
			Network: types.Network{
				Type: "private",
			},
		},
	}

	createGot, createErr := Update(Config{
		ClusterID: "myclusterid",
		Server: types.Server{
			ID:   "00000001",
			Size: "xs",
			Game: types.Game{
				Name:      "minecraft_java",
				Modloader: "fabric",
			},
			Network: types.Network{
				Type: "private",
			},
		},
	})

	if createErr != nil {
		t.Fatalf("TestUpdate() (create) returned an error: \n%v", createErr)
	}
	// Use cmp for more complex types.
	if diff := cmp.Diff(createWant, createGot); diff != "" {
		t.Fatalf("TestUpdate() (create) mismatch (-want +got):\n%s", diff)
	}

	// Check if the file exists.
	if _, pathErr := os.Stat(wd + configPath); errors.Is(pathErr, os.ErrNotExist) {
		t.Fatalf("TestUpdate() (create) file was not created: \n%v", pathErr)
	}

	// Test updating.
	var modifyWant Config = Config{
		ID:        "00000002",
		ClusterID: "mynewclusterid",
		Server: types.Server{
			ID:   "00000002",
			Size: "xl",
			Game: types.Game{
				Name:      "valheim",
				Modloader: "",
			},
			Network: types.Network{
				Type: "public",
			},
		},
	}

	modifyGot, modifyErr := Update(Config{
		ID:        "00000002",
		ClusterID: "mynewclusterid",
		Server: types.Server{
			ID:   "00000002",
			Size: "xl",
			Game: types.Game{
				Name:      "valheim",
				Modloader: "",
			},
			Network: types.Network{
				Type: "public",
			},
		},
	})

	if modifyErr != nil {
		t.Fatalf("TestUpdate() (modify) returned an error: \n%v", modifyErr)
	}
	// Use cmp for more complex types.
	if diff := cmp.Diff(modifyWant, modifyGot); diff != "" {
		t.Fatalf("TestUpdate() (modify) mismatch (-want +got):\n%s", diff)
	}

	// Check if the file exists.
	if _, pathErr := os.Stat(wd + configPath); errors.Is(pathErr, os.ErrNotExist) {
		t.Fatalf("TestUpdateCreate() (modify) file was not created: \n%v", pathErr)
	}
}

// cleanup removes the config file if it exists.
func cleanup() error {
	// Get the working directory.
	wd, wdErr := os.Getwd()
	if wdErr != nil {
		return fmt.Errorf("cleanup() error getting working directory: %v", wdErr)
	}

	// Remove the file if it exists.
	if _, pathErr := os.Stat(wd + configPath); !errors.Is(pathErr, os.ErrNotExist) {
		if err := os.Remove(wd + configPath); err != nil {
			return fmt.Errorf("cleanup() error removing file: %v", err)
		}
	}

	return nil
}
