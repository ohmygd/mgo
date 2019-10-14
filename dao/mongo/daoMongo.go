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
	"time"
)

var session *mgo.Session

type DaoMongo struct {
}

func init() {
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

	if username != "" && password != "" {
		dialInfo.Username = username
		dialInfo.Password = password
	}

	var err error

	session, err = mgo.DialWithInfo(dialInfo)

	if err != nil {
		panic(err.Error())
	}

	session.SetMode(mgo.Monotonic, true)
}

func (d *DaoMongo) NewSession() *mgo.Session {
	return session.Copy()
}
