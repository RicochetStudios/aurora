package types

// Game is details about the video game that the server is hosting.
type Game struct {
	Name      string
	Modloader string
}

// Network is configuration of the networking for the instance.
type Network struct {
	Type string
}

// Server is a set of useful details about a game server instance.
type Server struct {
	ID      string
	Size    string
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
