package main

import "testing"

func Test_AddUser(t *testing.T) {
	host := "0.0.0.0"
	port := "3306"
	database := "marathon"
	username := "root"
	password := "123456"
	mysqlMgr := NewMysqlManager(host, port, database, username, password)
	user := &User{
		Name: "dobbin",
	}
	_, err := mysqlMgr.AddUser(user)
	if err != nil {
		t.Error(err)
	}
}
