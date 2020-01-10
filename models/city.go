package models

import "github.com/ayannahindonesia/basemodel"

// City kota main type
type City struct {
	basemodel.BaseModel
	Name       string `json:"name" gorm:"column:name;type:varchar(255)"`
	ProvinceID int    `json:"province_id" gorm:"column:province_id"`
	Type       string `json:"type" gorm:"column:type;type:varchar(255)" sql:"DEFAULT:'kota'"`
}

// Create func
func (model *City) Create() error {
	return basemodel.Create(&model)
}

// Save func
func (model *City) Save() error {
	return basemodel.Save(&model)
}

// Delete func
func (model *City) Delete() error {
	return basemodel.Delete(&model)
}

// FindbyID func
func (model *City) FindbyID(id uint64) error {
	return basemodel.FindbyID(&model, id)
}

// PagedFindFilter func
func (model *City) PagedFindFilter(page int, rows int, orderby []string, sort []string, filter interface{}) (result basemodel.PagedFindResult, err error) {
	m := []City{}
	result, err = basemodel.PagedFindFilter(&m, page, rows, orderby, sort, filter)

	return result, err
}

// FindFilter func
func (model *City) FindFilter(order []string, sort []string, limit int, offset int, filter interface{}) (result interface{}, err error) {
	m := []City{}
	result, err = basemodel.FindFilter(&m, order, sort, limit, offset, filter)

	return result, err
}

// SingleFindFilter func
func (model *City) SingleFindFilter(filter interface{}) error {
	return basemodel.SingleFindFilter(&model, filter)
}
