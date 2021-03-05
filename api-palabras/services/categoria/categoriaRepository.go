package categoria

import (
	"errors"
	"time"

	echo "github.com/labstack/echo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"gitlab.com/adivinagame/backend/maxadivinabackend/api-palabras/config"
	dbConstantes "gitlab.com/adivinagame/backend/maxadivinabackend/db-mongo/constantes"
	db "gitlab.com/adivinagame/backend/maxadivinabackend/db-mongo/server"
	dbUtil "gitlab.com/adivinagame/backend/maxadivinabackend/db-mongo/util"
	apiUtil "gitlab.com/adivinagame/backend/maxadivinabackend/util-api/util"
)

// CollectionName nombre de la coleccion
const CollectionName = "categorias"

// CategoriaRepository objeto entidad
type CategoriaRepository struct {
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
	"descripcion":      dbUtil.TypeString,
	"estado":           dbUtil.TypeBoolean,
	"gratis":           dbUtil.TypeBoolean,
}

func (repository *CategoriaRepository) initContext() {
	conectionDB := db.ConectionDB(config.GetConnectionConfig().Database)
	repository.context, repository.err = db.NewContextWithIndex(CollectionName, &conectionDB, indexUnique)
	if repository.err == nil {
		repository.c = repository.context.DBCollection(CollectionName)
	}
}

// FindAll devuelve todos los registros
func (repository CategoriaRepository) FindAll(c echo.Context) (*[]Categoria, error) {
	repository.initContext()
	if repository.err != nil {
		return nil, repository.err
	}
	defer repository.context.Close()

	var objs []Categoria
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
func (repository CategoriaRepository) FindByID(id string) (*Categoria, error) {
	repository.initContext()
	if repository.err != nil {
		return nil, repository.err
	}
	defer repository.context.Close()

	if !bson.IsObjectIdHex(id) {
		return nil, errors.New(dbConstantes.ErrorDatabaseInvalidID)
	}

	var obj Categoria
	err := repository.c.FindId(bson.ObjectIdHex(id)).One(&obj)
	return &obj, err
}

// Count devuelve la cantidad de registros
func (repository CategoriaRepository) Count(c echo.Context) (int, error) {
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
func (repository CategoriaRepository) Create(obj *Categoria, usuarioCreacion string) error {
	repository.initContext()
	if repository.err != nil {
		return repository.err
	}
	defer repository.context.Close()

	obj.ID = bson.NewObjectId()
	obj.FechaCreacion = time.Now()
	obj.UsuarioCreacion = usuarioCreacion
	obj.FechaModificacion = time.Time{}
	err := repository.c.Insert(obj)
	return err
}

// Update actualiza un registro
func (repository CategoriaRepository) Update(id string, obj *Categoria, usuarioModificacion string) error {
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
			"descripcion":         obj.Descripcion,
			"fechaModificacion":   obj.FechaModificacion,
			"usuarioModificacion": obj.UsuarioModificacion,
		}})
	return err
}

// Delete elimina un registro
func (repository CategoriaRepository) Delete(id string) error {
	repository.initContext()
	defer repository.context.Close()

	if !bson.IsObjectIdHex(id) {
		return errors.New(dbConstantes.ErrorDatabaseInvalidID)
	}

	return repository.c.RemoveId(bson.ObjectIdHex(id))
}

// CreateMany agrega un array de objetos
func (repository CategoriaRepository) CreateMany(objs []Categoria) error {
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
func (repository CategoriaRepository) RemoveAll() error {
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
func (repository CategoriaRepository) DropCollection() error {
	repository.initContext()
	if repository.err != nil {
		return repository.err
	}
	defer repository.context.Close()

	// Delete Coleccion
	err := repository.c.DropCollection()
	return err
}
