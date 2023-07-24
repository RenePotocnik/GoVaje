package Logic

import (
	"vaja1/DataStructures"
)

func (c *Controller) GetUsers() (users []DataStructures.User, err error) {

	// Ker ni potrebno nobenih podatkov preverit ali z njimi nekaj naredit lahko direkt vrnemo rezultat klica na bazo
	return c.db.GetUsers()

}

func (c *Controller) GetUserById(userId int) (user DataStructures.User, err error) {

	// Ker ni potrebno nobenih podatkov preverit ali z njimi nekaj naredit lahko direkt vrnemo rezultat klica na bazo
	return c.db.GetUserById(userId)

}
