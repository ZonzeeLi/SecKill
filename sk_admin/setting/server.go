/**
    @Author:     ZonzeeLi
    @Project:    sk_admin
    @CreateDate: 1/7/2022
    @UpdateDate: xxx
    @Note:       服务启动
**/

package setting

import (
	"log"
	"sk_admin/router"

	"github.com/gin-gonic/gin"
)

func InitServer(host string) {
	r := gin.Default()
	router.Router(r)
	err := r.Run(host)
	if err != nil {
		log.Printf("Init http server. Error : %v", err)
	}
}
