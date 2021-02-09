package app

import (
	"github.com/lindseypoche/SELU_ACM/api/controllers"
)

func mapUrls() {

	router.GET("/ping", controllers.PingController.Ping)

	// blog endpoints
	router.POST("/blogs", controllers.BlogsController.Create)
	router.GET("/blogs/:blog_id", controllers.BlogsController.Get)
	router.GET("/blogs", controllers.BlogsController.GetAll)

	// user endpoints
	// router.POST("/user", controllers.UsersController.Create)
	// router.GET("/user:user_id", controllers.UsersController.Get)
}
