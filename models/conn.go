package models

import (
	"gopkg.in/mgo.v2"
)

type DB struct {
	Host       string
	DBname     string
	Connection *mgo.Database
}

func (db *DB) update(host, dbname string) {
	db.Host = host
	db.DBname = dbname
	session, _ := mgo.Dial(db.Host)
	db.Connection = session.DB(db.DBname)
}

func (db *DB) GetConnection() *mgo.Database {
	if db.Connection == nil {
		session, _ := mgo.Dial(db.Host)
		db.Connection = session.DB(db.DBname)
	}

	return db.Connection
}

func (db *DB) GetCollection(name string) *mgo.Collection {
	conn := db.GetConnection()
	return conn.C(name)
}
