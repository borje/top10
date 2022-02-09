package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Fund struct {
	gorm.Model
	Id   string `gorm:"primary_key"`
	Name string
}

func connectDb() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Fund{}, &FundHolding{})
	return db
}

type FundHolding struct {
	gorm.Model
	FundId         string
	Date           string
	Position       int
	Isin           string
	Name           string
	SizePercentage float64
}

type Repository interface {
}
