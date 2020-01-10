package models

import "github.com/ayannahindonesia/basemodel"

// Village kelurahan main type
type Village struct {
	basemodel.BaseModel
	Name       string `json:"name" gorm:"column:name;type:varchar(255)"`
	DistrictID int    `json:"district_id" gorm:"column:district_id"`
	ZipCode    int    `json:"zip_code" gorm:"column:zip_code"`
	AreaCode   string `json:"area_code" gorm:"column:area_code;type:varchar(255)"`
}

// Create func
func (model *Village) Create() error {
	return basemodel.Create(&model)
}

// Save func
func (model *Village) Save() error {
	return basemodel.Save(&model)
}

// Delete func
func (model *Village) Delete() error {
	return basemodel.Delete(&model)
}

// FindbyID func
func (model *Village) FindbyID(id uint64) error {
	return basemodel.FindbyID(&model, id)
}

// PagedFindFilter func
func (model *Village) PagedFindFilter(page int, rows int, orderby []string, sort []string, filter interface{}) (result basemodel.PagedFindResult, err error) {
	m := []Village{}
	result, err = basemodel.PagedFindFilter(&m, page, rows, orderby, sort, filter)

	return result, err
}

// FindFilter func
func (model *Village) FindFilter(order []string, sort []string, limit int, offset int, filter interface{}) (result interface{}, err error) {
	m := []Village{}
	result, err = basemodel.FindFilter(&m, order, sort, limit, offset, filter)

	return result, err
}

// SingleFindFilter func
func (model *Village) SingleFindFilter(filter interface{}) error {
	return basemodel.SingleFindFilter(&model, filter)
}
