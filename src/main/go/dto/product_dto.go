package dto

type ProductDto struct {
	Id           uint     `json:"id"`
	Name         string   `json:"name"`
	Price        float64  `json:"price"`
	Quantity     int      `json:"quantity"`
	Description  string   `json:"description"`
	PreviewImage string   `json:"previewImage"`
	Images       []string `json:"images"`
}
