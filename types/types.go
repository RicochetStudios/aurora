package types

// Game is details about the video game that the server is hosting.
type Game struct {
	Name      string `json:"name" xml:"name" form:"name"`
	Modloader string `json:"modloader" xml:"modloader" form:"modloader"`
}

// Network is configuration of the networking for the instance.
type Network struct {
	Type string `json:"type" xml:"type" form:"type"`
}

// Server is a set of useful details about a game server instance.
type Server struct {
	ID      string `json:"id" xml:"id" form:"id"`
	Size    string `json:"size" xml:"size" form:"size"`
	Game    Game
	Network Network
}

// // Server is a set of useful details about a game server instance.
// type Server struct {
// 	// ID: string // The tenant unique identifier of the game server instance.
// 	// Name: string // The tenant unique name of the game server instance.
// 	// TenantID string // The tenant this server instance belongs to.

// 	Container:
// }
