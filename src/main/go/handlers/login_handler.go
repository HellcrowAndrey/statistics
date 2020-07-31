package handlers

import (
	"../controllers"
	"../dto"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type LoginHandler struct {
	controller *controllers.LoginsController
}

func NewLoginHandler(controller *controllers.LoginsController) *LoginHandler {
	return &LoginHandler{controller: controller}
}

func (handler *LoginHandler) GetByAccountId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	result := vars["accountId"]
	accountId, err := strconv.ParseUint(result, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	account, err := handler.controller.GetByAccountId(uint(accountId))
	ResponseSender(w, account, http.StatusOK)
}

func (handler *LoginHandler) CreateLogin(w http.ResponseWriter, r *http.Request) {
	var payload dto.LoginDto
	err := json.NewDecoder(r.Body).Decode(&payload)
	result, err := handler.controller.CreateLogin(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ResponseSender(w, result, http.StatusCreated)
}
