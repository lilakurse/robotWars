package backend

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gookit/color"
	"robotWars/models"
	"strconv"
	"strings"
)

/*
 This function initializes the arena struct and its variables.
 Arguments:
 	bottomLeftX int - bottom left x of arena
 	bottomLeftY int - bottom left y of arena
 	topRightX int - top right x of arena
 	topRightY int -top right y of arena
 Return:
 	arena *models.Arena - reference to arena struct
*/
func InitializeArena(bottomLeftX int, bottomLeftY int, topRightX int, topRightY int) *models.Arena {
	arena := &models.Arena{}
	arena.BottomLeft.X = bottomLeftX
	arena.BottomLeft.Y = bottomLeftY
	arena.TopRight.X = topRightX
	arena.TopRight.Y = topRightY

	return arena
}

/*
 This function checks if the arena for robot wars can be created or not. If yes, it initializes the arena.
 Arguments:
 	topRightCoordinatesArena string - arenas top right coordinates in the form of "x y"
 Return:
 	arena *models.Arena - reference to arena struct
	error - any error that occurred during the function call
*/

func CreateArena(topRightCoordinatesArena string) (*models.Arena, error) {

	// Splits the string around each instance of an consecutive white space characters
	topRightPos := strings.Fields(topRightCoordinatesArena)

	// Arena cannot be created if split returned empty slice or one of the coordinates are missing.
	if len(topRightPos) != 2 {
		return nil, errors.New("Instruction invalid! Arena cannot be created!")
	}


	// Convert string topRightX to int
	topRightX, err := strconv.Atoi(topRightPos[0])
	if err != nil {
		return nil, err
	}
	// Convert string topRightY to int
	topRightY, err := strconv.Atoi(topRightPos[1])
	if err != nil {
		return nil, err
	}

	// Bottom left coordinates are assumed to be (0, 0) as given in the problem statement
	bottomLeftX := 0
	bottomLeftY := 0

	// Initialize the arena and return
	return InitializeArena(bottomLeftX, bottomLeftY, topRightX, topRightY), nil
}

/*
 This function initialize the robot struct and its variables.
 Arguments:
 	id int - robot ID
 	startX int - start X coordinate of the robot
	startY int - start Y coordinate of the robot
	orientation string - start orientation of the robot (N/S/W/E)
 Return:
 	robot *models.Robot - reference to robot struct
*/
func InitializeRobot(id int, startX int, startY int, orientation string) *models.Robot {
	robot := &models.Robot{}
	robot.Id = id
	robot.StartPos.X = startX
	robot.StartPos.Y = startY
	robot.StartOrientation = orientation
	robot.CurrPosition.X = startX
	robot.CurrPosition.Y = startY
	robot.CurrOrientation = orientation

	return robot
}

/*
 This method creates the robot on the war arena with the given initial position and orientation.
 Arguments:
	arena *models.Arena - arena for robot wars
	id int  -  robot ID
	line string - robot's initial position and orientation in the form of "X Y O"
 Return:
 	*models.Robot - reference to robot struct
	error - any error that occurred during the function call
*/
func CreateRobot(arena *models.Arena, id int, line string) (*models.Robot, error) {
	// Splits the string around each instance of an consecutive white space characters
	location := strings.Fields(line)

	// Robot cannot be created if split returned empty slice or one of the robots coordinates are missing.
	if len(location) != 3 {
		return nil, errors.New("Instruction invalid! Robot cannot be created!")
	}
	// Convert string posX to int
	posX, err := strconv.Atoi(location[0])
	if err != nil {
		return nil, err
	}
	// Convert string posY to int
	posY, err := strconv.Atoi(location[1])
	if err != nil {
		return nil, err
	}

	orientation := location[2]

	// Robot cannot be created if the initial position is outside the arena
	if posX < arena.BottomLeft.X || posY < arena.BottomLeft.Y || posX > arena.TopRight.X ||
		posY > arena.TopRight.Y {

		return nil, errors.New("Robot cannot be created outside the arena")
	}

	// Robot cannot be created if the initial orientation is not valid (N/S/W/E)
	if orientation != "N" && orientation != "S" && orientation != "E" && orientation != "W" {
		return nil, errors.New("Robot cannot be created without a valid orientation")
	}

	// Initialize the robot struct and return
	return InitializeRobot(id, posX, posY, orientation), nil
}

/*
 This function processes the series of instructions telling robot how to move within the arena.
 Arguments:
	arena *models.Arena - arena for robot wars
	robot *models.Robot - robot
	moves string - robots moves
 Return:
	error - any error that occurred during the function call
*/
func ProcessRobotMoves(arena *models.Arena, robot *models.Robot, moves string) error {
	// Iterate over all move instructions
	for _, currMove := range moves {
		switch currMove {
		// If the instruction is to turn left ('L'), call the function TurnLeft
		case 'L':
			robot.TurnLeft()
		// If the instruction is to turn right ('R'), call the function TurnRight
		case 'R':
			robot.TurnRight()
		// If the instruction is to move forward('M'), call the function MoveForward
		case 'M':
			err := robot.MoveForward(arena)
			if err != nil {
				return err
			}
		default:
			return fmt.Errorf("Invalid robot move: %c! Stopping the robot!", currMove)
		}
	}

	return nil
}

/*
 This function generates a JSON response containing the final position and orientation.
 Arguments:
	robots []*models.Robot - slice of robots struct
 Return:
 	output []byte - JSON
*/
func GenerateFinalPositionsJSON(robots []*models.Robot) []byte {
	// given a list of robots, generate a JSON response which contains their final positions and orientations
	output, err := json.MarshalIndent(robots, "", "    ")
	if err != nil {
		color.Error.Println(err)
		return nil
	}
	return output
}
