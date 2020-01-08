package handlers

import (
	"fmt"
	"geomapping/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/thedevsaddam/govalidator"
)

func CreateProvince(c echo.Context) error {
	defer c.Request().Body.Close()

	provinsi := models.Province{}

	payloadRules := govalidator.MapData{
		"name": []string{"required"},
	}

	validate := validateRequestPayload(c, payloadRules, &provinsi)
	if validate != nil {
		return returnInvalidResponse(http.StatusUnprocessableEntity, validate, "validation error")
	}

	err := provinsi.Create()
	if err != nil {
		return returnInvalidResponse(http.StatusInternalServerError, err, "Gagal membuat Provinsi baru")
	}

	return c.JSON(http.StatusCreated, provinsi)
}

func EditProvince(c echo.Context) error {
	defer c.Request().Body.Close()
	provinsiID, _ := strconv.ParseUint(c.Param("provinsi_id"), 10, 64)

	provinsi := models.Province{}
	err := provinsi.FindbyID(provinsiID)
	if err != nil {
		return returnInvalidResponse(http.StatusNotFound, err, fmt.Sprintf("provinsi %v tidak ditemukan", provinsiID))
	}

	payloadRules := govalidator.MapData{
		"name": []string{"required"},
	}

	validate := validateRequestPayload(c, payloadRules, &provinsi)
	if validate != nil {
		return returnInvalidResponse(http.StatusUnprocessableEntity, validate, "validation error")
	}

	err = provinsi.Save()
	if err != nil {
		return returnInvalidResponse(http.StatusInternalServerError, err, fmt.Sprintf("Gagal update Provinsi %v", provinsiID))
	}

	return c.JSON(http.StatusOK, provinsi)
}

func DeleteProvince(c echo.Context) error {
	defer c.Request().Body.Close()
	provinsiID, _ := strconv.ParseUint(c.Param("provinsi_id"), 10, 64)

	provinsi := models.Province{}
	err := provinsi.FindbyID(provinsiID)
	if err != nil {
		return returnInvalidResponse(http.StatusNotFound, err, fmt.Sprintf("provinsi %v tidak ditemukan", provinsiID))
	}

	err = provinsi.Delete()
	if err != nil {
		return returnInvalidResponse(http.StatusInternalServerError, err, fmt.Sprintf("Gagal Menghapus Provinsi %v", provinsiID))
	}

	return c.JSON(http.StatusOK, provinsi)
}
