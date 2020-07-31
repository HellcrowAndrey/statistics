package handlers

import (
	"../controllers"
	"../dto"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type ViewsHandler struct {
	controller *controllers.ViewsController
}

func NewViewsHandler(controller *controllers.ViewsController) *ViewsHandler {
	return &ViewsHandler{controller: controller}
}

func (handler *ViewsHandler) GetByAccountId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	result := vars["accountId"]
	accountId, err := strconv.ParseUint(result, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	payload, err := handler.controller.ReadByAccountId(uint(accountId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	ResponseSender(w, payload, http.StatusOK)
}

func (handler *ViewsHandler) CreateViews(w http.ResponseWriter, r *http.Request) {
	var payload dto.ViewDto
	err := json.NewDecoder(r.Body).Decode(&payload)
	result, err := handler.controller.CreateView(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ResponseSender(w, result, http.StatusCreated)
}
