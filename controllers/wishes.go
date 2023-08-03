package controllers

import (
	"be-wedding/models"
	"be-wedding/response"
	"be-wedding/services"

	"github.com/gofiber/fiber/v2"
)

func GetWishes(c *fiber.Ctx) error {

	page, limit := services.PageLimit(c.Query("page", "1"), c.Query("limit", "10"))

	name := c.Query("name")
	app := c.Query("app")

	data := models.GetWishes(page, limit, name, app)

	c.Status(200)
	return c.JSON(data)

}

func CreateWishes(c *fiber.Ctx) error {
	var (
		request = new(response.FormWishes)
	)

	if err := c.BodyParser(&request); err != nil {
		c.Status(400)
		return c.JSON(services.MessageError(err.Error()))
	}

	erro := models.CreateWishes(request.App_id, request.Name, request.Wishes)
	if erro != nil {
		c.Status(400)
		return c.JSON(services.MessageError(erro.Error()))
	}

	c.Status(200)
	return c.JSON(fiber.Map{
		"message": "Success Create Data",
		"request": request,
	})

}
