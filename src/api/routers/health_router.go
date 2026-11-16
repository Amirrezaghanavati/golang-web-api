package routers

import (
	"github.com/Amirrezaghanavati/golang-web-api/api/handlers"
	"github.com/gin-gonic/gin"
)

func HealthRouter(r *gin.RouterGroup) {
	newHealthHandler := handlers.NewHealthHandler()

	r.GET("/", newHealthHandler.Health)
}
