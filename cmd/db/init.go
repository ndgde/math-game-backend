package db

import (
	fdb "github.com/ndgde/flimsy-db/cmd/flimsydb"
	cm "github.com/ndgde/flimsy-db/cmd/flimsydb/common"
	"github.com/ndgde/flimsy-db/cmd/flimsydb/indexer"
)

// table example
// {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99}

type DB struct {
	inst *fdb.FlimsyDB
}

type Album struct {
	ID     int32   `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var db DB = DB{inst: nil}

func (db *DB) init() {
	dbInst := NewDB()

	col1, err := fdb.NewColumn("id", cm.Int32TType, int32(0), indexer.HashMapIndexerType, fdb.PrimaryKeyFlag&fdb.ImmutableFlag)
	if err != nil {
		panic(err)
	}

	col2, err := fdb.NewColumn("title", cm.StringTType, "Concert", indexer.BTreeIndexerType, 0)
	if err != nil {
		panic(err)
	}

	col3, err := fdb.NewColumn("artist", cm.StringTType, "Unknown artist", indexer.AbsentIndexerType, 0)
	if err != nil {
		panic(err)
	}

	col4, err := fdb.NewColumn("price", cm.Float64TType, float64(0), indexer.HashMapIndexerType, 0)
	if err != nil {
		panic(err)
	}

	scheme := fdb.Scheme{col1, col2, col3, col4}

	dbInst.CreateTable("albums", scheme)
}

func NewDB() *fdb.FlimsyDB { /* singleton */
	if db.inst != nil {
		return db.inst
	}

	db.inst = fdb.NewFlimsyDB()

	db.init()

	return db.inst
}
