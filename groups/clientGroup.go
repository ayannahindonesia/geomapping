package groups

import (
	"geomapping/middlewares"

	"github.com/labstack/echo"
)

func ClientGroup(e *echo.Echo) {
	g := e.Group("/client")
	middlewares.SetClientJWTmiddlewares(g, "client")

	g.GET("/provinsi", handlers.ClientProvinces)
	g.GET("/provinsi/:provinsi_id", handlers.ClientProvinceDetails)

}
