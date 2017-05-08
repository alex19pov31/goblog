package models

import (
	"../helpers"
)

var dbConnect = helpers.DB{
	Host:   "localhost",
	DBname: "testdb",
}
/*
func FindOne(model interface{}, collection *mgo.Collection, query bson.M) error {
	err := collection.Find(query).One(&model)

	return err
}*/