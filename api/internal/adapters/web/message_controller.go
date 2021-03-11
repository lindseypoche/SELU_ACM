package web

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/lindseypoche/SELU_ACM/api/internal/domain"
)

// MessageController handles blog routes and data
type MessageController interface {
	Get(*gin.Context)
	GetAll(*gin.Context)
	GetByAuthor(*gin.Context)
	GetFeatured(*gin.Context)
}

type messageController struct {
	messageService domain.MessageService
}

// NewMessageController creates a new controller for a blog
func NewMessageController(messageService domain.MessageService) MessageController {
	return &messageController{
		messageService: messageService,
	}
}

// Get gets a message with the specified id from the uri
func (c *messageController) Get(ctx *gin.Context) {

	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:8081")
	ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET")

	messageID := ctx.Param("event_id")
	if messageID == "" {
		ctx.JSON(http.StatusBadRequest, errors.New("no message id found"))
		return
	}

	message, getErr := c.messageService.GetMessage(messageID)
	if getErr != nil {
		ctx.JSON(http.StatusBadRequest, getErr)
		return
	}
	ctx.JSON(http.StatusOK, message)
}

// GetAll gets all messages
func (c *messageController) GetAll(ctx *gin.Context) {

	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:8081")
	ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET")

	messages, getErr := c.messageService.GetAllMessages()
	if getErr != nil {
		ctx.JSON(http.StatusNotFound, getErr)
		return
	}
	ctx.JSON(http.StatusOK, messages)
}

// Get gets a message with the specified id from the uri
func (c *messageController) GetByAuthor(ctx *gin.Context) {

	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:8081")
	ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET")

	authorID := ctx.Param("officer_id")
	// TODO: validate authorID
	if authorID == "" {
		ctx.JSON(http.StatusBadRequest, errors.New("no officer id found"))
		return
	}

	messages, getErr := c.messageService.GetMessagesByAuthor(authorID)
	if getErr != nil {
		ctx.JSON(http.StatusBadRequest, getErr)
		return
	}
	ctx.JSON(http.StatusOK, messages)
}

func (c *messageController) GetFeatured(ctx *gin.Context) {
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:8081")
	ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET")

	channel := ctx.Param("channel_id")
	if channel == "" {
		ctx.JSON(http.StatusBadRequest, errors.New("no channel id found"))
		return
	}

	featured, getErr := c.messageService.GetFeatured(channel)
	if getErr != nil {
		ctx.JSON(http.StatusNotFound, getErr)
		return
	}
	ctx.JSON(http.StatusOK, featured)
}
