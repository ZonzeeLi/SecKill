/**
    @Author:     ZonzeeLi
    @Project:    sk_layer
    @CreateDate: 1/10/2022
    @UpdateDate: xxx
    @Note:       etcd初始化
**/

package setting

import (
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"sk_layer/config"
	"sk_layer/logic"
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

	config.SecLayerCtx.EtcdConf = &config.EtcdConf{
		EtcdConn:          cli,
		EtcdSecProductKey: productKey,
	}
	logic.LoadProductFromEtcd()

}
