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

// GetRestaurants returns all the restaurants.
func (c RestaurantRepository) GetRestaurants() []model.Restaurant {
	restaurants := []model.Restaurant{}
	c.DB.Raw(`SELECT UuidFromBin(id) AS id,
	                 url,
	                 name,
	                 genre,
	                 tel,
	                 business_day_info,
	                 address,
	                 latitude,
	                 longitude,
	                 area
                FROM restaurants
               WHERE is_closed = 0
               ORDER BY area ASC, name ASC`).Scan(&restaurants)

	return restaurants
}

// GetRestaurantsByArea returns restaurants for the specified area.
func (c RestaurantRepository) GetRestaurantsByArea(area string, latitude string, longitude string) []model.SimpleRestaurant {
	restaurants := []model.SimpleRestaurant{}
	c.DB.Raw(`SELECT UuidFromBin(r.id) AS id,
	               r.url,
	               r.name,
	               r.genre,
	               r.tel,
	               r.business_day_info,
	               r.address,
	               r.latitude,
	               r.longitude,
		           GetDistance(r.latitude, r.longitude, ` + latitude + `, ` + longitude + `) AS distance,
	               r.area,
	               COUNT(p.restaurant_id) AS photo_count,
                   (
				       REPLACE(JSON_EXTRACT(r.business_day_info, CONCAT('$.', DAYOFWEEK(CURDATE()), ".Start")), '"', '') <= DATE_FORMAT(CONVERT_TZ(NOW(), '+00:00', '+09:00'), '%H%i')
                           AND
					   REPLACE(JSON_EXTRACT(r.business_day_info, CONCAT('$.', DAYOFWEEK(CURDATE()), ".End")), '"', '') >= DATE_FORMAT(CONVERT_TZ(NOW(), '+00:00', '+09:00'), '%H%i')
				   ) AS is_open
                  FROM restaurants AS r
                  LEFT JOIN photos AS p
                    ON r.id = p.restaurant_id
                 WHERE r.is_closed = 0
		           AND r.area = '` + area + `'
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
                 WHERE r.is_closed = 0
		   AND rg.genre_id = ` + genreID + `
                   AND REPLACE(JSON_EXTRACT(r.business_day_info, CONCAT('$.', DAYOFWEEK(CURDATE()), ".Start")), '"', '') <= DATE_FORMAT(CONVERT_TZ(NOW(), '+00:00', '+09:00'), '%H%i')
                   AND REPLACE(JSON_EXTRACT(r.business_day_info, CONCAT('$.', DAYOFWEEK(CURDATE()), ".End")), '"', '') >= DATE_FORMAT(CONVERT_TZ(NOW(), '+00:00', '+09:00'), '%H%i')
                 GROUP BY r.id
                 ORDER BY photo_count DESC`).Scan(&restaurants)

	return restaurants
}

// GetRestaurantsByGenreID returns the open restaurants for the specified genre.
func (c RestaurantRepository) GetRestaurantsByGenreID(genreID string, latitude string, longitude string) []model.SimpleRestaurant {
	restaurants := []model.SimpleRestaurant{}
	c.DB.Raw(`SELECT UuidFromBin(r.id) AS id,
	               r.url,
	               r.name,
	               r.genre,
	               r.tel,
	               r.business_day_info,
	               r.address,
	               r.latitude,
	               r.longitude,
		       GetDistance(r.latitude, r.longitude, ` + latitude + `, ` + longitude + `) AS distance,
	               r.area,
                   (
				       REPLACE(JSON_EXTRACT(r.business_day_info, CONCAT('$.', DAYOFWEEK(CURDATE()), ".Start")), '"', '') <= DATE_FORMAT(CONVERT_TZ(NOW(), '+00:00', '+09:00'), '%H%i')
                           AND
					   REPLACE(JSON_EXTRACT(r.business_day_info, CONCAT('$.', DAYOFWEEK(CURDATE()), ".End")), '"', '') >= DATE_FORMAT(CONVERT_TZ(NOW(), '+00:00', '+09:00'), '%H%i')
				   ) AS is_open
                  FROM restaurants AS r
                  JOIN restaurant_genres AS rg
                    ON r.id = rg.restaurant_id
                 WHERE r.is_closed = 0
		   AND rg.genre_id = ` + genreID + `
                 ORDER BY distance ASC`).Scan(&restaurants)

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
                 WHERE r.is_closed = 0
		   AND rd.drink_id = ` + drinkID + `
                   AND REPLACE(JSON_EXTRACT(r.business_day_info, CONCAT('$.', DAYOFWEEK(CURDATE()), ".Start")), '"', '') <= DATE_FORMAT(CONVERT_TZ(NOW(), '+00:00', '+09:00'), '%H%i')
                   AND REPLACE(JSON_EXTRACT(r.business_day_info, CONCAT('$.', DAYOFWEEK(CURDATE()), ".End")), '"', '') >= DATE_FORMAT(CONVERT_TZ(NOW(), '+00:00', '+09:00'), '%H%i')
                 GROUP BY r.id
                 ORDER BY photo_count DESC`).Scan(&restaurants)

	return restaurants
}

// GetRestaurantsByDrinkID returns the open restaurants for the specified drink.
func (c RestaurantRepository) GetRestaurantsByDrinkID(drinkID string, latitude string, longitude string) []model.SimpleRestaurant {
	restaurants := []model.SimpleRestaurant{}
	c.DB.Raw(`SELECT UuidFromBin(r.id) AS id,
	               r.url,
	               r.name,
	               r.genre,
	               r.tel,
	               r.business_day_info,
	               r.address,
	               r.latitude,
	               r.longitude,
				GetDistance(r.latitude, r.longitude, ` + latitude + `, ` + longitude + `) AS distance,
	               r.area,
                   (
				       REPLACE(JSON_EXTRACT(r.business_day_info, CONCAT('$.', DAYOFWEEK(CURDATE()), ".Start")), '"', '') <= DATE_FORMAT(CONVERT_TZ(NOW(), '+00:00', '+09:00'), '%H%i')
                           AND
					   REPLACE(JSON_EXTRACT(r.business_day_info, CONCAT('$.', DAYOFWEEK(CURDATE()), ".End")), '"', '') >= DATE_FORMAT(CONVERT_TZ(NOW(), '+00:00', '+09:00'), '%H%i')
				   ) AS is_open
                  FROM restaurants AS r
                  JOIN restaurant_drinks AS rd
                    ON r.id = rd.restaurant_id
		   AND rd.drink_id = ` + drinkID + `
                 ORDER BY distance ASC`).Scan(&restaurants)

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

// GetRestaurantsByDishID returns the open restaurants which have the specified dish.
func (c RestaurantRepository) GetRestaurantsByDishID(dishID string, latitude string, longitude string) []model.SimpleRestaurant {
	restaurants := []model.SimpleRestaurant{}
	c.DB.Raw(`SELECT UuidFromBin(r.id) AS id,
	               r.url,
	               r.name,
	               r.genre,
	               r.tel,
	               r.business_day_info,
	               r.address,
	               r.latitude,
	               r.longitude,
				GetDistance(r.latitude, r.longitude, ` + latitude + `, ` + longitude + `) AS distance,
	               r.area,
                   (
				       REPLACE(JSON_EXTRACT(r.business_day_info, CONCAT('$.', DAYOFWEEK(CURDATE()), ".Start")), '"', '') <= DATE_FORMAT(CONVERT_TZ(NOW(), '+00:00', '+09:00'), '%H%i')
                           AND
					   REPLACE(JSON_EXTRACT(r.business_day_info, CONCAT('$.', DAYOFWEEK(CURDATE()), ".End")), '"', '') >= DATE_FORMAT(CONVERT_TZ(NOW(), '+00:00', '+09:00'), '%H%i')
				   ) AS is_open
	          FROM dishes AS d
	          JOIN rankings AS rk
                    ON rk.dish_id = d.id
                  JOIN photos AS p
                    ON p.id = rk.photo_id
                  JOIN restaurants AS r
                    ON r.id = p.restaurant_id
                 WHERE r.is_closed = 0
                   AND d.id = ` + dishID + `
                 ORDER BY distance ASC`).Scan(&restaurants)

	return restaurants
}

// GetOpenRestaurantCount returns the number of open restaurants.
func (c RestaurantRepository) GetOpenRestaurantCount() []model.RestaurantCount {
	restaurantCounts := []model.RestaurantCount{}
	c.DB.Raw(`SELECT a.value AS area,
                     SUM(
                         REPLACE(JSON_EXTRACT(r.business_day_info, CONCAT('$.', DAYOFWEEK(CURDATE()), ".Start")), '"', '') <= DATE_FORMAT(CONVERT_TZ(NOW(), '+00:00', '+09:00'), '%H%i')
                         AND
                         REPLACE(JSON_EXTRACT(r.business_day_info, CONCAT('$.', DAYOFWEEK(CURDATE()), ".End")), '"', '') >= DATE_FORMAT(CONVERT_TZ(NOW(), '+00:00', '+09:00'), '%H%i')
                     ) AS count
                FROM areas AS a
               RIGHT JOIN restaurants AS r
                  ON a.value = r.area
               WHERE r.is_closed = 0
               GROUP BY area
               ORDER BY count DESC`).Scan(&restaurantCounts)

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
