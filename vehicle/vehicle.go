package vehicle

import (
	"encoding/json"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net/http"
	"os"
	"time"
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
