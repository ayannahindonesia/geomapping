package models

type (
	Provinsi struct {
		BaseModel
		Name string `json:"name" gorm:"column:name;type:varchar(255)"`
	}
)

func (p *Provinsi) Create() (*Provinsi, error) {
	err := Create(&p)
	return p, err
}

func (p *Provinsi) Save() (*Provinsi, error) {
	err := Save(&p)
	return p, err
}

func (p *Provinsi) Delete() (*Provinsi, error) {
	err := Delete(&p)
	return p, err
}

func (p *Provinsi) FindbyID(id int) (*Provinsi, error) {
	err := FindbyID(&p, id)
	return p, err
}

func (p *Provinsi) PagedFilterSearch(page int, rows int, orderby string, sort string, filter interface{}) (result PagedSearchResult, err error) {
	provinsi := []Provinsi{}
	result, err = PagedFilterSearch(&provinsi, page, rows, orderby, sort, filter)

	return result, err
}
