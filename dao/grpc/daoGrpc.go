/**
 * @Author: Machao
 * @Date: 2019-09-25 10:47
 * @To:
 */
package grpc

import (
	"github.com/ohmygd/mgo/config"
	"github.com/ohmygd/mgo/merror"
	"github.com/ohmygd/mgo/pc"
	"google.golang.org/grpc"
	"sync"
)

type DaoGrpc struct {
	Module string
}

var cons map[string]*grpc.ClientConn
var once sync.Once

func (d *DaoGrpc) GetConn() (*grpc.ClientConn, error) {
	if cons != nil && cons[d.Module] != nil {
		return cons[d.Module], nil
	}

	info := config.GetGrpcMsg(d.Module)
	if info == nil {
		return nil, merror.NewWM(pc.ErrorGrpcConfig, "grpc config lost.")
	}

	infoMap := info.(map[string]interface{})
	host := infoMap["host"]
	port := infoMap["port"]

	if host == nil || port == nil {
		return nil, merror.NewWM(pc.ErrorGrpcConfig, "grpc config lost.")
	}

	var con *grpc.ClientConn
	var err error
	con, err = grpc.Dial(host.(string)+":"+port.(string), grpc.WithInsecure())
	if err != nil {
		return nil, merror.NewWM(pc.ErrorGrpcConnect, err.Error())
	}

	once.Do(func() {
		cons = make(map[string]*grpc.ClientConn, 0)
	})

	cons[d.Module] = con

	return con, nil
}
