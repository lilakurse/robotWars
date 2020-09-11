package main

import (
	"bufio"
	"fmt"
	"github.com/gookit/color"
	"log"
	"os"
	"robotWars/backend"
	"robotWars/models"
)

/*
	This is the main function for the project. It interacts with the backend to move and navigate the robot,
	and generate an output, as a JSON response containing the final position and orientation of the robot.
*/

func main() {
	ipFilePath := "robots.txt"
	ipFile, err := os.Open(ipFilePath)
	if err != nil {
		color.Error.Println("Error in opening the input file!")
		log.Fatal(err)
	}
	defer ipFile.Close()
	scanner := bufio.NewScanner(ipFile)

	// Read the first line
	scanner.Scan()
	arenaCoordinates := scanner.Text()
	scanner.Scan()
	// Call the API function from backend to create an arena
	arena, err := backend.CreateArena(arenaCoordinates)

	// If arena was not created, the competition cannot continue
	// Log the error and exit
	if err != nil {
		color.Error.Println("Arena cannot be created! Exiting!")
		log.Fatal(err)
	}

	robotId := 0
	// Stores all robots
	var robots []*models.Robot

	for scanner.Scan() {
		// Call the API function from backend to create a robot
		robot, err := backend.CreateRobot(arena, robotId, scanner.Text())

		// If robot was not created, read next line, that contains moves for this robot,
		// and move to the next robot
		if err != nil {
			color.Error.Println(err)
			scanner.Scan()
			continue
		}
		robots = append(robots, robot)

		// Logs all robot's moves and navigations
		log.Println("+++++++++++++++++++++++++++++++++++++++++++")
		log.Printf("Robot %d: Position %v & Orientation: %s", robot.Id, robot.CurrPosition,
			robot.CurrOrientation)
		robotId++

		// Read next line
		scanner.Scan()
		// Call API function from backend to move the robot
		err = backend.ProcessRobotMoves(arena, robot, scanner.Text())

		// If the robot cannot move, show error and continue
		if err != nil {
			color.Error.Println(err)
			continue
		}
	}

	// Call API function from backend to generate final JSON output
	output := backend.GenerateFinalPositionsJSON(robots)
	fmt.Printf("RobotWars: %s\n", output)

}
