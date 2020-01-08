package models

import "github.com/ayannahindonesia/basemodel"

// Province provinsi main type
type Province struct {
	basemodel.BaseModel
	Name string `json:"name" gorm:"column:name;type:varchar(255)"`
}

// Create func
func (model *Province) Create() error {
	return basemodel.Create(&model)
}

// Save func
func (model *Province) Save() error {
	return basemodel.Save(&model)
}

// Delete func
func (model *Province) Delete() error {
	return basemodel.Delete(&model)
}

// FindbyID func
func (model *Province) FindbyID(id uint64) error {
	return basemodel.FindbyID(&model, id)
}

// PagedFindFilter func
func (model *Province) PagedFindFilter(page int, rows int, orderby []string, sort []string, filter interface{}) (result basemodel.PagedFindResult, err error) {
	m := []Province{}
	result, err = basemodel.PagedFindFilter(&m, page, rows, orderby, sort, filter)

	return result, err
}

// FindFilter func
func (model *Province) FindFilter(order []string, sort []string, limit int, offset int, filter interface{}) (result interface{}, err error) {
	m := []Province{}
	result, err = basemodel.FindFilter(&m, order, sort, limit, offset, filter)

	return result, err
}

// SingleFindFilter func
func (model *Province) SingleFindFilter(filter interface{}) error {
	return basemodel.SingleFindFilter(&model, filter)
}
