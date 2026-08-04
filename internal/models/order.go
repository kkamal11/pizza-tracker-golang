package models

import (
	"time"

	"gorm.io/gorm"
)

var (
	OrderStatus = []string{"Placed", "Processing", "Baking", "Shipped", "Delivered", "Cancelled"}
	PizzaType   = []string{"Margherita", "Pepperoni", "Hawaiian", "Veggie", "BBQ Chicken", "Meat Lovers", "Supreme"}
	PizzaSize   = []string{"Small", "Medium", "Large", "Extra Large"}
)

type OrderModel struct {
	DB *gorm.DB
}

type Order struct {
	ID          uint    `gorm:"primaryKey;size:32" json:"id"`
	CustomerID  uint    `gorm:"not null" json:"customer_id"`
	Phone       string  `gorm:"type:varchar(20);not null" json:"phone"`
	Address     string  `gorm:"type:varchar(255);not null" json:"address"`
	Status      string  `gorm:"type:varchar(50);not null" json:"status"`
	Items 	[]OrderItem `gorm:"foreignKey:OrderID" json:"items"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
}


type OrderItem struct {
	ID        uint   `gorm:"primaryKey;size:32" json:"id"`
	OrderID   uint   `gorm:"index" json:"order_id"`
	PizzaType string `gorm:"type:varchar(50);not null" json:"pizza_type"`
	PizzaSize string `gorm:"type:varchar(50);not null" json:"pizza_size"`
	Quantity  int    `gorm:"not null" json:"quantity"`
	Instructions string `json:"instructions"`
}


func (o *Order) BeforeCreate(tx *gorm.DB) {
	if o.ID == 0 {
		o.ID = uint(time.Now().UnixNano() / int64(time.Millisecond))
	}
}

func (oi *OrderItem) BeforeCreate(tx *gorm.DB) {
	if oi.ID == 0 {
		oi.ID = uint(time.Now().UnixNano() / int64(time.Millisecond))
	}
}

func (om *OrderModel) CreateOrder(order *Order) error {
	return om.DB.Create(order).Error
}

func (om *OrderModel) GetOrderByID(id uint) (*Order, error) {
	var order Order
	err := om.DB.Preload("Items").First(&order, id).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}