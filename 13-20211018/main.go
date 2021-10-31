package main

import (
	"gin-postgres/repo"
	"gin-postgres/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect Database
	db := repo.ConnectDatabase()
	defer db.Close()

	//repo.MockData()

	// Init Gin App
	app := gin.New()

	// Register router
	router.InitRouter(app)

	// Gin App listen port
	err := app.Run(":3000")
	if err != nil {
		panic(err)
	}
}
