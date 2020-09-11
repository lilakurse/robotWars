package models

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
Test the function for turning the robot left
*/
func TestRobot_TurnLeft(t *testing.T) {
	actualRobot := Robot{1, Coordinates{1, 3}, "N",
		Coordinates{1, 3}, "N"}
	actualRobot.TurnLeft()

	expectedRobot := Robot{1, Coordinates{1, 3}, "N",
		Coordinates{1, 3}, "W"}

	assert.Equal(t, actualRobot, expectedRobot, "Expected outputs are not equal")

}

/*
Test the function for turning the robot right
*/
func TestRobot_TurnRight(t *testing.T) {
	actualRobot := Robot{1, Coordinates{1, 3}, "N",
		Coordinates{1, 3}, "S"}
	actualRobot.TurnRight()

	expectedRobot := Robot{1, Coordinates{1, 3}, "N",
		Coordinates{1, 3}, "W"}

	assert.Equal(t, actualRobot, expectedRobot, "Expected outputs are not equal")
}

/*
Test the function for moving the robot forward
*/
func TestRobot_MoveForward(t *testing.T) {
	arena := &Arena{Coordinates{5,5}, Coordinates{0,0}}
	actualRobot := &Robot{1, Coordinates{1, 2}, "N",
		 Coordinates{1, 2}, "S"}

	actualRobot.MoveForward(arena)
	expectedRobot := &Robot{1, Coordinates{1, 2}, "N",
			Coordinates{1, 1}, "S"}

	assert.Equal(t, actualRobot, expectedRobot, "Expected outputs are not equal")

	// Test case when robot is moving outside the arena
	actualRobot.CurrPosition = Coordinates{1,0}

	// Create the error that we expect from MoveForward function
	expectedErr := errors.New("Robot cannot move outside the arena! Stopping the robot!")
	actualError := actualRobot.MoveForward(arena)
	// Assert that we get an error while moving the robot outside the arena
	assert.Equal(t, expectedErr, actualError, "Expected outputs are not equal")

}