package models

type (
	Kota struct {
		BaseModel
		Name     string `json:"name" gorm:"column:name;type:varchar(255)"`
		Provinsi int    `json:"provinsi" gorm:"column:provinsi_id"`
		Type     string `json:"type" gorm:"column:type;type:varchar(255)" sql:"DEFAULT:'kota'"`
	}
)

func (k *Kota) Create() (*Kota, error) {
	err := Create(&k)
	return p, err
}

func (k *Kota) Save() (*Kota, error) {
	err := Save(&k)
	return p, err
}

func (k *Kota) Delete() (*Kota, error) {
	err := Delete(&k)
	return p, err
}

func (k *Kota) FindbyID(id int) (*Kota, error) {
	err := FindbyID(&k, id)
	return p, err
}

func (k *Kota) PagedFilterSearch(page int, rows int, orderby string, sort string, filter interface{}) (result PagedSearchResult, err error) {
	kota := []Kota{}
	result, err = PagedFilterSearch(&kota, page, rows, orderby, sort, filter)

	return result, err
}
