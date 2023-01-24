package handler

import (
	"bank/errs"
	"bank/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type accountHandler struct {
	accSrv service.AccountService
}

func NewAccountHandler(accSrv service.AccountService) accountHandler {
	return accountHandler{accSrv: accSrv}
}

func (h accountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["customer_id"])
	if r.Header.Get("content-type") != "application/json" {
		handlerError(w, errs.NewValidationError("request body incorrect format"))
		return
	}
	request := service.NewAccountRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		handlerError(w, errs.NewValidationError("request body incorrect format"))
		return
	}
	response, err := h.accSrv.NewAccount(id, request)
	if err != nil {
		handlerError(w, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h accountHandler) GetAccounts(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["customer_id"])
	response, err := h.accSrv.GetAccounts(id)
	if err != nil {
		handlerError(w, err)
		return
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}
