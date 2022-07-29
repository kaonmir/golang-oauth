package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaonmir/OAuth/forms"
	"github.com/kaonmir/OAuth/models"
)

var userModel = new(models.User)

func GetUserById(c *gin.Context) {
	id := c.Query("id")
	password := c.Query("password")

	if id != "" {
		user, err := userModel.GetByID(id, password)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Error to retrieve user", "error": err.Error()})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "User founded!", "user": user})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
	c.Abort()
}

func CreateUser(c *gin.Context) {
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to read body", "error": err.Error()})
		c.Abort()
		return
	}

	var userPayload forms.UserSignup
	err = json.Unmarshal([]byte(value), &userPayload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error to parse body", "error": err.Error()})
		c.Abort()
		return
	}

	user, err := userModel.Signup(userPayload)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error to create user", "error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User created!", "user": user})
}
