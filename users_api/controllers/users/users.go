package users

import (
	"fmt"
	"net/http"

	"github.com/SHA056/bookstore/users_api/domain/users"
	"github.com/SHA056/bookstore/users_api/services"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}

func CreateUser(c *gin.Context) {
	var user users.User
	fmt.Println(user)
	err := c.ShouldBindJSON(&user)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(user)
	result, saveError := services.CreateUser(user)
	if saveError != nil {
		fmt.Println(saveError.Error())
		return
	}
	c.JSON(http.StatusCreated, result)

}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")

}
