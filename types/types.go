package types

// Game is details about the video game that the server is hosting.
type Game struct {
	Name      string `json:"name" yaml:"name" xml:"name" form:"name"`
	Modloader string `json:"modloader" yaml:"modloader" xml:"modloader" form:"modloader"`
}

// Network is configuration of the networking for the instance.
type Network struct {
	Type    string `json:"type" yaml:"type" xml:"type" form:"type"`
	Address string `json:"address" yaml:"address" xml:"address" form:"address"`
}

// Server is a set of useful details about a game server instance.
type Server struct {
	Size    string  `json:"size" yaml:"size" xml:"size" form:"size"`
	Game    Game    `json:"game" yaml:"game" xml:"game" form:"game"`
	Network Network `json:"network" yaml:"network" xml:"network" form:"network"`
	Status  string  `json:"status" yaml:"status" xml:"status" form:"status"`
}
