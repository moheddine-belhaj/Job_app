package config

import (
	// "log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var (
	db *gorm.DB
)

func Connect() {
	// Replace the empty string with the PostgreSQL connection string
	dsn := "host=localhost user=postgres dbname=store sslmode=disable password=0000" 
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
