package main

import (
	"log"

	"github.com/Amirrezaghanavati/golang-web-api/api"
	"github.com/Amirrezaghanavati/golang-web-api/config"
	"github.com/Amirrezaghanavati/golang-web-api/data/cache"
	"github.com/Amirrezaghanavati/golang-web-api/data/cache/db"
)

func main() {
	cfg := config.GetConfig()

	if err := cache.InitRedis(cfg); err != nil {
		log.Fatalf("initialize Redis: %v", err)
	}
	defer cache.CloseRedis()

	if err := db.InitDb(cfg); err != nil {
		log.Fatalf("initialize PostgreSQL: %v", err)
	}
	defer db.CloseDb()

	api.InitServer(cfg)
}
