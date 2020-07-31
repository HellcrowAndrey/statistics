package handlers

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type RestHandler struct {
	accountHandler  *AccountsHandler
	purchaseHandler *PurchasesHandler
	loginHandler    *LoginHandler
	viewsHandler    *ViewsHandler
}

func NewRestHandler(accountHandler *AccountsHandler, purchaseHandler *PurchasesHandler, loginHandler *LoginHandler, viewsHandler *ViewsHandler) *RestHandler {
	return &RestHandler{accountHandler: accountHandler, purchaseHandler: purchaseHandler, loginHandler: loginHandler, viewsHandler: viewsHandler}
}

func (handler *RestHandler) Handler() http.Handler {
	router := mux.NewRouter()
	router.
		HandleFunc("/v1/accounts/{userId}", handler.accountHandler.GetByUserId).
		Methods("GET")
	router.
		HandleFunc("/v1/accounts", handler.accountHandler.CreateAccount).
		Methods("POST")
	router.
		HandleFunc("/v1/purchase/{accountId}", handler.purchaseHandler.GetByAccountId).
		Methods("GET")
	router.
		HandleFunc("/v1/purchase", handler.purchaseHandler.CreatePurchase).
		Methods("POST")
	router.
		HandleFunc("/v1/logins/{accountId}", handler.loginHandler.GetByAccountId).
		Methods("GET")
	router.
		HandleFunc("/v1/logins", handler.loginHandler.CreateLogin).
		Methods("POST")
	router.Use(loggingMiddleware)
	http.Handle("/", router)
	return router
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(
			http.TimeFormat,
			"Method:", r.Method,
			"URL:", r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
