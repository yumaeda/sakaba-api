package controller

import (
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"sakaba.link/api/src/model"
)

var identityKey = "id"

// HomeController is a controller for Home API.
type HomeController struct{}

// Index returns welcome message in JSON format.
func (c *HomeController) Index(ctx *gin.Context) {
	claims := jwt.ExtractClaims(ctx)
	user, _ := ctx.Get(identityKey)
	ctx.JSON(http.StatusOK, gin.H{
		"userID":   claims[identityKey],
		"userName": user.(*model.User).UserName,
		"text":     "Hello World.",
	})
}
