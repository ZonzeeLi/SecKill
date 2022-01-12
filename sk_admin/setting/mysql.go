/**
    @Author:     ZonzeeLi
    @Project:    sk_admin
    @CreateDate: 1/7/2022
    @UpdateDate: xxx
    @Note:       mysql初始化
**/

package setting

import (
	"fmt"
	"sk_admin/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gohouse/gorose"
)

func InitMysql(hostMysql, portMysql, userMysql, pwdMysql, dbMysql string) {
	dsn := userMysql + ":" + pwdMysql + "@tcp(" + hostMysql + ":" + portMysql + ")/" + dbMysql + "?charset=utf8mb4&parseTime=true"
	fmt.Println(dsn)
	//dsn := "root:@tcp(127.0.0.1:3306)/seckill?charset=utf8mb4&parseTime=true"
	Db := &gorose.DbConfigSingle{
		Driver:          "mysql",
		EnableQueryLog:  true,
		SetMaxOpenConns: 300,
		SetMaxIdleConns: 10,
		Dsn:             dsn,
	}
	connection, err := gorose.Open(Db)
	if err != nil {
		fmt.Printf("failed connect mysql : %v\n", err)
		return
	}
	config.SecAdminConfCtx.DbConf = &config.DbConf{DbConn: *connection}
	list, _ := config.SecAdminConfCtx.DbConf.DbConn.Table("product").Get()
	fmt.Println(list)
}

//func InitMysql(hostMysql, portMysql, userMysql, pwdMysql, dbMysql string) {
//	dsn := "root:@tcp(127.0.0.1:3306)/seckill?charset=utf8mb4&parseTime=true"
//	//dsn := "root:@tcp(127.0.0.1:3306)/seckill?charset=utf8mb4&parseTime=true"
//	Db := &gorose.DbConfigSingle{
//		Driver:          "mysql",
//		EnableQueryLog:  true,
//		SetMaxOpenConns: 300,
//		SetMaxIdleConns: 10,
//		Dsn:             dsn,
//	}
//	connection, err := gorose.Open(Db)
//	if err != nil {
//		fmt.Printf("failed connect mysql : %v\n", err)
//		return
//	}
//
//	//db := connection.NewSession()
//	list, err := connection.Table("product").Get()
//
//	if err != nil {
//		fmt.Printf("failed get list : %v \n", err)
//	}
//	fmt.Println(list)
//
//	//config.SecAdminConfCtx.DbConf = &config.DbConf{
//	//	DbConn: *connection,
//	//}
//}
