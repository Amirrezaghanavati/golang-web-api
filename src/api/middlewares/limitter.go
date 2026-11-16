package middlewares

import (
	"net/http"

	"github.com/didip/tollbooth"
	"github.com/gin-gonic/gin"
)

func LimitByRequest() gin.HandlerFunc {
	lmt := tollbooth.NewLimiter(1, nil)
	return func(ctx *gin.Context) {
		httpError := tollbooth.LimitByRequest(lmt, ctx.Writer, ctx.Request)

		if httpError != nil {
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": httpError.Error(),
			})
			return
		} else {
			ctx.Next()
		}
	}
}
