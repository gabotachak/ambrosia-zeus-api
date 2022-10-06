package util

import (
	"ambrosia-zeus-api/cmd/api/model/storage"
	"fmt"
	"log"
	"time"

	mySqlDriver "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitializeDatabase() *gorm.DB {
	c := mySqlDriver.Config{
		User:      "root",
		Passwd:    "swarch2022ii",
		DBName:    "ambrosia_zeus_db",
		Addr:      "localhost:3307",
		Net:       "tcp",
		ParseTime: true,
		Loc:       time.UTC,
	}

	dsn := c.FormatDSN()
	fmt.Println(dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("DB connection failed. Error: %s", err)
		return nil
	}

	if !db.Migrator().HasTable(storage.User{}) || !db.Migrator().HasTable(storage.Credential{}) {
		db.Migrator().CreateTable(storage.User{}, storage.Credential{})
	}

	return db
}
