package mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/ohmygd/mgo/config"
	"time"
)

const (
	maxIdleConnsC int= 20
	maxOpenConnsC = 20
)

func init() {
	maxOpenConns := maxOpenConnsC
	if config.GetMysqlMsg("maxOpenConns") != nil {
		maxOpenConns = int(config.GetMysqlMsg("maxOpenConns").(float64))
	}

	maxIdleConns := maxIdleConnsC
	if config.GetMysqlMsg("maxIdleConns") != nil {
		maxIdleConns = int(config.GetMysqlMsg("maxIdleConns").(float64))
	}

	Db = getConn(maxOpenConns, maxIdleConns)
}

var Db *gorm.DB

type DaoMysql struct {
	InnerDb *gorm.DB
}

func (d *DaoMysql)GetConns() *gorm.DB {
	if d.InnerDb != nil {
		return d.InnerDb
	}

	return Db
}

func getConn(maxOpenConns, maxIdleConns int) *gorm.DB {
	user := config.GetMysqlMsg("user").(string)
	pwd := config.GetMysqlMsg("pwd").(string)
	host := config.GetMysqlMsg("host").(string)
	dbName := config.GetMysqlMsg("db").(string)
	port := config.GetMysqlMsg("port").(string)

	mysqlInfo := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pwd, host, port, dbName)

	db2, err := gorm.Open("mysql", mysqlInfo)
	if err != nil {
		panic("mysql conn error, err:" + err.Error())
		return nil
	}

	db2.DB().SetMaxOpenConns(maxOpenConns)
	db2.DB().SetMaxIdleConns(maxIdleConns)
	db2.DB().SetConnMaxLifetime(time.Minute)
	db2.SingularTable(true)
	db2.LogMode(true)

	return db2
}