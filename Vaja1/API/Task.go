package API

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (a *Controller) GetTasks(c *gin.Context) {

	// Dobimo naloge - Kličemo Logic in ne direkt baze!
	tasks, err := a.c.GetTasks()
	if err != nil {
		// Vrnemo error 400 - Bad request
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	// Avtomatsko serializiramo objekt user v JSON in ga pošljemo z HTTP kodo 200 - OK
	c.JSON(http.StatusOK, tasks)

}

func (a *Controller) GetTaskById(c *gin.Context) {
	// Dobimo Task id tipa string iz naslova in ga pretvorimo v int
	taskId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		// Vrnemo error 400 - Bad request
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	// Dobimo uporabnika glede na ID - Kličemo Logic in ne direkt baze!
	user, err := a.c.GetTaskById(taskId)
	if err != nil {
		// Vrnemo error 500 - Internal server error
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	// Avtomatsko serializiramo objekt user v JSON in ga pošljemo z HTTP kodo 200 - OK
	c.JSON(http.StatusOK, user)
}
