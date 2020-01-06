package migration

import (
	"fmt"
	"geomapping/geomapping"
	"geomapping/models"
	"strings"
)

func Seed() {
	seeder := geomapping.App.DB.Begin()
	defer seeder.Commit()

	if geomapping.App.ENV == "development" {
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

		client := []models.Client_config{
			models.Client_config{
				Name:   "admin",
				Key:    "admin",
				Role:   "admin",
				Secret: "adminkey",
			},
			models.Client_config{
				Name:   "client",
				Key:    "client",
				Role:   "client",
				Secret: "clientgeo",
			},
		}
		for _, clients := range client {
			clients.Create()
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
				"client_configs",
			}
		}

		tables := strings.Join(tableList, ", ")
		sqlQuery := fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY CASCADE", tables)
		err = geomapping.App.DB.Exec(sqlQuery).Error
		return err
	}

	return fmt.Errorf("define tables that you want to truncate")
}
