/**
    @Author:     ZonzeeLi
    @Project:    sk_proxy
    @CreateDate: 1/11/2022
    @UpdateDate: xxx
    @Note:       etcd初始化及读写监控
**/

package setting

import (
	"context"
	"encoding/json"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"sk_proxy/config"
	"time"
)

func InitEtcd(host, productKey string) {
	//c, err := client.New(client.Config{
	//	Endpoints:               []string{host},
	//	Transport:               client.DefaultTransport,
	//	HeaderTimeoutPerRequest: 5 * time.Second,
	//})
	//if err != nil {
	//	log.Printf("Connect etcd failed. Error: %v", err)
	//}
	//kapi := client.NewKeysAPI(config.SecAdminConfCtx.EtcdConf.EtcdConn)
	//o := client.SetOptions{Dir: true}
	//_, err = kapi.Set(context.Background(), productKey, "", &o)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//config.SecAdminConfCtx.EtcdConf = &config.EtcdConf{
	//	EtcdConn:          c,
	//	EtcdKeysApi:       kapi,
	//	EtcdSecProductKey: productKey,
	//}
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{host},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Printf("Connect etcd failed. Error : %v", err)
	}

	config.SecKillConfCtx.EtcdConf = &config.EtcdConf{
		EtcdConn:          cli,
		EtcdSecProductKey: productKey,
	}

	loadSecConf(cli)
	go watcherSecProductKey(cli, config.SecKillConfCtx.EtcdConf.EtcdSecProductKey)
}

//加载秒杀商品信息
func loadSecConf(cli *clientv3.Client) {
	rsp, err := cli.Get(context.Background(), config.SecKillConfCtx.EtcdConf.EtcdSecProductKey)
	if err != nil {
		log.Printf("get product info failed, err : %v", err)
		return
	}

	var secProductInfo []*config.SecProductInfoConf
	for _, v := range rsp.Kvs {
		err := json.Unmarshal(v.Value, &secProductInfo)
		if err != nil {
			log.Printf("unmarshal json failed, err : %v", err)
			return
		}
	}

	updateSecProductInfo(secProductInfo)
}

//更新秒杀商品信息
func updateSecProductInfo(secProductInfo []*config.SecProductInfoConf) {
	tmp := make(map[int]*config.SecProductInfoConf, 1024)
	for _, v := range secProductInfo {
		tmp[v.ProductId] = v
	}

	config.SecKillConfCtx.RWSecProductLock.Lock()
	config.SecKillConfCtx.SecProductInfoMap = tmp
	config.SecKillConfCtx.RWSecProductLock.Unlock()
}

//监听秒杀商品配置
func watcherSecProductKey(cli *clientv3.Client, key string) {
	for {
		rch := cli.Watch(context.Background(), key)
		var secProductInfo []*config.SecProductInfoConf
		var getConfSucc = true

		for wrsp := range rch {
			for _, ev := range wrsp.Events {
				//删除事件
				if ev.Type == mvccpb.DELETE {
					continue
				}

				//更新事件
				if ev.Type == mvccpb.PUT && string(ev.Kv.Key) == key {
					err := json.Unmarshal(ev.Kv.Value, &secProductInfo)
					if err != nil {
						getConfSucc = false
						continue
					}
				}
			}

			if getConfSucc {
				updateSecProductInfo(secProductInfo)
			}
		}
	}
}
