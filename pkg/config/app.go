package config

import (
	// "database/sql"

	"database/sql"
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
	// viper.SetConfigFile("app.env")
	// viper.ReadInConfig()
	// config := &Config{}
	// err := viper.Unmarshal(&config)
	// var dsn string = config.DBSource
	var dsn string = os.Getenv("DB_SOURCE")
	sqlDB, _ := sql.Open("mysql", dsn)
	d, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	// d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// sqlDB, err := d.DB()
	// sqlDB.SetMaxIdleConns(0)
	// // sqlDB.SetMaxOpenConns(100)
	// sqlDB.SetConnMaxLifetime(time.Minute * 4)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("database started")
	}
	db = d

}

func GetDB() *gorm.DB {
	return db
}
