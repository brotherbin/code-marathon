package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

const (
	DEFAULT_MAX_IDLE_CONNS = 10
	DEFAULT_MAX_OPEN_CONNS = 100
)

type MysqlManager struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

type UserModel struct {
	Id   string `json:"id" orm:"auto"`
	Name string `json:"name" orm:"column(name)"`
}

func (mm *MysqlManager) init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	ds := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", mm.Username, mm.Password, mm.Host, mm.Port, mm.Database)
	err := orm.RegisterDataBase("default", "mysql", ds, DEFAULT_MAX_IDLE_CONNS, DEFAULT_MAX_OPEN_CONNS)
	if err != nil {
		log.Panic(err)
	}
}
