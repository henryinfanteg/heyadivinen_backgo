package categoria

import (
	echo "github.com/labstack/echo"
)

// SetRouters setea los routers
func SetRouters(g *echo.Group) {
	var handler = CategoriaHandler{}

	g.GET("", handler.GetAll)
	g.GET("/:id", handler.GetByID)
	g.GET("/count", handler.Count)
	g.POST("", handler.Create)
	g.PUT("/:id", handler.Update)
	g.DELETE("/:id", handler.Delete)
}
