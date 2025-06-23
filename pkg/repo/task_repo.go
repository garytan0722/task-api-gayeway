package repo

import (
	"sync"
	"task/pkg/model"

	"github.com/google/uuid"
)

var (
	tasks = make(map[string]model.Task)
	mu    sync.RWMutex
)

func GetOneTask(id string) (model.Task, error) {
	mu.RLock()
	defer mu.RUnlock()

	task, exists := tasks[id]
	if !exists {
		return model.Task{}, ErrTaskNotFound
	}

	return task, nil
}

func GetTasks() []model.Task {
	mu.RLock()
	defer mu.RUnlock()

	result := []model.Task{}
	for _, task := range tasks {
		result = append(result, task)
	}

	return result
}

func CreateTask(task model.Task) model.Task {
	mu.Lock()
	defer mu.Unlock()

	task.ID = uuid.New().String()
	tasks[task.ID] = task

	return task
}

func UpdateTask(id string, task model.Task) (model.Task, error) {
	mu.Lock()
	defer mu.Unlock()

	if _, exists := tasks[id]; !exists {
		return model.Task{}, ErrTaskNotFound
	}

	task.ID = id
	tasks[id] = task

	return task, nil
}

func DeleteTask(id string) error {
	mu.Lock()
	defer mu.Unlock()

	if _, exists := tasks[id]; !exists {
		return ErrTaskNotFound
	}

	delete(tasks, id)

	return nil
}
