package web

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/lindseypoche/SELU_ACM/api/internal/domain"
)

// BlogController handles blog routes and data
type MessageController interface {
	Create(*gin.Context)
	Get(*gin.Context)
	GetAll(*gin.Context)
}

type messageController struct {
	messageService domain.MessageService
}

// NewBlogController creates a new controller for a blog
func NewMessageController(messageService domain.MessageService) MessageController {
	return &messageController{
		messageService: messageService,
	}
}

// Create creates a blog object using the data sent from the user
func (c *messageController) Create(ctx *gin.Context) {
	var message domain.Message
	if err := ctx.ShouldBindJSON(&message); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	result, saveErr := c.messageService.CreateMessage(message)
	if saveErr != nil {
		ctx.JSON(http.StatusBadRequest, saveErr)
		return
	}
	ctx.JSON(http.StatusCreated, result)
}

// Get gets a blog with the specified id from the uri
func (c *messageController) Get(ctx *gin.Context) {

	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:8081")
	ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET")

	messageID := ctx.Param("blog_id")
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

// GetAll gets all blogs
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
