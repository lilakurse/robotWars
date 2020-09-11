package models

import (
	"errors"
	"log"
)
// Struct for storing the robot
type Robot struct {
	Id               int         `json:"id"`	// robot id
	StartPos         Coordinates `json:"-"`		// start coordinates of the robot in the arena
	StartOrientation string      `json:"-"`		// start orientation of the robot in the arena
	CurrPosition     Coordinates `json:"position"`	// current coordinates of the robot in the arena
	CurrOrientation  string      `json:"orientation"`	// current orientation of the robot in the arena
}


// This function make the robot spin 90 degrees to the left without moving from its current spot
func (robot *Robot) TurnLeft() {
	// Logs all robot's moves and navigation
	log.Println("----------------")
	log.Println("Turning Left")
	log.Printf("Current Orientation: %s", robot.CurrOrientation)

	// Change the orientation depending on the current orientation of the robot
	switch robot.CurrOrientation {
	case "N":
		robot.CurrOrientation = "W"
	case "W":
		robot.CurrOrientation = "S"
	case "S":
		robot.CurrOrientation = "E"
	case "E":
		robot.CurrOrientation = "N"
	}

	log.Printf("New Orientation: %s", robot.CurrOrientation)
}

// This function makes the robot spin 90 degrees to the right without moving from its current spot
func (robot *Robot) TurnRight() {
	// Logs all robot's moves and navigations
	log.Println("----------------")
	log.Println("Turning Right")
	log.Printf("Current Orientation: %s", robot.CurrOrientation)

	// Change the orientation depending on the current orientation of the robot
	switch robot.CurrOrientation {
	case "N":
		robot.CurrOrientation = "E"
	case "W":
		robot.CurrOrientation = "N"
	case "S":
		robot.CurrOrientation = "W"
	case "E":
		robot.CurrOrientation = "S"
	}
	log.Printf("New Orientation: %s", robot.CurrOrientation)
}

/*
 This function moves robot forward one grid point and maintain the same heading orientation.
 Arguments:
	arena *Arena - arena for robot wars
 Return:
 	error - any error that occurred during the function call
*/
func (robot *Robot) MoveForward(arena *Arena) error {
	// Logs all robot's moves and navigation
	log.Println("----------------")
	log.Println("Moving forward")
	log.Printf("Current Position %v & Orientation: %s", robot.CurrPosition, robot.CurrOrientation)

	// Depending on the current orientation move the robot forward and update its current position
	switch robot.CurrOrientation {
	case "N":
		// If the new position is inside the arena move robot forward
		if robot.CurrPosition.Y < arena.TopRight.Y {
			robot.CurrPosition.Y += 1
		} else {
			return errors.New("Robot cannot move outside the arena! Stopping the robot!")
		}

	case "W":
		// If the new position is inside the arena move robot forward
		if robot.CurrPosition.X > arena.BottomLeft.X {
			robot.CurrPosition.X -= 1
		} else {
			return errors.New("Robot cannot move outside the arena! Stopping the robot!")
		}

	case "S":
		// If the new position is inside the arena move robot forward
		if robot.CurrPosition.Y > arena.BottomLeft.Y {
			robot.CurrPosition.Y -= 1
		} else {
			return errors.New("Robot cannot move outside the arena! Stopping the robot!")
		}
	case "E":
		// If the new position is inside the arena move robot forward
		if robot.CurrPosition.X < arena.TopRight.X {
			robot.CurrPosition.X += 1
		} else {
			return errors.New("Robot cannot move outside the arena! Stopping the robot!")
		}
	}
	log.Printf("New Position %v & Orientation: %s", robot.CurrPosition, robot.CurrOrientation)
	return nil
}
