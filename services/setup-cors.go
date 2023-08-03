package services

import (
	"be-wedding/config"
	"be-wedding/constants"
	"be-wedding/database"
	"be-wedding/table"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func GetDomain() string {

	db := database.DB

	var (
		data   []table.CorsDomain
		result string
	)

	if err := db.Table(constants.TABLE_CORS_DOMAIN).Find(&data).Error; err != nil {
		log.Println("Failed to retrieve data from the database:", err)
	}

	for i, domain := range data {
		if i > 0 {
			result += ", "
		}
		result += domain.Domain
	}

	return result

}

func SetupCors(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins:     GetDomain(),
		AllowHeaders:     config.AllowHeaders(),
		AllowMethods:     config.AllowMethods(),
		AllowCredentials: true,
	}))
}
