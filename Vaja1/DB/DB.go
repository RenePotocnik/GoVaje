package DB

import "vaja1/DataStructures"

type DB interface {
	Init() (err error)

	GetUsers() (users []DataStructures.User, err error)
	GetUserById(id int) (user DataStructures.User, err error)
}
