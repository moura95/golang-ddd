package middleware

import (
	"github.com/gin-gonic/gin"
)

func ActivityMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
	}
}
