/**
    @Author:     ZonzeeLi
    @Project:    sk_layer
    @CreateDate: 1/10/2022
    @UpdateDate: xxx
    @Note:       错误处理
**/

package srv_err

const (
	ErrServiceBusy = 1001 + iota
	ErrSecKillSucc
	ErrNotFoundProduct
	ErrSoldout
	ErrRetry
	ErrAlreadyBuy
)

const (
	ProductStatusSoldout = 2001 + iota
)
