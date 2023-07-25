package API

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (a *Controller) GetUsers(c *gin.Context) {
	// Get the tasks from the `Logic/` not the database directly
	users, err := a.c.GetUsers()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, users)
}

func (a *Controller) GetUserById(c *gin.Context) {
	// Get the ID from the URL and convert it to int
	userId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	user, err := a.c.GetUserById(userId)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)

}
