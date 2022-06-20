package repository

import (
	"sakaba.link/api/src/infrastructure"
	"sakaba.link/api/src/model"
)

// RankingRepository is responsible for reading from and writing to DB Table `rankings`.
type RankingRepository struct{}

// GetAllRankings returns all the rankings.
func (c *RankingRepository) GetAllRankings() []model.Ranking {
	db, closer, err := infrastructure.ConnectToDB()
	if err != nil {
		panic(err.Error())
	}
	defer closer()

	allRankings := []model.Ranking{}
	db.Raw(`SELECT dishes.name AS 'dish',
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
