package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net/http"
	"os"
	"time"
)

type Profiles struct {
	gorm.Model
	ID      string `gorm:"primary_key"`
	Name    string
	Age     string
	Company string
	IsAdmin string `gorm:"column:is_admin"`
}

type Vehicles struct {
	Vehicles []Vehicle `json:"items"`
	PageInfo PageInfo  `json:"pageInfo"`
}

type PageInfo struct {
	TotalRecord uint `json:"totalRecord"`
	TotalPage   uint `json:"totalPage"`
	Limit       uint `json:"limit"`
	Page        uint `json:"page"`
	PrevPage    uint `json:"prevPage"`
	NextPage    uint `json:"nextPage"`
}

type Vehicle struct {
	LicensePlateNumber string    `json:"licensePlateNumber"`
	VehicleModel       string    `json:"vehicleModel"`
	Vin                string    `json:"vin"`
	ControllerID       string    `json:"controllerId"`
	JudgementTime      time.Time `json:"judgementTime"`
	UserID             string    `json:"userId"`
	UserName           string    `json:"userName"`
	PhoneNumber        string    `json:"phoneNumber"`
	StartLatitude      int64     `json:"startLatitude"`
	StartLongitude     int64     `json:"startLongitude"`
	EndLatitude        int64     `json:"endLatitude"`
	EndLongitude       int64     `json:"endLongitude"`
	NotifyTime         time.Time `json:"notifyTime"`
	NotifyTarget       string    `json:"notifyTarget"`
	NotifyMsg          string    `json:"notifyMsg"`
	Memo               string
	CreatedBy          string    `json:"createdBy"`
	CreatorName        string    `json:"creatorName"`
	UpdatedBy          string    `json:"updatedBy"`
	UpdatorName        string    `json:"updatorName"`
	BusinessId         string    `json:"businessId"`
	BranchId           int64     `json:"branchId"`
	OperationTime      time.Time `json:"operationTime"`
	VehicleId          string    `json:"vehicleId"`
}

const dsn string = "root:1234@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/", getAbnormalOperationVehicleList).Methods("GET")
	//router.HandleFunc("/", GetProfileList).Methods("GET")
	//router.HandleFunc("/profile/{id}", GetProfile).Methods("GET")
	//router.HandleFunc("/insert", insertProfile).Methods("POST")
	//router.HandleFunc("/update", updateProfile).Methods("PATCH")
	//router.HandleFunc("/delete/{id}", deleteProfile).Methods("DELETE")

	http.ListenAndServe(":8080", httpHandler(router))
}

func getAbnormalOperationVehicleList(w http.ResponseWriter, r *http.Request) {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		panic("failed to connect database")
	}

	var vehicle []Vehicle
	db.Find(&vehicle)

	if len(vehicle) == 0 {
		json.NewEncoder(w).Encode(map[string]string{"result": "데이터가 없습니다."})
		return
	}

	var vehicles Vehicles
	vehicles.Vehicles = vehicle
	vehicles.PageInfo = PageInfo{
		TotalRecord: 1,
		TotalPage:   1,
		Limit:       15,
		Page:        1,
		PrevPage:    1,
		NextPage:    1,
	}

	json.NewEncoder(w).Encode(vehicles)
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

	json.NewEncoder(w).Encode(map[string]int{"status": http.StatusOK})
}

func httpHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print(r.RemoteAddr, " ", r.Proto, " ", r.Method, " ", r.URL)
		handler.ServeHTTP(w, r)
	})
}
