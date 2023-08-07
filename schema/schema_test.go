package schema

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetSchema(t *testing.T) {
	// Set desired result.
	var want Schema = Schema{
		Name:  "minecraft_java",
		Image: "itzg/minecraft-server:latest",
		URL:   "https://github.com/itzg/docker-minecraft-server",
		Ratio: "1-2",
		Sizes: Sizes{
			XS: Size{
				Resources: Resources{
					CPU:    "1000m",
					Memory: "2000Mi",
				},
				Players: 8,
			},
			S: Size{
				Resources: Resources{
					CPU:    "1500m",
					Memory: "4000Mi",
				},
				Players: 16,
			},
			M: Size{
				Resources: Resources{
					CPU:    "2000m",
					Memory: "8000Mi",
				},
				Players: 32,
			},
			L: Size{
				Resources: Resources{
					CPU:    "3000m",
					Memory: "16000Mi",
				},
				Players: 64,
			},
			XL: Size{
				Resources: Resources{
					CPU:    "4000m",
					Memory: "32000Mi",
				},
				Players: 128,
			},
		},
		Network: []Network{
			{
				Name:     "game",
				Port:     25565,
				Protocol: "tcp",
			},
		},
		Settings: []Setting{
			{
				Name:  "EULA",
				Value: "TRUE",
			},
			{
				Name:  "TYPE",
				Value: "{{ .Values.game.modLoader }}",
			},
			{
				Name:  "MAX_PLAYERS",
				Value: "{{ .size.players }}",
			},
			{
				Name:  "MOTD",
				Value: "{{ .Values.name }}",
			},
		},
		Volumes: []Volume{
			{
				Name:  "data",
				Path:  "/data",
				Class: "classic",
				Size:  "10Gi",
			},
		},
		Probes: Probes{
			Command: []string{"mc-health"},
			StartupProbe: Probe{
				FailureThreshold: 30,
				PeriodSeconds:    10,
			},
			ReadynessProbe: Probe{
				InitialDelaySeconds: 30,
				PeriodSeconds:       5,
				FailureThreshold:    20,
				SuccessThreshold:    3,
				TimeoutSeconds:      1,
			},
			LivenessProbe: Probe{
				InitialDelaySeconds: 30,
				PeriodSeconds:       5,
				FailureThreshold:    20,
				SuccessThreshold:    3,
				TimeoutSeconds:      1,
			},
		},
	}

	// Call the function to test.
	got, err := GetSchema("minecraft_java")

	// Error if results are incorrect.
	if err != nil {
		t.Fatalf("GetSchema() returned an error: \n%v", err)
	}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Fatalf("GetSchema() mismatch (-want +got):\n%s", diff)
	}
}
