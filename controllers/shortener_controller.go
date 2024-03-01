package controllers

import (
	"go-url-shortener/models"
	"time"

	"github.com/gin-gonic/gin"
)


func CreateShortURL(c *gin.Context) {
    var url models.URL

    c.BindJSON(&url)

    url.GenerateShortURL()

    models.CreateURL(&url)

    c.JSON(200, gin.H{"short_url": url.ShortURL, "Alias": url.Alias})
}

func RedirectShortURL(c *gin.Context) {
    shortURL := c.Param("short_url")
    url, err := models.GetURLByShortURL(shortURL)
    if err != nil {
        c.JSON(404, gin.H{"error": "URL not found"})
        return
    }
    
    url.AccessCount++
    now := time.Now()
    url.LastAccessed = &now
    url.AccessPlace = c.ClientIP()
    models.UpdateURL(&url)
    
    c.Redirect(301, url.LongURL)
}



func GetURLStatistics(c *gin.Context) {
    shortURL := c.Param("short_url")
    url, err := models.GetURLByShortURL(shortURL)
    if err != nil {
        c.JSON(404, gin.H{"error": "URL not found"})
        return
    }

    c.JSON(200, gin.H{
        "short_url":     url.ShortURL,
        "long_url":      url.LongURL,
        "access_count":  url.AccessCount,
        "last_accessed": url.LastAccessed,
        "access_place":  url.AccessPlace,
    })
}
