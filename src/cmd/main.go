package main

import (
	"github.com/Amirrezaghanavati/golang-web-api/api"
	"github.com/Amirrezaghanavati/golang-web-api/config"
	"github.com/Amirrezaghanavati/golang-web-api/data/cache"
)

func main() {
	cfg := config.GetConfig()
	cache.InitRedis(cfg)
	defer cache.CloseRedis()
	api.InitServer(cfg)
}
