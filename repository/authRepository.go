package repository

import (
	"jdsapp/config"
	"jdsapp/entity"
)

func Authentication(nik string, password string) (check bool, HashPassword string) {
	var result *entity.RequestAuth
	tx := config.DB.Table("users").
		Where("nik = ?", nik).
		First(&result)

	if tx.Error != nil {
		return false, "nil"
	}

	return true, result.Password
}
