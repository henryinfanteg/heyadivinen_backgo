package main

import (
	"fmt"
	"net/http"

	echo "github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"gitlab.com/adivinagame/backend/maxadivinabackend/api-palabras/config"
	"gitlab.com/adivinagame/backend/maxadivinabackend/api-palabras/routers"
	db "gitlab.com/adivinagame/backend/maxadivinabackend/db-mongo/server"
	apiUtil "gitlab.com/adivinagame/backend/maxadivinabackend/util-api/util"
	auditoriaConfig "gitlab.com/adivinagame/backend/maxadivinabackend/util-auditoria/config"
	loggerConfig "gitlab.com/adivinagame/backend/maxadivinabackend/util-logger/config"
	logger "gitlab.com/adivinagame/backend/maxadivinabackend/util-logger/util"
)

func init() {
	// Cargamos la configuracion inicial
	config.LoadConfigFile()

	// Mapeamos las conexiones
	conectionDB := db.ConectionDB(config.GetConnectionConfig().Database)
	loggerConfig.LoadConfigDB(&conectionDB)
	auditoriaConfig.LoadConfigDB(&conectionDB)
}

func main() {
	fmt.Println("")
	fmt.Println("*********************************************")
	fmt.Println("*************************** API PALABRAS ***")
	fmt.Println("*********************************************")
	fmt.Println("_____________________________________________")
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE, echo.OPTIONS},
		ExposeHeaders: []string{apiUtil.RequestID},
	}))
	e.Use(middlewareValidarPermiso)

	routers.InitRoutes(e)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Api - Palabras!")
	})

	e.Logger.Fatal(e.Start(":3000"))
}

func middlewareValidarPermiso(next echo.HandlerFunc) echo.HandlerFunc {
	// return a HandlerFunc
	return func(c echo.Context) error {
		// Mapeamos los headers al response
		c.Response().Header().Set(apiUtil.RequestID, c.Request().Header.Get(apiUtil.RequestID))

		// Obtenemos el nombre de la api
		nombreAPI := apiUtil.GetAPIToPath("palabras", c.Path())

		// Validamos los headers obligatorios
		if !apiUtil.ValidarHeaders(c.Request().Header) {
			go logger.PrintLog(nombreAPI, logger.ERROR, c.RealIP(), apiUtil.GetIPServer().String(), c.Request().Method, c.Request().Header, "prueba", c.Request().Host+c.Path(), http.StatusBadRequest, apiUtil.ErrorHeaderNotFound)
			return &echo.HTTPError{
				Code:    http.StatusBadRequest,
				Message: apiUtil.ErrorHeaderNotFound,
			}
		}
		return next(c)
	}
}
