package palabra

import (
	echo "github.com/labstack/echo"
)

// SetRouters setea los routers
func SetRouters(g *echo.Group) {
	var handler = PalabraHandler{}

	g.GET("", handler.GetAll)
	g.GET("/:id", handler.GetAllByCategoriaID)
	g.GET("/count", handler.Count)
	g.POST("", handler.Create)
	g.PUT("/:id", handler.Update)
	g.DELETE("/:id", handler.Delete)

}
