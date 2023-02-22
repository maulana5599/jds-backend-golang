package repository

import (
	"fmt"
	"jdsapp/config"
	"jdsapp/entity"
)

func Authentication(nik string, password string) (check bool, HashPassword string) {
	var result *entity.RequestAuth
	tx := config.DB.Table("users").
		Where("nik = ?", nik).
		First(&result)

	fmt.Println("HALOOOOOOOOOOOOOOoo")
	if tx.Error != nil {
		return false, "nil"
	}

	return true, result.Password
}
