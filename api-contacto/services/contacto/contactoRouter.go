package contacto

import (
	echo "github.com/labstack/echo/v4"
)

// SetRouters setea los routers
func SetRouters(g *echo.Group) {
	var handler = ContactoHandler{}

	g.POST("", handler.Create)
}
