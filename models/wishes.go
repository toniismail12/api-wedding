package models

import (
	"be-wedding/constants"
	"be-wedding/database"
	"be-wedding/response"
	"be-wedding/services"
	"be-wedding/table"

	"github.com/gofiber/fiber/v2"
)

func GetWishes(page, limit int, name, app string) interface{} {
	
	DB := database.DB

	offset := (page - 1) * limit

	var (
		total_rows int64
		data       []response.GetWishes
	)

	query := DB.Table(constants.TABLE_WISHES).
		Where("name LIKE ?", "%"+name+"%").
		Where("app_id = ?", app).
		Order("id desc")
	query.Count(&total_rows)
	query.Offset(offset).Limit(limit)
	query.Find(&data)

	datas := (fiber.Map{
		"message": "data wishes",
		"data":    data,
		"meta": fiber.Map{
			"limit": limit,
			"page":  page,
			"total": total_rows,
		},
	})

	return datas

}

func CreateWishes(app uint, name, wishes string) error {

	input := services.ValidationInputInt(app, "App Id")
	if input != nil {
		return input
	}
	input = services.ValidationInputString(name, "Name")
	if input != nil {
		return input
	}
	input = services.ValidationInputString(wishes, "Wishes")
	if input != nil {
		return input
	}

	DB := database.DB

	var (
		w table.Wishes
	)

	w.App_id = app
	w.Name = name
	w.Wishes = wishes

	err := DB.Create(&w).Error

	return err

}
