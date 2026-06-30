package crud

import (
	"encoding/json"

	"github.com/Mezali/go-task-tracker/internal/models"
)

var reModel []models.Task

// TODO: implement a way to know if there is a index missing,
// like if a user delete the item of Id 4 while the list is greater than 4
func IndexJson(File []byte) uint {
	var index uint

	json.Unmarshal(File, &reModel)

	index = uint(len(reModel))

	return index + 1
}

func List(File []byte) []string {

	json.Unmarshal(File, &reModel)

	// TODO: make this a map
	var List []string

	for _, value := range reModel {
		List = append(List, value.Description)
	}

	return List
}
