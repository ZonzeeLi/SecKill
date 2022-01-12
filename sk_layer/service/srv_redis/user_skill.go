/**
    @Author:     ZonzeeLi
    @Project:    sk_layer
    @CreateDate: 1/10/2022
    @UpdateDate: xxx
    @Note:       从队列中获取用户请求并进行处理再推入队列中
**/

package srv_redis

import (
	"crypto/md5"
	"fmt"
	"log"
	"sk_layer/config"
	"sk_layer/service/srv_err"
	"sk_layer/service/srv_user"
	"time"
)

func HandleUser() {
	log.Println("handle user running")
	for req := range config.SecLayerCtx.Read2HandleChan {
		log.Printf("begin process request : %v", req)
		res, err := HandleSkill(req)
		if err != nil {
			log.Printf("process request %v failed, err : %v", err)
			res = &config.SecResponse{
				Code: srv_err.ErrServiceBusy,
			}
		}
		fmt.Println("处理中~~ ", res)
		timer := time.NewTicker(time.Millisecond * time.Duration(config.SecLayerCtx.SecLayerConf.SendToWriteChanTimeout))
		select {
		case config.SecLayerCtx.Handle2WriteChan <- res:
		case <-timer.C:
			log.Printf("send to response chan timeout, res : %v", res)
			break
		}
	}
	return
}

func HandleSkill(req *config.SecRequest) (res *config.SecResponse, err error) {
	config.SecLayerCtx.RWSecProductLock.RLock()
	defer config.SecLayerCtx.RWSecProductLock.RUnlock()

	res = &config.SecResponse{}
	res.ProductId = req.ProductId
	res.UserId = req.UserId

	product, ok := config.SecLayerCtx.SecLayerConf.SecProductInfoMap[req.ProductId]
	if !ok {
		log.Printf("not found product : %v", req.ProductId)
		res.Code = srv_err.ErrNotFoundProduct
		return
	}

	if product.Status == srv_err.ProductStatusSoldout {
		res.Code = srv_err.ErrSoldout
		return
	}

	nowTime := time.Now().Unix()
	alreadySoldOut := product.SecLimit.Check(nowTime)
	if alreadySoldOut >= product.SoldMaxLimit {
		res.Code = srv_err.ErrRetry
		return
	}

	// 判断用户购买历史记录是否超过了单人购买的数量限制
	config.SecLayerCtx.HistoryMapLock.Lock()
	userHistory, ok := config.SecLayerCtx.HistoryMap[req.UserId]
	if !ok {
		userHistory = &srv_user.UserBuyHistory{
			History: make(map[int]int, 16),
		}
		config.SecLayerCtx.HistoryMap[req.UserId] = userHistory
	}
	historyCount := userHistory.GetProductBuyCount(req.ProductId)
	config.SecLayerCtx.HistoryMapLock.Unlock()

	if historyCount >= product.OnePersonBuyLimit {
		res.Code = srv_err.ErrAlreadyBuy
		return
	}

	// 判断当前产品卖出数量是否大于产品总数
	curSoldCount := config.SecLayerCtx.ProductCountMgr.Count(req.ProductId)
	if curSoldCount >= product.Total {
		res.Code = srv_err.ErrSoldout
		product.Status = srv_err.ProductStatusSoldout
		return
	}

	// 购买频率限制
	//curRate := rand.Float64()
	curRate := 0.1
	fmt.Println(curRate, product.BuyRate)
	if curRate > product.BuyRate {
		res.Code = srv_err.ErrRetry
		return
	}

	userHistory.Add(req.ProductId, 1)
	config.SecLayerCtx.ProductCountMgr.Add(req.ProductId, 1)

	//用户Id、商品id、当前时间、密钥
	res.Code = srv_err.ErrSecKillSucc
	tokenData := fmt.Sprintf("userId=%d&productId=%d&timestamp=%d&security=%s", req.UserId, req.ProductId, nowTime, config.SecLayerCtx.SecLayerConf.TokenPassWd)
	res.Token = fmt.Sprintf("%x", md5.Sum([]byte(tokenData))) //MD5加密
	res.TokenTime = nowTime

	return
}
