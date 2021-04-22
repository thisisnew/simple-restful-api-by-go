package billing

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

type Items struct {
	Billings Billings `json:"items"`
	PageInfo PageInfo `json:"pageInfo"`
}

type Billings struct {
	Business Business `json:"business"`
	Branch   Branch   `json:"branch"`

	ShortBillingCount      uint `json:"shortBillingCount"`
	ShortBillingAmount     uint `json:"shortBillingAmount"`
	ShortDepositAmount     uint `json:"shortDepositAmount"`
	InsuranceBillingCount  uint `json:"insuranceBillingCount"`
	InsuranceBillingAmount uint `json:"insuranceBillingAmount"`
	InsuranceDepositAmount uint `json:"insuranceDepositAmount"`
	LongBillingCount       uint `json:"longBillingCount"`
	LongBillingAmount      uint `json:"longBillingAmount"`
	LongDepositAmount      uint `json:"longDepositAmount"`
	MonthBillingCount      uint `json:"monthBillingCount"`
	MonthBillingAmount     uint `json:"monthBillingAmount"`
	MonthDepositAmount     uint `json:"monthDepositAmount"`
	BillingCount           uint `json:"billingCount"`
	BillingAmount          uint `json:"billingAmount"`
	DepositAmount          uint `json:"depositAmount"`
}

type PageInfo struct {
	TotalRecord uint `json:"totalRecord"`
	TotalPage   uint `json:"totalPage"`
	Limit       uint `json:"limit"`
	Page        uint `json:"page"`
	PrevPage    uint `json:"prevPage"`
	NextPage    uint `json:"nextPage"`
}

type Business struct {
	Id               string  `gorm:"column:id" json:"ID"`
	Name             string  `gorm:"column:name" json:"name"`
	BusinessRegNum   string  `gorm:"column:business_reg_num" json:"businessRegNum"`
	TotAddress       Address `gorm:"-" json:"address"`
	ZipCode          string  `gorm:"column:zip_code" json:"-"`
	Address          string  `gorm:"column:address" json:"-"`
	DetailAddress    string  `gorm:"column:detail_address" json:"-"`
	IsPrivateCompany bool    `gorm:"column:is_private_company" json:"isPrivateCompany"`
}

type Branch struct {
	Id                 string  `gorm:"column:id" json:"id"`
	BusinessId         string  `gorm:"column:business_id" json:"businessID"`
	Name               string  `gorm:"column:name" json:"name"`
	BusinessRegNum     string  `gorm:"column:business_reg_num" json:"businessRegNum"`
	CompanyBusiness    string  `gorm:"column:company_business" json:"companyBusiness"`
	BusinessType       string  `gorm:"column:business_type" json:"businessType"`
	PhoneNumber        string  `gorm:"column:phone_number" json:"phoneNumber"`
	FaxNumber          string  `gorm:"column:fax_number" json:"faxNumber"`
	TotAddress         Address `gorm:"-" json:"address"`
	ZipCode            string  `gorm:"column:zip_code" json:"-"`
	Address            string  `gorm:"column:address" json:"-"`
	DetailAddress      string  `gorm:"column:detail_address" json:"-"`
	ManagerName        string  `gorm:"column:manager_name" json:"managerName"`
	ManagerPhoneNumber string  `gorm:"column:manager_phone_number" json:"managerPhoneNumber"`
	ManagerEmail       string  `gorm:"column:manager_email" json:"managerEmail"`
}

type Address struct {
	ZipCode       string `gorm:"column:zip_code" json:"zipCode"`
	Address       string `gorm:"column:address" json:"address"`
	DetailAddress string `gorm:"column:detail_address" json:"detailedAddress"`
}

const dsn string = "root:1234@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"

func GetBillingList(w http.ResponseWriter, r *http.Request) {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		panic("failed to connect database")
	}

	var business Business
	db.Table("business").Find(&business)
	business.TotAddress = Address{
		ZipCode:       business.ZipCode,
		Address:       business.Address,
		DetailAddress: business.DetailAddress,
	}

	var branch Branch
	db.Table("branch").Find(&branch)
	branch.TotAddress = Address{
		ZipCode:       branch.ZipCode,
		Address:       branch.Address,
		DetailAddress: branch.DetailAddress,
	}

	var billings Billings
	billings.Business = business
	billings.Branch = branch

	billings.ShortBillingCount = 0
	billings.ShortBillingAmount = 0
	billings.ShortDepositAmount = 0
	billings.InsuranceBillingCount = 0
	billings.InsuranceBillingAmount = 0
	billings.InsuranceDepositAmount = 0
	billings.LongBillingCount = 0
	billings.LongBillingAmount = 0
	billings.LongDepositAmount = 0
	billings.MonthBillingCount = 0
	billings.MonthBillingAmount = 0
	billings.MonthDepositAmount = 0
	billings.BillingCount = 0
	billings.BillingAmount = 0
	billings.DepositAmount = 0

	var items Items
	items.Billings = billings
	items.PageInfo = PageInfo{
		TotalRecord: 1,
		TotalPage:   1,
		Limit:       15,
		Page:        1,
		PrevPage:    1,
		NextPage:    1,
	}

	json.NewEncoder(w).Encode(items)
}
