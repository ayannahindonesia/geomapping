package groups

import (
	"asira_geomapping/handlers"
	"asira_geomapping/middlewares"

	"github.com/labstack/echo"
)

func AdminGroup(e *echo.Echo) {
	g := e.Group("/admin")
	middlewares.SetClientJWTmiddlewares(g, "admin")

	// config info
	g.GET("/info", handlers.AsiraAppInfo)
	//Create Client Config
	g.POST("/client_config", handlers.CreateClientConfig)

}
