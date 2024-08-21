package middlewares

import (
	"alimed_go/pkg/configs"
	"alimed_go/pkg/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		head := ctx.GetHeader("Authorization")
		token := head[len("Bearer "):]
		if head == "" {
			configs.Response(ctx, configs.Res{
				Data:    nil,
				Error:   "Unauthorized",
				Success: false,
			}, http.StatusUnauthorized)
			ctx.Abort()
			return
		} else {
			user, err := helper.VerifyToken(token)
			if err != nil {
				configs.Response(ctx, configs.Res{
					Data:    nil,
					Error:   "Unauthorized",
					Success: false,
				}, http.StatusUnauthorized)
				ctx.Abort()
				return
			} else {
				ctx.Set("user", user)
				ctx.Next()
			}
		}
	}
}

func CustomValidator() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
