package web

func mapUrls() {

	// events endpoints
	httpRouter.GET("/events", messageHandler.GetAll)
	httpRouter.GET("/events/:event_id", messageHandler.Get)
	httpRouter.GET("/featured/:channel_id", messageHandler.GetFeatured)

	// officer endpoints
	httpRouter.GET("/officers/:officer_id/events", messageHandler.GetByAuthor)
	// httpRouter.GET("/officers/:officer_id", officerHandler.GetOfficer)
	// httpRouter.GET("/officers", authorHandler.GetAll)
}
