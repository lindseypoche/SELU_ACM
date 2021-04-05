package rest

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cmd-ctrl-q/SELU_ACM/api/internal/listing"
	"github.com/cmd-ctrl-q/SELU_ACM/api/internal/utils/errors/rest"
	"github.com/cmd-ctrl-q/SELU_ACM/api/internal/utils/http_utils"
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func Handler(l listing.Service) http.Handler {
	router := gin.Default()

	// cors
	router.Use(cors.New(cors.Config{
		AllowedOrigins: []string{"http://127.0.0.1"},
		AllowedMethods: []string{"GET"},
		AllowedHeaders: []string{"Origin"},
		ExposedHeaders: []string{"Content-Length"},
		// AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}))

	// load html/react templates
	router.Use(static.Serve("/", static.LocalFile("../../../../web/build", true)))

	// api routes
	api := router.Group("/api")
	api.GET("", getTestAPI())
	{
		// events
		events := api.Group("/events")
		events.GET("", getMessages(l))
		events.GET("/:message_id", getMessage(l))

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
	}

	return router
}

func getTestAPI() func(*gin.Context) {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, map[string]string{"status": "yerp"})
	}
}

// Get gets a message with the specified id from the uri
func getMessage(s listing.Service) func(*gin.Context) {
	return func(ctx *gin.Context) {

		messageID := ctx.Param("message_id")
		if messageID == "" {
			respErr := rest.NewRestError("uri parameter could not be found", 400, "bad request error", []interface{}{"uri paramter"})
			http_utils.RespondError(ctx.Writer, respErr)
			return
		}

		message, getErr := s.GetMessage(messageID)
		if getErr != nil {
			// ctx.JSON(http.StatusBadRequest, getErr) // old
			http_utils.RespondError(ctx.Writer, getErr)
			return
		}
		// ctx.JSON(http.StatusOK, message) // old
		http_utils.RespondJson(ctx.Writer, "GET", http.StatusOK, message)
	}
}

// GetAll gets all messages
func getMessages(s listing.Service) func(*gin.Context) {
	return func(ctx *gin.Context) {

		messages, getErr := s.GetAllMessages()
		if getErr != nil {
			http_utils.RespondError(ctx.Writer, getErr)
			return
		}
		// ctx.JSON(http.StatusOK, messages)
		http_utils.RespondJson(ctx.Writer, "GET", http.StatusOK, messages)
	}
}

// Get gets a message with the specified author name from the uri
func getMessagesByUsername(s listing.Service) func(*gin.Context) {
	return func(ctx *gin.Context) {

		authorName := ctx.Param("username")
		// TODO: validate username
		if authorName == "" {
			// ctx.JSON(http.StatusBadRequest, rest.NewBadRequestError("uri parameter 'username' could not be found")) // old
			respErr := rest.NewRestError("uri parameter could not be found", 400, "bad request error", []interface{}{"uri paramter"})
			http_utils.RespondError(ctx.Writer, respErr)
			return
		}

		messages, getErr := s.GetMessagesByUsername(authorName)
		if getErr != nil {
			// ctx.JSON(getErr.GetStatus(), getErr) // old
			http_utils.RespondError(ctx.Writer, getErr)
			return
		}
		// ctx.JSON(http.StatusOK, messages) // old
		http_utils.RespondJson(ctx.Writer, "GET", http.StatusOK, messages)
	}
}

// get a pinned message by the message id (pointless?)
func getPinnedMessage(s listing.Service) func(*gin.Context) {
	return func(ctx *gin.Context) {

		err := ctx.Request.ParseForm()
		if err != nil {
			restErr := rest.NewBadRequestError(fmt.Sprintf("query not readable: %v", err))
			http_utils.RespondError(ctx.Writer, restErr)
		}

		messageID := ctx.Request.FormValue("id")
		if messageID == "" {
			respErr := rest.NewRestError("uri parameter could not be found", 400, "bad_request", []interface{}{"uri paramter"})
			http_utils.RespondError(ctx.Writer, respErr)
			return
		}

		latest, getErr := s.GetPinnedMessage(messageID)
		if getErr != nil {
			// ctx.JSON(getErr.GetStatus(), getErr) // old
			http_utils.RespondError(ctx.Writer, getErr)
			return
		}

		http_utils.RespondJson(ctx.Writer, "GET", 200, latest)
	}
}

// get all pinned messages
func getPinnedMessages(s listing.Service) func(*gin.Context) {
	return func(ctx *gin.Context) {

		messages, getErr := s.GetAllPinnedMessages()
		if getErr != nil {
			http_utils.RespondError(ctx.Writer, getErr)
			return
		}
		// ctx.JSON(http.StatusOK, messages)
		http_utils.RespondJson(ctx.Writer, "GET", http.StatusOK, messages)
	}
}

// get the latest pinned message by the channel id
func getLatestPinnedMessage(s listing.Service) func(*gin.Context) {
	return func(ctx *gin.Context) {

		err := ctx.Request.ParseForm()
		if err != nil {
			restErr := rest.NewBadRequestError(fmt.Sprintf("query not readable: %v", err))
			http_utils.RespondError(ctx.Writer, restErr)
		}

		channelID := ctx.Request.FormValue("id")

		if channelID == "" {
			// ctx.JSON(http.StatusBadRequest, rest.NewBadRequestError("uri parameter 'id' could not be found")) // old
			respErr := rest.NewRestError("uri parameter could not be found", 400, "bad_request", []interface{}{"uri paramter"})
			http_utils.RespondError(ctx.Writer, respErr)
			return
		}

		featured, getErr := s.GetLatestPinned(channelID)
		if getErr != nil {
			// ctx.JSON(getErr.GetStatus(), getErr) // old
			http_utils.RespondError(ctx.Writer, getErr)
			return
		}
		// ctx.JSON(http.StatusOK, featured)
		http_utils.RespondJson(ctx.Writer, "GET", http.StatusOK, featured)
	}
}

func getAllLatestPinnedMessage(s listing.Service) func(*gin.Context) {
	return func(ctx *gin.Context) {

		allFeatured, getErr := s.GetAllLatestPinned()
		if getErr != nil {
			// ctx.JSON(getErr.GetStatus(), getErr) // old
			http_utils.RespondError(ctx.Writer, getErr)
			return
		}

		// ctx.JSON(http.StatusOK, featured)
		http_utils.RespondJson(ctx.Writer, "GET", http.StatusOK, allFeatured)
	}
}
