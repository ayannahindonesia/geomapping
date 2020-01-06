package router

import (
	"geomapping/groups"
	"geomapping/handlers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewRouter() *echo.Echo {
	e := echo.New()

	// ignore /api-asira_geomapping
	e.Pre(middleware.Rewrite(map[string]string{
		"/api-geomapping/*": "/$1",
	}))

	// e.GET("/test", handlers.Test)
	e.GET("/clientauth", handlers.ClientLogin)

	groups.ClientGroup(e)
	groups.AdminGroup(e)

	return e
}
