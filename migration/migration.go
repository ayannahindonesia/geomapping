package migration

import (
	"geomapping/asira"
	"geomapping/models"
	"fmt"
	"strings"
)

func Seed() {
	seeder := asira.App.DB.Begin()
	defer seeder.Commit()

	if asira.App.ENV == "development" {
		// seed bank types
		bankTypes := []models.BankType{
			models.BankType{
				Name:        "BPD",
				Description: "Layanan BPD",
			},
			models.BankType{
				Name:        "BPR",
				Description: "Layanan BPR",
			},
			models.BankType{
				Name:        "Koperasi",
				Description: "Layanan Koperasi",
			},
		}
		for _, bankType := range bankTypes {
			bankType.Create()
		}

	}
}

// truncate defined tables. []string{"all"} to truncate all tables.
func Truncate(tableList []string) (err error) {
	if len(tableList) > 0 {
		if tableList[0] == "all" {
			tableList = []string{
				"kelurahan",
				"kecamatan",
				"kota",
				"provinsi",
			}
		}

		tables := strings.Join(tableList, ", ")
		sqlQuery := fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY CASCADE", tables)
		err = asira.App.DB.Exec(sqlQuery).Error
		return err
	}

	return fmt.Errorf("define tables that you want to truncate")
}
