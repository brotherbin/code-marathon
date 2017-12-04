package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
)

const (
	DEFAULT_MAX_IDLE_CONNS = 10
	DEFAULT_MAX_OPEN_CONNS = 100
)

type MysqlManager struct {
	Host     string
	Port     string
	Database string
	Username string
	Password string
}

type User struct {
	Id   int64  `json:"id" orm:"auto"`
	Name string `json:"name" orm:"column(name)"`
}

func (u *User) TableName() string {
	return "m_user"
}

func NewMysqlManager(host, port, database, username, passowrd string) *MysqlManager {
	mysqlMgr := &MysqlManager{
		Host:     host,
		Port:     port,
		Database: database,
		Username: username,
		Password: passowrd,
	}
	mysqlMgr.init()
	return mysqlMgr
}

func (mm *MysqlManager) init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	ds := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", mm.Username, mm.Password, mm.Host, mm.Port, mm.Database)
	err := orm.RegisterDataBase("default", "mysql", ds, DEFAULT_MAX_IDLE_CONNS, DEFAULT_MAX_OPEN_CONNS)
	if err != nil {
		log.Panic(err)
	}
	orm.RegisterModel(new(User))
}

func (mm *MysqlManager) AddUser(user *User) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(user)
	return id, err
}
