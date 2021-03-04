package web

func mapUrls() {

	// blog endpoints
	httpRouter.GET("/blogs/:blog_id", messageHandler.Get) // message id
	httpRouter.GET("/blogs", messageHandler.GetAll)

	// officer endpoints
	httpRouter.GET("/officers/:officer_id/blogs", messageHandler.GetByAuthor)
	// httpRouter.GET("/officers/:officer_id", officerHandler.GetOfficer)
	// httpRouter.GET("/officers", authorHandler.GetAll)
}
