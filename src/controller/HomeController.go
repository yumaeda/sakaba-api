package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HomeController is a controller for Home API.
type HomeController struct{}

// Index returns welcome message in JSON format.
func (c *HomeController) Index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "hello world"})
}
