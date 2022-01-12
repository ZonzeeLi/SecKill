/**
    @Author:     ZonzeeLi
    @Project:    sk_proxy
    @CreateDate: 1/11/2022
    @UpdateDate: xxx
    @Note:       服务启动
**/

package setting

import (
	"github.com/gin-gonic/gin"
	"log"
	"sk_proxy/config"
	"sk_proxy/router"
	"strings"
)

//初始化Http服务
func InitServer(host string) {
	r := gin.Default()
	router.Router(r)
	err := r.Run(host)
	if err != nil {
		log.Printf("Init http server. Error : %v", err)
	}
}

//初始化服务配置项
func InitServiceConfig(ipSecAccessLimit, ipMinAccessLimit, userSecAccessLimit, userMinAccessLimit,
	writeProxy2layerGoroutineNum, readProxy2layerGoroutineNum int64, cookieSecretKey, referWhitelist string) {
	config.SecKillConfCtx.AccessLimitConf.IPSecAccessLimit = int(ipSecAccessLimit)
	config.SecKillConfCtx.AccessLimitConf.IPMinAccessLimit = int(ipMinAccessLimit)
	config.SecKillConfCtx.AccessLimitConf.UserSecAccessLimit = int(userSecAccessLimit)
	config.SecKillConfCtx.AccessLimitConf.UserMinAccessLimit = int(userMinAccessLimit)
	config.SecKillConfCtx.WriteProxy2LayerGoroutineNum = int(writeProxy2layerGoroutineNum)
	config.SecKillConfCtx.ReadProxy2LayerGoroutineNum = int(readProxy2layerGoroutineNum)
	config.SecKillConfCtx.CookieSecretKey = cookieSecretKey
	config.SecKillConfCtx.ReferWhiteList = strings.Split(referWhitelist, ",")
}
