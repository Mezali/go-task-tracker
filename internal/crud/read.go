package crud

import (
	"encoding/json"

	"github.com/Mezali/go-task-tracker/internal/models"
)

var task []models.Task

// TODO: implement a way to know if there is a index missing,
// like if a user delete the item of Id 4 while the list is greater than 4
func IndexJson(File []byte) uint {
	var index uint

	json.Unmarshal(File, &task)

	index = uint(len(task))

	return index + 1
}

func List(File []byte) []string {

	json.Unmarshal(File, &task)

	// TODO: make this a map
	var List []string
	
	for _, value := range task {
		List = append(List, value.Description)
	}
	
	return List
}

