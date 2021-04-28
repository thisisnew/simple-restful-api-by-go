package vehicle

import (
	"encoding/json"
	"net/http"
	connection "rest/db"
)

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
	Id                   string    `json:"ID"`
	Name                 string    `json:"name"`
	Vin                  string    `json:"VIN"`
	LicensePlateNumber   string    `json:"licensePlateNumber"`
	ControllerId         string    `json:"controllerId"`
	RegistrationDate     string    `json:"registrationDate"`
	ProductionDate       string    `json:"productionDate"`
	GarageAddress        string    `json:"garageAddress"`
	Color                string    `json:"color"`
	Transmission         string    `json:"transmission"`
	Business             Business  `json:"business" gorm:"-"`
	Location             string    `json:"location"`
	Model                Model     `json:"model" gorm:"-"`
	Insurance            Insurance `json:"insurance" gorm:"-"`
	Age                  uint      `json:"age"`
	Available            bool      `json:"available"`
	State                State     `json:"state" gorm:"-"`
	Payments             string    `json:"payments"`
	Branch               Branch    `json:"branch" gorm:"-"`
	Order                uint      `json:"order"`
	VehicleType          string    `json:"vehicleType"`
	ModelYear            uint      `json:"modelYear"`
	ReservationUsage     string    `json:"reservationUsage"`
	ReservationStatus    string    `json:"reservationStatus"`
	ReservationStartDate uint      `json:"reservationStartDate"`
	ReservationEndDate   uint      `json:"reservationEndDate"`
}

type Branch struct {
	Id                 uint   `json:"id"`
	BusinessID         string `json:"businessId"`
	Business           Business
	Name               string `json:"name"`
	BusinessRegNum     string `json:"businessRegNum"`
	CompanyBusiness    string `json:"companyBusiness"`
	BusinessType       string `json:"businessType"`
	PhoneNumber        string `json:"phoneNumber"`
	FaxNumber          string `json:"faxNumber"`
	ManagerName        string `json:"managerName"`
	ManagerPhoneNumber string `json:"managerPhoneNumber"`
	ManagerEmail       string `json:"managerEmail"`
	BankName           string `json:"bankName"`
	AccountNumber      string `json:"accountNumber"`
	AccountHolder      string `json:"accountHolder"`
}

type Business struct {
	Id           string `json:"ID"`
	Name         string `json:"name"`
	BusinessType string `json:"businessType"`
	PublicRide   bool   `json:"publicRide"`
	PrivateRide  bool   `json:"privateRide"`
	TestRide     bool   `json:"testRide"`
	AutoConfirm  bool   `json:"autoConfirm"`
	BusinessPath bool   `json:"businessPath"`
	PersonalPath bool   `json:"personalPath"`
	BusinessUse  bool   `json:"businessUse"`
	PersonalUse  bool   `json:"personalUse"`
}

type Insurance struct {
	Age      uint   `json:"age"`
	Company  string `json:"company"`
	Pd       string `json:"PD"`
	Bi1      string `json:"BI1"`
	Bi2      string `json:"BI2"`
	Db       string `json:"DB"`
	Uvd      string `json:"UVD"`
	Ibi      string `json:"IBI"`
	Ipd      string `json:"IPD"`
	Idb      string `json:"IDB"`
	Isd      string `json:"ISD"`
	JoinDate string `json:"joinDate"`
}

type Model struct {
	Id               string `json:"ID"`
	Name             string `json:"name"`
	Brand            string `json:"brand"`
	Standard         string `json:"standard"`
	StandardModelId  string `json:"standardModelId"`
	SeatingCapacity  string `json:"seatingCapacity"`
	FuelType         string `json:"fuelType"`
	FuelEfficiency   string `json:"fuelEfficiency"`
	FuelTankCapacity string `json:"fuelTankCapacity"`
	Displacement     string `json:"displacement"`
	Grade            string `json:"grade"`
	WarmUpTime       string `json:"warmUpTime"`
	ImageUrl         string `json:"imageUrl"`
}

type State struct {
	Latitude  uint   `json:"latitude"`
	Longitude uint   `json:"longitude"`
	Odometer  string `json:"odometer"`
	Dte       string `json:"DTE"`
	FuelLevel string `json:"fuelLevel"`
	Soc       uint   `json:"SOC"`
	Charging  bool   `json:"charging"`
	Battery   string `json:"battery"`
	Speed     string `json:"speed"`
	TpmsFl    string `json:"tpmsFl"`
	TpmsFr    string `json:"tpmsFr"`
	TpmsRl    string `json:"tpmsRl"`
	TpmsRr    string `json:"tpmsRr"`
}

func getAbnormalOperationVehicleList(w http.ResponseWriter, r *http.Request) {

	db := connection.GetDB()

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

func GetAvailableVehicleList(w http.ResponseWriter, r *http.Request) {

	db := connection.GetDB()

	var vehicle Vehicle
	db.Table("vehicle").Find(&vehicle)

	if vehicle.Id == "" {
		json.NewEncoder(w).Encode(map[string]string{"result": "데이터가 없습니다."})
		return
	}

	var model Model
	db.Table("model").Find(&model)

	var branch Branch
	db.Table("branch").Find(&branch)

	var business Business
	db.Table("business").Find(&business)

	var insurance Insurance
	db.Table("insurance").Find(&insurance)

	var state State
	db.Table("state").Find(&state)

	vehicle.Business = business
	vehicle.Insurance = insurance
	vehicle.Model = model
	vehicle.Branch = branch
	vehicle.State = state

	var vehicles Vehicles
	vehicles.Vehicles = []Vehicle{vehicle}
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
