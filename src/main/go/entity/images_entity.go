package entity

type Image struct {
	Id        uint   `gorm:"column:id"`
	ProductId uint   `gorm:"column:product_id"`
	Img       string `gorm:"column:image"`
}

func (image *Image) TableName() string {
	return "images"
}
