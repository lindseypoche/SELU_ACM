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
	router.POST("/users", controllers.UsersController.Create)
	router.GET("/users/:user_id", controllers.UsersController.Get)
	router.PUT("/users/:user_id", controllers.UsersController.Update)
	router.PATCH("/users/:user_id", controllers.UsersController.Update)
	router.DELETE("/users/:user_id", controllers.UsersController.Delete)
	router.POST("/users/login", controllers.UsersController.Login)
}
