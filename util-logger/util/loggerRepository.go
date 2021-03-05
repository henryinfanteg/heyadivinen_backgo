package util

import (
	"gopkg.in/mgo.v2"

	db "gitlab.com/adivinagame/backend/maxadivinabackend/db-mongo/server"
	config "gitlab.com/adivinagame/backend/maxadivinabackend/util-logger/config"
)

// LoggerRepository objeto
type LoggerRepository struct {
	c       *mgo.Collection
	context *db.Context
	err     error
}

func (repository *LoggerRepository) initContext(collectionName string) {
	repository.context, repository.err = db.NewContext(collectionName, config.GetConfigDB())
	if repository.err == nil {
		repository.c = repository.context.DBCollection(collectionName)
	}
}

// Create agrega un registro en la BD
func (repository LoggerRepository) Create(collectionName string, data interface{}) error {
	repository.initContext(collectionName + "Logs")
	if repository.err != nil {
		return repository.err
	}
	defer repository.context.Close()

	err := repository.c.Insert(data)
	return err
}
