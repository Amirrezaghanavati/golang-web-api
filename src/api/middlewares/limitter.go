package middlewares

import (
	"net/http"

	"github.com/Amirrezaghanavati/golang-web-api/api/helper"
	"github.com/didip/tollbooth"
	"github.com/gin-gonic/gin"
)

func LimitByRequest() gin.HandlerFunc {
	lmt := tollbooth.NewLimiter(1, nil)
	return func(ctx *gin.Context) {
		httpError := tollbooth.LimitByRequest(lmt, ctx.Writer, ctx.Request)

		if httpError != nil {
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests,
				helper.GenerateBaseResponseWithError(
					nil, false, -100, httpError))
			return
		} else {
			ctx.Next()
		}
	}
}
