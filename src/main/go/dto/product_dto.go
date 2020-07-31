package dto

type ProductDto struct {
	Id           uint
	Name         string
	Price        float64
	Quantity     int
	Description  string
	PreviewImage string
	Images       []string
}
