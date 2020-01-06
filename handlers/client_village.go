package handlers

import (
	"geomapping/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func ClientVillages(c echo.Context) error {
	defer c.Request().Body.Close()

	DistrictID, _ := strconv.Atoi(c.Param("kecamatan_id"))
	village := models.Village{}

	type Filter struct {
		DistrictID int `json:"district_id"`
	}

	result, err := village.GetAll(&Filter{
		DistrictID: DistrictID,
	})
	if err != nil {
		return returnInvalidResponse(http.StatusNotFound, err, "pencarian tidak ditemukan")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
		"data":   result,
	})
}

func ClientVillageDetails(c echo.Context) error {
	defer c.Request().Body.Close()

	VillageID, _ := strconv.Atoi(c.Param("kelurahan_id"))
	village := models.Village{}

	result, err := village.FindbyID(VillageID)
	if err != nil {
		return returnInvalidResponse(http.StatusNotFound, err, "pencarian tidak ditemukan")
	}

	return c.JSON(http.StatusOK, result)
}
