package handler

import (
	"bank/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type customerHandler struct {
	customer_service service.CustomerService
}

func InitializeCustomerHandler(customerService service.CustomerService) customerHandler {
	return customerHandler{customer_service: customerService}
}

func (handle customerHandler) GetCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := handle.customer_service.GetCustomers()
	if err != nil {
		handlerError(w, err)
		return
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func (handle customerHandler) GetCustomer(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["customer_id"])
	customer, err := handle.customer_service.GetCustomer(id)
	if err != nil {
		handlerError(w, err)
		return
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(customer)
}
