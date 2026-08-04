package models

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DBModel struct {
	Order OrderModel
}

func InitDB(databaseSource string) (*DBModel, error) {
	db, err := gorm.Open(sqlite.Open(databaseSource), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("Failed to migrate: %v", err)
	}
	
	err = db.AutoMigrate(&Order{}, &OrderItem{})
	if err != nil {
		return nil, fmt.Errorf("Failed to migrate: %v", err)
	}
	return &DBModel{
		Order: OrderModel{
			DB: db,
		},
	}, nil
}