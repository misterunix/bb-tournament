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
	BBprogram  string `json:"bbprogram"`  // full path to the basicbots program
	Robotspath string `json:"robotspath"` // full path to the robots directory
	DBpath     string `json:"dbpath"`     // full path to the database
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
