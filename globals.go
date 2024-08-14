package main

import sqlhelper "github.com/misterunix/sqlite-helper"

type therobots struct {
	ID           int     // ID
	OwnerID      int     // Owner ID
	Filename     string  // Filename of the robot
	FilenameHash string  // Filename hash
	Code         string  // Code of the robot
	CodeHash     string  // Code hash
	Count        int     // Number of times this robot has competed
	Points       float64 // Points scored by this robot
	Win          int     // Number of wins
	Tie          int     // Number of ties
	Loss         int     // Number of losses
}

type gameinfo struct {
	BBprogram    string      `json:"bbprogram"`  // full path to the basicbots program
	Robotspath   string      `json:"robotspath"` // full path to the robots directory
	DBpath       string      `json:"dbpath"`     // full path to the database
	Userspath    string      `json:"userspath"`  // full path to the users directory
	NumOfBots    int         `json:"numofbots"`  // number of bots
	RobotStorage []therobots `json:"robotstorage"`
}

// type Challenge struct {
// 	RobotsName []string `json:"robotname"` // List of robots
// 	Wins       []int    `json:"wins"`      // List of wins
// 	Ties       []int    `json:"ties"`      // List of ties
// 	Losses     []int    `json:"losses"`    // List of losses
// }

// type Robot struct {
// 	Filename string    `json:"filename"` // Filename of the robot
// 	Count    int       `json:"count"`    // Number of times this robot has competed
// 	Points   int       `json:"points"`   // Points scored by this robot
// 	Win      int       `json:"win"`      // Number of wins
// 	Tie      int       `json:"tie"`      // Number of ties
// 	Loss     int       `json:"loss"`     // Number of losses
// 	Battles  Challenge `json:"battles"`  // List of battles this robot has competed agaist other robots
// }

var db *sqlhelper.DbConfig

var GameInfo gameinfo

type match2 struct {
	ID      int
	Robot1  int
	Robot2  int
	Win1    int
	Win2    int
	Tie1    int
	Tie2    int
	Lose1   int
	Lose2   int
	Points1 float64
	Points2 float64
}

type match3 struct {
	ID      int
	Robot1  int
	Robot2  int
	Robot3  int
	Win1    int
	Win2    int
	Win3    int
	Tie1    int
	Tie2    int
	Tie3    int
	Lose1   int
	Lose2   int
	Lose3   int
	Points1 float64
	Points2 float64
	Points3 float64
}

type match4 struct {
	ID      int
	Robot1  int
	Robot2  int
	Robot3  int
	Robot4  int
	Win1    int
	Win2    int
	Win3    int
	Win4    int
	Tie1    int
	Tie2    int
	Tie3    int
	Tie4    int
	Lose1   int
	Lose2   int
	Lose3   int
	Lose4   int
	Points1 float64
	Points2 float64
	Points3 float64
	Points4 float64
}
