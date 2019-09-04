package models

type (
	Province struct {
		BaseModel
		Name string `json:"name" gorm:"column:name;type:varchar(255)"`
	}
)

func (p *Province) Create() (*Province, error) {
	err := Create(&p)
	return p, err
}

func (p *Province) Save() (*Province, error) {
	err := Save(&p)
	return p, err
}

func (p *Province) Delete() (*Province, error) {
	err := Delete(&p)
	return p, err
}

func (p *Province) FindbyID(id int) (*Province, error) {
	err := FindbyID(&p, id)
	return p, err
}

func (p *Province) PagedFilterSearch(page int, rows int, orderby string, sort string, filter interface{}) (result PagedSearchResult, err error) {
	province := []Province{}
	result, err = PagedFilterSearch(&province, page, rows, orderby, sort, filter)

	return result, err
}

func (p *Province) GetAll(filter interface{}) (result interface{}, err error) {
	province := []Province{}
	result, err = GetAll(&province, filter)

	return result, err
}

func (p *Province) FilterSearchSingle(filter interface{}) (*Province, error) {
	err := FilterSearchSingle(&p, filter)
	return p, err
}
