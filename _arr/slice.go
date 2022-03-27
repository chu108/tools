package _arr

import (
	"fmt"
	"strconv"
)

type Int interface {
	int | int32 | int64 | uint | uint32 | uint64
}

type Float interface {
	float32 | float64
}

//ToMap 转Map
func ToMap[T any](s []T) (res map[T]bool) {
	res = make(map[T]bool, 0)
	for i := 0; i < len(s); i++ {
		res[s[i]] = true
	}
	return
}

//Distinct 数组去重
func Distinct[T comparable](arr []T) (res []T) {
	temp := make(map[T]struct{}, 0)
	res = make([]T, 0, len(arr))
	for i := 0; i < len(arr); i++ {
		if _, ok := temp[arr[i]]; !ok {
			temp[arr[i]] = struct{}{}
			res = append(res, arr[i])
		}
	}
	return
}

//Exists 判断某个是否存在
func Exists[T comparable](arr []T, val T) bool {
	l := len(arr)
	if l == 0 {
		return false
	}
	for i := 0; i < l; i++ {
		if arr[i] == val {
			return true
		}
	}
	return false
}

//In 判断多个值是否存在，返回已存在的值数组
func In[T comparable](arr []T, vales ...T) (res []T) {
	al, vl := len(arr), len(vales)
	res = make([]T, 0, vl)
	if al == 0 || vl == 0 {
		return
	}
	valMap := make(map[T]struct{}, 0)
	for i := 0; i < vl; i++ {
		valMap[vales[i]] = struct{}{}
	}
	for i := 0; i < al; i++ {
		if _, ok := valMap[arr[i]]; ok {
			res = append(res, arr[i])
		}
	}
	return
}

// StrToInt 字符串数组转int数组
func StrToInt(str []string) (res []int, err error) {
	l := len(str)
	res = make([]int, 0, l)
	for i := 0; i < l; i++ {
		tmp, err := strconv.Atoi(str[i])
		if err != nil {
			return nil, err
		}
		res = append(res, tmp)
	}
	return
}

// IntToStr int数组转字符串数组
func IntToStr[T Int](str []T) (res []string, err error) {
	l := len(str)
	res = make([]string, 0, l)
	for i := 0; i < l; i++ {
		res = append(res, fmt.Sprintf("%d", str[i]))
	}
	return
}
