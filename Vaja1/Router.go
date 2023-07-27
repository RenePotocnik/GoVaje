package main

import (
	"github.com/gin-gonic/gin"
	"vaja1/API"
)

// Create a Router struct and bind functions to it
type Router struct {
	engine *gin.Engine
	api    API.Controller
}

func (r *Router) registerRoutes() (err error) {
	// Path to `/api/v1`
	api := r.engine.Group("/api/v1")

	todo := api.Group("/todo")
	r.registerTodoRoutes(todo)

	login := api.Group("/login")
	logout := api.Group("/logout")
	r.registerUserRoutes(login, logout)

	return
}

// registerTodoRoutes registers routes for the to-do group
func (r *Router) registerTodoRoutes(todo *gin.RouterGroup) {
	// Paths to `/api/v1/task/`
	todo.GET("/", r.api.GetTasks)
	todo.POST("/", r.api.CreateTask)

	// Paths to `/api/v1/task/:task_id`
	todo.GET("/:id", r.api.GetTaskById)
	todo.PUT("/:id", r.api.PutTaskById)
	todo.DELETE("/:id", r.api.DeleteTask)
}

func (r *Router) registerUserRoutes(login, logout *gin.RouterGroup) {
	// Path to `/api/v1/login`
	login.POST("/", r.api.Login)

	// Path to `/api/v1/logout`
	logout.POST("/", r.api.Logout)
}
