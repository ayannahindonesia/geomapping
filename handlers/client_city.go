package handlers

import (
	"asira_geomapping/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func ClientCities(c echo.Context) error {
	defer c.Request().Body.Close()

	ProvinceID, _ := strconv.Atoi(c.Param("provinsi_id"))
	cities := models.City{}

	type Filter struct {
		ProvinceID int `json:"province_id"`
	}
	result, err := cities.GetAll(&Filter{
		ProvinceID: ProvinceID,
	})
	if err != nil {
		return returnInvalidResponse(http.StatusNotFound, err, "pencarian tidak ditemukan")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
		"data":   result,
	})
}

func ClientCityDetails(c echo.Context) error {
	defer c.Request().Body.Close()

	CityID, _ := strconv.Atoi(c.Param("kota_id"))
	cities := models.City{}

	result, err := cities.FindbyID(CityID)
	if err != nil {
		return returnInvalidResponse(http.StatusNotFound, err, "pencarian tidak ditemukan")
	}

	return c.JSON(http.StatusOK, result)
}
