/**
    @Author:     ZonzeeLi
    @Project:    sk_proxy
    @CreateDate: 1/11/2022
    @UpdateDate: xxx
    @Note:       路由层
**/

package router

import (
	"github.com/gin-gonic/gin"
	"sk_proxy/controller"
)

func Router(r *gin.Engine) {
	//秒杀管理
	r.GET("/sec/info", controller.SecInfo)
	r.GET("/sec/list", controller.SecInfoList)
	r.POST("/sec/kill", controller.SecKill)

}
