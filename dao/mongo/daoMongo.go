/**
 * @Author: Machao
 * @Date: 2019-10-11 11:06
 * @To:
 */
package mongo

import (
	"github.com/ohmygd/mgo/config"
	"gopkg.in/mgo.v2"
	"strings"
	"sync"
	"time"
)

var sessionMap map[string]*mgo.Session
var once sync.Once

type DaoMongo struct {
	Db string
}

func (d *DaoMongo) getSession() (*mgo.Session, error) {
	if sessionMap != nil && sessionMap[d.Db] != nil {
		return sessionMap[d.Db], nil
	}

	hostInfo := config.GetMongoMsg("hosts").(string)
	username := config.GetMongoMsg("username").(string)
	password := config.GetMongoMsg("password").(string)
	poolLimit := int(config.GetMongoMsg("poolLimit").(float64))

	if hostInfo == "" {
		panic("mongo config lost.")
	}

	dialInfo := &mgo.DialInfo{
		Addrs:     strings.Split(hostInfo, ","),
		Direct:    false,
		Timeout:   time.Second * 2,
		PoolLimit: poolLimit,
	}

	var session *mgo.Session
	var err error

	session, err = mgo.DialWithInfo(dialInfo)
	if err != nil {
		panic(err.Error())
	}

	if username != "" && password != "" {
		err = session.DB(d.Db).Login(username, password)
		if err != nil {
			panic(err)
		}
	}

	session.SetMode(mgo.Monotonic, true)

	once.Do(func() {
		sessionMap = make(map[string]*mgo.Session, 0)
	})

	sessionMap[d.Db] = session

	return session, nil
}

func (d *DaoMongo) GetDb() *mgo.Database {
	session, err := d.getSession()
	if err != nil {
		panic(err)
	}

	return session.DB(d.Db)
}

// todo Session close 暂时没发现哪里需要用到

func (d *DaoMongo) NewSession() *mgo.Session {
	session, err := d.getSession()
	if err != nil {
		panic(err)
	}

	return session
}

func (d *DaoMongo) SessionCopy() *mgo.Session {
	session, err := d.getSession()
	if err != nil {
		panic(err)
	}
	return session.Copy()
}
