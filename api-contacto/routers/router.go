package routers

import (
	echo "github.com/labstack/echo/v4"

	contactoService "github.com/henryinfanteg/heyadivinen_backgo/api-contacto/services/contacto"
)

// PATH de la api
const PATH = "/api/contacto"

// InitRoutes inicializa las rutas
func InitRoutes(e *echo.Echo) {

	// create groups
	contactoGroup := e.Group(PATH + "/contacto")

	contactoService.SetRouters(contactoGroup)
}
