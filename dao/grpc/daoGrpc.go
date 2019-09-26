/**
 * @Author: Machao
 * @Date: 2019-09-25 10:47
 * @To:
 */
package grpc

import (
	"fmt"
	"github.com/ohmygd/mgo/config"
	"google.golang.org/grpc"
)

type DaoGrpc struct {
	Module string
}

var con *grpc.ClientConn

func (d *DaoGrpc) GetConn() *grpc.ClientConn {
	if con != nil {
		return con
	}

	info := config.GetGrpcMsg(d.Module)
	if info == nil {
		panic("grpc config lost.")
	}

	infoMap := info.(map[string]interface{})
	host := infoMap["host"]
	port := infoMap["port"]

	if host == nil || port == nil {
		panic("grpc config host or port lost.")
	}

	var err error
	con, err = grpc.Dial(host.(string) + ":" + port.(string), grpc.WithInsecure())
	if err != nil {
		panic(fmt.Sprintf("did not connect: %v", err))
	}

	return con
}
