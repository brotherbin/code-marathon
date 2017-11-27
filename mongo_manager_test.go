package main

import (
	"testing"
)

func TestAddPerson(t *testing.T) {
	p := &Person{
		Name: "dobbin",
	}
	mgoMgr := NewMongoManager("0.0.0.0:27017", "marathon")
	err := mgoMgr.AddPerson(p)
	if err != nil {
		t.Error(err)
	}
}
