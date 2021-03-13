package routers

import (
	echo "github.com/labstack/echo"

	categoryService "github.com/henryinfanteg/heyadivinen_backgo/api-words/services/category"
)

// PATH de la api
const PATH = "/api/words"

// InitRoutes inicializa las rutas
func InitRoutes(e *echo.Echo) {

	// create groups
	categoryGroup := e.Group(PATH + "/categories")

	categoryService.SetRouters(categoryGroup)
}
