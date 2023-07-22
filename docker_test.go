package main

import (
	"testing"
)

// TestNewServerPort calls NewServerPort with a protocol and port,
// checking for a valid ServerPort type to be returned.
func TestNewServerPort(t *testing.T) {
	var want ServerPort = ServerPort{Protocol: "tcp", Port: "8080"}
	result := NewServerPort("tcp", "8080")
	if want != result {
		t.Fatalf(`NewServerPort("tcp", "8080") = %q, want match for %#q`, result, want)
	}
}

// TestNewServerMountPath calls NewServerMount with a path,
// checking for a valid ServerMount returned.
func TestNewServerMountPath(t *testing.T) {
	var want ServerMount = "/data"
	result, err := NewServerMount("/data")
	if want != result || err != nil {
		t.Fatalf(`NewServerMount("/data") = %q, %v, want match for %#q, nil`, result, err, want)
	}
}

// TestNewServerMountInvalidPath calls NewServerMount with an invalid path,
// checking for an error in return.
func TestNewServerMountInvalidPath(t *testing.T) {
	// Missing the directory path '/'.
	_, err := NewServerMount("data")
	if err == nil {
		t.Fatalf(`expected a mount target invalid error while testing TestNewServerMountInvalidPath, got %T`, err)
	}
}

// TestNewServerEnvVar calls NewServerEnvVar with a name and value,
// checking for a valid ServerEnvVar in return.
func TestNewServerEnvVar(t *testing.T) {
	var want ServerEnvVar = ServerEnvVar{Name: "EULA", Value: "TRUE"}
	result, err := NewServerEnvVar("EULA", "TRUE")
	if want != result || err != nil {
		t.Fatalf(`NewServerMount("EULA", "TRUE") = %q, %v, want match for %#q, nil`, result, err, want)
	}
}

// TestNewServerEnvVarInvalidName calls NewServerEnvVar with an invalid name,
// checking for an error in return.
func TestNewServerEnvVarInvalidName(t *testing.T) {
	// Add invalid characters to the env var name.
	_, err := NewServerEnvVar("$EULA!", "TRUE")
	if err == nil {
		t.Fatalf(`expected a name invalid error while testing TestNewServerEnvVarInvalidName, got %T`, err)
	}
}
