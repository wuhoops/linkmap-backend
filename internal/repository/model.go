package repository

import (
	"backend/internal/core/domain/database"

	"gorm.io/gorm"
)

var UserModel *gorm.DB

func assignModel() {
	UserModel = DB.Model(new(database.User))
}
