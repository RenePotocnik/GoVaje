package DB

import "vaja1/DataStructures"

type DB interface {
	Init() (err error)

	Login(user DataStructures.User) (err error)
	GetUserByUsername(username string) (user DataStructures.User, err error)

	GetTasks(userId int) (tasks []DataStructures.Task, err error)
	GetTaskById(id, userId int) (task DataStructures.Task, err error)
	PutTaskById(id int, task DataStructures.Task) (err error)

	CreateTask(task DataStructures.Task, userId int) (err error)
	DeleteTask(taskId int) (err error)
}
