package controllers

import (
	"../dto"
	"../services"
	"../utils"
)

type PurchasesController struct {
	purchaseService *services.PurchaseService
}

func NewPurchasesController(purchaseService *services.PurchaseService) *PurchasesController {
	return &PurchasesController{purchaseService: purchaseService}
}

func (controller *PurchasesController) GetByAccountId(accountId uint) ([]*dto.PurchaseDto, error) {
	purchases, err := controller.purchaseService.ReadByAccountId(accountId)
	if err != nil {
		return nil, err
	}
	return utils.FromPurchaseArray(purchases), nil
}

func (controller *PurchasesController) CreatePurchase(payload *dto.PurchaseDto) (*dto.PurchaseDto, error) {
	purchase := utils.ToPurchase(payload)
	result, err := controller.purchaseService.CreatePurchase(purchase)
	if err != nil {
		return nil, err
	}
	return utils.FromPurchase(result), err
}
