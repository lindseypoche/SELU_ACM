package web

func mapUrls() {

	httpRouter.GET("/ping", PingController.Ping)

	// blog endpoints
	httpRouter.POST("/blogs", blogHandler.Create)
	httpRouter.GET("/blogs/:blog_id", blogHandler.Get)
	httpRouter.GET("/blogs", blogHandler.GetAll)

	// user endpoints
	httpRouter.POST("/users", userHandler.Create)
	httpRouter.GET("/users/:user_id", userHandler.Get)
}
