package routers

import (
	echo "github.com/labstack/echo"

	categoriaService "gitlab.com/adivinagame/backend/maxadivinabackend/api-palabras/services/categoria"
	palabraService "gitlab.com/adivinagame/backend/maxadivinabackend/api-palabras/services/palabra"
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
