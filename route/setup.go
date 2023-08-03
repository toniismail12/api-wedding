package route

import (
	"be-wedding/controllers"
	"be-wedding/services"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {

	// setup cors
	services.SetupCors(app)

	app.Get("/", controllers.AppName)
	app.Static("/docs", "./docs")
	
	// wishes
	app.Get("/wishes", controllers.GetWishes)
	app.Post("/wishes", controllers.CreateWishes)

}
