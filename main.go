package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "sakaba.link/api/infrastructure"
)

func main() {
    router := gin.Default()
    router.GET("/", func(c *gin.Context) {
        infrastructure.NewDatabase()
        c.JSON(http.StatusOK, gin.H{ "message": "hello world" })
    })
    router.Run(":8080")
}
