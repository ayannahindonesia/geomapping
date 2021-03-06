package middlewares

import (
	"fmt"
	"geomapping/geomapping"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func SetClientJWTmiddlewares(g *echo.Group, role string) {
	jwtConfig := geomapping.App.Config.GetStringMap(fmt.Sprintf("%s.jwt", geomapping.App.ENV))

	g.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS512",
		SigningKey:    []byte(jwtConfig["jwt_secret"].(string)),
	}))

	switch role {
	case "client":
		g.Use(validateJWTclient)
		break
	default:
		g.Use(validateJWTadmin)
		break
	}
}

func validateJWTadmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user")
		token := user.(*jwt.Token)

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			if claims["role"] == "admin" {
				return next(c)
			} else {
				return echo.NewHTTPError(http.StatusForbidden, fmt.Sprintf("%s", "invalid role"))
			}
		}

		return echo.NewHTTPError(http.StatusForbidden, fmt.Sprintf("%s", "invalid token"))
	}
}

func validateJWTclient(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user")
		token := user.(*jwt.Token)

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			if claims["role"] == "client" {
				return next(c)
			} else {
				return echo.NewHTTPError(http.StatusForbidden, fmt.Sprintf("%s", "invalid role"))
			}
		}

		return echo.NewHTTPError(http.StatusForbidden, fmt.Sprintf("%s", "invalid token"))
	}
}
