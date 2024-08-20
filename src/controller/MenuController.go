package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yumaeda/sakaba-api/src/model"
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

// AddMenu adds a new menu.
func (c MenuController) AddMenu(ctx *gin.Context) {
	var json model.Menu
	if err := ctx.ShouldBindJSON(&json); err == nil {
		dbError := c.Repository.AddMenu(json.ID, json.RestaurantID, json.Name, json.NameJpn, json.Price)
		if dbError == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"statusCode": 200,
				"id":         json.ID,
			})
			return
		}
	}

	ctx.JSON(http.StatusBadRequest, gin.H{
		"statusCode": 400,
		"error":      "Menu insertion failed.",
	})
}

func (c MenuController) SetMenu(ctx *gin.Context) {
	var json model.MenuUpdate
	if err := ctx.ShouldBindJSON(&json); err == nil {
		if err := c.Repository.SetMenu(json.ID, json.Column, json.Value); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"statusCode": 200,
			})
			return
		}
	}

	ctx.JSON(http.StatusBadRequest, gin.H{
		"statusCode": 400,
		"error":      "Menu update failed.",
	})
}

// DeleteMenu deletes the specified menu.
func (c MenuController) DeleteMenu(ctx *gin.Context) {
	var json model.MenuUpdate
	if err := ctx.ShouldBindJSON(&json); err == nil {
		if err := c.Repository.DeleteMenu(json.ID); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"statusCode": 200,
			})
			return
		}
	}

	ctx.JSON(http.StatusBadRequest, gin.H{
		"statusCode": 400,
		"error":      "Menu deletion failed.",
	})
}
