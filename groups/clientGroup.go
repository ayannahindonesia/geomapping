package groups

import (
	"asira_geomapping/handlers"
	"asira_geomapping/middlewares"

	"github.com/labstack/echo"
)

func ClientGroup(e *echo.Echo) {
	g := e.Group("/client")
	middlewares.SetClientJWTmiddlewares(g, "client")

	g.GET("/provinsi/kota/kecamatan/:kecamatan_id/kelurahan", handlers.ClientVillages)
	g.GET("/kelurahan/:kelurahan_id", handlers.ClientVillageDetails)
}
