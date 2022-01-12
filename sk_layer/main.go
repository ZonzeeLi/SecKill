/**
    @Author:     ZonzeeLi
    @Project:    sk_layer
    @CreateDate: 1/10/2022
    @UpdateDate: xxx
    @Note:       秒杀系统逻辑层
**/

package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"sk_layer/setting"
)

var cfgFile string
var Verbose bool

func main() {

	var RootCmd = &cobra.Command{
		Use:   "SKLayer Server",
		Short: "SKLayer Server",
		Long:  "SKLayer Server",
		Run: func(cmd *cobra.Command, args []string) {
			serviceMap := viper.GetStringMap("service")
			writeProxy2layerGoroutineNum, _ := serviceMap["write_proxy2layer_goroutine_num"].(int64)
			readLayer2proxyGoroutineNum, _ := serviceMap["read_layer2proxy_goroutine_num"].(int64)
			handleUserGoroutineNum, _ := serviceMap["handle_user_goroutine_num"].(int64)
			read2HandleChanSize, _ := serviceMap["read2handle_chan_size"].(int64)
			handle2WriteChanSize, _ := serviceMap["handle2write_chan_size"].(int64)
			maxRequestWaitTimeout, _ := serviceMap["max_request_wait_timeout"].(int64)
			sendToWriteChanTimeout, _ := serviceMap["send_to_write_chan_timeout"].(int64)
			sendToHandleChanTimeout, _ := serviceMap["send_to_handle_chan_timeout"].(int64)
			secKillTokenPassWd, _ := serviceMap["seckill_token_passwd"].(string)
			setting.InitService(writeProxy2layerGoroutineNum, readLayer2proxyGoroutineNum, handleUserGoroutineNum, read2HandleChanSize,
				handle2WriteChanSize, maxRequestWaitTimeout, sendToWriteChanTimeout, sendToHandleChanTimeout, secKillTokenPassWd)

			redisMap := viper.GetStringMap("redis")
			hostRedis, _ := redisMap["host"].(string)
			pwdRedis, _ := redisMap["password"].(string)
			dbRedis, _ := redisMap["db"].(int)
			proxy2layerQueueName, _ := redisMap["proxy2layer_queue_name"].(string)
			layer2proxyQueueName, _ := redisMap["layer2proxy_queue_name"].(string)
			setting.InitRedis(hostRedis, pwdRedis, dbRedis, proxy2layerQueueName, layer2proxyQueueName)

			etcdMap := viper.GetStringMap("etcd")
			hostEtcd, _ := etcdMap["host"].(string)
			productKey, _ := etcdMap["product_key"].(string)
			setting.InitEtcd(hostEtcd, productKey)

			setting.RunService()
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
