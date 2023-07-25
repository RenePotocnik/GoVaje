package DB

import "vaja1/DataStructures"

type DB interface {
	Init() (err error)

	GetUsers() (users []DataStructures.User, err error)
	GetUserById(id int) (user DataStructures.User, err error)

	GetTasks() (tasks []DataStructures.Task, err error)
	GetTaskById(id int) (task DataStructures.Task, err error)

	CreateTask(task DataStructures.Task) (err error)
	DeleteTask(taskId int) (err error)
}
