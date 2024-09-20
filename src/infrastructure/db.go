package infrastructure

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DatabaseConfig is a configuration of the Database connection.
type DatabaseConfig struct {
	Password string `json:"db.password"`
	Host     string `json:"db.host"`
	Name     string `json:"db.name"`
	User     string `json:"db.user"`
}

// ConnectToDB connects to the TiDB based on the configuration and returns pointer to the connection.
func ConnectToDB() (*gorm.DB, func(), error) {
	secretManagerJSON := os.Getenv("TIDB_CONFIG_JSON")
	dbConfig := DatabaseConfig{}
	json.Unmarshal([]byte(secretManagerJSON), &dbConfig)

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:4000)/%s?tls=true&charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}

	sqlDB, sqlDBErr := db.DB()
	if sqlDBErr != nil {
		return nil, nil, sqlDBErr
	}
	sqlDB.SetMaxIdleConns(0)
	sqlDB.SetMaxOpenConns(6)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, func() {
		sqlDB.Close()
	}, nil
}

// UUIDToBin converts UUID to Binary format.
func UUIDToBin(uuid string) string {
	db, closer, err := ConnectToDB()
	if err != nil {
		panic(err.Error())
	}
	defer closer()

	rows, err := db.Raw("SELECT UUID_TO_BIN('" + uuid + "', 1)").Rows()
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var bin string
	for rows.Next() {
		if err := rows.Scan(&bin); err != nil {
			panic("Error occurrs while retrieving a value")
		}
	}

	return bin
}

// Get SQL statement for calculating the distance between two points
func GetDistanceSQL(lat1 string, lon1 string, lat2 string, lon2 string) string {
	template := `6371.009 * 2 * ATAN2(
		SQRT(
			SIN(({{lat1}} - {{lat2}}) * PI() / 180 / 2) *
			SIN(({{lat1}} - {{lat2}}) * PI() / 180 / 2) +
			SIN(({{lon1}} - {{lon2}}) * PI() / 180 / 2) *
			SIN(({{lon1}} - {{lon2}}) * PI() / 180 / 2) *
			COS({{lat2}} * PI() / 180) *
			COS({{lat1}} * PI() / 180)
		),
		SQRT(
			1 - (
				SIN(({{lat1}} - {{lat2}}) * PI() / 180 / 2) *
				SIN(({{lat1}} - {{lat2}}) * PI() / 180 / 2) +
				SIN(({{lon1}} - {{lon2}}) * PI() / 180 / 2) *
				SIN(({{lon1}} - {{lon2}}) * PI() / 180 / 2) *
				COS({{lat2}} * PI() / 180) *
				COS({{lat1}} * PI() / 180)
			)
		)
	)`

	result := strings.ReplaceAll(template, "{{lat1}}", lat1)
	result = strings.ReplaceAll(result, "{{lon1}}", lon1)
	result = strings.ReplaceAll(result, "{{lat2}}", lat2)
	result = strings.ReplaceAll(result, "{{lon2}}", lon2)

	return result
}
