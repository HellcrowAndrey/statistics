package controllers

import (
	"../dto"
	"../services"
	"../utils"
)

type LoginsController struct {
	service *services.LoginService
}

func NewLoginsController(service *services.LoginService) *LoginsController {
	return &LoginsController{service: service}
}

func (controller *LoginsController) GetByAccountId(accountId uint) (*dto.LoginDto, error) {
	login, err :=  controller.service.ReadByAccountId(accountId)
	if err != nil {
		return nil, err
	}
	return utils.FromLogin(login), nil
}

func (controller *LoginsController) CreateLogin(payload *dto.LoginDto) (*dto.LoginDto, error) {
	login := utils.ToLogin(payload)
	data, err := controller.service.CreateLogin(login)
	if err != nil {
		return nil, err
	}
	return utils.FromLogin(data), nil
}
