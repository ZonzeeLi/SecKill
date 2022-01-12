/**
    @Author:     ZonzeeLi
    @Project:    sk_proxy
    @CreateDate: 1/11/2022
    @UpdateDate: xxx
    @Note:       错误处理
**/

package srv_err

import "errors"

const (
	ErrInvalidRequest = 1001 + iota
	ErrNotFoundProductId
	ErrUserCheckAuthFailed
	ErrUserServiceBusy
	ErrActiveNotStart
	ErrActiveAlreadyEnd
	ErrActiveSaleOut
	ErrProcessTimeout
	ErrClientClosed
)

const (
	ErrServiceBusy = 1001 + iota
	ErrSecKillSucc
	ErrNotFoundProduct
	ErrSoldout
	ErrRetry
	ErrAlreadyBuy
)

var errMsg = map[int]string{
	ErrServiceBusy:     "服务器错误",
	ErrSecKillSucc:     "抢购成功",
	ErrNotFoundProduct: "没有该商品",
	ErrSoldout:         "商品售罄",
	ErrRetry:           "请重试",
	ErrAlreadyBuy:      "已经抢购",
}

func GetErrMsg(code int) error {
	return errors.New(errMsg[code])
}
