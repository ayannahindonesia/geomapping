package handlers

import (
	"geomapping/geomapping"
	"net/http"

	"github.com/labstack/echo"
)

func AsiraAppInfo(c echo.Context) error {
	defer c.Request().Body.Close()

	type AppInfo struct {
		AppName string                 `json:"app_name"`
		Version string                 `json:"version"`
		ENV     string                 `json:"env"`
		Config  map[string]interface{} `json:"configs"`
	}

	var show AppInfo

	show.AppName = geomapping.App.Name
	show.Version = geomapping.App.Version
	show.ENV = geomapping.App.ENV
	show.Config = geomapping.App.Config.GetStringMap(geomapping.App.ENV)

	return c.JSON(http.StatusOK, show)
}
