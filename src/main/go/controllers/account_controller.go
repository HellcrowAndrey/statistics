package controllers

import (
	"../dto"
	"../services"
	"../utils"
)

type AccountsController struct {
	accountService *services.AccountService
}

func NewAccountController(service *services.AccountService) *AccountsController {
	return &AccountsController{accountService: service}
}

func (controller *AccountsController) GetByUserId(userId uint) (*dto.AccountDto, error) {
	account, err := controller.accountService.ReadByUserId(userId)
	if err == nil {
		return utils.FromAccount(account), nil
	}
	return nil, err
}

func (controller *AccountsController) CreateAccount(payload *dto.AccountDto) (*dto.AccountDto, error) {
	account := utils.ToAccount(payload)
	result, err := controller.accountService.CreateAccount(account)
	if err == nil {
		return utils.FromAccount(result), nil
	}
	return nil, err
}
