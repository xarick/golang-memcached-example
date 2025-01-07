package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/xarick/golang-memcached-example/cache"
	"github.com/xarick/golang-memcached-example/db"
	"github.com/xarick/golang-memcached-example/handlers"
)

// Config
const (
	dbSource    = "postgres://username:password@localhost:5432/shortener_db?sslmode=disable"
	memcachedIP = "127.0.0.1:11211"
)

func main() {
	// DB va Cache ulash
	if err := db.InitDB(dbSource); err != nil {
		log.Fatalf("DB ulanmadi: %v", err)
	}
	cache.InitCache(memcachedIP)

	r := gin.Default()
	r.POST("/create", handlers.CreateShortURL)
	r.GET("/:shortCode", handlers.HandleShortURL)

	log.Println("Server http://localhost:8040 da ishlamoqda...")
	if err := r.Run(":8040"); err != nil {
		log.Fatalf("Serverda xatolik: %v", err)
	}
}
