package category

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Category struct {
	ID                  bson.ObjectId    `bson:"_id,omitempty" json:"_id"`
	Id                  string           `bson:"id" json:"id" validate:"required"`
	Description         string           `bson:"description" json:"description" validate:"required"`
	Status              bool             `bson:"status" json:"status"`
	Words			    []Word 	         `bson:"words" json:"words"`
	Free              	bool             `bson:"free" json:"free"`
	Price               float64          `bson:"price" json:"price"`
	Icon                string           `bson:"icon" json:"icon"`
	CreationDate        time.Time        `bson:"creationDate" json:"creationDate,omitempty"`
	CreationUser        string           `bson:"creationUser" json:"creationUser,omitempty"`
	DateModify          time.Time        `bson:"dateModify" json:"dateModify,omitempty"`
	UserModify          string           `bson:"userModify" json:"userModify,omitempty"`
}

type Word struct {
	ID                   bson.ObjectId   `bson:"_id,omitempty" json:"_id"`
	Description          string          `bson:"description" json:"description" validate:"required"`
	Clue				 string          `bson:"clue" json:"clue"`
	CreationDate         time.Time       `bson:"creationDate" json:"creationDate,omitempty"`
	CreationUser         string          `bson:"creationUser" json:"creationUser,omitempty"`
	DateModify           time.Time       `bson:"dateModify" json:"dateModify,omitempty"`
	UserModify           string          `bson:"userModify" json:"userModify,omitempty"`
}

