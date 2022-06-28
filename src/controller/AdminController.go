package controller

import (
	"net/http"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/yumaeda/sakaba-api/src/model"
)

var identityKey = "id"

// AdminController is a controller for Admin API.
type AdminController struct{}

// Index returns welcome message in JSON format.
func (c AdminController) Index(ctx *gin.Context) {
	claims := jwt.ExtractClaims(ctx)
	user, exists := ctx.Get(identityKey)
	if exists {
		ctx.JSON(http.StatusOK, gin.H{
			"userID":   claims[identityKey],
			"userName": user.(*model.User).UserName,
			"text":     "Hello Admin.",
		})
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "Cannot find claims in the current context.",
		})
	}
}
