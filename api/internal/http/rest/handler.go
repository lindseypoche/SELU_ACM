package rest

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cmd-ctrl-q/SELU_ACM/api/internal/listing"
	"github.com/cmd-ctrl-q/SELU_ACM/api/internal/utils/errors/rest"
	"github.com/cmd-ctrl-q/SELU_ACM/api/internal/utils/http_utils"
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
)

func Handler(l listing.Service) http.Handler {
	router := gin.Default()

	// cors
	router.Use(cors.New(cors.Config{
		// AllowedOrigins: []string{"http://web:3000"}, // allow from docker web container
		AllowedOrigins: []string{"http://localhost:8080", "http://localhost:8082", "http://localhost:3000"}, // allow all
		AllowedMethods: []string{"GET"},
		AllowedHeaders: []string{"Origin"},
		ExposedHeaders: []string{"Content-Length"},
		// AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}))

	// load html/react templates in the web docker container
	// router.Use(static.Serve("/", static.LocalFile("web:/web/build", true)))

	// api routes
	api := router.Group("/api")
	api.GET("", getTestAPI())
	{
		// events
		events := api.Group("/events")
		events.GET("", getMessages(l))
		events.GET("/:message_id", getMessage(l))
		events.GET("/:message_id/comments", getCommentsByMessageRefID(l))

		// pins
		pinned := api.Group("/pins")
		pinned.GET("", getPinnedMessages(l))
		pinned.GET("/message", getPinnedMessage(l))
		pinned.GET("/channel", getLatestPinnedMessage(l))
		// pinned.GET("/all_latest", getAllLatestPinnedMessage(l)) // ie latest from all channels

		// authors
		author := api.Group("/author")
		author.GET("/:username/events", getMessagesByUsername(l)) //api/author/quantacake/events

		// officers / members
		member := api.Group("/member")
		member.GET("", getOfficers(l))
	}

	return router
}

func getTestAPI() func(*gin.Context) {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "yerp")
	}
}

// Get gets a message with the specified id from the uri
func getMessage(s listing.Service) func(*gin.Context) {
	return func(c *gin.Context) {

		messageID := c.Param("message_id")
		if messageID == "" {
			respErr := rest.NewRestError("uri parameter could not be found", 400, "bad request error", []interface{}{"uri paramter"})
			http_utils.RespondError(c.Writer, respErr)
			return
		}

		message, getErr := s.GetMessage(messageID)
		if getErr != nil {
			// c.JSON(http.StatusBadRequest, getErr) // old
			http_utils.RespondError(c.Writer, getErr)
			return
		}
		// c.JSON(http.StatusOK, message) // old
		http_utils.RespondJson(c.Writer, "GET", http.StatusOK, message)
	}
}

// GetAll gets all messages
func getMessages(s listing.Service) func(*gin.Context) {
	return func(c *gin.Context) {

		// messages, getErr := s.GetAllMessages() // original get all messages in messages collection
		messages, getErr := s.GetByStartTime()
		if getErr != nil {
			http_utils.RespondError(c.Writer, getErr)
			return
		}
		// c.JSON(http.StatusOK, messages)
		http_utils.RespondJson(c.Writer, "GET", http.StatusOK, messages)
	}
}

// Get gets a message with the specified author name from the uri
func getMessagesByUsername(s listing.Service) func(*gin.Context) {
	return func(c *gin.Context) {

		authorName := c.Param("username")
		// TODO: validate username
		if authorName == "" {
			// c.JSON(http.StatusBadRequest, rest.NewBadRequestError("uri parameter 'username' could not be found")) // old
			respErr := rest.NewRestError("uri parameter could not be found", 400, "bad request error", []interface{}{"uri paramter"})
			http_utils.RespondError(c.Writer, respErr)
			return
		}

		messages, getErr := s.GetMessagesByUsername(authorName)
		if getErr != nil {
			// c.JSON(getErr.GetStatus(), getErr) // old
			http_utils.RespondError(c.Writer, getErr)
			return
		}
		// c.JSON(http.StatusOK, messages) // old
		http_utils.RespondJson(c.Writer, "GET", http.StatusOK, messages)
	}
}

// Get all comments by the message reference id (event id)
func getCommentsByMessageRefID(s listing.Service) func(*gin.Context) {
	return func(c *gin.Context) {

		refID := c.Param("message_id")
		comments, getErr := s.GetComments(refID)
		if getErr != nil {
			http_utils.RespondError(c.Writer, getErr)
		}

		http_utils.RespondJson(c.Writer, "GET", http.StatusOK, comments)
	}
}

// get a pinned message by the message id (pointless?)
func getPinnedMessage(s listing.Service) func(*gin.Context) {
	return func(c *gin.Context) {

		err := c.Request.ParseForm()
		if err != nil {
			restErr := rest.NewBadRequestError(fmt.Sprintf("query not readable: %v", err))
			http_utils.RespondError(c.Writer, restErr)
		}

		messageID := c.Request.FormValue("id")
		if messageID == "" {
			respErr := rest.NewRestError("uri parameter could not be found", 400, "bad_request", []interface{}{"uri paramter error"})
			http_utils.RespondError(c.Writer, respErr)
			return
		}

		// id := c.Query("id") // use this rather than ParseForm()

		latest, getErr := s.GetPinnedMessage(messageID)
		if getErr != nil {
			http_utils.RespondError(c.Writer, getErr)
			return
		}

		http_utils.RespondJson(c.Writer, "GET", 200, latest)
	}
}

// get all pinned messages
func getPinnedMessages(s listing.Service) func(*gin.Context) {
	return func(c *gin.Context) {

		messages, getErr := s.GetAllPinnedMessages()
		if getErr != nil {
			http_utils.RespondError(c.Writer, getErr)
			return
		}
		// c.JSON(http.StatusOK, messages)
		http_utils.RespondJson(c.Writer, "GET", http.StatusOK, messages)
	}
}

// get the latest pinned message by the channel id
func getLatestPinnedMessage(s listing.Service) func(*gin.Context) {
	return func(c *gin.Context) {

		err := c.Request.ParseForm()
		if err != nil {
			restErr := rest.NewBadRequestError(fmt.Sprintf("query not readable: %v", err))
			http_utils.RespondError(c.Writer, restErr)
		}

		channelID := c.Request.FormValue("id")
		if channelID == "" {
			// c.JSON(http.StatusBadRequest, rest.NewBadRequestError("uri parameter 'id' could not be found")) // old
			respErr := rest.NewRestError("uri parameter could not be found", 400, "bad_request", []interface{}{"uri paramter"})
			http_utils.RespondError(c.Writer, respErr)
			return
		}

		featured, getErr := s.GetLatestPinned(channelID)
		if getErr != nil {
			// c.JSON(getErr.GetStatus(), getErr) // old
			http_utils.RespondError(c.Writer, getErr)
			return
		}
		// c.JSON(http.StatusOK, featured)
		http_utils.RespondJson(c.Writer, "GET", http.StatusOK, featured)
	}
}

func getAllLatestPinnedMessage(s listing.Service) func(*gin.Context) {
	return func(c *gin.Context) {

		allFeatured, getErr := s.GetAllLatestPinned()
		if getErr != nil {
			// c.JSON(getErr.GetStatus(), getErr) // old
			http_utils.RespondError(c.Writer, getErr)
			return
		}

		// c.JSON(http.StatusOK, featured)
		http_utils.RespondJson(c.Writer, "GET", http.StatusOK, allFeatured)
	}
}

// GetAll gets all messages
func getOfficers(s listing.Service) func(*gin.Context) {
	return func(c *gin.Context) {

		// messages, getErr := s.GetAllMessages() // original get all messages in messages collection
		officers, getErr := s.GetAllOfficers()
		if getErr != nil {
			http_utils.RespondError(c.Writer, getErr)
			return
		}
		// c.JSON(http.StatusOK, messages)
		http_utils.RespondJson(c.Writer, "GET", http.StatusOK, officers)
	}
}
