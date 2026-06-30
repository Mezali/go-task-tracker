package crud

import (
	"encoding/json"

	"github.com/Mezali/go-task-tracker/internal/models"
)

var deModel []models.Task

func DeleteTask(File []byte, TaskId uint) []byte {
	json.Unmarshal(File, &deModel)

	for index, value := range deModel {
		if value.Id == TaskId {
			deModel = append(deModel[:index], deModel[index+1:]...)
			break
		}
	}
	deModel, _ := json.MarshalIndent(deModel, "", "  ")
	
	return deModel
}
