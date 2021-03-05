package util

import (
	"strconv"

	echo "github.com/labstack/echo"
)

// GetParametroSort obtiene el parametro de ordenamiento
func GetParametroSort(c echo.Context) string {
	return c.QueryParam("sort")
}

// GetParametrosPaginacion obtiene los parametros de paginacion
func GetParametrosPaginacion(c echo.Context) (int, int) {
	pageNumber, _ := strconv.Atoi(c.QueryParam("page[number]"))
	pageSize, _ := strconv.Atoi(c.QueryParam("page[size]"))
	return pageNumber, pageSize
}
