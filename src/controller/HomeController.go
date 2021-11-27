package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HomeController struct{}

func (c *HomeController) Index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "hello world"})
}
