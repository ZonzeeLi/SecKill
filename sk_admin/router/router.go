/**
    @Author:     ZonzeeLi
    @Project:    sk_admin
    @CreateDate: 1/7/2022
    @UpdateDate: xxx
    @Note:       路由层
**/

package router

import (
	"sk_admin/controller/activity"
	"sk_admin/controller/product"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	//商品
	r.GET("/product/list", product.GetPorductList)
	r.POST("/product/create", product.CreateProduct)

	//活动
	r.GET("/activity/list", activity.GetActivityList)
	r.POST("/activity/create", activity.CreateActivity)

}
