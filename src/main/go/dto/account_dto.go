package dto

import (
	"time"
)

type AccountDto struct {
	Id        uint
	CreatedAt time.Time
	UserId    uint
	Purchases []*PurchaseDto
	Logins    []*LoginDto
	Views     []*ViewDto
}
