package main

import (
	"log"

	"github.com/xarick/golang-memcached-example/cache"
	"github.com/xarick/golang-memcached-example/config"
	"github.com/xarick/golang-memcached-example/db"
	"github.com/xarick/golang-memcached-example/routes"
)

func main() {
	cfg := config.LoadConfig()

	if err := db.InitDB(cfg.DBSource); err != nil {
		log.Fatalf("DB ulanmadi: %v", err)
	}

	cache.InitCache(cfg.MemcachedIP)

	r := routes.SetupRouter()

	if err := r.Run(cfg.RunPort); err != nil {
		log.Fatalf("Serverda xatolik: %v", err)
	}
}
