package palabra

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Palabra entidad
type Palabra struct {
	ID                   bson.ObjectId `bson:"_id,omitempty" json:"_id"`
	Palabra              string        `bson:"palabra" json:"palabra" validate:"required"`
	CategoriaId          string        `bson:"categoriaId" json:"categoriaId"`
	Pista				 string        `bson:"pista" json:"pista"`
	FechaCreacion        time.Time     `bson:"fechaCreacion" json:"fechaCreacion,omitempty"`
	UsuarioCreacion      string        `bson:"usuarioCreacion" json:"usuarioCreacion,omitempty"`
	FechaModificacion    time.Time     `bson:"fechaModificacion" json:"fechaModificacion,omitempty"`
	UsuarioModificacion  string        `bson:"usuarioModificacion" json:"usuarioModificacion,omitempty"`
}
