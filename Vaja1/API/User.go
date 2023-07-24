package API

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (a *Controller) GetUsers(c *gin.Context) {

	// Dobimo uporabnike - Kličemo Logic in ne direkt baze!
	users, err := a.c.GetUsers()
	if err != nil {
		// Vrnemo error 400 - Bad request
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	// Avtomatsko serializiramo objekt user v JSON in ga pošljemo z HTTP kodo 200 - OK
	c.JSON(http.StatusOK, users)

}

func (a *Controller) GetUserById(c *gin.Context) {

	// Dobimo user_id tipa string iz naslova in ga pretvorimo v int
	userId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		// Vrnemo error 400 - Bad request
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	// Dobimo uporabnika glede na ID - Kličemo Logic in ne direkt baze!
	user, err := a.c.GetUserById(userId)
	if err != nil {
		// Vrnemo error 500 - Internal server error
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	// Avtomatsko serializiramo objekt user v JSON in ga pošljemo z HTTP kodo 200 - OK
	c.JSON(http.StatusOK, user)

}
