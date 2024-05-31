package util

import (
	"log"
	"os"
	"time"

	"github.com/gabotachak/ambrosia-zeus-api/cmd/api/model/storage"
	mySqlDriver "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitializeDatabase() *gorm.DB {
	godotenv.Load()

	c := mySqlDriver.Config{
		User:                 goDotEnvVariable("DB_USER"),
		Passwd:               goDotEnvVariable("DB_PASS"),
		DBName:               goDotEnvVariable("DB_NAME"),
		Addr:                 goDotEnvVariable("DB_HOST"),
		Net:                  "tcp",
		ParseTime:            true,
		Loc:                  time.UTC,
		AllowNativePasswords: true,
	}

	dsn := c.FormatDSN()

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("DB connection failed. Error: %s", err)
		return nil
	}

	db.Migrator().CreateTable(storage.Credential{}, storage.User{})

	return db
}

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
