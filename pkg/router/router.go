package router

import (
	"net/http"
	"task/pkg/controller"
	"task/pkg/middleware"
)

func SetupRouter() http.Handler {
	mux := http.NewServeMux()

	taskController := controller.NewTaskController()

	mux.HandleFunc(TasksPath, middleware.Chain(taskController.HandleRequest, middleware.RequestLogger, middleware.Auth))

	mux.HandleFunc(TaskByIDPath, middleware.Chain(taskController.HandleRequest, middleware.RequestLogger, middleware.Auth))

	return mux
}
