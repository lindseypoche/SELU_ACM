package web

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/lindseypoche/SELU_ACM/api/internal/domain"
)

// BlogController handles blog routes and data
type BlogController interface {
	Create(*gin.Context)
	Get(*gin.Context)
	GetAll(*gin.Context)
}

type blogController struct {
	blogService domain.BlogService
}

// NewBlogController creates a new controller for a blog
func NewBlogController(blogService domain.BlogService) BlogController {
	return &blogController{
		blogService: blogService,
	}
}

// Create creates a blog object using the data sent from the user
func (c *blogController) Create(ctx *gin.Context) {
	var blog domain.Blog
	if err := ctx.ShouldBindJSON(&blog); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	result, saveErr := c.blogService.CreateBlog(blog)
	if saveErr != nil {
		ctx.JSON(http.StatusBadRequest, saveErr)
		return
	}
	ctx.JSON(http.StatusCreated, result)
}

// Get gets a blog with the specified id from the uri
func (c *blogController) Get(ctx *gin.Context) {

	blogID := ctx.Param("blog_id")
	if blogID == "" {
		ctx.JSON(http.StatusBadRequest, errors.New("no blog id found"))
		return
	}

	blog, getErr := c.blogService.GetBlog(blogID)
	if getErr != nil {
		ctx.JSON(http.StatusBadRequest, getErr)
		return
	}

	ctx.JSON(http.StatusOK, blog)
}

// GetAll gets all blogs
func (c *blogController) GetAll(ctx *gin.Context) {

	blogs, getErr := c.blogService.GetAllBlogs()
	if getErr != nil {
		ctx.JSON(http.StatusNotFound, getErr)
		return
	}

	ctx.JSON(http.StatusOK, blogs)
}
