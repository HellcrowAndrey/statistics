package entity

import (
	"time"
)

type Purchase struct {
	Id        uint      `gorm:"column:id"`
	CreatedAt time.Time `gorm:"column:created_at"`
	OrderId   uint		`gorm:"column:order_id"`
	AccountId uint      `gorm:"column:account_id"`
	Products  []*Product `gorm:"foreignkey:PurchaseId"`
}

func (purchase *Purchase) TableName() string {
	return "purchase"
}
