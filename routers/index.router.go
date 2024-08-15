package routers

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func Route(router *gin.Engine, database *sql.DB) {
	r := router.Group("api/v1")
	ChatRoute(r)
	AppointmentRoute(r)
	DoctorRoute(r)
	UserRoute(r, database)
}
