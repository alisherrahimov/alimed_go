package routers

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type ROUTES struct {
	Url         string
	Handler     func(gin *gin.Context, db *sql.DB)
	TypeRequest string
	Middleware  gin.HandlerFunc
}

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)

func Route(router *gin.Engine, database *sql.DB) {
	r := router.Group("api/v1")
	ChatRoute(r)
	AppointmentRoute(r)
	DoctorRoute(r)
	UserRoute(r, database)
	AuthRoute(r, database)
}
