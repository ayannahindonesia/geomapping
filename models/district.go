package models

type (
	District struct {
		BaseModel
		Name   string `json:"name" gorm:"column:name;type:varchar(255)"`
		CityID int    `json:"city_id" gorm:"column:city_id"`
	}
)

func (p *District) Create() (*District, error) {
	err := Create(&p)
	return p, err
}

func (p *District) Save() (*District, error) {
	err := Save(&p)
	return p, err
}

func (p *District) Delete() (*District, error) {
	err := Delete(&p)
	return p, err
}

func (p *District) FindbyID(id int) (*District, error) {
	err := FindbyID(&p, id)
	return p, err
}

func (p *District) PagedFilterSearch(page int, rows int, orderby string, sort string, filter interface{}) (result PagedSearchResult, err error) {
	district := []District{}
	result, err = PagedFilterSearch(&district, page, rows, orderby, sort, filter)

	return result, err
}

func (p *District) GetAll(filter interface{}) (result interface{}, err error) {
	dist := []District{}
	result, err = GetAll(&dist, filter)

	return result, err
}

func (p *District) FilterSearchSingle(filter interface{}) (*District, error) {
	err := FilterSearchSingle(&p, filter)
	return p, err
}
