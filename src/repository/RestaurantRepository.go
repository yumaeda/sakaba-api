package repository

import (
	"github.com/google/uuid"
	"github.com/yumaeda/sakaba-api/src/infrastructure"
	"github.com/yumaeda/sakaba-api/src/model"
	"gorm.io/gorm"
)

// RestaurantRepository is responsible for reading from and writing to DB Table `restaurants`.
type RestaurantRepository struct {
	TiDB *gorm.DB
}

// GetRestaurants returns all the restaurants.
func (c RestaurantRepository) GetRestaurants() []model.Restaurant {
	restaurants := []model.Restaurant{}
	c.TiDB.Raw(`SELECT BIN_TO_UUID(id, 1) AS id,
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
	c.TiDB.Raw(`SELECT BIN_TO_UUID(r.id, 1) AS id,
	               r.url,
	               r.name,
	               r.genre,
	               r.tel,
	               r.business_day_info,
	               r.address,
	               r.latitude,
	               r.longitude,` +
		infrastructure.GetDistanceSQL("r.latitude", "r.longitude", latitude, longitude) + ` AS distance,
	               r.area,
	               COUNT(p.restaurant_id) AS photo_count,
                   (
				       REPLACE(JSON_EXTRACT(r.business_day_info, CONCAT('$."', DAYOFWEEK(CURDATE()), '".Start')), '"', '') <= DATE_FORMAT(CONVERT_TZ(NOW(), '+00:00', '+09:00'), '%H%i')
                           AND
					   REPLACE(JSON_EXTRACT(r.business_day_info, CONCAT('$."', DAYOFWEEK(CURDATE()), '".End')), '"', '') >= DATE_FORMAT(CONVERT_TZ(NOW(), '+00:00', '+09:00'), '%H%i')
				   ) AS is_open
                  FROM restaurants AS r
                  LEFT JOIN photos AS p
                    ON r.id = p.restaurant_id
                 WHERE r.is_closed = 0
		           AND r.area = '` + area + `'
                 GROUP BY r.id
                 ORDER BY distance ASC`).Scan(&restaurants)

	return restaurants
}

// GetRestaurantsByGenreID returns the open restaurants for the specified genre.
func (c RestaurantRepository) GetRestaurantsByGenreID(genreID string, latitude string, longitude string) []model.SimpleRestaurant {
	restaurants := []model.SimpleRestaurant{}
	c.TiDB.Raw(`SELECT BIN_TO_UUID(r.id, 1) AS id,
	               r.url,
	               r.name,
	               r.genre,
	               r.tel,
	               r.business_day_info,
	               r.address,
	               r.latitude,
	               r.longitude,` +
		infrastructure.GetDistanceSQL("r.latitude", "r.longitude", latitude, longitude) + ` AS distance,
	               r.area,
	               COUNT(p.restaurant_id) AS photo_count,
                   (
                       REPLACE(JSON_EXTRACT(r.business_day_info, CONCAT('$."', DAYOFWEEK(CURDATE()), '".Start')), '"', '') <= DATE_FORMAT(CONVERT_TZ(NOW(), '+00:00', '+09:00'), '%H%i')
                           AND
                       REPLACE(JSON_EXTRACT(r.business_day_info, CONCAT('$."', DAYOFWEEK(CURDATE()), '".End')), '"', '') >= DATE_FORMAT(CONVERT_TZ(NOW(), '+00:00', '+09:00'), '%H%i')
				   ) AS is_open
                  FROM restaurants AS r
                  JOIN restaurant_genres AS rg
                    ON r.id = rg.restaurant_id
                  LEFT JOIN photos AS p
                    ON r.id = p.restaurant_id
                 WHERE r.is_closed = 0
                   AND rg.genre_id = ` + genreID + `
                 GROUP BY r.id
                 ORDER BY distance ASC`).Scan(&restaurants)

	return restaurants
}

// GetRestaurantsByDrinkID returns the open restaurants for the specified drink.
func (c RestaurantRepository) GetRestaurantsByDrinkID(drinkID string, latitude string, longitude string) []model.SimpleRestaurant {
	restaurants := []model.SimpleRestaurant{}
	c.TiDB.Raw(`SELECT BIN_TO_UUID(r.id, 1) AS id,
	               r.url,
	               r.name,
	               r.genre,
	               r.tel,
	               r.business_day_info,
	               r.address,
	               r.latitude,
	               r.longitude,` +
		infrastructure.GetDistanceSQL("r.latitude", "r.longitude", latitude, longitude) + ` AS distance,
	               r.area,
	               COUNT(p.restaurant_id) AS photo_count,
                   (
				       REPLACE(JSON_EXTRACT(r.business_day_info, CONCAT('$."', DAYOFWEEK(CURDATE()), '".Start')), '"', '') <= DATE_FORMAT(CONVERT_TZ(NOW(), '+00:00', '+09:00'), '%H%i')
                           AND
					   REPLACE(JSON_EXTRACT(r.business_day_info, CONCAT('$."', DAYOFWEEK(CURDATE()), '".End')), '"', '') >= DATE_FORMAT(CONVERT_TZ(NOW(), '+00:00', '+09:00'), '%H%i')
				   ) AS is_open
                  FROM restaurants AS r
                  JOIN restaurant_drinks AS d
                    ON r.id = d.restaurant_id
                  LEFT JOIN photos AS p
                    ON r.id = p.restaurant_id
                 WHERE r.is_closed = 0
                   AND d.drink_id = ` + drinkID + `
                 GROUP BY r.id
                 ORDER BY distance ASC`).Scan(&restaurants)

	return restaurants
}

// GetRestaurantsByDishID returns the open restaurants which have the specified dish.
func (c RestaurantRepository) GetRestaurantsByDishID(dishID string, latitude string, longitude string) []model.SimpleRestaurant {
	restaurants := []model.SimpleRestaurant{}
	c.TiDB.Raw(`SELECT BIN_TO_UUID(r.id, 1) AS id,
	               r.url,
	               r.name,
	               r.genre,
	               r.tel,
	               r.business_day_info,
	               r.address,
	               r.latitude,
	               r.longitude,` +
		infrastructure.GetDistanceSQL("r.latitude", "r.longitude", latitude, longitude) + ` AS distance,
	               r.area,
	               COUNT(p.restaurant_id) AS photo_count,
                   (
				       REPLACE(JSON_EXTRACT(r.business_day_info, CONCAT('$."', DAYOFWEEK(CURDATE()), '".Start')), '"', '') <= DATE_FORMAT(CONVERT_TZ(NOW(), '+00:00', '+09:00'), '%H%i')
                           AND
					   REPLACE(JSON_EXTRACT(r.business_day_info, CONCAT('$."', DAYOFWEEK(CURDATE()), '".End')), '"', '') >= DATE_FORMAT(CONVERT_TZ(NOW(), '+00:00', '+09:00'), '%H%i')
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
                 GROUP BY r.id
                 ORDER BY distance ASC`).Scan(&restaurants)

	return restaurants
}

// GetOpenRestaurantCount returns the number of open restaurants within 3km.
func (c RestaurantRepository) GetOpenRestaurantCount(latitude string, longitude string) []model.RestaurantCount {
	restaurantCounts := []model.RestaurantCount{}
	c.TiDB.Raw(`SELECT a.value AS area,
	                   a.name AS name,
                       SUM(
                         REPLACE(JSON_EXTRACT(r.business_day_info, CONCAT('$."', DAYOFWEEK(CURDATE()), '".Start')), '"', '') <= DATE_FORMAT(CONVERT_TZ(NOW(), '+00:00', '+09:00'), '%H%i')
                         AND
                         REPLACE(JSON_EXTRACT(r.business_day_info, CONCAT('$."', DAYOFWEEK(CURDATE()), '".End')), '"', '') >= DATE_FORMAT(CONVERT_TZ(NOW(), '+00:00', '+09:00'), '%H%i')
                       ) AS count
                  FROM areas AS a
                 RIGHT JOIN restaurants AS r
                    ON a.value = r.area
                 WHERE r.is_closed = 0
				   AND (` + infrastructure.GetDistanceSQL("r.latitude", "r.longitude", latitude, longitude) + `) < 1
                 GROUP BY a.value, a.name
                 ORDER BY count DESC`).Scan(&restaurantCounts)

	return restaurantCounts
}

// AddRestaurant adds a new restaurant.
func (c RestaurantRepository) AddRestaurant(URL string, name string, genre string, tel string, businessDayInfo string, address string, latitude string, longitude string, area string) (string, error) {
	id := uuid.New().String()
	restaurant := model.Restaurant{
		ID:              infrastructure.UUIDToBin(id),
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
	dbError := c.TiDB.Create(&restaurant).Error

	return id, dbError
}
