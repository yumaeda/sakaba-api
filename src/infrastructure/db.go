package infrastructure

import (
    "encoding/json"
    "fmt"
    "os"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

type Database struct {
    DB *gorm.DB
}

type DatabaseConfig struct {
    Password string `json:"db.password"`
    Host string `json:"db.host"`
    Name string `json:"db.name"`
    User string `json:"db.user"`
}

/*
 * TODO: Retrieve DB credentials from AWS Secret Manager.
 * https://dev.to/aws-builders/aws-secret-manager-with-a-simple-golang-ecs-task-2d98
 */
func NewDatabase() Database {
    springAppJson := os.Getenv("SPRING_APPLICATION_JSON")
    dbConfig := DatabaseConfig{} 
    json.Unmarshal([]byte(springAppJson), &dbConfig)

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
    fmt.Println("Database connection established.")
    return Database {
        DB: db,
    }
}
