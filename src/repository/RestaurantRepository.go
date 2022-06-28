package repository

import (
	"github.com/google/uuid"
	"github.com/yumaeda/sakaba-api/src/infrastructure"
	"github.com/yumaeda/sakaba-api/src/model"
	"gorm.io/gorm"
)

// RestaurantRepository is responsible for reading from and writing to DB Table `restaurants`.
type RestaurantRepository struct {
	DB *gorm.DB
}

// GetOpenRestaurants returns open restaurants.
func (c RestaurantRepository) GetOpenRestaurants() []model.RestaurantView {
	restaurants := []model.RestaurantView{}
	c.DB.Raw(`SELECT UuidFromBin(r.id) AS id,
	               r.url,
	               r.image_name,
	               r.name,
	               r.genre,
	               r.tel,
	               r.business_day_info,
	               r.address,
	               r.latitude,
	               r.longitude,
	               r.area,
	               r.comment,
	               r.takeout_available,
	               COUNT(p.restaurant_id) AS photo_count
                  FROM restaurants AS r
                  LEFT JOIN photos AS p
                    ON r.id = p.restaurant_id
                 WHERE is_closed = 0
                   AND REPLACE(JSON_EXTRACT(r.business_day_info, CONCAT('$.', DAYOFWEEK(CURDATE()), ".Start")), '"', '') <= DATE_FORMAT(CONVERT_TZ(NOW(), '+00:00', '+09:00'), '%H%i')
                   AND REPLACE(JSON_EXTRACT(r.business_day_info, CONCAT('$.', DAYOFWEEK(CURDATE()), ".End")), '"', '') >= DATE_FORMAT(CONVERT_TZ(NOW(), '+00:00', '+09:00'), '%H%i')
                 GROUP BY r.id
                 ORDER BY photo_count DESC`).Scan(&restaurants)

	return restaurants
}

// GetOpenRestaurantsByGenreID returns the open restaurants for the specified genre.
func (c RestaurantRepository) GetOpenRestaurantsByGenreID(genreID string) []model.RestaurantView {
	restaurants := []model.RestaurantView{}
	c.DB.Raw(`SELECT UuidFromBin(r.id) AS id,
	               r.url,
	               r.name,
	               r.genre,
	               r.tel,
	               r.business_day_info,
	               r.address,
	               r.latitude,
	               r.longitude,
	               r.area,
	               r.comment,
	               r.takeout_available,
	               COUNT(p.restaurant_id) AS photo_count
                  FROM restaurants AS r
                  JOIN restaurant_genres AS rg
                    ON r.id = rg.restaurant_id
                  LEFT JOIN photos AS p
                    ON r.id = p.restaurant_id
                 WHERE is_closed = 0
		   AND rg.genre_id = ` + genreID + `
                   AND REPLACE(JSON_EXTRACT(r.business_day_info, CONCAT('$.', DAYOFWEEK(CURDATE()), ".Start")), '"', '') <= DATE_FORMAT(CONVERT_TZ(NOW(), '+00:00', '+09:00'), '%H%i')
                   AND REPLACE(JSON_EXTRACT(r.business_day_info, CONCAT('$.', DAYOFWEEK(CURDATE()), ".End")), '"', '') >= DATE_FORMAT(CONVERT_TZ(NOW(), '+00:00', '+09:00'), '%H%i')
                 GROUP BY r.id
                 ORDER BY photo_count DESC`).Scan(&restaurants)

	return restaurants
}

// GetOpenRestaurantsByDrinkID returns the open restaurants for the specified drink.
func (c RestaurantRepository) GetOpenRestaurantsByDrinkID(drinkID string) []model.RestaurantView {
	restaurants := []model.RestaurantView{}
	c.DB.Raw(`SELECT UuidFromBin(r.id) AS id,
	               r.url,
	               r.name,
	               r.genre,
	               r.tel,
	               r.business_day_info,
	               r.address,
	               r.latitude,
	               r.longitude,
	               r.area,
	               r.comment,
	               r.takeout_available,
	               COUNT(p.restaurant_id) AS photo_count
                  FROM restaurants AS r
                  JOIN restaurant_drinks AS rd
                    ON r.id = rd.restaurant_id
                  LEFT JOIN photos AS p
                    ON r.id = p.restaurant_id
                 WHERE is_closed = 0
		   AND rd.drink_id = ` + drinkID + `
                   AND REPLACE(JSON_EXTRACT(r.business_day_info, CONCAT('$.', DAYOFWEEK(CURDATE()), ".Start")), '"', '') <= DATE_FORMAT(CONVERT_TZ(NOW(), '+00:00', '+09:00'), '%H%i')
                   AND REPLACE(JSON_EXTRACT(r.business_day_info, CONCAT('$.', DAYOFWEEK(CURDATE()), ".End")), '"', '') >= DATE_FORMAT(CONVERT_TZ(NOW(), '+00:00', '+09:00'), '%H%i')
                 GROUP BY r.id
                 ORDER BY photo_count DESC`).Scan(&restaurants)

	return restaurants
}

// GetOpenRestaurantsByDishID returns the open restaurants which have the specified dish.
func (c RestaurantRepository) GetOpenRestaurantsByDishID(dishID string) []model.RestaurantView {
	restaurants := []model.RestaurantView{}
	c.DB.Raw(`SELECT UuidFromBin(r.id) AS id,
	               r.url,
	               r.name,
	               r.genre,
	               r.tel,
	               r.business_day_info,
	               r.address,
	               r.latitude,
	               r.longitude,
	               r.area,
	               r.comment,
	               r.takeout_available,
	               0 AS photo_count
	          FROM dishes AS d
	          JOIN rankings AS rk
                    ON rk.dish_id = d.id
                  JOIN photos AS p
                    ON p.id = rk.photo_id
                  JOIN restaurants AS r
                    ON r.id = p.restaurant_id
                 WHERE r.is_closed = 0
                   AND d.id = ` + dishID + `
                   AND REPLACE(JSON_EXTRACT(r.business_day_info, CONCAT('$.', DAYOFWEEK(CURDATE()), ".Start")), '"', '') <= DATE_FORMAT(CONVERT_TZ(NOW(), '+00:00', '+09:00'), '%H%i')
                   AND REPLACE(JSON_EXTRACT(r.business_day_info, CONCAT('$.', DAYOFWEEK(CURDATE()), ".End")), '"', '') >= DATE_FORMAT(CONVERT_TZ(NOW(), '+00:00', '+09:00'), '%H%i')
                 ORDER BY rk.rank ASC`).Scan(&restaurants)

	return restaurants
}

// GetOpenRestaurantCount returns the number of open restaurants.
func (c RestaurantRepository) GetOpenRestaurantCount() []model.RestaurantCount {
	restaurantCounts := []model.RestaurantCount{}
	c.DB.Raw(`SELECT area,
                       COUNT(area) AS count
                  FROM restaurants
                 WHERE is_closed = 0
                   AND REPLACE(JSON_EXTRACT(business_day_info, CONCAT('$.', DAYOFWEEK(CURDATE()), ".Start")), '"', '') <= DATE_FORMAT(CONVERT_TZ(NOW(), '+00:00', '+09:00'), '%H%i')
                   AND REPLACE(JSON_EXTRACT(business_day_info, CONCAT('$.', DAYOFWEEK(CURDATE()), ".End")), '"', '') >= DATE_FORMAT(CONVERT_TZ(NOW(), '+00:00', '+09:00'), '%H%i')
                 GROUP BY area
                 ORDER BY COUNT(area) DESC`).Scan(&restaurantCounts)

	return restaurantCounts
}

// AddRestaurant adds a new restaurant.
func (c RestaurantRepository) AddRestaurant(URL string, name string, genre string, tel string, businessDayInfo string, address string, latitude string, longitude string, area string) error {
	id := uuid.New()
	restaurant := model.Restaurant{
		ID:              infrastructure.UUIDToBin(id.String()),
		URL:             URL,
		Name:            name,
		Genre:           genre,
		Tel:             tel,
		BusinessDayInfo: businessDayInfo,
		Address:         address,
		Latitude:        latitude,
		Longitude:       longitude,
		Area:            area,
	}
	dbError := c.DB.Create(&restaurant).Error

	return dbError
}
