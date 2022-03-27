package _arr

import (
	"github.com/chu108/tools/_num"
	"github.com/chu108/tools/_str"
)

func StrToInt64(req []string) (res []int64) {
	length := len(req)
	res = make([]int64, 0, length)
	for i := 0; i < length; i++ {
		temp := _str.ToInt64(req[i])
		res = append(res, temp)
	}
	return
}

func StrToMap(req []string) (res map[string]bool) {
	length := len(req)
	res = make(map[string]bool, length)
	for i := 0; i < length; i++ {
		res[req[i]] = true
	}
	return
}

func StrToSql(req []string) (query, args []string) {
	length := len(req)
	query = make([]string, 0, length)
	args = make([]string, 0, length)
	for i := 0; i < length; i++ {
		query = append(query, "?")
		args = append(args, req[i])
	}
	return
}

func Int64ToStr(req []int64) (res []string) {
	length := len(req)
	res = make([]string, 0, length)
	for i := 0; i < length; i++ {
		res = append(res, _num.Int64ToStr(req[i]))
	}
	return
}

func Int64ToMap(req []int64) (res map[int64]bool) {
	length := len(req)
	res = make(map[int64]bool, length)
	for i := 0; i < length; i++ {
		res[req[i]] = true
	}
	return
}

func Int64ToSql(req []int64) (query []string, args []int64) {
	length := len(req)
	query = make([]string, 0, length)
	args = make([]int64, 0, length)
	for i := 0; i < length; i++ {
		query = append(query, "?")
		args = append(args, req[i])
	}
	return
}

func IntToStr(req []int) (res []string) {
	length := len(req)
	res = make([]string, 0, length)
	for i := 0; i < length; i++ {
		res = append(res, _num.IntToStr(req[i]))
	}
	return
}

func IntToMap(req []int) (res map[int]bool) {
	length := len(req)
	res = make(map[int]bool, length)
	for i := 0; i < length; i++ {
		res[req[i]] = true
	}
	return
}

func IntToSql(req []int) (query []string, args []int) {
	length := len(req)
	query = make([]string, 0, length)
	args = make([]int, 0, length)
	for i := 0; i < length; i++ {
		query = append(query, "?")
		args = append(args, req[i])
	}
	return
}

func InStrArr(val string, arr []string) bool {
	for i := 0; i < len(arr); i++ {
		if val == arr[i] {
			return true
		}
	}
	return false
}

func InInt64Arr(val int64, arr []int64) bool {
	for i := 0; i < len(arr); i++ {
		if val == arr[i] {
			return true
		}
	}
	return false
}

func InIntArr(val int, arr []int) bool {
	for i := 0; i < len(arr); i++ {
		if val == arr[i] {
			return true
		}
	}
	return false
}
