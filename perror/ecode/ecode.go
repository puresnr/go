package ecode

import "errors"

// Ecode represents an error type identified by a string code.
// It is designed to work with the standard errors.Is function.
// Two Ecode instances are considered equal by errors.Is if they are both of type Ecode
// and were created with the same code string.
// Ecode is intended to be a source of errors, not a wrapper around other errors,
// so it does not implement the Unwrap method.
type Ecode struct{ error }

func New(ecode string) *Ecode { return &Ecode{error: errors.New(ecode)} }

// Is 实现了 errors.Is 的接口。
// 如果目标错误也是 Ecode 类型且 code 相同，则返回 true。
func (e *Ecode) Is(target error) bool {
	// 判断 target 是否为 *Ecode 类型
	et, ok := target.(*Ecode)
	if !ok {
		// 如果 target 不是 *Ecode 类型，则它们不相等
		return false
	}

	// 此时 target 是 *Ecode 类型 (或一个 nil 的 *Ecode 指针)。
	// 必须处理指针为 nil 的情况，否则会 panic。
	if e == nil || et == nil {
		return e == et // 只有当两者都为 nil 时才相等
	}

	// 两个指针都不是 nil，可以安全地比较 code。
	return e.Error() == et.Error()
}