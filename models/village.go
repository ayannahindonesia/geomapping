package models

type (
	Village struct {
		BaseModel
		Name       string `json:"name" gorm:"column:name;type:varchar(255)"`
		DistrictID int    `json:"district_id" gorm:"column:district_id"`
		ZipCode    int    `json:"zip_code" gorm:"column:zip_code"`
		AreaCode   string `json:"area_code" gorm:"column:area_code;type:varchar(255)"`
	}
)

func (p *Village) Create() (*Village, error) {
	err := Create(&p)
	return p, err
}

func (p *Village) Save() (*Village, error) {
	err := Save(&p)
	return p, err
}

func (p *Village) Delete() (*Village, error) {
	err := Delete(&p)
	return p, err
}

func (p *Village) FindbyID(id int) (*Village, error) {
	err := FindbyID(&p, id)
	return p, err
}

func (p *Village) PagedFilterSearch(page int, rows int, orderby string, sort string, filter interface{}) (result PagedSearchResult, err error) {
	kelurahan := []Village{}
	result, err = PagedFilterSearch(&kelurahan, page, rows, orderby, sort, filter)

	return result, err
}

func (p *Village) GetAll(filter interface{}) (result interface{}, err error) {
	vill := []Village{}
	result, err = GetAll(&vill, filter)

	return result, err
}

func (p *Village) FilterSearchSingle(filter interface{}) (*Village, error) {
	err := FilterSearchSingle(&p, filter)
	return p, err
}
