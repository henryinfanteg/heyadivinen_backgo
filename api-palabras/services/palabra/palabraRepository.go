package palabra

import (
	"errors"
	"time"

	echo "github.com/labstack/echo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/henryinfanteg/heyadivinen_backgo/tree/master/tree/master/api-palabras/config"
	dbConstantes "github.com/henryinfanteg/heyadivinen_backgo/tree/master/db-mongo/constantes"
	db "github.com/henryinfanteg/heyadivinen_backgo/tree/master/db-mongo/server"
	dbUtil "github.com/henryinfanteg/heyadivinen_backgo/tree/master/db-mongo/util"
	apiUtil "github.com/henryinfanteg/heyadivinen_backgo/tree/master/util-api/util"
)

// CollectionName nombre de la coleccion
const CollectionName = "palabras"

// PalabraRepository objeto entidad
type PalabraRepository struct {
	c       *mgo.Collection
	context *db.Context
	err     error
}

// Parametros
var typeFields = map[string]string{
	"palabra":             dbUtil.TypeString,
	"categoriaId":          dbUtil.TypeString,
}

func (repository *PalabraRepository) initContext() {
	conectionDB := db.ConectionDB(config.GetConnectionConfig().Database)
	repository.context, repository.err = db.NewContext(CollectionName, &conectionDB)
	if repository.err == nil {
		repository.c = repository.context.DBCollection(CollectionName)
	}
}

// FindAll devuelve todos los registros
func (repository PalabraRepository) FindAll(c echo.Context) (*[]Palabra, error) {
	repository.initContext()
	if repository.err != nil {
		return nil, repository.err
	}
	defer repository.context.Close()

	var objs []Palabra
	var err error

	// Obtenemos el parametro
	params := dbUtil.GetParametrosFiltro(c.QueryParams(), typeFields)
	sort := apiUtil.GetParametroSort(c)
	pageNumber, pageSize := apiUtil.GetParametrosPaginacion(c)

	query := dbUtil.CreateQuery(repository.c, params, sort, pageNumber, pageSize)
	err = query.All(&objs)

	return &objs, err
}

// FindByID devuelve un registro filtrado por Id
func (repository PalabraRepository) FindByID(id string) (*Palabra, error) {
	repository.initContext()
	if repository.err != nil {
		return nil, repository.err
	}
	defer repository.context.Close()

	if !bson.IsObjectIdHex(id) {
		return nil, errors.New(dbConstantes.ErrorDatabaseInvalidID)
	}

	var obj Palabra
	err := repository.c.FindId(bson.ObjectIdHex(id)).One(&obj)
	return &obj, err
}

// FindByUsername devuelve una lista de registros filtrado por categoriaId
func (repository PalabraRepository) FindAllByCategoriaID(c echo.Context, categoriaId string) (*[]Palabra, error) {
	repository.initContext()
	if repository.err != nil {
		return nil, repository.err
	}
	defer repository.context.Close()

	var objs []Palabra
	var err error

	// Obtenemos el parametro
	params := dbUtil.GetParametrosFiltro(c.QueryParams(), typeFields)
	params["categoriaId"] = categoriaId

	query := dbUtil.CreateQueryRandom(repository.c, params)
	err = query.All(&objs)
	return &objs, err
}

// Count devuelve la cantidad de registros
func (repository PalabraRepository) Count(c echo.Context) (int, error) {
	repository.initContext()
	if repository.err != nil {
		return 0, repository.err
	}
	defer repository.context.Close()

	// Obtenemos el parametro
	params := dbUtil.GetParametrosFiltro(c.QueryParams(), typeFields)

	count, err := repository.c.Find(params).Count()
	return count, err
}

// Create inserta un nuevo registro
func (repository PalabraRepository) Create(obj *Palabra, usuarioCreacion string) error {
	repository.initContext()
	if repository.err != nil {
		return repository.err
	}
	defer repository.context.Close()

	obj.ID = bson.NewObjectId()
	obj.FechaCreacion = time.Now()
	obj.UsuarioCreacion = usuarioCreacion
	err := repository.c.Insert(obj)
	return err
}

// Update actualiza un registro
func (repository PalabraRepository) Update(id string, obj *Palabra, usuarioModificacion string) error {
	repository.initContext()
	defer repository.context.Close()

	if !bson.IsObjectIdHex(id) {
		return errors.New(dbConstantes.ErrorDatabaseInvalidID)
	}

	obj.FechaModificacion = time.Now()
	obj.UsuarioModificacion = usuarioModificacion
	obj.ID = bson.ObjectIdHex(id)

	// err := repository.c.UpdateId(obj.ID, &obj)
	err := repository.c.UpdateId(obj.ID,
		bson.M{"$set": bson.M{
			"palabra":              obj.Palabra,
			"categoriaId":          obj.CategoriaId,
			"pista": 				obj.Pista,
			"fechaModificacion":    obj.FechaModificacion,
			"usuarioModificacion":  obj.UsuarioModificacion,
		}})
	return err
}

// Delete elimina un registro
func (repository PalabraRepository) Delete(id string) error {
	repository.initContext()
	defer repository.context.Close()

	if !bson.IsObjectIdHex(id) {
		return errors.New(dbConstantes.ErrorDatabaseInvalidID)
	}

	return repository.c.RemoveId(bson.ObjectIdHex(id))
}

// CreateMany agrega un array de objetos
func (repository PalabraRepository) CreateMany(objs []Palabra) error {
	repository.initContext()
	if repository.err != nil {
		return repository.err
	}
	defer repository.context.Close()

	// Clear DB
	// repository.c.RemoveAll(bson.M{})
	repository.c.DropCollection()

	for _, obj := range objs {
		err := repository.c.Insert(obj)
		if err != nil {
			return err
		}
	}
	return nil
}