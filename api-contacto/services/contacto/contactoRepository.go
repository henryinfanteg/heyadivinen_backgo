package contacto

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/henryinfanteg/heyadivinen_backgo/api-contacto/config"
	db "github.com/henryinfanteg/heyadivinen_backgo/db-mongo/server"
	dbUtil "github.com/henryinfanteg/heyadivinen_backgo/db-mongo/util"
)

// CollectionName nombre de la coleccion
const CollectionName = "contacto"

// ContactoRepository objeto entidad
type ContactoRepository struct {
	c       *mgo.Collection
	context *db.Context
	err     error
}

// Parametros
var typeFields = map[string]string{
	"descripcion":      dbUtil.TypeString,
	"estado":           dbUtil.TypeBoolean,
	"gratis":           dbUtil.TypeBoolean,
}

func (repository *ContactoRepository) initContext() {
	conectionDB := db.ConectionDB(config.GetConnectionConfig().Database)
	repository.context, repository.err = db.NewContext(CollectionName, &conectionDB)
	if repository.err == nil {
		repository.c = repository.context.DBCollection(CollectionName)
	}
}


// Create inserta un nuevo registro
func (repository ContactoRepository) Create(obj *Contacto, usuarioCreacion string) error {
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

// CreateMany agrega un array de objetos
func (repository ContactoRepository) CreateMany(objs []Contacto) error {
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
func (repository ContactoRepository) RemoveAll() error {
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
func (repository ContactoRepository) DropCollection() error {
	repository.initContext()
	if repository.err != nil {
		return repository.err
	}
	defer repository.context.Close()

	// Delete Coleccion
	err := repository.c.DropCollection()
	return err
}
