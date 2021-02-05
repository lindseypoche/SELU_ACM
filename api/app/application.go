package app

import (
	"os"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func getPort() string {
	if os.Getenv("PORT") != "" {
		return os.Getenv("PORT")
	}
	return ":8080"
}

func StartApplication() {
	mapUrls()
	router.Run(":" + getPort())
}
