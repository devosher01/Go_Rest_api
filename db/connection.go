package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBS = "host=localhost user=devosher password=mysecretpassword dbname=gorm port=5432 "
var DB *gorm.DB

func DBConection() {
	var error error
	DB, error = gorm.Open(postgres.Open(DBS), &gorm.Config{})
	if error != nil {
		panic("failed to connect database")
	} else {
		println("Connection Opened to Database")
	}
}
