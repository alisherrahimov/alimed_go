package routers

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func UserRoute(router *gin.RouterGroup, database *sql.DB) {

	userV1 := router.Group("users")

	routes := []ROUTES{}

	for _, route := range routes {
		if route.TypeRequest == GET {
			userV1.GET(route.Url, func(ctx *gin.Context) {
				route.Handler(ctx, database)
			})
		}
		if route.TypeRequest == POST {
			userV1.POST(route.Url, func(ctx *gin.Context) {
				route.Handler(ctx, database)
			})
		}
		if route.TypeRequest == PUT {
			userV1.PUT(route.Url, func(ctx *gin.Context) {
				route.Handler(ctx, database)
			})
		}
		if route.TypeRequest == DELETE {
			userV1.DELETE(route.Url, func(ctx *gin.Context) {
				route.Handler(ctx, database)
			})
		}
	}

}
