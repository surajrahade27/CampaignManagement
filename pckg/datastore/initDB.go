package datastore

import (
	"campaing_management/pckg/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const DNS = "root:Password@1@tcp(localhost:3306)/FairPrise?parseTime=true"

func InitialMigration() {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}
	var tables = []interface{}{&models.Campaing{},
		&models.Product{},
		&models.Store{}}

	DB.AutoMigrate(tables...)
}
