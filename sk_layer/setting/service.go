/**
    @Author:     ZonzeeLi
    @Project:    sk_layer
    @CreateDate: 1/10/2022
    @UpdateDate: xxx
    @Note:       服务配置并启动redis队列
**/

package setting

import (
	"sk_layer/config"
	"sk_layer/service/srv_product"
	"sk_layer/service/srv_redis"
	"sk_layer/service/srv_user"
)

func InitService(writeProxy2layerGoroutineNum, readLayer2proxyGoroutineNum, handleUserGoroutineNum, read2HandleChanSize,
	handle2WriteChanSize, maxRequestWaitTimeout, sendToWriteChanTimeout, sendToHandleChanTimeout int64, secKillTokenPassWd string) {

	config.AppConfig.WriteGoroutineNum = int(writeProxy2layerGoroutineNum)
	config.AppConfig.ReadGoroutineNum = int(readLayer2proxyGoroutineNum)
	config.AppConfig.HandleUserGoroutineNum = int(handleUserGoroutineNum)
	config.AppConfig.Read2HandleChanSize = int(read2HandleChanSize)
	config.AppConfig.Handle2WriteChanSize = int(handle2WriteChanSize)
	config.AppConfig.MaxRequestWaitTimeout = int(maxRequestWaitTimeout)
	config.AppConfig.SendToWriteChanTimeout = int(sendToWriteChanTimeout)
	config.AppConfig.SendToHandleChanTimeout = int(sendToHandleChanTimeout)
	config.AppConfig.TokenPassWd = secKillTokenPassWd

	config.SecLayerCtx.SecLayerConf = config.AppConfig
	config.SecLayerCtx.Read2HandleChan = make(chan *config.SecRequest, config.AppConfig.Read2HandleChanSize)
	config.SecLayerCtx.Handle2WriteChan = make(chan *config.SecResponse, config.AppConfig.Handle2WriteChanSize)
	config.SecLayerCtx.HistoryMap = make(map[int]*srv_user.UserBuyHistory, 10000)
	config.SecLayerCtx.ProductCountMgr = srv_product.NewProductCountMgr()
}

func RunService() {
	//启动处理线程
	srv_redis.RunProcess()
}
