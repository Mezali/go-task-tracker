package models

import "time"

type Task struct {
	Id           uint      `json:"id"`
	Description  string    `json:"description"`
	IsDone       bool      `json:"isDone"`
	IsInProgress bool      `json:"isInProgress"`
	CreateAt     time.Time `json:"createAt"`
	UpdateAt     time.Time `json:"updateAt"`
}
