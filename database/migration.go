package database

import (
	"be-wedding/table"
)

func Migration() {

	DB.AutoMigrate(
		&table.CorsDomain{},
		&table.Wishes{},
	)

}
