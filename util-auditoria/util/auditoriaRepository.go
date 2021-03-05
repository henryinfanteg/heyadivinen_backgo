package util

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	db "gitlab.com/adivinagame/backend/maxadivinabackend/db-mongo/server"
	config "gitlab.com/adivinagame/backend/maxadivinabackend/util-auditoria/config"
)

// AuditoriaRepository objeto
type AuditoriaRepository struct {
	c       *mgo.Collection
	context *db.Context
	err     error
}

// Auditoria objeto
type Auditoria struct {
	ID         bson.ObjectId `bson:"_id" json:"_id"`
	Fecha      time.Time     `bson:"fecha" json:"fecha"`
	Usuario    string        `bson:"usuario" json:"usuario"`
	Accion     string        `bson:"accion" json:"accion"`
	RequestID  string        `bson:"requestId" json:"requestId"`
	AppID      string        `bson:"appId" json:"appId"`
	IDRegistro interface{}   `bson:"idRegistro" json:"idRegistro"`
	Body       interface{}   `bson:"body" json:"body"`
	Comentario string        `bson:"comentario" json:"comentario"`
}

func (repository *AuditoriaRepository) initContext(collectionName string) {
	repository.context, repository.err = db.NewContext(collectionName, config.GetConfigDB())
	if repository.err == nil {
		repository.c = repository.context.DBCollection(collectionName)
	}
}

// Create agrega un registro en la BD
func (repository AuditoriaRepository) Create(collectionName string, usuario string, requestID string, appID string, idRegistro interface{}, body interface{}, accion string, comentario string) error {
	repository.initContext(collectionName + "Aud")
	if repository.err != nil {
		return repository.err
	}
	defer repository.context.Close()

	var auditoria Auditoria

	auditoria.ID = bson.NewObjectId()
	auditoria.Fecha = time.Now()
	auditoria.Usuario = usuario
	auditoria.Accion = accion
	auditoria.RequestID = requestID
	auditoria.AppID = appID
	auditoria.IDRegistro = idRegistro
	auditoria.Body = body
	auditoria.Comentario = comentario
	err := repository.c.Insert(auditoria)
	return err
}
