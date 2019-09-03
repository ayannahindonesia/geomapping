package groups

import (
	"asira_geomapping/handlers"
	"asira_geomapping/middlewares"

	"github.com/labstack/echo"
)

func ClientGroup(e *echo.Echo) {
	g := e.Group("/client")
	middlewares.SetClientJWTmiddlewares(g, "client")

	g.GET("/provinsi", handlers.ClientProvinces)
	g.GET("/provinsi/:provinsi_id", handlers.ClientProvinceDetails)
	g.GET("/provinsi/:provinsi_id/kota", handlers.ClientCities)
	g.GET("/kota/:kota_id", handlers.ClientCityDetails)
	g.GET("/kota/:kota_id/kecamatan", handlers.ClientDistricts)
	g.GET("/kecamatan/:kecamatan_id", handlers.ClientDistrictDetails)
	g.GET("/kecamatan/:kecamatan_id/kelurahan", handlers.ClientVillages)
	g.GET("/kelurahan/:kelurahan_id", handlers.ClientVillageDetails)
}
