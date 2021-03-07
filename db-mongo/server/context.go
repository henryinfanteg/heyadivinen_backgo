package server

import (
	"errors"
	"gopkg.in/mgo.v2"

	"github.com/henryinfanteg/heyadivinen_backgo/db-mongo/constantes"
)

// Context contexto que maneja la session de la BD
type Context struct {
	Session *mgo.Session
}

// Close cierra una session
func (c *Context) Close() {
	c.Session.Close()
}

// Variables
var (
	databaseName string
)

// DBCollection referencia la collection a trabajar
func (c *Context) DBCollection(name string) *mgo.Collection {
	return c.Session.DB(databaseName).C(name)
}

// NewContext crea una nueva instancia del contexto
func NewContext(collectionName string, conection *ConectionDB) (*Context, error) {
	databaseName = conection.Database
	session, errSession := getSession(conection)
	if errSession != nil {
		return nil, errors.New(constantes.ErrorDatabaseConexion)
	}

	session = session.Copy()
	c := &Context{
		Session: session,
	}
	return c, nil
}

// NewContextWithIndex crea una nueva instancia del contexto con un indice
func NewContextWithIndex(collectionName string, conection *ConectionDB, indexs ...mgo.Index) (*Context, error) {
	databaseName = conection.Database
	session, errSession := getSession(conection)
	if errSession != nil {
		return nil, errors.New(constantes.ErrorDatabaseConexion)
	}

	session = session.Copy()
	c := &Context{
		Session: session,
	}

	for _, index := range indexs {
		if err := addIndex(collectionName, index); err != nil {
			return c, err
		}
	}
	return c, nil
}

// addIndex agrega un indice a una collection
func addIndex(collectionName string, index mgo.Index) error {
	c := session.DB(databaseName).C(collectionName)
	return c.EnsureIndex(index)
}
