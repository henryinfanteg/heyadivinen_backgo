package categoria

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Categoria entidad
type Categoria struct {
	ID                  bson.ObjectId    `bson:"_id,omitempty" json:"_id"`
	Id                  string           `bson:"id" json:"id" validate:"required"`
	Descripcion         string           `bson:"descripcion" json:"descripcion" validate:"required"`
	Estado              bool             `bson:"estado" json:"estado"`
	Gratis              bool             `bson:"gratis" json:"gratis"`
	Precio              float64          `bson:"precio" json:"precio"`
	Icon                string           `bson:"icon" json:"icon"`
	FechaCreacion       time.Time        `bson:"fechaCreacion" json:"fechaCreacion,omitempty"`
	UsuarioCreacion     string           `bson:"usuarioCreacion" json:"usuarioCreacion,omitempty"`
	FechaModificacion   time.Time        `bson:"fechaModificacion" json:"fechaModificacion,omitempty"`
	UsuarioModificacion string           `bson:"usuarioModificacion" json:"usuarioModificacion,omitempty"`
}

