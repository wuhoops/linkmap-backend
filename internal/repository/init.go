package repository

import (
	"context"
	"goHexagonalBlog/internal/core/domain/database"
	"goHexagonalBlog/internal/util/config"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewDB(ctx context.Context) (*gorm.DB, error) {
	dsn := config.C.DBConn
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	DB = db
	if err := migrate(); err != nil {
		logrus.Fatal("UNABLE TO MIGRATE GORM MODEL")
	}

	assignModel()
	logrus.Debugln("INITIALIZE MYSQL CONNECTION")
	return db, err
}

func migrate() error {
	if err := DB.AutoMigrate(
		new(database.User),
		new(database.Card),
	); err != nil {
		return err
	}
	return nil
}
