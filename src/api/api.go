package api

import (
	"fmt"

	"github.com/Amirrezaghanavati/golang-web-api/api/routers"
	"github.com/Amirrezaghanavati/golang-web-api/config"
	"github.com/gin-gonic/gin"
)

func InitServer(){
	cfg := config.GetConfig()

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	v1 := r.Group("/api/v1")
	{
		routers.Health(v1.Group("/health"))
	}
	
	r.Run(fmt.Sprintf(":%s", cfg.Server.Port))
}