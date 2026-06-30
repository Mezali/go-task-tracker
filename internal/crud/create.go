package crud

import (
	"encoding/json"
	"time"

	"github.com/Mezali/go-task-tracker/internal/models"
)

// This function will take a task string and add it to the json database
func CreateTask(TaskFile []models.Task, NewTask string, Index uint) []byte {
	var modelTask models.Task
	modelTask.Id = Index
	modelTask.Description = NewTask
	modelTask.CreateAt = time.Now()

	TaskFile = append(TaskFile, modelTask)

	file, _ := json.MarshalIndent(TaskFile, "", "  ")
	return file
}
