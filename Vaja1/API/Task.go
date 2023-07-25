package API

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"vaja1/DataStructures"
)

func (a *Controller) GetTasks(c *gin.Context) {
	// Get the tasks from the `Logic/` not the database directly
	tasks, err := a.c.GetTasks()
	if err != nil {
		// Return error 400 - Bad request if there was an error while getting the tasks
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	// Serialise the task object to JSON and send it with HTTP code 200 - OK
	c.JSON(http.StatusOK, tasks)

}

func (a *Controller) GetTaskById(c *gin.Context) {
	// Get the ID from the URL and convert it to int
	taskId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// Return error 400 - Bad request if the ID is not a number
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	user, err := a.c.GetTaskById(taskId)
	if err != nil {
		// Return error 500 - Internal server error if the task is not found or some other error occurs
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	// Serialise the task object to JSON and send it with HTTP code 200 - OK
	c.JSON(http.StatusOK, user)
}

func (a *Controller) PutTaskById(c *gin.Context) {
	// Get the ID from the URL and convert it to int
	taskId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// Return error 400 - Bad request if the ID is not a number
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	var task DataStructures.Task
	// Read the data from the request into the object
	err = c.BindJSON(&task)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	err = a.c.PutTaskById(taskId, task)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.String(http.StatusOK, "Task updated")
}

func (a *Controller) CreateTask(c *gin.Context) {
	var task DataStructures.Task
	// Read the data from the request into the object
	err := c.BindJSON(&task)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	err = a.c.CreateTask(task)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.String(http.StatusOK, "Task created")
}

func (a *Controller) DeleteTask(c *gin.Context) {
	// Get the ID from the URL and convert it to int
	taskId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// Return error 400 - Bad request if the ID is not a number
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	err = a.c.DeleteTask(taskId)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.String(http.StatusOK, "Task deleted")
}
