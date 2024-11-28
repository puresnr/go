package gosync

import (
	"github.com/puresnr/pgo/gosafe"
	"sync"
	"time"
)

// GoWait 用于并发的执行一组函数, 并等待所有函数执行完毕
func GoWait(funcs ...func()) {
	var wg sync.WaitGroup

	for _, f := range funcs {
		wg.Add(1)

		gosafe.GoP(func(ef func()) {
			defer wg.Done()

			ef()
		}, f)
	}

	wg.Wait()
}

// GoWaitWithTimeout 函数用于等待一组函数执行完成或超时。
//
// 参数：
//
//	timeout：超时时间，单位为秒。
//	funcs：需要等待执行完成的函数列表。
//
// 返回值：
//
//	如果等待超时则返回 true，否则返回 false。
//
// 说明：
//
//	函数通过 select 语句同时监听超时事件和函数执行完成事件。
//	如果超时事件发生，则返回 true；如果所有函数执行完成事件先发生，则返回 false。
func GoWaitWithTimeout(timeout uint, funcs ...func()) bool {
	select {
	case <-time.After(time.Duration(timeout) * time.Second):
		return true
	case <-func() <-chan struct{} {
		dc := make(chan struct{})
		go func() { GoWait(funcs...); close(dc) }()
		return dc
	}():
	}

	return false
}
