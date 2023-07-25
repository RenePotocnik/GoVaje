package Logic

import (
	"vaja1/DataStructures"
)

func (c *Controller) GetUsers() (users []DataStructures.User, err error) {
	return c.db.GetUsers()
}

func (c *Controller) GetUserById(userId int) (user DataStructures.User, err error) {
	return c.db.GetUserById(userId)
}
