package controller

import (
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"sakaba.link/api/src/model"
)

var identityKey = "id"

// AdminController is a controller for Admin API.
type AdminController struct{}

// Index returns welcome message in JSON format.
func (c *AdminController) Index(ctx *gin.Context) {
	claims := jwt.ExtractClaims(ctx)
	user, _ := ctx.Get(identityKey)
	ctx.JSON(http.StatusOK, gin.H{
		"userID":   claims[identityKey],
		"userName": user.(*model.User).UserName,
		"text":     "Hello Admin.",
	})
}
