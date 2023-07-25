package Logic

import "vaja1/DataStructures"

func (c *Controller) GetTasks() (tasks []DataStructures.Task, err error) {
	// Because no checks are needed, we can directly return the result of the call to the database
	return c.db.GetTasks()
}

func (c *Controller) GetTaskById(taskId int) (task DataStructures.Task, err error) {
	return c.db.GetTaskById(taskId)
}

func (c *Controller) CreateTask(task DataStructures.Task) (err error) {
	return c.db.CreateTask(task)
}

func (c *Controller) DeleteTask(taskId int) (err error) {
	return c.db.DeleteTask(taskId)
}
