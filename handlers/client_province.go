package handlers

import (
	"geomapping/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// ClientProvinces client province list
func ClientProvinces(c echo.Context) error {
	defer c.Request().Body.Close()

	provinces := models.Province{}

	var filter struct{}
	result, err := provinces.FindFilter([]string{}, []string{}, 0, 0, &filter)
	if err != nil {
		return returnInvalidResponse(http.StatusNotFound, err, "data tidak ditemukan")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": true,
		"data":   result,
	})
}

// ClientProvinceDetails province details
func ClientProvinceDetails(c echo.Context) error {
	defer c.Request().Body.Close()

	ProvinceID, _ := strconv.ParseUint(c.Param("provinsi_id"), 10, 64)
	provinces := models.Province{}

	err := provinces.FindbyID(ProvinceID)
	if err != nil {
		return returnInvalidResponse(http.StatusNotFound, err, "data tidak ditemukan")
	}

	return c.JSON(http.StatusOK, provinces)
}
