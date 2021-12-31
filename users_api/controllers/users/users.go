package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/SHA056/bookstore/users_api/domain/users"
	"github.com/SHA056/bookstore/users_api/services"
	"github.com/SHA056/bookstore/users_api/utils/errors"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("user id should be a number")
		c.JSON(err.Status, err)
		return
	}
	c.String(http.StatusNotImplemented, "implement me!")
}

func CreateUser(c *gin.Context) {
	var user users.User
	fmt.Println(user)
	err := c.ShouldBindJSON(&user)
	if err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveError := services.CreateUser(user)
	if saveError != nil {
		c.JSON(saveError.Status, saveError)
		// fmt.Println(saveError.Error)
		return
	}
	c.JSON(http.StatusCreated, result)

}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")

}
