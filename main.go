package main

import (
	"github.com/gorilla/mux"
	"net/http"
	util "rest/util"
	"rest/vehicle"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", vehicle.GetAvailableVehicleList).Methods("GET")
	//router.HandleFunc("/profiles", GetProfileList).Methods("GET")
	//router.HandleFunc("/", billing.GetBillingList).Methods("GET")
	//router.HandleFunc("/", billing.GetBillingList).Methods("GET")
	//router.HandleFunc("/profiles/{id}", GetProfile).Methods("GET")
	//router.HandleFunc("/profiles", insertProfile).Methods("POST")
	//router.HandleFunc("/profiles/{id}", updateProfile).Methods("PATCH")
	//router.HandleFunc("/profiles/{id}", deleteProfile).Methods("DELETE")

	http.ListenAndServe(":8080", util.HttpHandler(router))
}
