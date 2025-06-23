package types

import (
	"encoding/json"
	"fmt"
)

type TaskStatus int

const (
	TaskStatusUnSpecified TaskStatus = -1
	TaskStatusInCompleted TaskStatus = 0
	TaskStatusCompleted   TaskStatus = 1
)

var taskStatusTransitions = map[TaskStatus][]TaskStatus{
	TaskStatusInCompleted: {TaskStatusCompleted},
	TaskStatusCompleted:   {TaskStatusInCompleted},
}

func (s TaskStatus) String() string {
	switch s {
	case TaskStatusInCompleted:
		return "InCompleted"
	case TaskStatusCompleted:
		return "Completed"
	default:
		return "UnSpecified"
	}
}

func (s TaskStatus) In(status ...TaskStatus) bool {
	for _, v := range status {
		if s == v {
			return true
		}
	}

	return false
}

func MakeTaskStatus(status int) TaskStatus {
	switch status {
	case 0:
		return TaskStatusInCompleted
	case 1:
		return TaskStatusCompleted
	default:
		return TaskStatusUnSpecified
	}
}

func (s TaskStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s *TaskStatus) UnmarshalJSON(data []byte) error {
	var raw interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	switch val := raw.(type) {
	case float64:
		*s = MakeTaskStatus(int(val))
	case string:
		switch val {
		case "Completed":
			*s = TaskStatusCompleted
		case "InCompleted":
			*s = TaskStatusInCompleted
		default:
			*s = TaskStatusUnSpecified
		}
	default:
		return fmt.Errorf("invalid task status type: %T", val)
	}

	return nil
}

func (s TaskStatus) CanTransitionTo(next TaskStatus) bool {
	allowed, ok := taskStatusTransitions[s]
	if !ok {
		return false
	}

	for _, to := range allowed {
		if next == to {
			return true
		}
	}
	return false
}
