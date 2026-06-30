package crud

import (
	"encoding/json"
	"time"

	"github.com/Mezali/go-task-tracker/internal/models"
)

var upModel []models.Task

func Update(File []byte, TaskId uint, description string) []byte {
	json.Unmarshal(File, &upModel)

	for index, value := range upModel {
		if value.Id == TaskId {
			upModel[index].Description = description
		}
	}

	upModel, _ := json.MarshalIndent(upModel, "", "  ")
	return upModel
}

func MarkDone(File []byte, TaskId uint, isOrNot bool) []byte {
	json.Unmarshal(File, &upModel)

	for index, value := range upModel {
		if value.Id == TaskId {
			upModel[index].IsDone = isOrNot
			upModel[index].UpdateAt = time.Now()
		}
	}

	upModel, _ := json.MarshalIndent(upModel, "", "  ")
	return upModel
}

func MarkProgress(File []byte, TaskId uint, isOrNot bool) []byte {
	json.Unmarshal(File, &upModel)

	for index, value := range upModel {
		if value.Id == TaskId {
			upModel[index].IsInProgress = isOrNot
			upModel[index].UpdateAt = time.Now()
		}
	}

	upModel, _ := json.MarshalIndent(upModel, "", "  ")
	return upModel
}
