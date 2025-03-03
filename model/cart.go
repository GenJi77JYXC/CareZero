package model

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	UserID int64
}

type CartItem struct {
	gorm.Model
	CartID    int64
	ProductID int64
	Quantity  int64
}
