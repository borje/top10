package main

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Fund struct {
	gorm.Model
	AvanzaId string //`	gorm:"primary_key"`
	Name     string
}

type FundHolding struct {
	gorm.Model
	FundId         uint
	Date           string
	Position       int
	Isin           string
	Name           string
	SizePercentage float64
}

var DB *gorm.DB

func connectDb() {
	log.Println("Connecting Database")
	var err error
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB.AutoMigrate(&Fund{})
	DB.AutoMigrate(&FundHolding{})
}

type Repository interface {
}
