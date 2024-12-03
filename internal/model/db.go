package model

import (
	"bingyan-freshman-task0/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	dsn := config.Config.DB.Dsn
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&User{}); err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&Post{}); err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&Comment{}); err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&Like{}); err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&Node{}); err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&Follow{}); err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&Body{}); err != nil {
		panic(err)
	}
}
