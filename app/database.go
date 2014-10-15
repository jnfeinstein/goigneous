package app

import (
	"database/sql"
	"github.com/coopernurse/gorp"
	_ "github.com/lib/pq"
	. "goigneous/app/models"
	"goigneous/config"
)

type Database struct {
	Connection *gorp.DbMap
}

func MakeDb() (*Database, error) {
	db, err := sql.Open("postgres", config.PostgresArgs())
	if err != nil {
		return nil, err
	}
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	dbmap.AddTableWithName(Document{}, "documents").SetKeys(true, "id")
	return &Database{dbmap}, nil
}

func (db *Database) Add(doc *Document) error {
	return db.Connection.Insert(doc)
}

func (db *Database) Get(id int) (*Document, error) {
	obj, err := db.Connection.Get(Document{}, id)
	if obj == nil || err != nil {
		return nil, err // Can't convert interface to document if there was an error
	}

	return obj.(*Document), err
}

func (db *Database) Update(doc *Document) (int, error) {
	count, err := db.Connection.Update(doc)
	return int(count), err
}

func (db *Database) Remove(id int) error {
	_, err := db.Connection.Delete(&Document{id, ""})
	return err
}
