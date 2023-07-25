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

	// user := api.Group("/user")
	// r.registerUserRoutes(user)

	todo := api.Group("/todo")
	r.registerTodoRoutes(todo)

	return

}

func (r *Router) registerUserRoutes(user *gin.RouterGroup) {
	// Path to `/api/v1/user`
	user.GET("/", r.api.GetUsers)

	// Path to `/api/v1/user/:user_id`
	user.GET("/:user_id", r.api.GetUserById)
}

// registerTodoRoutes registers routes for the to-do group
func (r *Router) registerTodoRoutes(todo *gin.RouterGroup) {
	// Path to `/api/v1/task`
	todo.GET("/", r.api.GetTasks)

	// Path to `/api/v1/task/:task_id`
	todo.GET("/:task_id", r.api.GetTaskById)
	todo.POST("/", r.api.CreateTask)
}
