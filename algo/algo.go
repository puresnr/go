package algo

import (
	"golang.org/x/exp/constraints"
)

// 查找元素位置时, 找不到该元素的返回值. 在 C++ 里, 这个类型是 size_t, 是个无符号数, 但是 GO 里 len() 返回的是 int, 所以这里如果返回无符号数,
// 实际使用时就会特别麻烦, 需要不停的强转, 所以只好也定义为 int 了. 然后因为是 int 了, 索性就定义为一个不可能的下标, 即负数, 这样实际使用时,
// 直接和这个值进行比较即可, 而不用模仿 C++ 里的 std::find 行为(在函数找不到元素时, 返回指向范围末尾的迭代器), 返回 slice 的长度,
// 毕竟在有代码提示的情况下, aslice.InvalidIdx 怎么着也比 len(变量名) 方便输入, 而且虽然 len() 的开销几乎可以忽略, 但是也不是完全没有开销.
const InvalidIdx = -1

func Find_idx[T comparable](sli []T, t T) int {
	for i := range sli {
		if sli[i] == t {
			return i
		}
	}

	return InvalidIdx
}

func FindIf_idx[T any](sli []T, compare func(v T) bool) int {
	for i := range sli {
		if compare(sli[i]) {
			return i
		}
	}

	return InvalidIdx
}

func Find[T comparable](sli []T, t T) *T {
	for i := range sli {
		if sli[i] == t {
			return &sli[i]
		}
	}

	return nil
}

func FindIf[T any](sli []T, compare func(v T) bool) *T {
	for i := range sli {
		if compare(sli[i]) {
			return &sli[i]
		}
	}

	return nil
}

func Erase[T comparable](sli *[]T, t T) {
	if idx := Find_idx(*sli, t); idx != InvalidIdx {
		*sli = append((*sli)[:idx], (*sli)[idx+1:]...)
	}
}

func EraseIf[T any](sli *[]T, compare func(v T) bool) {
	if idx := FindIf_idx(*sli, compare); idx != InvalidIdx {
		*sli = append((*sli)[:idx], (*sli)[idx+1:]...)
	}
}

func Exist[T comparable](sli []T, t T) bool {
	return Find(sli, t) != nil
}

func ExistIf[T any](sli []T, compare func(v T) bool) bool {
	return FindIf(sli, compare) != nil
}

// Reverse reverses the order of the elements in sli
func Reverse[T any](sli []T) {
	for i, j := 0, len(sli)-1; i != len(sli)/2; func() { i++; j-- }() {
		sli[i], sli[j] = sli[j], sli[i]
	}
}

// ReverseCopy copies the elements from sli to another slice in reverse order
// return: new slice in reverse order
func ReverseCopy[T any](sli []T) []T {
	rsli := make([]T, len(sli))
	for i := 0; i != len(sli); i++ {
		rsli[len(sli)-1-i] = sli[i]
	}
	return rsli
}

// Min 返回两个值中较小的一个，相同时，返回第一个
func Min[T constraints.Ordered](v1, v2 T) T {
	if v1 > v2 {
		return v2
	}

	return v1
}

// Max 返回两个值中较大的一个，相同时，返回第一个
func Max[T constraints.Ordered](v1, v2 T) T {
	if v1 < v2 {
		return v2
	}

	return v1
}

// Pick 根据 compare 的结果返回元素, 如果是 true 返回前者, 否则返回后者
// todo: Pick 无法实现短路特性, 传参总是会被计算
func Pick[T any](t1 T, t2 T, compare func() bool) T {
	if compare() {
		return t1
	}
	return t2
}

func Empty_slice[T any](sli []T) bool {
	return len(sli) == 0
}

func Empty_map[TK comparable, TV any](m map[TK]TV) bool {
	return len(m) == 0
}

func Empty_str(str string) bool {
	return len(str) == 0
}

func Empty_chan[T any](c chan T) bool {
	return len(c) == 0
}
