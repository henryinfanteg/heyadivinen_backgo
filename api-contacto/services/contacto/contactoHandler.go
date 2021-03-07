package contacto

import (
	"net/http"
	"strings"

	echo "github.com/labstack/echo/v4"

	dbConstantes "github.com/henryinfanteg/heyadivinen_backgo/db-mongo/constantes"
	apiUtil "github.com/henryinfanteg/heyadivinen_backgo/util-api/util"
	auditoriaUtil "github.com/henryinfanteg/heyadivinen_backgo/util-auditoria/util"
	authUtil "github.com/henryinfanteg/heyadivinen_backgo/util-auth/util"
	logger "github.com/henryinfanteg/heyadivinen_backgo/util-logger/util"
)

// ContactoHandler objeto
type ContactoHandler struct{}

// Variables globales
const (
	ID      = "id"
	APIName = "Contacto"
)

// Create inserta un nuevo registro
func (ContactoHandler) Create(c echo.Context) error {
	nombreMetodo := "Create"
	defer c.Request().Body.Close()

	// Obtenemos el token y username
	token := authUtil.GetTokenByHeader(c.Request().Header.Get(echo.HeaderAuthorization))
	username, errUsername := authUtil.GetUsernameByToken(token)

	// Validamos el username
	if errUsername != nil {
		return c.JSON(http.StatusUnauthorized, authUtil.ErrorInvalidToken)
	}

	// Transformamos el body a la entidad
	var obj Contacto
	err := apiUtil.ConvertBodyToEntity(c.Request().Body, &obj)
	if err != nil {
		// Logueamos
		go logger.PrintResponse(CollectionName, c.RealIP(), apiUtil.GetIPServer().String(), apiUtil.ERROR, c.Request().Header, username, APIName, nombreMetodo, c.Path(), nil, http.StatusBadRequest, err.Error())
		return c.NoContent(http.StatusBadRequest)
	}

	// Logueamos
	go logger.PrintRequest(CollectionName, c.RealIP(), apiUtil.GetIPServer().String(), c.Request().Header, username, APIName, nombreMetodo, c.Path(), obj)

	// Validamos el request
	if err := apiUtil.ValidateStruct(obj); err != nil {
		// Logueamos
		go logger.PrintResponse(CollectionName, c.RealIP(), apiUtil.GetIPServer().String(), apiUtil.ERROR, c.Request().Header, username, APIName, nombreMetodo, c.Path(), nil, http.StatusBadRequest, err.Error())
		return c.JSONBlob(http.StatusBadRequest, []byte(err.Error()))
	}

	// Agregamos a la BD
	var contactoRepository = ContactoRepository{}
	if err := contactoRepository.Create(&obj, username); err != nil {
		if strings.Contains(err.Error(), dbConstantes.CodeDuplicateKey) {
			// Logueamos
			go logger.PrintResponse(CollectionName, c.RealIP(), apiUtil.GetIPServer().String(), apiUtil.ERROR, c.Request().Header, username, APIName, nombreMetodo, c.Path(), nil, http.StatusConflict, err.Error())
			return c.JSON(http.StatusConflict, dbConstantes.ErrorDatabaseDuplicateKey)
		}
		// Logueamos
		go logger.PrintResponse(CollectionName, c.RealIP(), apiUtil.GetIPServer().String(), apiUtil.ERROR, c.Request().Header, username, APIName, nombreMetodo, c.Path(), nil, http.StatusInternalServerError, err.Error())
		return c.NoContent(http.StatusInternalServerError)
	}

	// Auditamos
	go auditoriaUtil.Auditar(CollectionName, username, c.Request().Header, obj.ID, &obj, auditoriaUtil.ADD)

	// Logueamos
	go logger.PrintResponse(CollectionName, c.RealIP(), apiUtil.GetIPServer().String(), apiUtil.INFO, c.Request().Header, username, APIName, nombreMetodo, c.Path(), obj, http.StatusCreated, "")
	return c.JSON(http.StatusCreated, &obj)
}

