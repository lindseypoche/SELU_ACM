package web

func mapUrls() {

	// blog endpoints
	// httpRouter.POST("/blogs", messageHandler.Create) // no POST req from website. only discord client
	httpRouter.GET("/blogs/:blog_id", messageHandler.Get)
	httpRouter.GET("/blogs", messageHandler.GetAll)

	// user endpoints
	// httpRouter.POST("/users", userHandler.Create)
	// httpRouter.GET("/users/:user_id", userHandler.Get)
}
