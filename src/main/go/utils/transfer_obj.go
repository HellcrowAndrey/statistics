package utils

import (
	"../dto"
	"../entity"
)

// ===========================================================
// ================= CONVERT FROM ENTITY =====================
// ===========================================================

func FromAccount(data *entity.Account) *dto.AccountDto {
	purchases := FromPurchaseArray(data.Purchases)
	logins := FromLoginArray(data.Logins)
	views := FromViewsArray(data.Views)
	return &dto.AccountDto{
		Id:        data.Id,
		CreatedAt: data.CreatedAt,
		Purchases: purchases,
		Logins:    logins,
		Views:     views,
	}
}

func FromPurchase(data *entity.Purchase) *dto.PurchaseDto {
	array := data.Products
	var products []*dto.ProductDto
	for _, product := range array {
		products = append(products, FromProduct(product))
	}
	return &dto.PurchaseDto{
		Id:        data.Id,
		CreatedAt: data.CreatedAt,
		OrderId:   data.OrderId,
		Products:  products,
	}
}

func FromLogin(data *entity.Login) *dto.LoginDto {
	return &dto.LoginDto{
		Id:        data.Id,
		CreatedAt: data.CreatedAt,
		Ip:        data.Ip,
		Device:    data.Device,
		Location:  data.Location,
	}
}

func FromView(data *entity.View) *dto.ViewDto {
	array := data.Products
	var products []*dto.ProductDto
	for _, product := range array {
		products = append(products, FromProduct(product))
	}
	return &dto.ViewDto{
		Id:        data.Id,
		CreatedAt: data.CreatedAt,
		Products:  products,
	}
}

func FromProduct(data *entity.Product) *dto.ProductDto {
	array := data.Images
	var images []string
	for _, i := range array {
		images = append(images, i.Img)
	}
	return &dto.ProductDto{
		Id:           data.Id,
		Name:         data.Name,
		Price:        data.Price,
		Quantity:     data.Quantity,
		Description:  data.Description,
		PreviewImage: data.PreviewImage,
		Images:       images,
	}
}

func FromViewsArray(data []*entity.View) []*dto.ViewDto {
	var views []*dto.ViewDto
	for _, purchase := range data {
		views = append(views, FromView(purchase))
	}
	return views
}

func FromPurchaseArray(data []*entity.Purchase) []*dto.PurchaseDto {
	var purchases []*dto.PurchaseDto
	for _, purchase := range data {
		purchases = append(purchases, FromPurchase(purchase))
	}
	return purchases
}

func FromLoginArray(data []*entity.Login) []*dto.LoginDto {
	var logins []*dto.LoginDto
	for _, purchase := range data {
		logins = append(logins, FromLogin(purchase))
	}
	return logins
}

// ===========================================================
// ================= CONVERT TO ENTITY =======================
// ===========================================================

func ToAccount(data *dto.AccountDto) *entity.Account {
	purchases := ToPurchasesArray(data.Purchases)
	logins := ToLoginsArray(data.Logins)
	views := ToViewsArray(data.Views)
	return &entity.Account{
		Id:        data.Id,
		CreatedAt: data.CreatedAt,
		UserId:    data.UserId,
		Purchases: purchases,
		Logins:    logins,
		Views:     views,
	}
}

func ToPurchase(data *dto.PurchaseDto) *entity.Purchase {
	array := data.Products
	var products []*entity.Product
	for _, product := range array {
		products = append(products, ToProduct(product))
	}
	return &entity.Purchase{
		Id:        data.Id,
		CreatedAt: data.CreatedAt,
		OrderId:   data.OrderId,
		AccountId: data.AccountId,
		Products:  products,
	}
}

func ToLogin(data *dto.LoginDto) *entity.Login {
	return &entity.Login{
		Id:        data.Id,
		CreatedAt: data.CreatedAt,
		AccountId: data.AccountId,
		Ip:        data.Ip,
		Device:    data.Device,
		Location:  data.Location,
	}
}

func ToView(data *dto.ViewDto) *entity.View {
	array := data.Products
	var products []*entity.Product
	for _, product := range array {
		products = append(products, ToProduct(product))
	}
	return &entity.View{
		Id:        data.Id,
		CreatedAt: data.CreatedAt,
		AccountId: data.AccountId,
		Products:  products,
	}
}

func ToProduct(data *dto.ProductDto) *entity.Product {
	return &entity.Product{
		Id:           data.Id,
		Name:         data.Name,
		Price:        data.Price,
		Quantity:     data.Quantity,
		Description:  data.Description,
		PreviewImage: data.PreviewImage,
		Images:       ToImages(data.Images),
	}
}

func ToImages(data []string) []*entity.Image {
	var images []*entity.Image
	for _, i := range data {
		images = append(images, &entity.Image{Img: i})
	}
	return images
}

func ToPurchasesArray(data []*dto.PurchaseDto) []*entity.Purchase {
	var purchases []*entity.Purchase
	for _, purchase := range data {
		purchases = append(purchases, ToPurchase(purchase))
	}
	return purchases
}

func ToLoginsArray(data []*dto.LoginDto) []*entity.Login {
	var logins []*entity.Login
	for _, login := range data {
		logins = append(logins, ToLogin(login))
	}
	return logins
}

func ToViewsArray(data []*dto.ViewDto) []*entity.View {
	var views []*entity.View
	for _, view := range data {
		views = append(views, ToView(view))
	}
	return views
}
