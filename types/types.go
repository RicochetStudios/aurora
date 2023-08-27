package types

// Game is details about the video game that the server is hosting.
type Game struct {
	Name      string `json:"name" yaml:"name" xml:"name" form:"name"`                     // Name of the video game.
	Modloader string `json:"modloader" yaml:"modloader" xml:"modloader" form:"modloader"` // Software used to load mods into the game, or vanilla (no modloader).
}

// Network contains details about connecting to the server.
type Network struct {
	Type    string `json:"type" yaml:"type" xml:"type" form:"type"`             // Whether the server is available on the internet (public) or connected to via vpn (private).
	Address string `json:"address" yaml:"address" xml:"address" form:"address"` // Public or private IP of the server.
}

// Server is a set of useful details about a game server instance.
type Server struct {
	Size    string  `json:"size" yaml:"size" xml:"size" form:"size"`             // Scale of the server. Effects the resources allocated.
	Game    Game    `json:"game" yaml:"game" xml:"game" form:"game"`             // Details about the video game that the server is hosting.
	Network Network `json:"network" yaml:"network" xml:"network" form:"network"` // Networking configuration of the server.
	Status  string  `json:"status" yaml:"status" xml:"status" form:"status"`     // Condition of the server.
}

// Instance is a single item of a game server and an id.
type Instance struct {
	ID     string `json:"id" yaml:"id" xml:"id" form:"id"`                 // The identifier of the instance.
	Server Server `json:"server" yaml:"server" xml:"server" form:"server"` // Details about a game server instance.
}
