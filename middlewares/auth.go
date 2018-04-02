package middlewares

import (
	"github.com/gin-gonic/gin"
)

// Auth -
func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
	}
}
