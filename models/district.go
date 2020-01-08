package models

import "github.com/ayannahindonesia/basemodel"

// District kecamatan main type
type District struct {
	basemodel.BaseModel
	Name   string `json:"name" gorm:"column:name;type:varchar(255)"`
	CityID int    `json:"city_id" gorm:"column:city_id"`
}

// Create func
func (model *District) Create() error {
	return basemodel.Create(&model)
}

// Save func
func (model *District) Save() error {
	return basemodel.Save(&model)
}

// Delete func
func (model *District) Delete() error {
	return basemodel.Delete(&model)
}

// FindbyID func
func (model *District) FindbyID(id uint64) error {
	return basemodel.FindbyID(&model, id)
}

// PagedFindFilter func
func (model *District) PagedFindFilter(page int, rows int, orderby []string, sort []string, filter interface{}) (result basemodel.PagedFindResult, err error) {
	m := []District{}
	result, err = basemodel.PagedFindFilter(&m, page, rows, orderby, sort, filter)

	return result, err
}

// FindFilter func
func (model *District) FindFilter(order []string, sort []string, limit int, offset int, filter interface{}) (result interface{}, err error) {
	m := []District{}
	result, err = basemodel.FindFilter(&m, order, sort, limit, offset, filter)

	return result, err
}

// SingleFindFilter func
func (model *District) SingleFindFilter(filter interface{}) error {
	return basemodel.SingleFindFilter(&model, filter)
}
