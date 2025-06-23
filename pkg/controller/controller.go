package controller

import "net/http"

type Handler func(w http.ResponseWriter, r *http.Request)

type Controller interface {
	RegisterHandler(path string, handler Handler)
	HandleRequest(w http.ResponseWriter, r *http.Request)
}

type taskController struct {
	handlers map[string]Handler
}

func NewTaskController() *taskController {
	c := &taskController{
		handlers: make(map[string]Handler),
	}

	c.RegisterHandler(http.MethodGet, GetTasksHandler)
	c.RegisterHandler(http.MethodPost, CreateTaskHandler)
	c.RegisterHandler(http.MethodPut, UpdateTaskHandler)
	c.RegisterHandler(http.MethodDelete, DeleteTaskHandler)

	return c

}

func (c *taskController) RegisterHandler(path string, handler Handler) {
	c.handlers[path] = handler
}

func (c *taskController) HandleRequest(w http.ResponseWriter, r *http.Request) {
	handler, exists := c.handlers[r.Method]
	if !exists {
		http.NotFound(w, r)
		return
	}

	handler(w, r)
}
