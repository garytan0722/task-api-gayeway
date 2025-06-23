package router

import (
	"net/http"
	"task/pkg/controller"
	"task/pkg/middleware"
)

const ()

func SetupRouter() http.Handler {
	mux := http.NewServeMux()

	taskController := controller.NewTaskController()

	mux.HandleFunc("/tasks", middleware.Chain(taskController.HandleRequest, middleware.RequestLogger, middleware.Auth))

	mux.HandleFunc("/tasks/", middleware.Chain(taskController.HandleRequest, middleware.RequestLogger, middleware.Auth))

	return mux
}
