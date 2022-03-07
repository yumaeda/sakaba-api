package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"sakaba.link/api/src/repository"
)

type RankingController struct{}

func (c *RankingController) GetAllRankings(ctx *gin.Context) {
	rankingRepository := repository.RankingRepository{}
	allRankings := rankingRepository.GetAllRankings()

	ctx.JSON(http.StatusOK, gin.H{
		"statusCode": 200,
		"body":       allRankings,
	})
}
