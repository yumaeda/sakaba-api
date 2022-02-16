package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"sakaba.link/api/src/repository"
)

type DishController struct{}

func (c *DishController) GetAllDishes(ctx *gin.Context) {
	dishRepository := repository.DishRepository{}
	allDishes := dishRepository.GetAllDishes()

	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"body":       allDishes,
	})
}

func (c *DishController) GetDishById(ctx *gin.Context) {
	dishRepository := repository.DishRepository{}
	dish := dishRepository.GetDishById(ctx.Param("id"))

	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"body":       dish,
	})
}
