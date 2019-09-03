package models

type (
	City struct {
		BaseModel
		Name       string `json:"name" gorm:"column:name;type:varchar(255)"`
		ProvinceID int    `json:"province_id" gorm:"column:province_id"`
		Type       string `json:"type" gorm:"column:type;type:varchar(255)" sql:"DEFAULT:'kota'"`
	}
)

func (k *City) Create() (*City, error) {
	err := Create(&k)
	return k, err
}

func (k *City) Save() (*City, error) {
	err := Save(&k)
	return k, err
}

func (k *City) Delete() (*City, error) {
	err := Delete(&k)
	return k, err
}

func (k *City) FindbyID(id int) (*City, error) {
	err := FindbyID(&k, id)
	return k, err
}

func (k *City) PagedFilterSearch(page int, rows int, orderby string, sort string, filter interface{}) (result PagedSearchResult, err error) {
	city := []City{}
	result, err = PagedFilterSearch(&city, page, rows, orderby, sort, filter)

	return result, err
}

func (k *City) GetAll(filter interface{}) (result interface{}, err error) {
	cities := []City{}
	result, err = GetAll(&cities, filter)

	return result, err
}
