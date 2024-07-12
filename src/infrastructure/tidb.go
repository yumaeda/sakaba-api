package infrastructure

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// TiDBConfig is a configuration of the TiDB connection.
type TiDBConfig struct {
	Password string `json:"tidb.password"`
	Host     string `json:"tidb.host"`
	Name     string `json:"tidb.name"`
	User     string `json:"tidb.user"`
	Port     string `json:"tidb.port"`
}

// ConnectToTiDB connects to the TiDB based on the configuration and returns pointer to the connection.
func ConnectToTiDB() (*gorm.DB, func(), error) {
	secretManagerJSON := os.Getenv("TIDB_CONFIG_JSON")
	dbConfig := TiDBConfig{}
	json.Unmarshal([]byte(secretManagerJSON), &dbConfig)

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
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
