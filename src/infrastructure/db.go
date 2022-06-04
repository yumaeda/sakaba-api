package infrastructure

import (
	"encoding/json"
	"fmt"
	"os"
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

// ConnectToDB connects to the Database based on the configuration and returns pointer to the connection.
func ConnectToDB() *gorm.DB {
	secretManagerJSON := os.Getenv("APP_CONFIG_JSON")
	dbConfig := DatabaseConfig{}
	json.Unmarshal([]byte(secretManagerJSON), &dbConfig)

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the specified database.")
	}

	sqlDB, sqlDBErr := db.DB()
	if sqlDBErr != nil {
		panic("Failed to retrieve inner database.")
	}
	sqlDB.SetMaxIdleConns(0)
	sqlDB.SetMaxOpenConns(3)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}

// CloseDB closes the current Database connection.
func CloseDB(db *gorm.DB) {
	targetDB, err := db.DB()
	if err == nil {
		targetDB.Close()
	}
}

// UUIDToBin converts UUID to Binary format.
func UUIDToBin(uuid string) string {
	db := ConnectToDB()
	rows, err := db.Raw("SELECT UuidToBin('" + uuid + "')").Rows()
	if err != nil {
		panic("UuidToBin() is not specified")
	}

	defer rows.Close()
	var bin string
	for rows.Next() {
		if err := rows.Scan(&bin); err != nil {
			panic("Error occurrs while retrieving a value")
		}
	}
	CloseDB(db)

	return bin
}
