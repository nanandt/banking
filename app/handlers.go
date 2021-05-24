package app

import (
	"banking/service"
	"encoding/json"
	"encoding/xml"
	"net/http"
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
