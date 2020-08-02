package handlers

import (
	"../controllers"
	"../dto"
	log "../logger"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type PurchasesHandler struct {
	controller *controllers.PurchasesController
}

func NewPurchasesHandler(controller *controllers.PurchasesController) *PurchasesHandler {
	return &PurchasesHandler{controller: controller}
}

func (handler *PurchasesHandler) GetByAccountId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	result := vars["accountId"]
	accountId, err := strconv.ParseUint(result, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	purchase, err := handler.controller.GetByAccountId(uint(accountId))
	log.Debug("Create new information ", purchase)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	ResponseSender(w, purchase, http.StatusOK)
}

func (handler *PurchasesHandler) CreatePurchase(w http.ResponseWriter, r *http.Request) {
	var payload dto.PurchaseDto
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	purchase, err := handler.controller.CreatePurchase(&payload)
	log.Debug("Create new information ", purchase)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ResponseSender(w, purchase, http.StatusCreated)
}
