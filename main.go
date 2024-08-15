package main

import (
	"alimed_go/configs"
	"alimed_go/db"
	"alimed_go/routers"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	//read env file

	configs.LoadEnv()
	database := db.ConnectDatabase()

	port := os.Getenv("PORT")

	app := gin.New()

	// router
	routers.Route(app, database)

	// run server
	app.Run(":" + port)

}
