package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xarick/golang-memcached-example/cache"
	"github.com/xarick/golang-memcached-example/db"
	"github.com/xarick/golang-memcached-example/utils"
)

func CreateShortURL(c *gin.Context) {
	longURL := c.PostForm("url")
	if longURL == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "URL kiriting!"})
		return
	}

	shortCode := utils.GenerateShortCode()

	if err := db.SaveURL(shortCode, longURL); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Bazaga yozilmadi"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"short_url": "http://127.0.0.1:8040/sh/" + shortCode})
}

func HandleShortURL(c *gin.Context) {
	shortCode := c.Param("shortCode")

	longURL, err := cache.GetFromCache(shortCode)
	if err != nil {
		fmt.Println("Bazadan qidirish boshlandi... Cache error: ", err)

		longURL, err = db.GetURL(shortCode)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "URL topilmadi"})
			return
		}
		err = cache.SetToCache(shortCode, longURL, 300)
		if err != nil {
			fmt.Println("Cache ga saqlashda xatolik, error: ", err)
		}
	}

	c.JSON(http.StatusOK, gin.H{"long_url": longURL})
}
