package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/Mezali/go-task-tracker/internal/commands"
	"github.com/Mezali/go-task-tracker/internal/crud"
	"github.com/Mezali/go-task-tracker/internal/models"
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

var JsonFileName string = "tasks.json"
var Arguments []string
var TaskFile []models.Task

// Run all validations and checks before starting
func init() {

	// check if user provides an Argument
	if len(os.Args) == 1 {
		printUsage()
		os.Exit(0)
	}

	Arguments = os.Args

	_, error := os.Stat(JsonFileName) // find the damm file
	if error != nil && Arguments[1] != "init" {
		fmt.Println("Database not found...\nPlease run: task-tracker init")
		os.Exit(0)
	}

	file, error := os.ReadFile(JsonFileName)
	if error != nil && os.Args[1] != "init" {
		log.Fatalf("Error reading the file: %v", error)
	}

	// Checks if the json is valid
	if !json.Valid(file) && os.Args[1] != "init" {
		fmt.Println("Invalid Json...\nPlease create a new one by running: task-tracker init")
		os.Remove(JsonFileName)
	}
}

func main() {

	// Open the json file
	File, err := os.ReadFile(JsonFileName)

	if err != nil && Arguments[1] != "init" {
		log.Fatalf("Error opening the file: %v", err)
	}

	json.Unmarshal(File, &TaskFile)

	switch Arguments[1] {
	case "init": // Init the database

		commands.InitDb(JsonFileName)

	case "add": // Create a task

		// We need first to read all the database to index the new task
		index := crud.IndexJson(File) // DONE

		// Now we need to create the task
		File = crud.CreateTask(TaskFile, Arguments[2], uint(index))

		// Write in the file
		os.WriteFile(JsonFileName, File, 0644)

	case "list": // List or Search for a task
		List := crud.List(File)
		for index, value := range List {
			fmt.Printf("%v. %v\n", index+1, value)
		}

	case "update":
		

	case "delete":
		// os.Args returns a String, so convert it into a int
		delIndex, _ := strconv.Atoi(Arguments[2])
		// Delete the task in the struct
		File = crud.DeleteTask(File, uint(delIndex))
		// Write to file
		os.WriteFile(JsonFileName, File, 0664)
	default:
		printUsage()
	}

	os.Exit(0)
}
