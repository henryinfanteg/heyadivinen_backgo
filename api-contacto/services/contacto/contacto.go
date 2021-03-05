package contacto

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Contacto entidad
type Contacto struct {
	ID                  bson.ObjectId    `bson:"_id,omitempty" json:"_id"`
	Correo              string           `bson:"correo" json:"correo" validate:"required"`
	Asunto              string             `bson:"asunto" json:"asunto" validate:"required"`
	Descripcion         string           `bson:"descripcion" json:"descripcion" validate:"required"`
	Estado         		string           `bson:"estado" json:"estado" validate:"required"`
	FechaCreacion       time.Time        `bson:"fechaCreacion" json:"fechaCreacion,omitempty"`
	UsuarioCreacion     string           `bson:"usuarioCreacion" json:"usuarioCreacion,omitempty"`
	FechaModificacion   time.Time        `bson:"fechaModificacion" json:"fechaModificacion,omitempty"`
	UsuarioModificacion string           `bson:"usuarioModificacion" json:"usuarioModificacion,omitempty"`
}

