package database

import (
	"fmt"

	"email.v1/config"
	"email.v1/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDBPostgres(cfg *config.AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta",
		cfg.DBHOST, cfg.DBUSERNAME, cfg.DBPASSWORD, cfg.DBNAME, cfg.DBPORT)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		PrepareStmt: true,
	})
	if err != nil {
		panic(err)
	}

	return db
}

func DBMigration(db *gorm.DB) {
	db.AutoMigrate(models.Users{})
	db.AutoMigrate(models.TokenVerif{})
}