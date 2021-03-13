package category

import (
	"errors"
	"time"

	echo "github.com/labstack/echo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/henryinfanteg/heyadivinen_backgo/api-words/config"
	dbConstantes "github.com/henryinfanteg/heyadivinen_backgo/db-mongo/constantes"
	db "github.com/henryinfanteg/heyadivinen_backgo/db-mongo/server"
	dbUtil "github.com/henryinfanteg/heyadivinen_backgo/db-mongo/util"
	apiUtil "github.com/henryinfanteg/heyadivinen_backgo/util-api/util"
)

// CollectionName nombre de la coleccion
const CollectionName = "categories"

// CategoryRepository objeto entidad
type CategoryRepository struct {
	c       *mgo.Collection
	context *db.Context
	err     error
}

var indexUnique = mgo.Index{
	Key:    []string{"id"},
	Unique: true,
}

// Parametros
var typeFields = map[string]string{
	"description": dbUtil.TypeString,
	"status":      dbUtil.TypeBoolean,
	"free":        dbUtil.TypeBoolean,
}

func (repository *CategoryRepository) initContext() {
	conectionDB := db.ConectionDB(config.GetConnectionConfig().Database)
	repository.context, repository.err = db.NewContextWithIndex(CollectionName, &conectionDB, indexUnique)
	if repository.err == nil {
		repository.c = repository.context.DBCollection(CollectionName)
	}
}

// FindAll devuelve todos los registros
func (repository CategoryRepository) FindAll(c echo.Context) (*[]Category, error) {
	repository.initContext()
	if repository.err != nil {
		return nil, repository.err
	}
	defer repository.context.Close()

	var objs []Category
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
func (repository CategoryRepository) FindByID(id string) (*Category, error) {
	repository.initContext()
	if repository.err != nil {
		return nil, repository.err
	}
	defer repository.context.Close()

	if !bson.IsObjectIdHex(id) {
		return nil, errors.New(dbConstantes.ErrorDatabaseInvalidID)
	}

	var obj Category
	err := repository.c.FindId(bson.ObjectIdHex(id)).One(&obj)
	return &obj, err
}

// Count devuelve la cantidad de registros
func (repository CategoryRepository) Count(c echo.Context) (int, error) {
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
func (repository CategoryRepository) Create(obj *Category, creationUser string) error {
	repository.initContext()
	if repository.err != nil {
		return repository.err
	}
	defer repository.context.Close()

	obj.ID = bson.NewObjectId()
	obj.CreationDate = time.Now()
	obj.CreationUser = creationUser
	obj.DateModify = time.Time{}
	err := repository.c.Insert(obj)
	return err
}

// Update actualiza un registro
func (repository CategoryRepository) Update(id string, obj *Category, userModify string) error {
	repository.initContext()
	defer repository.context.Close()

	if !bson.IsObjectIdHex(id) {
		return errors.New(dbConstantes.ErrorDatabaseInvalidID)
	}

	obj.DateModify = time.Now()
	obj.UserModify = userModify
	obj.ID = bson.ObjectIdHex(id)

	// err := repository.c.UpdateId(obj.ID, &obj)
	err := repository.c.UpdateId(obj.ID,
		bson.M{"$set": bson.M{
			"description": obj.Description,
			"dateModify":  obj.DateModify,
			"userModify":  obj.UserModify,
		}})
	return err
}

// Delete elimina un registro
func (repository CategoryRepository) Delete(id string) error {
	repository.initContext()
	defer repository.context.Close()

	if !bson.IsObjectIdHex(id) {
		return errors.New(dbConstantes.ErrorDatabaseInvalidID)
	}

	return repository.c.RemoveId(bson.ObjectIdHex(id))
}

// CreateMany agrega un array de objetos
func (repository CategoryRepository) CreateMany(objs []Category) error {
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

// RemoveAll elimina los registros de una coleccion
func (repository CategoryRepository) RemoveAll() error {
	repository.initContext()
	if repository.err != nil {
		return repository.err
	}
	defer repository.context.Close()

	// Clear DB
	_, err := repository.c.RemoveAll(nil)
	return err
}

// DropCollection elimina la coleccion completa
func (repository CategoryRepository) DropCollection() error {
	repository.initContext()
	if repository.err != nil {
		return repository.err
	}
	defer repository.context.Close()

	// Delete Coleccion
	err := repository.c.DropCollection()
	return err
}
