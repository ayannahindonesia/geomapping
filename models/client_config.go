package models

import "github.com/ayannahindonesia/basemodel"

// ClientConfig main type
type ClientConfig struct {
	basemodel.BaseModel
	Name   string `json:"name" gorm:"column:name"`
	Key    string `json:"key" gorm:"column:key"`
	Secret string `json:"secret" gorm:"column:secret"`
	Role   string `json:"role" gorm:"column:role"`
}

// Create func
func (model *ClientConfig) Create() error {
	return basemodel.Create(&model)
}

// Save func
func (model *ClientConfig) Save() error {
	return basemodel.Save(&model)
}

// Delete func
func (model *ClientConfig) Delete() error {
	return basemodel.Delete(&model)
}

// FindbyID func
func (model *ClientConfig) FindbyID(id uint64) error {
	return basemodel.FindbyID(&model, id)
}

// PagedFindFilter func
func (model *ClientConfig) PagedFindFilter(page int, rows int, orderby []string, sort []string, filter interface{}) (result basemodel.PagedFindResult, err error) {
	m := []ClientConfig{}
	result, err = basemodel.PagedFindFilter(&m, page, rows, orderby, sort, filter)

	return result, err
}

// FindFilter func
func (model *ClientConfig) FindFilter(order []string, sort []string, limit int, offset int, filter interface{}) (result interface{}, err error) {
	m := []ClientConfig{}
	result, err = basemodel.FindFilter(&m, order, sort, limit, offset, filter)

	return result, err
}

// SingleFindFilter func
func (model *ClientConfig) SingleFindFilter(filter interface{}) error {
	return basemodel.SingleFindFilter(&model, filter)
}
