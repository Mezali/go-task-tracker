package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type Task struct {
	Id          uint      `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreateAt    time.Time `json:"createAt"`
	UpdateAt    time.Time `json:"UpdateAt"`
}

var JsonFileName string = "tasks.json"

// Run all validations and checks before starting
func init() {

	// check if user provides an Argument
	if len(os.Args) == 1 {
		fmt.Fprintln(os.Stderr, "Task Tracker Program.\nPlease provide an argument.")
		os.Exit(64) // code 64 is for invalid or no argument.
	}

	_, error := os.Stat(JsonFileName) // find the damm file
	if error != nil {
		fmt.Println("Database not found...\nPlease run: task-track init")
		os.Exit(0)
	}

	// check if the json file is valid
	file, error := os.ReadFile(JsonFileName)
	if error != nil {
		log.Fatalf("Error reading the file: %v", error)
	}

	if !json.Valid(file) {
		fmt.Println("Json invalid...\nCreating new one.")
		os.WriteFile(JsonFileName, []byte("[]"), 0644)
	}
}

func main() {

	

	os.Exit(0)
}
