package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yumaeda/sakaba-api/src/repository"
)

// AreaController is a controller for Area API.
type AreaController struct {
	Repository repository.AreaRepository
}

// GetAllAreas returns all the areas.
func (c AreaController) GetAllAreas(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"body":       c.Repository.GetAllAreas(),
	})
}
