/**
    @Author:     ZonzeeLi
    @Project:    sk_proxy
    @CreateDate: 1/11/2022
    @UpdateDate: xxx
    @Note:       秒级限制
**/

package srv_limit

type Limit struct {
	secLimit TimeLimit
	minLimit TimeLimit
}

//秒限制
type SecLimit struct {
	count   int
	curTime int64
}

//一秒内访问的次数
func (p *SecLimit) Count(nowTime int64) (curCount int) {
	if p.curTime != nowTime {
		p.count = 1
		p.curTime = nowTime
		curCount = p.count
		return
	}

	p.count++
	curCount = p.count
	return
}

//检查用户访问的次数
func (p *SecLimit) Check(nowTime int64) int {
	if p.curTime != nowTime {
		return 0
	}
	return p.count
}
