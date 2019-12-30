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

	newProvince, err := provinsi.Create()
	if err != nil {
		return returnInvalidResponse(http.StatusInternalServerError, err, "Gagal membuat Provinsi baru")
	}

	return c.JSON(http.StatusCreated, newProvince)
}

func EditProvince(c echo.Context) error {
	defer c.Request().Body.Close()
	provinsi_id, _ := strconv.Atoi(c.Param("provinsi_id"))

	provinsi := models.Province{}
	_, err := provinsi.FindbyID(provinsi_id)
	if err != nil {
		return returnInvalidResponse(http.StatusNotFound, err, fmt.Sprintf("provinsi %v tidak ditemukan", provinsi_id))
	}

	payloadRules := govalidator.MapData{
		"name": []string{"required"},
	}

	validate := validateRequestPayload(c, payloadRules, &provinsi)
	if validate != nil {
		return returnInvalidResponse(http.StatusUnprocessableEntity, validate, "validation error")
	}

	_, err = provinsi.Save()
	if err != nil {
		return returnInvalidResponse(http.StatusInternalServerError, err, fmt.Sprintf("Gagal update Provinsi %v", provinsi_id))
	}

	return c.JSON(http.StatusOK, provinsi)
}

func DeleteProvince(c echo.Context) error {
	defer c.Request().Body.Close()
	provinsi_id, _ := strconv.Atoi(c.Param("provinsi_id"))

	provinsi := models.Province{}
	_, err := provinsi.FindbyID(provinsi_id)
	if err != nil {
		return returnInvalidResponse(http.StatusNotFound, err, fmt.Sprintf("provinsi %v tidak ditemukan", provinsi_id))
	}

	_, err = provinsi.Delete()
	if err != nil {
		return returnInvalidResponse(http.StatusInternalServerError, err, fmt.Sprintf("Gagal Menghapus Provinsi %v", provinsi_id))
	}

	return c.JSON(http.StatusOK, provinsi)
}
