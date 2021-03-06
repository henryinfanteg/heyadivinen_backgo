package category

import (
	"net/http"
	"strings"

	echo "github.com/labstack/echo"
	mgo "gopkg.in/mgo.v2"

	dbConstantes "github.com/henryinfanteg/heyadivinen_backgo/db-mongo/constantes"
	apiUtil "github.com/henryinfanteg/heyadivinen_backgo/util-api/util"
	auditoriaUtil "github.com/henryinfanteg/heyadivinen_backgo/util-auditoria/util"
	logger "github.com/henryinfanteg/heyadivinen_backgo/util-logger/util"
)

// CategoryHandler objeto
type CategoryHandler struct{}

// Variables globales
const (
	ID      = "id"
	APIName = "Categories"
)

// GetAll obtiene todos los registros
func (CategoryHandler) GetAll(c echo.Context) error {
	nombreMetodo := "GetAll"

	// Logueamos
	go logger.PrintRequest(CollectionName, c.RealIP(), apiUtil.GetIPServer().String(), c.Request().Header, "prueba", APIName, nombreMetodo, c.Path(), nil)

	// Buscamos en la BD
	var categoryRepository = CategoryRepository{}
	objs, err := categoryRepository.FindAll(c)
	if err != nil {
		// Logueamos
		go logger.PrintResponse(CollectionName, c.RealIP(), apiUtil.GetIPServer().String(), apiUtil.ERROR, c.Request().Header, "prueba", APIName, nombreMetodo, c.Path(), nil, http.StatusInternalServerError, err.Error())
		return c.NoContent(http.StatusInternalServerError)
	}

	// Logueamos
	go logger.PrintResponse(CollectionName, c.RealIP(), apiUtil.GetIPServer().String(), apiUtil.INFO, c.Request().Header, "prueba", APIName, nombreMetodo, c.Path(), nil, http.StatusOK, "")
	return c.JSON(http.StatusOK, objs)
}

// GetByID obtiene un registro por Id
func (CategoryHandler) GetByID(c echo.Context) error {
	nombreMetodo := "GetById"

	// Logueamos
	go logger.PrintRequest(CollectionName, c.RealIP(), apiUtil.GetIPServer().String(), c.Request().Header, "prueba", APIName, nombreMetodo, c.Path(), nil)

	// Obtenemos los parametros
	id := c.Param(ID)
	// Buscamos en la BD
	var categoryRepository = CategoryRepository{}
	obj, err := categoryRepository.FindByID(id)
	if err != nil {
		if err == mgo.ErrNotFound || err.Error() == dbConstantes.ErrorDatabaseInvalidID {
			// Logueamos
			go logger.PrintResponse(CollectionName, c.RealIP(), apiUtil.GetIPServer().String(), apiUtil.INFO, c.Request().Header, "prueba", APIName, nombreMetodo, c.Path(), nil, http.StatusNoContent, err.Error())
			return c.JSON(http.StatusNoContent, dbConstantes.ErrorDatabaseRecordNotFound)
		}
		// Logueamos
		go logger.PrintResponse(CollectionName, c.RealIP(), apiUtil.GetIPServer().String(), apiUtil.ERROR, c.Request().Header, "prueba", APIName, nombreMetodo, c.Path(), nil, http.StatusInternalServerError, err.Error())
		return c.NoContent(http.StatusInternalServerError)
	}

	// Logueamos
	go logger.PrintResponse(CollectionName, c.RealIP(), apiUtil.GetIPServer().String(), apiUtil.INFO, c.Request().Header, "prueba", APIName, nombreMetodo, c.Path(), nil, http.StatusOK, "")
	return c.JSON(http.StatusOK, &obj)
}

// Count obtiene la cantidad de registros
func (CategoryHandler) Count(c echo.Context) error {
	nombreMetodo := "Count"

	// Logueamos
	go logger.PrintRequest(CollectionName, c.RealIP(), apiUtil.GetIPServer().String(), c.Request().Header, "prueba", APIName, nombreMetodo, c.Path(), nil)

	// Buscamos en la BD
	var categoryRepository = CategoryRepository{}
	count, err := categoryRepository.Count(c)
	if err != nil {
		// Logueamos
		go logger.PrintResponse(CollectionName, c.RealIP(), apiUtil.GetIPServer().String(), apiUtil.ERROR, c.Request().Header, "prueba", APIName, nombreMetodo, c.Path(), nil, http.StatusInternalServerError, err.Error())
		return c.NoContent(http.StatusInternalServerError)
	}

	// Logueamos
	go logger.PrintResponse(CollectionName, c.RealIP(), apiUtil.GetIPServer().String(), apiUtil.INFO, c.Request().Header, "prueba", APIName, nombreMetodo, c.Path(), nil, http.StatusOK, "")
	return c.JSON(http.StatusOK, count)
}

// Create inserta un nuevo registro
func (CategoryHandler) Create(c echo.Context) error {
	nombreMetodo := "Create"
	defer c.Request().Body.Close()

	// Transformamos el body a la entidad
	var obj Category
	err := apiUtil.ConvertBodyToEntity(c.Request().Body, &obj)
	if err != nil {
		// Logueamos
		go logger.PrintResponse(CollectionName, c.RealIP(), apiUtil.GetIPServer().String(), apiUtil.ERROR, c.Request().Header, "prueba", APIName, nombreMetodo, c.Path(), nil, http.StatusBadRequest, err.Error())
		return c.NoContent(http.StatusBadRequest)
	}

	// Logueamos
	go logger.PrintRequest(CollectionName, c.RealIP(), apiUtil.GetIPServer().String(), c.Request().Header, "prueba", APIName, nombreMetodo, c.Path(), obj)

	// Validamos el request
	if err := apiUtil.ValidateStruct(obj); err != nil {
		// Logueamos
		go logger.PrintResponse(CollectionName, c.RealIP(), apiUtil.GetIPServer().String(), apiUtil.ERROR, c.Request().Header, "prueba", APIName, nombreMetodo, c.Path(), nil, http.StatusBadRequest, err.Error())
		return c.JSONBlob(http.StatusBadRequest, []byte(err.Error()))
	}

	// Agregamos a la BD
	var categoryRepository = CategoryRepository{}
	if err := categoryRepository.Create(&obj, "prueba"); err != nil {
		if strings.Contains(err.Error(), dbConstantes.CodeDuplicateKey) {
			// Logueamos
			go logger.PrintResponse(CollectionName, c.RealIP(), apiUtil.GetIPServer().String(), apiUtil.ERROR, c.Request().Header, "prueba", APIName, nombreMetodo, c.Path(), nil, http.StatusConflict, err.Error())
			return c.JSON(http.StatusConflict, dbConstantes.ErrorDatabaseDuplicateKey)
		}
		// Logueamos
		go logger.PrintResponse(CollectionName, c.RealIP(), apiUtil.GetIPServer().String(), apiUtil.ERROR, c.Request().Header, "prueba", APIName, nombreMetodo, c.Path(), nil, http.StatusInternalServerError, err.Error())
		return c.NoContent(http.StatusInternalServerError)
	}

	// Auditamos
	go auditoriaUtil.Auditar(CollectionName, "prueba", c.Request().Header, obj.ID, &obj, auditoriaUtil.ADD)

	// Logueamos
	go logger.PrintResponse(CollectionName, c.RealIP(), apiUtil.GetIPServer().String(), apiUtil.INFO, c.Request().Header, "prueba", APIName, nombreMetodo, c.Path(), obj, http.StatusCreated, "")
	return c.JSON(http.StatusCreated, &obj)
}

// Update actualiza un registro
func (CategoryHandler) Update(c echo.Context) error {
	nombreMetodo := "Update"
	defer c.Request().Body.Close()

	// Obtenemos los parametros
	id := c.Param(ID)

	// Transformamos en requestBody en la entidad
	var obj Category
	err := apiUtil.ConvertBodyToEntity(c.Request().Body, &obj)
	if err != nil {
		// Logueamos
		go logger.PrintResponse(CollectionName, c.RealIP(), apiUtil.GetIPServer().String(), apiUtil.ERROR, c.Request().Header, "prueba", APIName, nombreMetodo, c.Path(), nil, http.StatusBadRequest, err.Error())
		return c.NoContent(http.StatusBadRequest)
	}

	// Logueamos
	go logger.PrintRequest(CollectionName, c.RealIP(), apiUtil.GetIPServer().String(), c.Request().Header, "prueba", APIName, nombreMetodo, c.Path(), obj)

	// Validamos el request
	if err := apiUtil.ValidateStruct(obj); err != nil {
		// Logueamos
		go logger.PrintResponse(CollectionName, c.RealIP(), apiUtil.GetIPServer().String(), apiUtil.ERROR, c.Request().Header, "prueba", APIName, nombreMetodo, c.Path(), nil, http.StatusBadRequest, err.Error())
		return c.JSONBlob(http.StatusBadRequest, []byte(err.Error()))
	}

	// Actualizamos la BD
	var categoryRepository = CategoryRepository{}
	if err := categoryRepository.Update(id, &obj, "prueba"); err != nil {
		if err == mgo.ErrNotFound || err.Error() == dbConstantes.ErrorDatabaseInvalidID {
			// Logueamos
			go logger.PrintResponse(CollectionName, c.RealIP(), apiUtil.GetIPServer().String(), apiUtil.INFO, c.Request().Header, "prueba", APIName, nombreMetodo, c.Path(), nil, http.StatusNoContent, err.Error())
			return c.JSON(http.StatusNoContent, dbConstantes.ErrorDatabaseRecordNotFound)
		}
		// Logueamos
		go logger.PrintResponse(CollectionName, c.RealIP(), apiUtil.GetIPServer().String(), apiUtil.ERROR, c.Request().Header, "prueba", APIName, nombreMetodo, c.Path(), nil, http.StatusInternalServerError, err.Error())
		return c.NoContent(http.StatusInternalServerError)
	}

	// Auditamos
	go auditoriaUtil.Auditar(CollectionName, "prueba", c.Request().Header, id, &obj, auditoriaUtil.UPDATE)

	// Logueamos
	go logger.PrintResponse(CollectionName, c.RealIP(), apiUtil.GetIPServer().String(), apiUtil.INFO, c.Request().Header, "prueba", APIName, nombreMetodo, c.Path(), obj, http.StatusOK, "")
	return c.JSON(http.StatusOK, &obj)
}

// Delete elimina un registro
func (CategoryHandler) Delete(c echo.Context) error {
	nombreMetodo := "Delete"

	// Logueamos
	go logger.PrintRequest(CollectionName, c.RealIP(), apiUtil.GetIPServer().String(), c.Request().Header, "prueba", APIName, nombreMetodo, c.Path(), nil)

	// Obtenemos los parametros
	id := c.Param(ID)

	// Eliminamos de la BD
	var categoryRepository = CategoryRepository{}
	if err := categoryRepository.Delete(id); err != nil {
		if err == mgo.ErrNotFound {
			// Logueamos
			go logger.PrintResponse(CollectionName, c.RealIP(), apiUtil.GetIPServer().String(), apiUtil.ERROR, c.Request().Header, "prueba", APIName, nombreMetodo, c.Path(), nil, http.StatusNoContent, err.Error())
			return c.JSON(http.StatusNoContent, dbConstantes.ErrorDatabaseRecordNotFound)
		}
		// Logueamos
		go logger.PrintResponse(CollectionName, c.RealIP(), apiUtil.GetIPServer().String(), apiUtil.ERROR, c.Request().Header, "prueba", APIName, nombreMetodo, c.Path(), nil, http.StatusInternalServerError, err.Error())
		return c.NoContent(http.StatusInternalServerError)
	}

	// Auditamos
	go auditoriaUtil.Auditar(CollectionName, "prueba", c.Request().Header, id, nil, auditoriaUtil.DELETE)

	// Logueamos
	go logger.PrintResponse(CollectionName, c.RealIP(), apiUtil.GetIPServer().String(), apiUtil.INFO, c.Request().Header, "prueba", APIName, nombreMetodo, c.Path(), nil, http.StatusOK, "")
	return c.NoContent(http.StatusOK)
}
