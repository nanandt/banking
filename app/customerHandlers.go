package app

import (
	"banking/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// type Customer struct {
// 	Name    string `json:"full_name" xml:"name"`
// 	City    string `json:"city" xml:"city"`
// 	ZipCode string `json:"zip_code" xml:"zipcode"`
// }

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(rw http.ResponseWriter, r *http.Request) {
	// customers := []Customer{
	// 	{"Rizky", "Kalimantan", "19281"},
	// 	{"Fatih", "Tegal", "929292"},
	// }
	status := r.URL.Query().Get("status")

	customers, err := ch.service.GetAllCustomer(status)

	if err != nil {
		writeResponse(rw, err.Code, err.AsMessage())
	} else {
		writeResponse(rw, http.StatusOK, customers)
	}
}

func (ch *CustomerHandlers) getCustomer(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["customer_id"]

	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		writeResponse(rw, err.Code, err.AsMessage())
	} else {
		writeResponse(rw, http.StatusOK, customer)
	}
}

func writeResponse(rw http.ResponseWriter, code int, data interface{}) {
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(code)
	err := json.NewEncoder(rw).Encode(data)
	if err != nil {
		panic(err)
	}

}
