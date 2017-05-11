package models

import (
	"goblog/helpers"
)

var Config = helpers.LoadConfig("config.json")

var dbConnect = helpers.DB{
	Host:   Config.DBHost,
	DBname: Config.DBName,
}

/*
func FindOne(model interface{}, collection *mgo.Collection, query bson.M) error {
	err := collection.Find(query).One(&model)

	return err
}*/
