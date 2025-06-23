package model

import "task/pkg/model/types"

type Task struct {
	ID     string           `json:"id"`
	Name   string           `json:"name"`
	Status types.TaskStatus `json:"status"`
}
