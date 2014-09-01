package models

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type (
	Geo struct {
		Type        string     `json:"type" bson:"type"`
		Coordinates [2]float64 `json:"coordinates" bson:"coordinates"`
	}

	Poi struct {
		Id    bson.ObjectId `json:"id" bson:"_id"`
		Layer string        `json:"layer", bson:"layar"`
		Loc   Geo           `json:"loc", bson:"loc"`
	}
)

func Collection(session *mgo.Session, name string) *mgo.Collection {
	return session.DB(DbSingaporeFacts).C(name)
}

func GetPoisByLayer(session *mgo.Session, layers []string) (pois []Poi, err error) {
	err = Collection(session, "pois").Find(bson.M{"layer": bson.M{"$in": layers}}).All(&pois)
	return
}
