package migration

import (
	"fmt"
	"geomapping/asira"
	"geomapping/models"
	"strings"
)

func Seed() {
	seeder := asira.App.DB.Begin()
	defer seeder.Commit()

	if asira.App.ENV == "development" {
		// seed province
		province := []models.Province{
			models.Province{
				Name: "DKI Jakarta",
			},
			models.Province{
				Name: "Jawa Barat",
			},
			models.Province{
				Name: "Sumatera Selatan",
			},
		}
		for _, prov := range province {
			prov.Create()
		}

		// seed province
		city := []models.City{
			models.City{
				Name:       "Jakarta Selatan",
				ProvinceID: 1,
			},
			models.City{
				Name:       "Jakarta Barat",
				ProvinceID: 1,
			},
			models.City{
				Name:       "Jakarta Pusat",
				ProvinceID: 1,
			},
			models.City{
				Name:       "Jakarta Utara",
				ProvinceID: 1,
			},
			models.City{
				Name:       "Bandung",
				ProvinceID: 2,
			},
			models.City{
				Name:       "Palembang",
				ProvinceID: 3,
			},
		}
		for _, k := range city {
			k.Create()
		}

		district := []models.District{
			models.District{
				Name:   "Mampang Prapatan",
				CityID: 1,
			},
			models.District{
				Name:   "Setiabudi",
				CityID: 1,
			},
			models.District{
				Name:   "Sekip",
				CityID: 6,
			},
		}
		for _, kec := range district {
			kec.Create()
		}

		village := []models.Village{
			models.Village{
				Name:       "Mampang Prapatan",
				DistrictID: 1,
				ZipCode:    12790,
				AreaCode:   "31.74.03.1001",
			},
			models.Village{
				Name:       "Kuningan Barat",
				DistrictID: 1,
				ZipCode:    12710,
				AreaCode:   "31.74.03.1005",
			},
			models.Village{
				Name:       "Sekip Jaya",
				DistrictID: 3,
				ZipCode:    30128,
				AreaCode:   "16.71.09.1003",
			},
		}
		for _, kel := range village {
			kel.Create()
		}
	}
}

// truncate defined tables. []string{"all"} to truncate all tables.
func Truncate(tableList []string) (err error) {
	if len(tableList) > 0 {
		if tableList[0] == "all" {
			tableList = []string{
				"villages",
				"districts",
				"cities",
				"provinces",
			}
		}

		tables := strings.Join(tableList, ", ")
		sqlQuery := fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY CASCADE", tables)
		err = asira.App.DB.Exec(sqlQuery).Error
		return err
	}

	return fmt.Errorf("define tables that you want to truncate")
}
