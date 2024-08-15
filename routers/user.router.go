package routers

import (
	"alimed_go/controllers"
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)

type ROUTES struct {
	Url         string
	Handler     func(gin *gin.Context, db *sql.DB)
	TypeRequest string
}

func UserRoute(router *gin.RouterGroup, database *sql.DB) {

	userV1 := router.Group("users")
	userV1.GET("/get", func(ctx *gin.Context) {
		controllers.Get(ctx, database)
	})
	// routes := []ROUTES{
	// 	{
	// 		Url:         "/get",
	// 		Handler:     controllers.Get,
	// 		TypeRequest: "GET",
	// 	},
	// 	{
	// 		Url:         "/get/:id",
	// 		Handler:     controllers.GetById,
	// 		TypeRequest: "GET",
	// 	},
	// 	{
	// 		Url:         "/add",
	// 		Handler:     controllers.Post,
	// 		TypeRequest: "POST",
	// 	},
	// 	{
	// 		Url:         "/delete/:id",
	// 		Handler:     controllers.Delete,
	// 		TypeRequest: "DELETE",
	// 	},
	// 	{
	// 		Url:         "/put/:id",
	// 		Handler:     controllers.Put,
	// 		TypeRequest: "PUT",
	// 	},
	// }

	// for _, route := range routes {
	// 	if route.TypeRequest == GET {
	// 		userV1.GET(route.Url, func(ctx *gin.Context) {
	// 			route.Handler(ctx, database)
	// 		})
	// 	}
	// 	if route.TypeRequest == POST {
	// 		userV1.POST(route.Url, func(ctx *gin.Context) {
	// 			route.Handler(ctx, database)
	// 		})
	// 	}
	// 	if route.TypeRequest == PUT {
	// 		userV1.PUT(route.Url, func(ctx *gin.Context) {
	// 			route.Handler(ctx, database)
	// 		})
	// 	}
	// 	if route.TypeRequest == DELETE {
	// 		userV1.DELETE(route.Url, func(ctx *gin.Context) {
	// 			route.Handler(ctx, database)
	// 		})
	// 	}
	// }

}
