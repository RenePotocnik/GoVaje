package Logic

import (
	"fmt"
	"vaja1/DataStructures"
)

func (c *Controller) GetTasks(userID int) (tasks []DataStructures.Task, err error) {
	// Because no checks are needed, we can directly return the result of the call to the database
	return c.db.GetTasks(userID)
}

func (c *Controller) GetTaskById(taskId int) (task DataStructures.Task, err error) {
	return c.db.GetTaskById(taskId)
}

func (c *Controller) PutTaskById(taskId int, task DataStructures.Task) (err error) {
	return c.db.PutTaskById(taskId, task)
}

func (c *Controller) CreateTask(task DataStructures.Task) (err error) {
	return c.db.CreateTask(task)
}

func (c *Controller) DeleteTask(taskId int) (err error) {
	// Check if the task exists
	if _, err = c.db.GetTaskById(taskId); err != nil {
		fmt.Println("Task with that ID does not exist")
		return err
	}
	return c.db.DeleteTask(taskId)
}
