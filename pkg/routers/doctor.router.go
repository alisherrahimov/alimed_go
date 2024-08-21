package routers

import "github.com/gin-gonic/gin"

func DoctorRoute(router *gin.RouterGroup) {
	doctorV1 := router.Group("doctor")
	{
		doctorV1.GET("/login", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"token": "token",
			})
		})
	}
}
