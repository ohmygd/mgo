package config

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"go.etcd.io/etcd/clientv3"
	"log"
	"strconv"
	"sync"
	"time"
)

var (
	// 并发安全map
	allC sync.Map

	etcdName string
)

func init() {
	// 获取etcd配置
	hosts, res := getConf()

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   hosts,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Println("connect failed, err:" + err.Error())
		return
	}

	setConfig(cli, res)

	watchConfig(cli, res)
}

func watchConfig(cli *clientv3.Client, res []string) {
	for _, v := range res {
		go watchConfigBase(cli, v)
	}
}

func watchConfigBase(cli *clientv3.Client, key string) {
	for {
		rch := cli.Watch(context.Background(), getKey(key))
		for v := range rch {
			for _, ev := range v.Events {

				var res map[string]interface{}
				err := json.Unmarshal(ev.Kv.Value, &res)
				if err == nil {
					allC.Store(key, res)
				} else {
					log.Printf("watch etcd config json unmarshal error. err: %s", err.Error())
				}
			}
		}
	}
}

func setConfig(cli *clientv3.Client, res []string) {
	for _, v := range res {
		setConfigBase(cli, v)
	}
}

func setConfigBase(cli *clientv3.Client, key string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, getKey(key))
	cancel()
	if err != nil {
		panic(fmt.Sprintf("config %s get failed, err:%s", key, err))
		return
	}

	if len(resp.Kvs) == 0 {
		return
	}

	var c map[string]interface{}

	for _, ev := range resp.Kvs {
		err = json.Unmarshal(ev.Value, &c)
		if err != nil {
			panic(fmt.Sprintf("config %s json unmarshal err:%s", key, err.Error()))
		}

		// 设置c内容
		allC.Store(key, c)
	}
}

func getKey(key string) string {
	return "/" + etcdName + "/" + key;
}

func getConf() (res []string, cf []string) {
	viper.SetConfigName("etcd")
	viper.AddConfigPath("config")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		viper.AddConfigPath("../config")
		err = viper.ReadInConfig()
		if err != nil {
			panic("etcd config lost, err: " + err.Error())
		}
	}

	hosts := viper.Get("hosts")
	if hosts == nil {
		panic("etcd config lost")
	}

	name := viper.Get("name")
	if name == nil {
		panic("etcd name config lost")
	}

	etcdName = name.(string)

	hostsInfo := hosts.([]interface{})
	for _, v := range hostsInfo {
		res = append(res, v.(string))
	}

	viper.SetConfigName("cf")
	err = viper.ReadInConfig()
	if err != nil {
		viper.AddConfigPath("../config")
		err = viper.ReadInConfig()
		if err != nil {
			panic("cf config lost, err: " + err.Error())
		}
	}
	cfName := viper.Get("configFile")
	if cfName == nil {
		panic("cf config lost")
	}
	cfNameInfo := cfName.([]interface{})
	for _, v := range cfNameInfo {
		cf = append(cf, v.(string))
	}

	return
}

func GetCodeMsg(code int) string {
	cS := strconv.Itoa(code)

	codeInfo, ok := allC.Load("code")
	if !ok {
		return "unGet error"
	}

	res := codeInfo.(map[string]interface{})

	if res[cS] != nil {
		return res[cS].(string)
	}

	return ""
}

func GetMysqlMsg(key string) interface{} {
	mysqlInfo, ok := allC.Load("mysql")
	if !ok {
		return nil
	}

	res := mysqlInfo.(map[string]interface{})

	return res[key]
}

func GetRedisMsg(key string) interface{} {
	redisInfo, ok := allC.Load("redis")
	if !ok {
		return nil
	}

	res := redisInfo.(map[string]interface{})

	return res[key]
}

func GetConfigMsg(key string) interface{} {
	configInfo, ok := allC.Load("config")
	if !ok {
		return nil
	}

	res := configInfo.(map[string]interface{})

	return res[key]
}

func GetHttpMsg(key string) interface{} {
	httpInfo, ok := allC.Load("http")
	if !ok {
		return nil
	}

	res := httpInfo.(map[string]interface{})

	return res[key]
}
