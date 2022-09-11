package util

import (
	"ambrosia-zeus-api/cmd/api/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func InitializeDatabase() *gorm.DB {
	dsn := "root:swarch2022ii@tcp(127.0.0.1:3307)/ambrosia_zeus_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("DB connection failed. Error: %s", err)
		return nil
	}

	if !db.Migrator().HasTable(model.User{}) || !db.Migrator().HasTable(model.Credential{}) {
		db.Migrator().CreateTable(model.User{}, model.Credential{})
	}

	return db
}
