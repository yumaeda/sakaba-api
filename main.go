package main

import (
    "fmt"
    "os"
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    spring_app_json = os.Getenv("SPRING_APPLICATION_JSON")
    engine:= gin.Default()
    engine.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "hello world: " + spring_app_json,
        })
    })
    engine.Run(":8080")
}
