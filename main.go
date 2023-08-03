package main

import (
	"be-wedding/config"
	"be-wedding/database"
	"be-wedding/route"
	"be-wedding/services"
)

func main() {

	// database connect
	database.Connect()

	// start fiber
	app := services.CreateApp()

	// route
	route.Setup(app)

	// get port
	port := config.Env("PORT")

	// run
	app.Listen(":" + port)

}
