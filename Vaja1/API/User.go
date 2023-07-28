package API

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vaja1/DataStructures"
)

func (a *Controller) Login(c *gin.Context) {
	var user DataStructures.User

	// Bind the data from the request body to the user struct (username, password)
	err := c.BindJSON(&user)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	token, err := a.c.Login(user)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	// Set the JWT token in the authorisation header
	c.Writer.Header().Set("Authorization", token)
	c.JSON(http.StatusOK, token)
}

func (a *Controller) Logout(c *gin.Context) {
	// Set the JWT token in the authorisation header to an empty string
	c.Writer.Header().Set("Authorization", "")
	c.JSON(http.StatusOK, "logged out")
}
