package config

import (
	"log"
	"os"
	"time"

	mgo "gopkg.in/mgo.v2"
)

// Mongo has connection params ready to use
var Mongo *mongoStore

type mongoStore struct {
	session *mgo.Session
	db      string
}

type mongoConfig struct {
	DB      string `required:"true"`
	URL     string `required:"true"`
	Logging bool   `default:"false"`
}

func newMongoConfig() *mongoConfig {
	var c mongoConfig
	err := ReadConfig("MONGO", &c)
	if err != nil {
		panic(err)
	}
	return &c
}

func (m mongoStore) NewSession() (*mgo.Session, *mgo.Database) {
	session := m.session.Copy()
	return session, session.DB(m.db)
}

func initMongoDB() {
	c := newMongoConfig()
	session, err := mgo.Dial(c.URL)
	if err != nil {
		panic(err)
	}
	if c.Logging {
		enableLogging()
	}
	session.DB(c.URL)
	Mongo = &mongoStore{session, c.DB}
	ensureIndexes()
}

func ensureIndexes() {
	session, db := Mongo.NewSession()
	defer session.Close()
	ensureSearchIndex(db)
	ensureSearchExpiresAt(db)
	ensureBusSearchIndex(db)
	ensureBusResultIndex(db)
	ensureBusSearchExpirationIndex(db)
	ensureBusResultExpirationIndex(db)
	ensureFlightSearchIndex(db)
	ensureFlightResultIndex(db)
	ensureFlightSearchExpirationIndex(db)
	ensureFlightResultExpirationIndex(db)
}

func ensureSearchIndex(db *mgo.Database) {
	index := mgo.Index{
		Key:        []string{"integer_id"},
		Unique:     true,
		DropDups:   true,
		Background: true,
	}
	db.C("searches").EnsureIndex(index)
}

func ensureSearchExpiresAt(db *mgo.Database) {
	index := mgo.Index{
		Key:         []string{"created_at"},
		Background:  true,
		ExpireAfter: 3600 * time.Second,
	}
	db.C("searches").EnsureIndex(index)
}

func ensureBusSearchIndex(db *mgo.Database) {
	index := mgo.Index{
		Key:        []string{"route_id", "departure"},
		Unique:     true,
		DropDups:   true,
		Background: true,
	}
	db.C("bus_searches").EnsureIndex(index)
}

func ensureBusSearchExpirationIndex(db *mgo.Database) {
	index := mgo.Index{
		Key:         []string{"expires_at"},
		ExpireAfter: 0,
		Background:  true,
	}
	db.C("bus_searches").EnsureIndex(index)
}

func ensureBusResultExpirationIndex(db *mgo.Database) {
	index := mgo.Index{
		Key:         []string{"expires_at"},
		ExpireAfter: 0,
		Background:  true,
	}
	db.C("bus_results").EnsureIndex(index)
}

func ensureBusResultIndex(db *mgo.Database) {
	index := mgo.Index{
		Key:        []string{"slug"},
		DropDups:   true,
		Background: true,
	}
	db.C("bus_results").EnsureIndex(index)
}

func ensureFlightSearchIndex(db *mgo.Database) {
	index := mgo.Index{
		Key:        []string{"route_id", "departure"},
		Unique:     true,
		DropDups:   true,
		Background: true,
	}
	db.C("flight_searches").EnsureIndex(index)
}

func ensureFlightResultIndex(db *mgo.Database) {
	index := mgo.Index{
		Key:        []string{"slug"},
		DropDups:   true,
		Background: true,
	}
	db.C("flight_results").EnsureIndex(index)
}

func ensureFlightSearchExpirationIndex(db *mgo.Database) {
	index := mgo.Index{
		Key:         []string{"expires_at"},
		ExpireAfter: 0,
		Background:  true,
	}
	db.C("flight_searches").EnsureIndex(index)
}

func ensureFlightResultExpirationIndex(db *mgo.Database) {
	index := mgo.Index{
		Key:         []string{"expires_at"},
		ExpireAfter: 0,
		Background:  true,
	}
	db.C("flight_results").EnsureIndex(index)
}

func enableLogging() {
	mgo.SetDebug(true)
	var aLogger *log.Logger
	aLogger = log.New(os.Stderr, "", log.LstdFlags)
	mgo.SetLogger(aLogger)
}
