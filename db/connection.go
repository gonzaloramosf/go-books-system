package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var DSN = "host=localhost user=root password=root dbname=books port=5432"

func ConnectionDB() {
	d, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	DB = d
}
