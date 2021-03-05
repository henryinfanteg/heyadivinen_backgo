package server

import (
	"time"
	"gopkg.in/mgo.v2"
)

var session *mgo.Session

func createDBSession(conection *ConectionDB) error {
	var err error
	session, err = mgo.DialWithInfo(&mgo.DialInfo{
		Username: conection.Username,
		Password: conection.Password,
		Database: conection.Database,
		Addrs: conection.Host,
		Timeout: 3 * time.Second,
	})
	return err
}

func getSession(conection *ConectionDB) (*mgo.Session, error) {
	var err error
	if session == nil {
		err = createDBSession(conection)
		// return session, err
	}
	return session, err
}

func InitDatabase(conection *ConectionDB) {
	createDBSession(conection)
}
