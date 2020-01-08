package handlers

import (
	"geomapping/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func ClientDistricts(c echo.Context) error {
	defer c.Request().Body.Close()

	CityID, _ := strconv.Atoi(c.Param("kota_id"))
	district := models.District{}

	type Filter struct {
		CityID int `json:"city_id"`
	}

	result, err := district.FindFilter([]string{}, []string{}, 0, 0, &Filter{
		CityID: CityID,
	})
	if err != nil {
		return returnInvalidResponse(http.StatusNotFound, err, "pencarian tidak ditemukan")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
		"data":   result,
	})
}

func ClientDistrictDetails(c echo.Context) error {
	defer c.Request().Body.Close()

	DistrictID, _ := strconv.ParseUint(c.Param("kecamatan_id"), 10, 64)
	district := models.District{}

	err := district.FindbyID(DistrictID)
	if err != nil {
		return returnInvalidResponse(http.StatusNotFound, err, "data tidak ditemukan")
	}

	return c.JSON(http.StatusOK, district)
}
