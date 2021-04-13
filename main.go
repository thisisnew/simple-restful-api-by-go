package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type Profiles struct {
	gorm.Model
	ID      string `gorm:"primary_key"`
	Name    string
	Age     string
	Company string
	IsAdmin string `gorm:"column:is_admin"`
}

const dsn = "root:1234@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", GetProfileList).Methods("GET")
	router.HandleFunc("/profile/{id}", GetProfile).Methods("GET")
	router.HandleFunc("/insert", insertProfile).Methods("POST")
	router.HandleFunc("/update", updateProfile).Methods("POST")
	router.HandleFunc("/delete/{id}", deleteProfile).Methods("DELETE")

	http.ListenAndServe(":8080", httpHandler(router))
}

func GetProfileList(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	var profiles []Profiles

	db.Find(&profiles)

	if len(profiles) == 0 {
		json.NewEncoder(w).Encode(map[string]string{"result": "데이터가 없습니다."})
		return
	}

	json.NewEncoder(w).Encode(profiles)
}

func GetProfile(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	p := mux.Vars(r)
	id := p["id"]

	profiles := Profiles{}
	db.First(&profiles, id)

	if profiles.Name == "" {
		json.NewEncoder(w).Encode(map[string]string{"result": "데이터가 없습니다."})
		return
	}

	json.NewEncoder(w).Encode(profiles)
}

func insertProfile(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	newProfile := Profiles{}
	err = json.NewDecoder(r.Body).Decode(&newProfile)

	result := db.Create(&newProfile)

	if result.RowsAffected > 0 {
		json.NewEncoder(w).Encode(map[string]int{"status": http.StatusOK})
		return
	}

	json.NewEncoder(w).Encode(map[string]int{"status": http.StatusConflict})
}

func updateProfile(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	modifyProfile := Profiles{}
	err = json.NewDecoder(r.Body).Decode(&modifyProfile)

	var profiles Profiles
	result := db.Model(&profiles).Where("id = ?", modifyProfile.ID).Updates(Profiles{Name: modifyProfile.Name, Age: modifyProfile.Age, Company: modifyProfile.Company, IsAdmin: modifyProfile.IsAdmin})

	if result.RowsAffected == 0 {
		json.NewEncoder(w).Encode(map[string]int{"status": http.StatusNotModified})
	} else {
		json.NewEncoder(w).Encode(map[string]int{"status": http.StatusCreated})
	}
}

func deleteProfile(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	p := mux.Vars(r)
	id := p["id"]

	var profiles Profiles
	db.Where("id = ?", id).Delete(&profiles)
	db.Commit()

	json.NewEncoder(w).Encode(map[string]int{"status": http.StatusOK})
}

func httpHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print(r.RemoteAddr, " ", r.Proto, " ", r.Method, " ", r.URL)
		handler.ServeHTTP(w, r)
	})
}