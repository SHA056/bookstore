package app

import (
	"github.com/SHA056/bookstore/users_api/controllers/ping"
	"github.com/SHA056/bookstore/users_api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	// router.GET("/users/:user_id", controllers.GetUser)
	router.GET("/users/search", users.SearchUser)
	router.POST("/users", users.CreateUser)
}
