RobotWars
==========

This project provides an API which can be used for “Robot Wars” competition.
The API supports to move and navigate the robots in the given arena.
Given an input file in a specific format, it processes the instructions for defining the boundary of the arena,
initial positions and orientations of the robots, and the instructions telling the robots how to move in the arena.

As output, it should produce a JSON response indicating the final position and orientation of the robots.

Input format
=============
```
5 5             # top-right corner of the arena
1 2 N           # initial position of robot-1 and its orientation (N:North / S:South / W:West / E:East)
LMLMLMLMM       # instructions to move the robot (L:Left / R:Right / M:Move forward)
3 3 E           # initial position of robot-2 and its orientation (N:North / S:South / W:West / E:East)
MMRMMRMRRM      # instructions to move the robot (L:Left / R:Right / M:Move forward)
....
```

Output format
==============
```
[
    {
        "id": 0,
        "position": {
            "X": 1,
            "Y": 4
        },
        "orientation" : "N"
    },
    {
        ...
    }
]
```
Implementation Details
=======================

**Programming Language**: GoLang

**Version**: go 1.14 (refer to go.mod file)

**Requirements**:

    github.com/stretchr/testify v1.6.1 (refer to go.mod file)
    
    github.com/gookit/color v1.2.7 (refer to go.mod file)


Instructions to Run the System
================================

To run the code:

1) Go to `robotWars` folder on the command prompt

2) Build the `main.go` file
```
    $ go build main.go
```
3) Run the `main.go` file
```
   $ ./main
```

Instructions for System Tests
=============================

**Option 1:** Go to each package and individually run the tests

1) Go to `robotWars/models/` folder and run the test 
```
   $ go test
```
3) Go to `robotWars/backend/` folder and run the test
```
   $ go test
```
**Option 2:** Run all the tests from `robotWars` folder
```
   $ go test ./...
```
