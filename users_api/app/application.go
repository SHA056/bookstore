package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// First layer of application
// Only layer that interacts with the http server
// it would be easy to switch it from gin/gonics to any other http server
func StartApplication() {
	mapUrls()
	router.Run(":8080")
}
