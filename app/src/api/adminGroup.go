package api

import (
    "github.com/labstack/echo"
    "api/handlers"
)

func AdminGroup(g *echo.Group) {
    g.GET("/", handlers.MainAdmin)
    g.GET("/logout") handlers.Logout)
}
