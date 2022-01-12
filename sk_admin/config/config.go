/**
    @Author:     ZonzeeLi
    @Project:    sk_admin
    @CreateDate: 1/7/2022
    @UpdateDate: xxx
    @Note:       配置层及全局变量
**/

package config

import (
	"github.com/gohouse/gorose"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var SecAdminConfCtx = &SecAdminConf{}

type SecAdminConf struct {
	DbConf   *DbConf
	EtcdConf *EtcdConf
}

//数据库配置
type DbConf struct {
	DbConn gorose.Connection //链接
}

//Etcd配置
type EtcdConf struct {
	EtcdConn          *clientv3.Client //链接
	EtcdSecProductKey string           //商品键
}
