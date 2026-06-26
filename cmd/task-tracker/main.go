package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Mezali/go-task-tracker/internal/commands"
)

func printUsage() {
	usage := `Usage: task-tracker <COMMAND> [ARGUMENT]

Commands:
  init                         		 Initialize the json Database
  add <description>		    	 Add a new task
  list <status>          		 List all tasks and filter
  update <id> [description]		 Update a description
  delete <id>              		 Delete a task
  mark-done <id>         		 Mark a task as done
  mark-in-progress <id>			 Mark a task as in progress
`
	fmt.Print(usage)
}

type Task struct {
	Id           uint      `json:"id"`
	Description  string    `json:"description"`
	Status       string    `json:"status"`
	IsDone       bool      `json:"isDone"`
	IsInProgress bool      `json:"isInProgress"`
	CreateAt     time.Time `json:"createAt"`
	UpdateAt     time.Time `json:"UpdateAt"`
}

var JsonFileName string = "tasks.json"
var Arguments []string

// Run all validations and checks before starting
func init() {

	// check if user provides an Argument
	if len(os.Args) == 1 {
		printUsage()
		os.Exit(0)
	}

	Arguments = os.Args

	_, error := os.Stat(JsonFileName) // find the damm file
	if error != nil && os.Args[1] != "init" {
		fmt.Println("Database not found...\nPlease run:\ntask-tracker init")
		os.Exit(0)
	}

	// check if the json file is valid
	file, error := os.ReadFile(JsonFileName)
	if error != nil && os.Args[1] != "init" {
		log.Fatalf("Error reading the file: %v", error)
	}

	// Checks if the json is valid
	if !json.Valid(file) && os.Args[1] != "init" {
		fmt.Println("Invalid Json...\nPlease create a new one by running:\ntask-tracker init")
		os.Remove(JsonFileName)
	}
}

func main() {

	switch Arguments[1] {
	case "init":
		commands.InitDb(JsonFileName)
	 case "add":
		fmt.Println("teste")

	default:
		printUsage()
	}

	os.Exit(0)
}
