/**
    @Author:     ZonzeeLi
    @Project:    sk_admin
    @CreateDate: 1/7/2022
    @UpdateDate: xxx
    @Note:       秒杀系统管理层
**/

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sk_admin/setting"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var Verbose bool

func main() {

	var RootCmd = &cobra.Command{
		Use:   "SKAdmin Server",
		Short: "SKAdmin Server",
		Long:  "SKAdmin Server",
		Run: func(cmd *cobra.Command, args []string) {
			mysqlMap := viper.GetStringMap("mysql")
			hostMysql, _ := mysqlMap["host"].(string)
			portMysql, _ := mysqlMap["port"].(string)
			userMysql, _ := mysqlMap["user"].(string)
			pwdMysql, _ := mysqlMap["pass_wd"].(string)
			dbMysql, _ := mysqlMap["db"].(string)
			setting.InitMysql(hostMysql, portMysql, userMysql, pwdMysql, dbMysql)

			etcdMap := viper.GetStringMap("etcd")
			hostEtcd, _ := etcdMap["host"].(string)
			productKey, _ := etcdMap["product_key"].(string)
			setting.InitEtcd(hostEtcd, productKey)

			httpMap := viper.GetStringMap("http")
			hostHttp, _ := httpMap["host"].(string)
			setting.InitServer(hostHttp)
		},
	}

	cobra.OnInitialize(initConfig)
	RootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file")
	RootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

}

func initConfig() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName("conf")
		viper.SetConfigType("toml") //配置文件扩展名
		viper.AddConfigPath("./")
		viper.AddConfigPath(dir)
		viper.AutomaticEnv()
	}

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		fmt.Println(err)
	}
}
