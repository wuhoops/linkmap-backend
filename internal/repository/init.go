package repository

import (
	"backend/internal/core/domain/database"
	"backend/internal/util/config"
	"github.com/redis/go-redis/v9"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewDB() (*gorm.DB, error) {
	dsn := config.C.DBConn
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	DB = db
	if err := migrate(); err != nil {
		logrus.Fatal("UNABLE TO MIGRATE GORM MODEL")
	}

	logrus.Debugln("INITIALIZE DB CONNECTION")
	return db, err
}

func InitRedis() (*redis.Client, error) {
	opt, err := redis.ParseURL(config.C.Redis)
	if err != nil {
		panic(err)
	}

	client := redis.NewClient(opt)
	logrus.Debugln("INITIALIZE REDIS CONNECTION")
	return client, nil
}

func migrate() error {
	if err := DB.AutoMigrate(
		new(database.User),
		new(database.Card),
		new(database.Social),
	); err != nil {
		return err
	}
	return nil
}
