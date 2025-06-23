package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"task/pkg/model"
	"task/pkg/model/types"
	"task/pkg/repo"
)

func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks := repo.GetTasks()

	err := respondToJSON(w, http.StatusOK, tasks)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to respond: %v", err), http.StatusInternalServerError)
	}
}

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task model.Task
	body, _ := io.ReadAll(r.Body)

	defer r.Body.Close()
	err := json.Unmarshal(body, &task)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)

		return
	}

	if task.Status == types.TaskStatusUnSpecified {
		http.Error(w, "task status is invalid", http.StatusBadRequest)

		return
	}

	created := repo.CreateTask(task)

	if err := respondToJSON(w, http.StatusCreated, created); err != nil {
		http.Error(w, fmt.Sprintf("failed to respond: %v", err), http.StatusInternalServerError)
	}
}

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/tasks/")
	var task model.Task

	body, _ := io.ReadAll(r.Body)
	defer r.Body.Close()

	err := json.Unmarshal(body, &task)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)

		return
	}

	currentTask, err := repo.GetOneTask(id)
	if err != nil {
		if errors.Is(err, repo.ErrTaskNotFound) {
			http.Error(w, "task not found", http.StatusNotFound)

			return
		}
		http.Error(w, "failed to get task", http.StatusInternalServerError)

		return
	}

	if !currentTask.Status.CanTransitionTo(task.Status) {
		http.Error(w, "task status cannot transition", http.StatusBadRequest)

		return
	}

	updated, err := repo.UpdateTask(id, task)
	if err != nil {
		if errors.Is(err, repo.ErrTaskNotFound) {
			http.Error(w, "task not found", http.StatusNotFound)

			return
		}
		http.Error(w, "failed to update task", http.StatusInternalServerError)

		return
	}

	if err := respondToJSON(w, http.StatusOK, updated); err != nil {
		http.Error(w, fmt.Sprintf("failed to respond: %v", err), http.StatusInternalServerError)
	}
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/tasks/")

	if err := repo.DeleteTask(id); err != nil {
		if errors.Is(err, repo.ErrTaskNotFound) {
			http.Error(w, "task not found", http.StatusNotFound)

			return
		}
		http.Error(w, "failed to delete task", http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func respondToJSON(w http.ResponseWriter, code int, payload any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(payload)
	if err != nil {
		return fmt.Errorf("failed to encode response: %w", err)
	}

	return nil
}
