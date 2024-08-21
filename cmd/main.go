package main

import (
	"alimed_go/pkg/configs"
	"alimed_go/pkg/db"
	"alimed_go/pkg/routers"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	// gin-swagger middleware
	_ "alimed_go/docs"
)

// @title Go + Gin Todo API
// @version 1.0
// @description This is a sample server todo server. You can visit the GitHub repository at https://github.com/LordGhostX/swag-gin-demo

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:9090
// @BasePath /
// @query.collection.format multi

func main() {
	//read env file

	configs.LoadEnv()
	configs.InitCache()

	database := db.ConnectDatabase()

	port := os.Getenv("PORT")

	app := gin.New()

	// router
	routers.Route(app, database)

	// run server
	err := app.Run(":" + port)
	if err != nil {
		return
	}

}
