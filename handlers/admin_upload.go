package handlers

import (
	"encoding/csv"
	"fmt"
	"geomapping/models"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

// AdminUpload uploads csv sent by api
func AdminUpload(c echo.Context) error {
	defer c.Request().Body.Close()

	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	src, err := file.Open()
	if err != nil {
		return returnInvalidResponse(http.StatusInternalServerError, "", fmt.Sprint(err))
	}
	defer src.Close()

	// Destination
	dst, err := os.Create("uploads/file.csv")
	if err != nil {
		return returnInvalidResponse(http.StatusInternalServerError, "", fmt.Sprint(err))
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return returnInvalidResponse(http.StatusInternalServerError, "", fmt.Sprint(err))
	}

	// Open the file
	csvfile, err := os.Open("uploads/file.csv")
	if err != nil {
		return returnInvalidResponse(http.StatusInternalServerError, "", fmt.Sprint(err))
	}

	// Parse the file
	r := csv.NewReader(csvfile)
	//r := csv.NewReader(bufio.NewReader(csvfile))
	i := 0
	// Iterate through the records
	for {
		if i%50 == 0 {
			time.Sleep(2 * time.Second)
		}
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"status":  true,
				"message": "Data Berhasil Di input",
			})
		}

		type Filter struct {
			Name string `json:"name"`
		}

		//provinsi
		prov := models.Province{}
		err = prov.SingleFindFilter(&Filter{
			Name: record[7],
		})
		if err != nil {
			prov.Name = record[7]
			prov.Create()
		}

		//Kota
		type FilterCity struct {
			Name       string `json:"name"`
			ProvinceID uint64 `json:"province_id"`
		}
		city := models.City{}
		err = city.SingleFindFilter(&FilterCity{
			Name:       record[6],
			ProvinceID: prov.ID,
		})
		if err != nil {
			city.Name = record[6]
			city.ProvinceID = int(prov.ID)
			city.Type = record[5]
			city.Create()
		}

		//Kecamatan
		type FilterDistrict struct {
			Name   string `json:"name"`
			CityID uint64 `json:"city_id"`
		}
		district := models.District{}
		err = district.SingleFindFilter(&FilterDistrict{
			Name:   record[4],
			CityID: city.ID,
		})
		if err != nil {
			district.Name = record[4]
			district.CityID = int(city.ID)
			district.Create()
		}

		//Kelurahan
		type FilterVillage struct {
			Name       string `json:"name"`
			DistrictID uint64 `json:"district_id"`
		}
		village := models.Village{}
		err = village.SingleFindFilter(&FilterVillage{
			Name:       record[2],
			DistrictID: district.ID,
		})
		if err != nil {
			village.Name = record[2]
			village.ZipCode, _ = strconv.Atoi(record[1])
			village.AreaCode = record[3]
			village.DistrictID = int(district.ID)
			village.Create()
		}
		i++
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "Data Berhasil Di input",
	})
}
