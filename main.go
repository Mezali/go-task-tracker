package main

import (
	"fmt"
	// "log"
	"os"
)

func main() {
	jsonFileName := "tasks.json"

	// checking if user passes an argument before start
	if len(os.Args) == 1 {
		fmt.Fprintln(os.Stderr, "Task Tracker Program.\nPlease provide an argument.")
		os.Exit(64) // code 64 is for invalid or no argument.
	}

	// check if json file exists, if not, create one
	_, error := os.Stat(jsonFileName) // find the damm file
	if error != nil {
	    fmt.Println("Database not found...\nCreating one!")
		os.Create(jsonFileName)
	}

	os.Exit(0)
}
