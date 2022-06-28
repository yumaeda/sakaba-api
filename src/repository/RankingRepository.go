package repository

import (
	"github.com/yumaeda/sakaba-api/src/model"
	"gorm.io/gorm"
)

// RankingRepository is responsible for reading from and writing to DB Table `rankings`.
type RankingRepository struct {
	DB *gorm.DB
}

// GetAllRankings returns all the rankings.
func (c RankingRepository) GetAllRankings() []model.Ranking {
	allRankings := []model.Ranking{}
	c.DB.Raw(`SELECT dishes.name AS 'dish',
	               rankings.rank AS 'rank',
	               restaurants.name AS 'restaurant',
	               UuidFromBin(restaurants.id) AS 'restaurant_id',
	               photos.name AS 'photo',
	               restaurants.url AS 'restaurant_url'
                  FROM rankings
                  JOIN dishes
                    ON rankings.dish_id = dishes.id
                  JOIN photos
                    ON rankings.photo_id = photos.id
                  JOIN restaurants
                    ON photos.restaurant_id = restaurants.id
                 ORDER BY dishes.name ASC, rankings.rank ASC`).Scan(&allRankings)

	return allRankings
}
