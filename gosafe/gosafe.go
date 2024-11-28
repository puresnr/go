// package gosafe 提供一个安全的方式来开启协程, 当 panic 时, 不会导致进程退出, 调用方式: 直接把待执行函数传进去即可, 比如Go(f(){})
package gosafe

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"os"
	"runtime"
)

func PrintPanicStack(extras ...interface{}) {
	i := 0
	funcName, file, line, ok := runtime.Caller(i)
	for ok {
		fmt.Fprintf(os.Stderr, "frame %v:[func:%v,file:%v,line:%v]\n", i, runtime.FuncForPC(funcName).Name(), file, line)
		i++
		funcName, file, line, ok = runtime.Caller(i)
	}
	for k := range extras {
		fmt.Fprintf(os.Stderr, "EXTRAS#%v DATA:%v\n", k, spew.Sdump(extras[k]))
	}
}

func onExit(extras ...interface{}) {
	if err := recover(); err != nil {
		PrintPanicStack(extras)
		fmt.Fprintf(os.Stderr, "<recover from panic: %s>\n", err)
	}
}

// Go 用于执行无入参的函数, panic 时不会开启新的协程重新执行该函数
// 入参:
//
//	f: 待执行函数
func Go(f func()) {
	go func() {
		defer onExit()

		f()
	}()
}

// GoP 用于执行有入参的函数, panic 时不会开启新的协程重新执行该函数
// 入参:
//
//	f: 待执行函数
//	p: 待执行函数使用的入参, 如果入参数量多于 1 个, 需要把所有入参包装为一个结构体, 使用该结构体作为入参
func GoP[T any](f func(T), p T) {
	go func() {
		defer onExit()

		f(p)
	}()
}

func onExitR(f func(), extras ...interface{}) {
	if err := recover(); err != nil {
		PrintPanicStack(extras)
		fmt.Fprintf(os.Stderr, "<recover from panic: %s>\n", err)

		GoR(f)
	}
}

// GoR 用于执行无入参的函数, panic 时会开启新的协程重新执行该函数
// 入参:
//
//	f: 待执行函数
func GoR(f func()) {
	go func() {
		defer onExitR(f)

		f()
	}()
}

func onExitPR[T any](f func(T), p T, extras ...interface{}) {
	if err := recover(); err != nil {
		PrintPanicStack(extras)
		fmt.Fprintf(os.Stderr, "<recover from panic: %s>\n", err)

		GoPR(f, p)
	}
}

// GoPR 用于执行有入参的函数, panic 时会开启新的协程重新执行该函数
// 入参:
//
//	f: 待执行函数
//	p: 待执行函数使用的入参, 如果入参数量多于 1 个, 需要把所有入参包装为一个结构体, 使用该结构体作为入参
func GoPR[T any](f func(T), p T) {
	go func() {
		defer onExitPR(f, p)

		f(p)
	}()
}
