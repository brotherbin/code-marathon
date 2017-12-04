package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type MongoDBManager struct {
	ConnectionURI string
	Database      string
}

const (
	C_PERSON = "person"
)

type Person struct {
	ID   string `bson:"_id,omitempty"`
	Name string `bson:"name"`
}

func NewMongoManager(connURI, database string) *MongoDBManager {
	return &MongoDBManager{
		ConnectionURI: connURI,
		Database:      database,
	}
}

func (mgoMgr *MongoDBManager) Execute(collection, database string, operate func(*mgo.Collection) error) error {
	session, err := mgo.DialWithTimeout(mgoMgr.ConnectionURI, 1*time.Second)
	if err != nil {
		return err
	}
	session.SetMode(mgo.Monotonic, true)
	c := session.DB(database).C(collection)
	return operate(c)
}

func (mgoMgr *MongoDBManager) AddPerson(p *Person) error {
	insert := func(c *mgo.Collection) error {
		if p.ID == "" {
			p.ID = bson.NewObjectId().Hex()
		}
		return c.Insert(p)
	}
	return mgoMgr.Execute(C_PERSON, mgoMgr.Database, insert)
}
