package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yumaeda/sakaba-api/src/repository"
)

// MenuController is a controller for Menu API.
type MenuController struct {
	Repository repository.MenuRepository
}

// GetMenuByRestaurantID returns menus for the specified restaurant.
func (c MenuController) GetMenusByRestaurantID(ctx *gin.Context) {
	menus := c.Repository.GetMenus(ctx.Param("id"))

	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"body":       menus,
	})
}
