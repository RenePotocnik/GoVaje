package Logic

import "vaja1/DataStructures"

func (c *Controller) GetTasks() (tasks []DataStructures.Task, err error) {

	// Ker ni potrebno nobenih podatkov preverit ali z njimi nekaj naredit lahko direkt vrnemo rezultat klica na bazo
	return c.db.GetTasks()

}

func (c *Controller) GetTaskById(taskId int) (task DataStructures.Task, err error) {

	// Ker ni potrebno nobenih podatkov preverit ali z njimi nekaj naredit lahko direkt vrnemo rezultat klica na bazo
	return c.db.GetTaskById(taskId)

}
