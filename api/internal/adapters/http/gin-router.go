package router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Context is a mock for gin.Context.
type Context interface {
	GET(uri string, f func(ctx *gin.Context))
	POST(uri string, f func(ctx *gin.Context))
	SERVE(port string)
}

var (
	ginDispatcher = gin.Default() // gin.Engine
)

// Note: ginRouter does not implement the Router interface
type ginRouter struct{}

// NewGinRouter creates a context object for a new gin router
func NewGinRouter() Context {
	return &ginRouter{}
}

// GET handles get requests
func (r *ginRouter) GET(uri string, f func(ctx *gin.Context)) {
	ginDispatcher.GET(uri, f)
}

// POST handles post requests
func (r *ginRouter) POST(uri string, f func(ctx *gin.Context)) {
	ginDispatcher.POST(uri, f)
}

// SERVE serves the gin router
func (r *ginRouter) SERVE(port string) {
	fmt.Printf("Gin HTTP server running on port %v", port)
	http.ListenAndServe(port, ginDispatcher)

}
