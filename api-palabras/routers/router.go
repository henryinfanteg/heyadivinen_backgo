package routers

import (
	echo "github.com/labstack/echo"

	categoriaService "github.com/henryinfanteg/heyadivinen_backgo/api-palabras/services/categoria"
	palabraService "github.com/henryinfanteg/heyadivinen_backgo/api-palabras/services/palabra"
)

// PATH de la api
const PATH = "/api/palabras"

// InitRoutes inicializa las rutas
func InitRoutes(e *echo.Echo) {

	// create groups
	palabrasGroup := e.Group(PATH + "/palabras")
	categoriaGroup := e.Group(PATH + "/categorias")

	palabraService.SetRouters(palabrasGroup)
	categoriaService.SetRouters(categoriaGroup)
}
