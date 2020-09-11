package models

// Struct for storing the coordinates
type Coordinates struct {
	X int
	Y int
}

// Struct for storing the arena
type Arena struct {
	TopRight   Coordinates
	BottomLeft Coordinates
}
