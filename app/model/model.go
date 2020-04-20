package model

import (
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Page struct {
	Start      int      `json:"start,omitempty"`
	End        int      `json:"end,omitempty"`
	Tag        []string `bson:"tag,omitempty" json:"Tag,omitempty"`
	Location   Location `bson:"location,omitempty" json:"location,omitempty"`
	Distance   int      `bson:"distance,omitempty" json:"distance,omitempty"`
	Fovorites  []string `bson:"fovorites,omitempty" json:"fovorites,omitempty"`
	NameCreate string   `bson:"nameCreate,omitempty" json:"nameCreate,omitempty"`
}

type Rating struct {
	User       string `bson:"user,omitempty" json:"user,omitempty"`
	Evaluation int    `bson:"eval,omitempty" json:"eval,omitempty"`
}

type Location struct {
	Type        string    `bson:"type,omitempty" json:"type,omitempty"`
	Coordinates []float64 `bson:"coordinates,omitempty" json:"coordinates,omitempty"`
}

type Platfroms struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title      string             `bson:"title,omitempty" json:"title,omitempty"`
	NameCreate string             `bson:"nameCreate,omitempty" json:"nameCreate,omitempty"`
	Info       string             `bson:"info,omitempty" json:"info,omitempty"`
	Rating     []Rating           `bson:"rating,omitempty" json:"rating,omitempty"`
	Urlimage   string             `bson:"urlimage,omitempty" json:"urlimage,omitempty"`
	Views      int                `bson:"views,omitempty" json:"views,omitempty"`
	Fovorites  []string           `bson:"fovorites,omitempty" json:"fovorites,omitempty"`
	Time       string             `bson:"time,omitempty" json:"time,omitempty"`
	Tag        []string           `bson:"tag,omitempty" json:"Tag,omitempty"`
	Location   Location           `bson:"location,omitempty" json:"location,omitempty"`
}

type Tag_plafrom struct {
	Tag []string `bson:"tag,omitempty" json:"tag,omitempty"`
}

type Default_auth struct {
	R *http.Request
}
