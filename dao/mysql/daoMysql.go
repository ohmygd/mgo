package mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/ohmygd/mgo/config"
	"sync"
	"time"
)

const (
	maxIdleConnsC int = 20
	maxOpenConnsC     = 20
)

func init() {

}

var (
	Db    map[string]*gorm.DB
	conMu sync.RWMutex
	once  sync.Once
)

type DaoMysql struct {
	InnerDb *gorm.DB
	Db      string
}

func (d *DaoMysql) GetConns() *gorm.DB {
	if d.InnerDb != nil {
		return d.InnerDb
	}

	conMu.RLock()
	if db, ok := Db[d.Db]; ok && db != nil {
		return db
	}
	conMu.RUnlock()

	maxOpenConns := maxOpenConnsC
	if config.GetMysqlMsg("maxOpenConns") != nil {
		maxOpenConns = int(config.GetMysqlMsg("maxOpenConns").(float64))
	}

	maxIdleConns := maxIdleConnsC
	if config.GetMysqlMsg("maxIdleConns") != nil {
		maxIdleConns = int(config.GetMysqlMsg("maxIdleConns").(float64))
	}

	db := getConn(maxOpenConns, maxIdleConns)
	once.Do(func() {
		Db = make(map[string]*gorm.DB, 0)
	})

	conMu.Lock()
	Db[d.Db] = db
	conMu.Unlock()

	return db
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
	// 全局禁用表名复数
	db2.SingularTable(true) // 如果设置为true,`User`的默认表名为`user`,使用`TableName`设置的表名不受影响

	env := config.GetConfigMsg("env")
	logModel := false
	if env == nil || env == "dev" || env == "test" {
		logModel = true
	}
	db2.LogMode(logModel)

	return db2
}
