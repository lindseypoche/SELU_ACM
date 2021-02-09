package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lindseypoche/SELU_ACM/api/domain/blogs"
	"github.com/lindseypoche/SELU_ACM/api/services"
	"github.com/pkg/errors"
)

var (
	BlogsController blogsControllerInterface = &blogsController{}
)

type blogsControllerInterface interface {
	Create(*gin.Context)
	Get(*gin.Context)
	GetAll(*gin.Context)
}

type blogsController struct{}

// Create creates a blog object using the data sent from the user
func (c *blogsController) Create(ctx *gin.Context) {
	var blog blogs.Blog
	if err := ctx.ShouldBindJSON(&blog); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	result, saveErr := services.BlogsService.CreateBlog(blog)
	if saveErr != nil {
		ctx.JSON(http.StatusBadRequest, saveErr)
		return
	}
	ctx.JSON(http.StatusCreated, result)
}

// Get gets a blog with the specified id from the uri
func (c *blogsController) Get(ctx *gin.Context) {

	blogID := ctx.Param("blog_id")
	if blogID == "" {
		ctx.JSON(http.StatusBadRequest, errors.New("no blog id found"))
		return
	}

	blog, getErr := services.BlogsService.GetBlog(blogID)
	if getErr != nil {
		ctx.JSON(http.StatusBadRequest, getErr)
		return
	}

	ctx.JSON(http.StatusOK, blog)
}

// GetAll gets all blogs
func (c *blogsController) GetAll(ctx *gin.Context) {

	blogs, getErr := services.BlogsService.GetAllBlogs()
	if getErr != nil {
		ctx.JSON(http.StatusNotFound, getErr)
		return
	}

	ctx.JSON(http.StatusOK, blogs)
}
