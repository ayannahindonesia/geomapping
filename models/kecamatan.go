package models

type (
	Kecamatan struct {
		BaseModel
		Name string `json:"name" gorm:"column:name;type:varchar(255)"`
		Kota int    `json:"kota" gorm:"column:kota_id"`
	}
)

func (p *Kecamatan) Create() (*Kecamatan, error) {
	err := Create(&p)
	return p, err
}

func (p *Kecamatan) Save() (*Kecamatan, error) {
	err := Save(&p)
	return p, err
}

func (p *Kecamatan) Delete() (*Kecamatan, error) {
	err := Delete(&p)
	return p, err
}

func (p *Kecamatan) FindbyID(id int) (*Kecamatan, error) {
	err := FindbyID(&p, id)
	return p, err
}

func (p *Kecamatan) PagedFilterSearch(page int, rows int, orderby string, sort string, filter interface{}) (result PagedSearchResult, err error) {
	kecamatan := []Kecamatan{}
	result, err = PagedFilterSearch(&kecamatan, page, rows, orderby, sort, filter)

	return result, err
}
