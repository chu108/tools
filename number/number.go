package number

import (
	"fmt"
	"github.com/chu108/tools/config"
	"math"
	"strconv"
	"time"
)

type Num struct {
}

func NewNum() *Num {
	return &Num{}
}

/**
num: 数字
retain：保留位数，精度
*/
func (*Num) FloatToInt64(num float64, retain int) int64 {
	return int64(num * math.Pow10(retain))
}

/**
num: 数字
retain：保留位数，精度
*/
func (*Num) Int64ToFloat64(num int64, retain int) float64 {
	return float64(num) / math.Pow10(retain)
}

//fload64保留两位小数
func (*Num) FloatDecimal(num float64) float64 {
	value, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", num), 64)
	return value
}

//int64转字符串
func (*Num) Int64ToStr(i int64) string {
	return strconv.FormatInt(i, 10)
}

//int转字符串
func (*Num) IntToStr(i int) string {
	return strconv.Itoa(i)
}

/*
float转字符串
'b' (-ddddp±ddd，二进制指数)
'e' (-d.dddde±dd，十进制指数)
'E' (-d.ddddE±dd，十进制指数)
'f' (-ddd.dddd，没有指数)
'g' ('e':大指数，'f':其它情况)
'G' ('E':大指数，'f':其它情况)
*/
func (*Num) FloatToStr(f float64) string {
	return strconv.FormatFloat(f, 'E', -1, 64)
}

//时间戳转日期字符串
func (*Num) Int64ToDateStr(i int64) string {
	return time.Unix(i, 0).Format(config.LayoutDate)
}

//时间戳转时间字符串
func (*Num) Int64ToTimeStr(i int64) string {
	return time.Unix(i, 0).Format(config.LayoutTime)
}

//INT最大值
func (*Num) MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//INT64最大值
func (*Num) MaxInt64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

//FLOAT最大值
func (*Num) MaxFloat(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

//INT最小值
func (*Num) MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

//INT64最小值
func (*Num) MinInt64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

//FLOAT最小值
func (*Num) MinFloat(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
