package ptime

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

// CalcDiffDay 计算两个日期之间的自然天数差
//
// 参数:
// - d1: string 类型，表示第一个日期字符串，格式为 "2006-01-02"。
// - d2: string 类型，表示第二个日期字符串，格式为 "2006-01-02"。
//
// 返回值:
// - int 类型，表示 d2 和 d1 日期之间的天数差，如果 d2 在 d1 之前，则返回负数。
// - error 类型，如果日期字符串格式不正确，则返回错误信息。
//
// 说明:
// - 使用 time.Parse 函数将日期字符串解析为 time.Time 类型。
// - 使用 time.Time 的 Sub 方法计算两个时间之间的差值，并将其转换为小时数。
// - 将小时数除以 24 得到天数差，并返回结果。
// 直接把日期转换成 utc 时间, 那么一天一定是 24 小时, 这样就避免了时区或者夏令时(一天可能有 23 或者  25 小时)带来的影响了
func CalcDiffDay(d1, d2 string) (int, error) {
	t1, err := time.Parse("2006-01-02", d1)
	if err != nil {
		return 0, err
	}

	t2, err := time.Parse("2006-01-02", d2)
	if err != nil {
		return 0, err
	}

	return int(t2.Sub(t1).Hours()) / 24, nil
}

var (
	monthSumDays = map[bool][]uint{
		false: {0, 31, 59, 90, 120, 151, 181, 212, 243, 273, 304, 334},
		true:  {0, 31, 60, 91, 121, 152, 182, 213, 244, 274, 305, 335},
	}
	monthDays = map[bool][]uint{
		false: {31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31},
		true:  {31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31},
	}
)

// IsLeapYear 函数用于判断给定的年份是否为闰年。
//
// 参数:
// - year: uint 类型，表示要判断的年份。
//
// 返回值:
// - bool 类型，如果年份是闰年则返回 true，否则返回 false。
//
// 说明:
// - 闰年的判断依据是：能被4整除但不能被100整除的年份是闰年，或者能被400整除的年份也是闰年。
func IsLeapYear(year uint) bool { return (year%4 == 0 && year%100 != 0) || (year%400 == 0) }

// CountLeapYears 计算两个年份之间的闰年数量（包括起始年，不包括结束年）
//
// 参数:
// - syear: uint 类型，表示起始年份。
// - eyear: uint 类型，表示结束年份。
//
// 返回值:
// - uint 类型，表示闰年的数量。
// - error 类型，如果起始年份不小于结束年份或起始年份小于1，则返回错误信息。
//
// 注意事项:
// - 函数使用整数除法计算闰年数量，基于闰年的定义（能被4整除但不能被100整除的年份是闰年，除非它也能被400整除）。
func CountLeapYears(syear, eyear uint) (uint, error) {
	if syear >= eyear || syear < 1 {
		return 0, errors.New("invalid year range")
	}

	ts, te := syear-1, eyear-1

	return (te/4 - ts/4) - (te/100 - ts/100) + (te/400 - ts/400), nil
}

// YearDay 函数用于计算给定日期在一年中的天数（从年初开始计算）。
//
// 参数:
// - md: string 类型，表示月份和日期的字符串，格式为 "MM-DD"，其中 MM 和 DD 都是两位数。
// - isleap: bool 类型，表示年份是否为闰年。
//
// 返回值:
// - uint 类型，表示给定日期在一年中的天数。
// - error 类型，如果日期格式不正确、月份或日期无效，则返回错误信息。
//
// 说明:
// - 函数首先检查日期字符串的格式是否正确，即是否为 "MM-DD" 格式。
// - 然后，检查月份和日期是否有效，月份必须在 01 到 12 之间，日期必须在 01 到对应月份的天数之间（闰年二月为 29 天，非闰年为 28 天）。
// - 如果月份或日期以零开头，会去掉零。
// - 最后，根据月份和日期计算该日期在一年中的天数，并返回结果。
func YearDay(md string, isleap bool) (uint, error) {
	mds := strings.Split(md, "-")
	if len(mds) != 2 {
		return 0, errors.New("invalid date format")
	}

	m, d := mds[0], mds[1]
	if len(m) != 2 || len(d) != 2 {
		return 0, errors.New("invalid date format")
	}

	if m[0] == '0' {
		m = m[1:]
	}
	if d[0] == '0' {
		d = d[1:]
	}

	im, _ := strconv.ParseUint(m, 10, 64)
	if im == 0 || im > 12 {
		return 0, errors.New("invalid month")
	}
	id, _ := strconv.ParseUint(d, 10, 64)
	if id == 0 || id > uint64(monthDays[isleap][im-1]) {
		return 0, errors.New("invalid date")
	}

	return monthSumDays[isleap][im-1] + uint(id), nil
}
