##  秒杀系统 SecKill

高并发、秒杀、抢票系统项目等是现在热门的实战项目，这次我寻找了一个github上比较不错的[SecKill](https://github.com/yxhsea/SecKill)项目，这里介绍一下如何使用运行该项目。

## 项目介绍
该部分摘自原作者README
### 秒杀接入层
1. 从Etcd中加载秒杀活动数据到内存当中。
2. 监听Etcd中的数据变化，实时加载数据到内存中。
3. 从Redis中加载黑名单数据到内存当中。
4. 设置白名单。
5. 对用户请求进行黑名单限制。
6. 对用户请求进行流量限制、秒级限制、分级限制。
7. 将用户数据进行签名校验、检验参数的合法性。
8. 接收逻辑层的结果实时返回给用户。

### 秒杀逻辑层
1. 从Etcd中加载秒杀活动数据到内存当中。
2. 监听Etcd中的数据变化，实时加载数据到内存中。
3. 处理Redis队列中的请求。
4. 限制用户对商品的购买次数。
5. 对商品的抢购频次进行限制。
5. 对商品的抢购概率进行限制。
6. 对合法的请求给予生成抢购资格Token令牌。

### 秒杀管理层
1. 添加商品数据。
2. 添加抢购活动数据。
3. 将数据同步到Etcd。
4. 将数据同步到数据库。

## 重要提示

1. 虽然是几年前开发的，但是内部调用环环相扣，逻辑复杂度和实用性都比较强，学习了一下这个项目后，因为本身项目的逻辑挺庞大，再修改框架的话需要花几天时间，所以我基本没有改动原来的代码框架，包括原先的依赖，只是将旧版本失效的用新的代替，而且将一些个人认为作者的小失误修改完善了一下，在本地成功运行测试，本人更多的学习了这个项目的思想和处理方式，并没有做过多的二次开发。
2. 如果想要二次开发成自己的框架，或者使用自己的依赖工具，请核对内部函数function的处理方法去找相应的版本，否则可能出现初始化或者是导入依赖的错误。
3. 因为该项目还有一些点并没有开发完成，但是一些参数还是添加写入了，比如：活动购买并没有处理剩余数量、查询显示并不完整等等，有些参数没有使用不代表没有意义，如果开发者想要自己完善该部分，请仔细理解该项目的逻辑处理然后根据自己的需求再进行补充。

## 目录结构

```sk_admin(秒杀管理层)
├─config          (配置包)
├─controller      (api层)
│  ├─activity     (活动处理)
│  └─product      (商品处理)
├─model           (数据库处理)
├─router          (路由层)
├─service         (服务层)
└─setting         (初始化)
```

```sk_layer(秒杀逻辑层)
├─config    	    (配置包)
├─logic           (逻辑处理)
├─service		      (服务层)
│  ├─srv_err      (错误处理)
│  ├─srv_limit    (频率限制)
│  ├─srv_product  (商品管理)
│  ├─srv_redis    (redis队列读写及用户数据处理)
│  └─srv_user     (用户购买历史记录)
└─setting  	      (初始化)
```

```sk_proxy(秒杀接入层)
├─config		      (配置包)
├─controller	    (api)
├─router		      (路由层)
├─service		      (逻辑处理)
│  ├─srv_err      (错误处理)
│  ├─srv_limit    (黑名单限制、反作弊、频率限制)
│  ├─srv_redis    (redis读写)
│  └─srv_sec      (秒杀接口处理逻辑)
└─setting		      (初始化)
```

## 主要功能

- 查询活动商品：管理层可以从数据库Mysql中查询具体的商品和活动信息，接入层可以根据id查询或者返回商品信息的状态信息。
- 添加活动商品：管理层可以添加商品和活动信息到数据库Mysql并同步到Etcd。
- Etcd读写监控：接入和逻辑层从Etcd中加载数据信息，并且监听数据变化更新，管理层同步Etcd。
- Redis队列读写：接入和逻辑层利用Redis的队列移除获取请求，并处理商品信息和请求。
- 安全校验：对用户数据进行签名校验、检验参数合法性，并做白名单和黑名单限制。
- 流量限制：对用户及用户IP进行秒级和分级的请求频率限制。

## 系统架构图
该图摘自原作者仓库

![系统架构图](https://github.com/ZonzeeLi/SecKill/blob/master/framework.png)

## 项目初始化
以下内容介绍项目依赖中的特殊说明，以及数据库和Etcd的使用工具。

#### 依赖
- [github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)
- [github.com/go-redis/redis](https://github.com/go-redis/redis)
- [github.com/spf13/cobra](https://github.com/spf13/cobra)
- [github.com/spf13/viper](github.com/spf13/viper)
- [go.etcd.io/etcd/client/v3](https://pkg.go.dev/go.etcd.io/etcd/client/v3#section-readme)
- [github.com/go-sql-driver/mysql v1.5.0](https://github.com/go-sql-driver/mysql) 
- [github.com/gohouse/gorose v1.0.5](https://github.com/gohouse/gorose)
- [github.com/Unknwon/com v0.0.0-20180617003950-da59b551951d](https://github.com/Unknwon/com)

注：这里sql驱动使用的是v1.5.0的版本，原因是gorose原作者使用的是v1版本，现在gorose的作者已经不维护v1了，v2的可以继续使用而且是很不错的操作数据库依赖，v1版本的驱动如果是v1.4.0的话则会报错，这个问题我排查了半天，早知道改依赖了....另外com包其实很多都是常规基本的函数操作，用go语言原生的包也没问题。所有的工具都可以二次开发，请自行修改。

#### 工具
- Etcdkeeper web端Etcd可视化工具，支持v2和v3
- Redis Desktop Manager 桌面端Redis可视化工具
- Postman api管理工具

注：具体使用请参考工具的详细教程。

#### 数据库
实现建立好mysql数据库生成表activity和product，不用事先往表中存放数据，数据由接口post存入，因为项目代码中要同步Etcd，并没有对mysql中的数据进行监控读取，sql代码中数据只是举例。
```mysql
SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for activity
-- ----------------------------
DROP TABLE IF EXISTS `activity`;
CREATE TABLE `activity` (
  `activity_id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '活动Id',
  `activity_name` varchar(50) NOT NULL DEFAULT '' COMMENT '活动名称',
  `product_id` int(11) unsigned NOT NULL COMMENT '商品Id',
  `start_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '活动开始时间',
  `end_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '活动结束时间',
  `total` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '商品数量',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '活动状态',
  `sec_speed` int(5) unsigned NOT NULL DEFAULT '0' COMMENT '每秒限制多少个商品售出',
  `buy_limit` int(5) unsigned NOT NULL COMMENT '购买限制',
  `buy_rate` decimal(2,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '购买限制',
  PRIMARY KEY (`activity_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='@活动数据表';

-- ----------------------------
-- Records of activity
# -- ----------------------------
# INSERT INTO `activity` VALUES ('1', '香蕉大甩卖', '1', '530871061', '530872061', '20', '0', '1', '1', '0.20');
# INSERT INTO `activity` VALUES ('2', '苹果大甩卖', '2', '530871061', '530872061', '20', '0', '1', '1', '0.20');
# INSERT INTO `activity` VALUES ('3', '桃子大甩卖', '3', '1530928052', '1530989052', '20', '0', '1', '1', '0.20');
# INSERT INTO `activity` VALUES ('4', '梨子大甩卖', '4', '1530928052', '1530989052', '20', '0', '1', '1', '0.20');

-- ----------------------------
-- Table structure for product
-- ----------------------------
DROP TABLE IF EXISTS `product`;
CREATE TABLE `product` (
  `product_id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '商品Id',
  `product_name` varchar(50) NOT NULL DEFAULT '' COMMENT '商品名称',
  `total` int(5) unsigned NOT NULL DEFAULT '0' COMMENT '商品数量',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '商品状态',
  PRIMARY KEY (`product_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='@商品数据表';

-- ----------------------------
-- Records of product
-- ----------------------------
# INSERT INTO `product` VALUES ('1', '香蕉', '100', '1');
# INSERT INTO `product` VALUES ('2', '苹果', '100', '1');
# INSERT INTO `product` VALUES ('3', '桃子', '100', '1');
# INSERT INTO `product` VALUES ('4', '梨子', '100', '1');
```

#### Etcd
如果下载安装最新的Etcd，默认不开启v2服务，命令需要添加如下启动项：
```
etcd --enable-v2=true
```

## Start
按照顺序对三层进行分别开发介绍，注意需要将`sk_admin`、`sk_layer`、`sk_proxy`分别打开运行，即将clone或者download下来的文件里的三个子项目`sk_admin`、`sk_layer`、`sk_proxy`在ide分别打开，开三个终端启动。

### sk_admin
#### 端口介绍
```
POST 127.0.0.1:8081/product/create 添加商品
{
    "product_name": "桃子",
    "total": 100,
    "status": 0
}
```
```
POST 127.0.0.1:8081/activity/create 添加活动
{
    "activity_name": "桃子大甩卖",
    "product_id": 1,
    "start_time": 1641945600,
    "end_time": 1642118400,
    "total": 30,
    "status": 0,
    "speed": 1,
    "buy_limit": 10,
    "buy_rate": 0.2
}
```
```
GET 127.0.0.1:8081/product/list 查询商品信息列表
```
```
GET 127.0.0.1:8081/activity/list 查询活动信息列表
```

#### 运行说明
修改配置文件`conf.toml`中Mysql的配置文件，如果Etcd没有改配置的话，则不用修改Etcd部分的配置。
```conf.toml
[http]
host = "127.0.0.1:8081"

[mysql]   //主要修改该部分
host = "127.0.0.1"
port = "3306"
user = "root"
pass_wd = ""
db = "seckill"

[etcd]
host = "127.0.0.1:2379"
product_key = "sec_kill_product"
```

### sk_layer
#### 运行说明
该层主要是进行队列的监听读写，没有端口使用及其他太多的改动。
```conf.toml
[service]
write_proxy2layer_goroutine_num = 16
read_layer2proxy_goroutine_num = 16
handle_user_goroutine_num = 16
read2handle_chan_size = 100000
handle2write_chan_size = 100000
max_request_wait_timeout = 30
#单位是毫秒
send_to_write_chan_timeout= 100
send_to_handle_chan_timeout = 100
#token秘钥
seckill_token_passwd = "fIOxU7iik65vVvBGtNcnrjL4E9MdRpTfzzxE3dx6b7BAHN5etUdSzRW5yjzHzF"

[redis]
host = "127.0.0.1:6379"
password = ""
db = 0
proxy2layer_queue_name = "sec_queue"   ### 接入层->逻辑层
layer2proxy_queue_name = "recv_queue"  ### 逻辑层->接入层

[etcd]
host = "127.0.0.1:2379"
product_key = "sec_kill_product"
```

### sk_proxy
#### 端口介绍
```
GET 127.0.0.1:8082/sec/info?product_id=2 秒杀接入_id查询商品
```
```
GET 127.0.0.1:8082/sec/list 秒杀接入_查询商品列表
```
```
POST 127.0.0.1:8082/sec/kill 秒杀接入请求 
{
    "product_id": 2,							// 商品ID
    "user_id": 1,								// 用户ID
    "source": "127.0.0.1",    					// 用户IP地址
    "auth_code": "userauthcode",
    "sec_time": 1641890448,						
    "nance": "dsdsdjkdjskdjksdjhuieurierei"
}
Header添加"AuthSign"和"Referer"，Referer为引用的URL，作者可能是在规划项目的时候想用这个获取URL，这里直接添加用户IP即可，如上面的json数据，则添加"127.0.0.1"，也就是和"source"对应，"AuthSign"由md5生成，生成例子在项目中附加代码generate.go中
```

![Header](https://github.com/ZonzeeLi/SecKill/blob/master/Header.png)

#### 运行说明
该配置文件中改动也不大，主要就是白名单部分补充允许请求通过的URL，这个地方的处理逻辑写的并不完整，该代码中的逻辑是判断请求的获取到的`Referer`是否在`refer_whitelist`中，如果在的话则用户检查通过。当然实际中不能把所有的URL都存起来，不过我并没有改动原作者的写法，可能是想让开发者自己补充处理逻辑，所以这个地方我理解是在部署的时候可能会使用代理来做。
```conf.toml
[http]
host = "127.0.0.1:8082"

[redis]
host = "127.0.0.1:6379"
password = ""
db = 0
proxy2layer_queue_name = "sec_queue"          ### 接入层->逻辑层
layer2proxy_queue_name = "recv_queue"         ### 逻辑层->接入层
id_black_list_hash = "id_black_list_hash"     ### 用户黑名单Hash表
ip_black_list_hash = "ip_black_list_hash"     ### IP黑名单Hash表
id_black_list_queue = "id_black_list_queue"   ### 用户黑名单队列
ip_black_list_queue = "ip_black_list_queue"   ### IP黑名单队列

[etcd]
host = "127.0.0.1:2379"
product_key = "sec_kill_product"
black_list_key = "sec_kill_backlist"

[service]
#频率控制阈值
ip_sec_access_limit = 50
user_sec_access_limit = 1
ip_min_access_limit = 500
user_min_access_limit = 10
#cookie 秘钥
cookie_secretkey = "WK5wJOiuYaXRUlPsxo3LZEbpCNSyvm8T"  
#白名单(用","分割)
refer_whitelist= "127.0.0.1"           	// 注意这个地方，白名单中请求可通过的URL
#goroutine数量控制
write_proxy2layer_goroutine_num = 16
read_proxy2layer_goroutine_num = 16
```
