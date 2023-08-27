package config

import (
	"errors"
	"fmt"
	"os"
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
		ID:        "00000001",
		ClusterID: "myclusterid",
	}

	createGot, createErr := Update(Config{
		ID:        "00000001",
		ClusterID: "myclusterid",
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
	}

	modifyGot, modifyErr := Update(Config{
		ID:        "00000002",
		ClusterID: "mynewclusterid",
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

// TestGetId calls GetId,
// checking for an id in return.
func TestGetId(t *testing.T) {
	// Cleanup at the end of the test.
	t.Cleanup(func() {
		if err := cleanup(); err != nil {
			t.Fatalf("TestGetId() error cleaning up:\n%v", err)
		}
	})

	var want string = "00000001"

	// Setup the config file.
	if _, err := Update(Config{
		ID: "00000001",
	}); err != nil {
		t.Fatalf("TestGetId() error setting up the config: \n%v", err)
	}

	got, err := GetId()

	if err != nil {
		t.Fatalf("TestGetId() error getting the id from the config: \n%v", err)
	}
	if got != want {
		t.Fatalf("TestGetId() = %q, %v, want match for %#q, nil", got, err, want)
	}
}

// TestUpdateId calls UpdateId with an id,
// checking the id was updated.
func TestUpdateId(t *testing.T) {
	// Cleanup at the end of the test.
	t.Cleanup(func() {
		if err := cleanup(); err != nil {
			t.Fatalf("TestUpdateId() error cleaning up:\n%v", err)
		}
	})

	var want string = "00000002"

	// Setup the config file.
	if _, err := Update(Config{
		ID: "00000001",
	}); err != nil {
		t.Fatalf("TestUpdateId() error setting up the config: \n%v", err)
	}

	// Update the id.
	got, err := UpdateId("00000002")

	if err != nil {
		t.Fatalf("TestUpdateId() error updating the id in the config: \n%v", err)
	}
	if got.ID != want {
		t.Fatalf("TestUpdateId() = %q, %v, want match for %#q, nil", got, err, want)
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
