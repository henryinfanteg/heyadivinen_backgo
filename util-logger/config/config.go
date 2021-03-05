package config

import (
	db "gitlab.com/adivinagame/backend/maxadivinabackend/db-mongo/server"
)

var conectionDB db.ConectionDB

// LoadConfigDB metodo para cargar la coneccion a la BD
func LoadConfigDB(conection *db.ConectionDB) {
	conectionDB = *conection
}

// GetConfigDB obtiene la configuracion de la BD
func GetConfigDB() *db.ConectionDB {
	return &conectionDB
}
