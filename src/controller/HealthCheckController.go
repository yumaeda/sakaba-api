package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheckController is a controller for health check API.
type HealthCheckController struct{}

// GetStatus returns health status message in JSON format.
func (c HealthCheckController) GetStatus(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "API server is up and running."})
}
