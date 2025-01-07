package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xarick/golang-memcached-example/cache"
	"github.com/xarick/golang-memcached-example/db"
)

// Qisqa URL yaratish
func CreateShortURL(c *gin.Context) {
	longURL := c.PostForm("url")
	if longURL == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "URL kiriting!"})
		return
	}

	// Tasodifiy kod yaratish
	shortCode := generateShortCode()

	// Bazaga yozish
	if err := db.SaveURL(shortCode, longURL); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Bazaga yozilmadi"})
		return
	}

	// Javob qaytarish
	c.JSON(http.StatusOK, gin.H{"short_url": "http://localhost:8080/" + shortCode})
}

// Qisqa URL orqali yo'naltirish
func HandleShortURL(c *gin.Context) {
	shortCode := c.Param("shortCode")

	// Memcached orqali qidirish
	longURL, err := cache.GetFromCache(shortCode)
	if err != nil {
		// Bazadan qidirish
		longURL, err = db.GetURL(shortCode)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "URL topilmadi"})
			return
		}
		cache.SetToCache(shortCode, longURL, 3600)
	}

	c.Redirect(http.StatusFound, longURL)
}

// Qisqa kod yaratish
func generateShortCode() string {
	t := time.Now().UnixNano()
	return string([]byte{byte(t & 0xff), byte((t >> 8) & 0xff), byte((t >> 16) & 0xff)})
}
