package router

import (
	"asira_geomapping/groups"
	"asira_geomapping/handlers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewRouter() *echo.Echo {
	e := echo.New()

	// ignore /api-asira_geomapping
	e.Pre(middleware.Rewrite(map[string]string{
		"/asira_geomapping/*": "/$1",
	}))

	// e.GET("/test", handlers.Test)
	e.GET("/clientauth", handlers.ClientLogin)

	groups.AdminGroup(e)
	groups.ClientGroup(e)

	return e
}
