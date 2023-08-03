package services

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func MessageError(err string) interface{} {

	log.Println(err)
	datas := (fiber.Map{
		"message": err,
	})

	return datas
}
