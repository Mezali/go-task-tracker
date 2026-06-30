package crud

import (
	"encoding/json"

	"github.com/Mezali/go-task-tracker/internal/models"
)

var model []models.Task

func DeleteTask(File []byte, TaskId uint) []byte {
	json.Unmarshal(File, &model)

	for index, value := range model {
		if value.Id == TaskId {
			model = append(model[:index], model[index+1:]...)
			break
		}
	}
	model, _ := json.MarshalIndent(model, "", "  ")
	
	return model
}
