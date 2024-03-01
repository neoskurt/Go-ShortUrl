package main

import (
    "go-url-shortener/models"
    "github.com/gin-gonic/gin"

    "go-url-shortener/routes"
)

func main() {
    models.InitDB()
    defer models.CloseDB()

    r := gin.Default()

    router  := routes.InitializeRoutes(r)
    router.Run()
}

