package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Fund struct {
	gorm.Model
	AvanzaId string //`	gorm:"primary_key"`
	Name     string
	Holdings []FundHolding
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

func connectDb() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Fund{})
	db.AutoMigrate(&FundHolding{})
	return db
}

type Repository interface {
}
