package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

func CORS(allowOrigins string) gin.HandlerFunc {
	policy := cors.New(cors.Options{
		AllowedOrigins: []string{allowOrigins},
		AllowedHeaders: []string{"*"},
	})

	return func(ctx *gin.Context) {
		continued := false
		policy.ServeHTTP(ctx.Writer, ctx.Request, func(http.ResponseWriter, *http.Request) {
			continued = true
			ctx.Next()
		})
		if !continued {
			ctx.Abort()
		}
	}
}
