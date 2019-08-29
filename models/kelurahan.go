package models

type (
	Kelurahan struct {
		BaseModel
		Name      string `json:"name" gorm:"column:name;type:varchar(255)"`
		Kecamatan int    `json:"kecamatan" gorm:"column:kecamatan_id"`
		ZipCode   int    `json:"zip_code" gorm:"column:zip_code"`
		AreaCode  string `json:"area_code" gorm:"column:area_code;type:varchar(255)"`
	}
)

func (p *Kelurahan) Create() (*Kelurahan, error) {
	err := Create(&p)
	return p, err
}

func (p *Kelurahan) Save() (*Kelurahan, error) {
	err := Save(&p)
	return p, err
}

func (p *Kelurahan) Delete() (*Kelurahan, error) {
	err := Delete(&p)
	return p, err
}

func (p *Kelurahan) FindbyID(id int) (*Kelurahan, error) {
	err := FindbyID(&p, id)
	return p, err
}

func (p *Kelurahan) PagedFilterSearch(page int, rows int, orderby string, sort string, filter interface{}) (result PagedSearchResult, err error) {
	kelurahan := []Kelurahan{}
	result, err = PagedFilterSearch(&kelurahan, page, rows, orderby, sort, filter)

	return result, err
}
