package handlers

import (
	"asira_geomapping/models"
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/labstack/echo"
)

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
		if i%100 == 0 {
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
		_, err = prov.FilterSearchSingle(&Filter{
			Name: record[7],
		})
		if err != nil {
			prov.Name = record[7]
			prov.Create()
		}

		//Kota
		city := models.City{}
		_, err = city.FilterSearchSingle(&Filter{
			Name: record[6],
		})
		if err != nil {
			city.Name = record[6]
			city.ProvinceID = int(prov.BaseModel.ID)
			city.Type = record[5]
			city.Create()
		}

		//Kecamatan
		district := models.District{}
		_, err = district.FilterSearchSingle(&Filter{
			Name: record[4],
		})
		if err != nil {
			district.Name = record[4]
			district.CityID = int(city.BaseModel.ID)
			district.Create()
		}

		//Kelurahan
		village := models.Village{}
		_, err = village.FilterSearchSingle(&Filter{
			Name: record[2],
		})
		if err != nil {
			village.Name = record[4]
			village.ZipCode, _ = strconv.Atoi(record[1])
			village.AreaCode = record[3]
			village.DistrictID = int(district.BaseModel.ID)
			village.Create()
		}
		i++
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  true,
		"message": "Data Berhasil Di input",
	})
}
