package routers

import (
	"alimed_go/pkg/controllers"
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func AuthRoute(r *gin.RouterGroup, database *sql.DB) {
	authV1 := r.Group("auth")

	routes := []ROUTES{
		{
			Url:         "/login",
			Handler:     controllers.Login,
			TypeRequest: POST,
			// Middleware:  middlewares.Validator(models.User{}),
		},
		{
			Url:         "/login/by-email",
			Handler:     controllers.LoginByEmail,
			TypeRequest: POST,
		},
		{
			Url:         "/otp",
			Handler:     controllers.VerifyOTP,
			TypeRequest: POST,
		},
		{
			Url:         "/ask-data",
			Handler:     controllers.AskSomeData,
			TypeRequest: PUT,
		},
		{
			Url:         "/delete-data",
			Handler:     controllers.DeleteSomeData,
			TypeRequest: DELETE,
			// Middleware:  middlewares.Validator(models.LoginRequest{}),
		},
	}

	for _, route := range routes {
		if route.TypeRequest == GET {
			authV1.GET(route.Url, route.Middleware, func(ctx *gin.Context) {
				route.Handler(ctx, database)
			})
		}
		if route.TypeRequest == POST {
			authV1.POST(route.Url, route.Middleware, func(ctx *gin.Context) {
				route.Handler(ctx, database)
			})
		}
		if route.TypeRequest == PUT {
			authV1.PUT(route.Url, route.Middleware, func(ctx *gin.Context) {
				route.Handler(ctx, database)
			})
		}
		if route.TypeRequest == DELETE {
			authV1.DELETE(route.Url, route.Middleware, func(ctx *gin.Context) {
				route.Handler(ctx, database)
			})
		}
	}

}
