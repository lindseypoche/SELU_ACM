package app

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lindseypoche/SELU_ACM/api/clients/mongo"
)

var (
	router = gin.Default()
)

func getPort() string {
	if os.Getenv("PORT") != "" {
		return ":" + os.Getenv("PORT")
	}
	return ":8080"
}

// StartApplication ...
func StartApplication() {

	mongo.Init()

	mapUrls()

	srv := &http.Server{
		Addr:         "127.0.0.1" + getPort(),
		Handler:      router,
		WriteTimeout: 500 * time.Millisecond,
		ReadTimeout:  2 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
