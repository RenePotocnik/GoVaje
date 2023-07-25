package main

import (
	"github.com/gin-gonic/gin"
	"vaja1/API"
)

// Naredimo Router objekt, da na njega obesimo funkcije
type Router struct {
	engine *gin.Engine
	api    API.Controller
}

func (r *Router) registerRoutes() (err error) {

	// Pot /api/v1
	api := r.engine.Group("/api/v1")

	// Pot /api/v1/user
	user := api.Group("/user")
	r.registerUserRoutes(user)

	todo := api.Group("/todo")
	r.registerTodoRoutes(todo)

	return

}

func (r *Router) registerUserRoutes(user *gin.RouterGroup) {

	// Pot /api/v1/user
	user.GET("/", r.api.GetUsers)

	// Pot /api/v1/user/:user_id
	user.GET("/:user_id", r.api.GetUserById)

}

func (r *Router) registerTodoRoutes(todo *gin.RouterGroup) {

	// Pot /api/v1/task
	todo.GET("/", r.api.GetTasks)

	// Pot /api/v1/task/:task_id
	todo.GET("/:task_id", r.api.GetTaskById)

	todo.POST("/", r.api.CreateTask)

}
