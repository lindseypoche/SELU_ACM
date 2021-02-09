package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	PingController pingControllerInterface = &pingController{}
)

type pingControllerInterface interface {
	Ping(*gin.Context)
}

type pingController struct{}

func (c *pingController) Ping(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}
