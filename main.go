package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"rest/company"
	util "rest/util"
)

func main() {

	router := mux.NewRouter()
	//router.HandleFunc("/", vehicle.GetAvailableVehicleList).Methods("GET")
	router.HandleFunc("/", company.GetCompanyList).Methods("GET")
	//router.HandleFunc("/profiles", GetProfileList).Methods("GET")
	//router.HandleFunc("/", billing.GetBillingList).Methods("GET")
	//router.HandleFunc("/", billing.GetBillingList).Methods("GET")
	//router.HandleFunc("/profiles/{id}", GetProfile).Methods("GET")
	//router.HandleFunc("/profiles", insertProfile).Methods("POST")
	//router.HandleFunc("/profiles/{id}", updateProfile).Methods("PATCH")
	//router.HandleFunc("/profiles/{id}", deleteProfile).Methods("DELETE")

	http.ListenAndServe(":8080", util.HttpHandler(router))
}
