package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseConnection() {
	var err error
	dsn := "maulana:maulana186@tcp(127.0.0.1:3306)/jdsapp?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Connection refused")
		log.Println(err.Error())
	}

	// initiate authority
	fmt.Println("Connection successfully")
}
