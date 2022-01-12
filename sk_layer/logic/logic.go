/**
    @Author:     ZonzeeLi
    @Project:    sk_layer
    @CreateDate: 1/10/2022
    @UpdateDate: xxx
    @Note:       etcd加载内存到本地并监控
**/

package logic

import (
	"context"
	"encoding/json"
	"go.etcd.io/etcd/api/v3/mvccpb"
	"log"
	"sk_layer/config"
	"sk_layer/service/srv_limit"
	"time"
)

//从Etcd中加载商品数据
func LoadProductFromEtcd() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	//从etcd获取商品数据
	rsp, err := config.SecLayerCtx.EtcdConf.EtcdConn.Get(ctx, config.SecLayerCtx.EtcdConf.EtcdSecProductKey)
	if err != nil {
		log.Printf("get [%s] from etcd failed. Error : %v", config.SecLayerCtx.EtcdConf.EtcdSecProductKey, err)
		return
	}

	//结构转换
	var secProductInfo []*config.SecProductInfoConf
	for _, v := range rsp.Kvs {
		err = json.Unmarshal(v.Value, &secProductInfo)
		if err != nil {
			log.Printf("Unmarshal sec product info failed. Error : %v", err)
			return
		}
	}

	UpdateSecProductInfo(secProductInfo)
	log.Printf("update product info success, data : %v", secProductInfo)

	InitSecProductWatcher()
	log.Printf("init ectd watcher success.")

	return
}

//更新商品信息
func UpdateSecProductInfo(secProductInfo []*config.SecProductInfoConf) {
	tmp := make(map[int]*config.SecProductInfoConf, 1024)

	for _, v := range secProductInfo {
		productInfo := v
		productInfo.SecLimit = &srv_limit.SecLimit{}
		tmp[v.ProductId] = productInfo
	}

	config.SecLayerCtx.RWSecProductLock.Lock()
	config.SecLayerCtx.SecLayerConf.SecProductInfoMap = tmp
	config.SecLayerCtx.RWSecProductLock.Unlock()
}

//监控商品变化
func InitSecProductWatcher() {
	go watchSecProductKey()
}

func watchSecProductKey() {
	key := config.SecLayerCtx.EtcdConf.EtcdSecProductKey

	var err error
	for {
		rch := config.SecLayerCtx.EtcdConf.EtcdConn.Watch(context.Background(), key)
		var secProductInfo []*config.SecProductInfoConf
		var getConfSucc = true

		for wrsp := range rch {
			for _, ev := range wrsp.Events {
				//删除事件
				if ev.Type == mvccpb.DELETE {
					log.Printf("key[%s] 's config deleted", key)
					continue
				}

				//更新事件
				if ev.Type == mvccpb.PUT && string(ev.Kv.Key) == key {
					err = json.Unmarshal(ev.Kv.Value, &secProductInfo)
					if err != nil {
						log.Printf("key [%s], Unmarshal[%s]. Error : %v", key, err)
						getConfSucc = false
						continue
					}
				}
				log.Printf("get config from etcd, %s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			}

			if getConfSucc {
				log.Printf("get config from etcd success, %v", secProductInfo)
				UpdateSecProductInfo(secProductInfo)
			}
		}
	}
}
