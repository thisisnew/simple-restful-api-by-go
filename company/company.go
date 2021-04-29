package company

import (
	"encoding/json"
	"net/http"
	database "rest/db"
)

type Companies struct {
	Companies []Company `json:"items"`
	PageInfo  PageInfo  `json:"pageInfo"`
}

type PageInfo struct {
	TotalRecord uint `json:"totalRecord"`
	TotalPage   uint `json:"totalPage"`
	Limit       uint `json:"limit"`
	Page        uint `json:"page"`
	PrevPage    uint `json:"prevPage"`
	NextPage    uint `json:"nextPage"`
}

type Company struct {
	Id              uint    `json:"ID"`
	CreatedBy       string  `json:"createdBy"`
	CreatorName     string  `json:"creatorName"`
	Name            string  `json:"name"`
	BusinessId      string  `json:"businessID"`
	BranchId        string  `json:"branchID"`
	CompanyType     uint    `json:"companyType"`
	BusinessRegNum  string  `json:"businessRegNum"`
	CeoName         string  `json:"ceoName"`
	PhoneNumber     string  `json:"phoneNumber"`
	Homepage        string  `json:"homepage"`
	CompanyBusiness string  `json:"companyBusiness"`
	RepBusinessType string  `json:"repBusinessType"`
	BusinessType    string  `json:"businessType"`
	Address         Address `json:"address" gorm:"-"`
}

type Address struct {
	Id            uint   `json:"-"`
	ZipCode       string `gorm:"column:zip_code" json:"zipCode"`
	Address       string `gorm:"column:address" json:"address"`
	DetailAddress string `gorm:"column:detail_address" json:"detailedAddress"`
}

func GetCompanyList(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()

	var companies []Company
	db.Table("company").Find(&companies)

	var address []Address
	db.Table("address").Find(&address)

	for _, company := range companies {
		for _, addr := range address {
			var companyId = company.Id
			var addrId = addr.Id
			if companyId == addrId {
				company.Address = Address{
					Id:            addr.Id,
					ZipCode:       addr.ZipCode,
					Address:       addr.Address,
					DetailAddress: addr.DetailAddress,
				}
			}
		}
	}

	var items Companies
	items.Companies = companies
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
