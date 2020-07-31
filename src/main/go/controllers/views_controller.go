package controllers

import (
	"../dto"
	"../services"
	"../utils"
)

type ViewsController struct {
	service *services.ViewService
}

func NewViewsController(service *services.ViewService) *ViewsController {
	return &ViewsController{service: service}
}

func (controller *ViewsController) ReadByAccountId(accountId uint) ([]*dto.ViewDto, error) {
	views, err := controller.service.ReadByAccountId(accountId)
	if err != nil {
		return nil, err
	}
	return utils.FromViewsArray(views), nil
}

func (controller *ViewsController) CreateView(payload *dto.ViewDto) (*dto.ViewDto, error) {
	views := utils.ToView(payload)
	result, err := controller.service.CreateView(views)
	if err != nil {
		return nil, err
	}
	return utils.FromView(result), nil
}
