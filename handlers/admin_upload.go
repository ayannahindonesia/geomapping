package handlers

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func AdminUpload(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusOK, err)
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

	//read csv
	// Open the file
	csvfile, err := os.Open("uploads/file.csv")
	if err != nil {
		return returnInvalidResponse(http.StatusInternalServerError, "", fmt.Sprint(err))
	}

	// Parse the file
	r := csv.NewReader(csvfile)
	//r := csv.NewReader(bufio.NewReader(csvfile))

	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			return c.JSON(http.StatusForbidden, err)
		}
		if err != nil {
			return c.JSON(http.StatusForbidden, err)
		}
	}

	return c.JSON(http.StatusOK, r)
}
