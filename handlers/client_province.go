package handlers

import (
	"asira_geomapping/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func ClientProvinces(c echo.Context) error {
	defer c.Request().Body.Close()

	provinces := models.Province{}

	var filter struct{}
	result, err := provinces.GetAll(&filter)
	if err != nil {
		return returnInvalidResponse(http.StatusNotFound, err, "data tidak ditemukan")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
		"data":   result,
	})
}

func ClientProvinceDetails(c echo.Context) error {
	defer c.Request().Body.Close()

	ProvinceID, _ := strconv.Atoi(c.Param("provinsi_id"))
	provinces := models.Province{}

	result, err := provinces.FindbyID(ProvinceID)
	if err != nil {
		return returnInvalidResponse(http.StatusNotFound, err, "data tidak ditemukan")
	}

	return c.JSON(http.StatusOK, result)
}
