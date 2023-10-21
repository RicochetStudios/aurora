package models

import (
	"github.com/RicochetStudios/aurora/types"
)

// ClusterId is the model for a cluster id.
type ClusterId struct {
	ClusterId string `json:"clusterId" validate:"required"`
}

// GameServer is the model for a game server.
type GameServer struct {
	Name string `json:"name" validate:"required"`
	Size string `json:"size" validate:"required"`
	Game types.Game `json:"game" validate:"required"`
	Network types.Network `json:"network" validate:"required"`
}