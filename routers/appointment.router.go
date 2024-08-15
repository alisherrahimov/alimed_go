package routers

import "github.com/gin-gonic/gin"

func AppointmentRoute(router *gin.RouterGroup) {
	appointmentV1 := router.Group("appointment")
	{
		appointmentV1.GET("/login", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"token": "token",
			})
		})
	}

}
