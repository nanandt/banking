package app

import (
	"banking/service"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	ZipCode string `json:"zip_code" xml:"zipcode"`
}

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(rw http.ResponseWriter, r *http.Request) {
	// customers := []Customer{
	// 	{"Rizky", "Kalimantan", "19281"},
	// 	{"Fatih", "Tegal", "929292"},
	// }

	customers, _ := ch.service.GetAllCustomer()

	if r.Header.Get("Content-Type") == "application/xml" {
		rw.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(rw).Encode(customers)
	} else {
		rw.Header().Add("Content-Type", "application/json")

		json.NewEncoder(rw).Encode(customers)
	}
}

func (ch *CustomerHandlers) getCustomer(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["customer_id"]

	customer, err := ch.service.GetCustomer(id)
	rw.WriteHeader(err.Code)
	if err != nil {
		fmt.Fprintf(rw, err.Message)
	} else {
		json.NewEncoder(rw).Encode(customer)
	}

}
