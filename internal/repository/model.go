package repository

import (
	"goHexagonalBlog/internal/core/domain/database"

	"gorm.io/gorm"
)

var UserModel *gorm.DB

func assignModel() {
	UserModel = DB.Model(new(database.User))
}
