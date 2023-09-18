package config

import (
	// "database/sql"

	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

type Config struct {
	DBDriver string `mapstructure:"DB_DRIVER"`
	DBSource string `mapstructure:"DB_SOURCE"`
}

func Connect() {

	var dsn string = os.Getenv("DB_SOURCE")
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	_ = d.Exec("CREATE DATABASE IF NOT EXISTS bookStore;")
	dsn = os.Getenv("DB_SOURCE_DB")
	database_after_creation, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	} else {
		fmt.Println("database started")
	}
	db = database_after_creation
}

func GetDB() *gorm.DB {
	return db
}
