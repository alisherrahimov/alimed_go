package routers

import "github.com/gin-gonic/gin"

func ChatRoute(router *gin.RouterGroup) {
	chatV1 := router.Group("chat")
	{
		chatV1.GET("/login", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"token": "token",
			})
		})
	}
}
