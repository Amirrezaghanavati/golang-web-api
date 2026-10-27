package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (healthHandler *HealthHandler) Health(c *gin.Context){
	c.JSON(http.StatusOK, "working")
}