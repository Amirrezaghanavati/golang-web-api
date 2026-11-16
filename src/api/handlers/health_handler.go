package handlers

import (
	"net/http"

	"github.com/Amirrezaghanavati/golang-web-api/api/helper"
	"github.com/gin-gonic/gin"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (healthHandler *HealthHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, helper.GenerateBaseResponse("working", true, 0))
}
