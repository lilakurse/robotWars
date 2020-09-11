package backend

import (
	"github.com/stretchr/testify/assert"
	"robotWars/models"
	"testing"
)

/*
Test the function for creating the robots war arena
*/
func TestCreateArena(t *testing.T) {
	var tests = []struct {
		input    string
		expected *models.Arena
	}{
		// If the input string is empty
		{
			input:    "",
			expected: nil,
		},
		// If the input string is valid
		{
			input:    "3 3",
			expected: &models.Arena{TopRight: models.Coordinates{3, 3}, BottomLeft: models.Coordinates{0, 0}},
		},
		{
			input:    "5 5",
			expected: &models.Arena{TopRight: models.Coordinates{5, 5}, BottomLeft: models.Coordinates{0, 0}},
		},
		// If the input string is not in a proper format
		{
			input:    "M 3",
			expected: nil,
		},
		// If the input string is missing any expected elements
		{
			input:    "2",
			expected: nil,
		},
	}

	for _, test := range tests {
		arena, _ := CreateArena(test.input)
		assert.Equal(t, test.expected, arena, "Expected outputs are not equal")
	}
}

/*
Test the function for initializing the robots war arena
*/
func TestInitializeArena(t *testing.T) {
	var tests = []struct {
		inputBottomRightX int
		inputBottomRightY int
		inputTopRightX    int
		inputTopRightY    int

		expected *models.Arena
	}{
		// valid input
		{0, 0, 0, 0,
			&models.Arena{TopRight: models.Coordinates{0, 0}, BottomLeft: models.Coordinates{0, 0}},
		},
		// valid input
		{0, 0, 5, 5,
			&models.Arena{TopRight: models.Coordinates{5, 5}, BottomLeft: models.Coordinates{0, 0}},
		},
	}
	for _, test := range tests {
		assert.Equal(t, InitializeArena(test.inputBottomRightX, test.inputBottomRightY,
			test.inputTopRightY, test.inputTopRightX), test.expected, "Expected outputs are not equal")
	}
}


/*
Test the function for creating the robot on the arena
*/
func TestCreateRobot(t *testing.T) {
	arena := &models.Arena{TopRight: models.Coordinates{5, 5}, BottomLeft: models.Coordinates{0, 0}}

	var tests = []struct {
		id       int
		line     string
		expected *models.Robot
	}{
		// valid initial position of the robot
		{1, "1 2 N", &models.Robot{1, models.Coordinates{1, 2}, "N",
			models.Coordinates{1, 2}, "N"}},
		{2, "1 6 N", nil}, // Initial position outside the arena
		{3, "6 N", nil}, // Missing element in the input
		{4, "N 6 N", nil}, // Input not in the expected form
		{5, "1 6 G", nil}, // Initial orientation is not valid
	}

	for _, test := range tests {
		robot, _ := CreateRobot(arena, test.id, test.line)
		assert.Equal(t, test.expected, robot, "Expected outputs are not equal")

	}
}

/*
Test the function for initializing the robot on the arena
*/
func TestInitializeRobot(t *testing.T) {
	var test = struct {
		id          int
		startX      int
		startY      int
		orientation string
		expected    *models.Robot
	}{
		1, 1, 3, "N", &models.Robot{1, models.Coordinates{1, 3}, "N",
			models.Coordinates{1, 3}, "N"},
	}
	result := InitializeRobot(test.id, test.startX, test.startY, test.orientation)
	assert.Equal(t, test.expected, result, "Expected outputs are not equal")

}
