package main

import (
	"go-food-delivery-api/src/configs"
	"go-food-delivery-api/src/models"
)

func main() {
	if db, err := configs.GetGormInstance(); err != nil {
		panic("Can't connect to db via GORM: " + err.Error())
	} else {
		models := []interface{}{
			&models.Role{},
			&models.Employee{},
			&models.Account{},
			&models.Area{},
			&models.Table{},
			&models.Order{},
			&models.Categories{},
			&models.Discount{},
			&models.Product{},
			&models.Bill{},
		}
		db.AutoMigrate(models...)
	}
}
