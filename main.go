package main

import (
	"encoding/json"
	"fmt"
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

func main() {
	// jsonFileName := "tasks.json"

	task := Task{
		Id:          1,
		Description: "Demo",
		Status:      "Done",
		CreateAt:    time.Now(),
		UpdateAt:    time.Now(),
	}

	jsonData, _ := json.Marshal(task)
	fmt.Println(string(jsonData))

	// // checking if user passes an argument before start
	// if len(os.Args) == 1 {
	// 	fmt.Fprintln(os.Stderr, "Task Tracker Program.\nPlease provide an argument.")
	// 	os.Exit(64) // code 64 is for invalid or no argument.
	// }

	// // check if json file exists, if not, create one
	// _, error := os.Stat(jsonFileName) // find the damm file
	// if error != nil {
	// 	fmt.Println("Database not found...\nCreating one!")
	// 	os.Create(jsonFileName)
	// }

	os.Exit(0)
}
