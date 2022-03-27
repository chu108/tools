package _map

func IntBoolToKeys(req map[int]bool) (res []int) {
	length := len(req)
	res = make([]int, 0, length)
	for k := range req {
		res = append(res, k)
	}
	return
}

func IntIntToKeys(req map[int]int) (res []int) {
	length := len(req)
	res = make([]int, 0, length)
	for k := range req {
		res = append(res, k)
	}
	return
}

func IntStrToKeys(req map[int]string) (res []int) {
	length := len(req)
	res = make([]int, 0, length)
	for k := range req {
		res = append(res, k)
	}
	return
}

func Int64BoolToKeys(req map[int64]bool) (res []int64) {
	length := len(req)
	res = make([]int64, 0, length)
	for k := range req {
		res = append(res, k)
	}
	return
}

func Int64Int64ToKeys(req map[int64]int64) (res []int64) {
	length := len(req)
	res = make([]int64, 0, length)
	for k := range req {
		res = append(res, k)
	}
	return
}

func Int64StrToKeys(req map[int64]string) (res []int64) {
	length := len(req)
	res = make([]int64, 0, length)
	for k := range req {
		res = append(res, k)
	}
	return
}

func StrBoolToKeys(req map[string]bool) (res []string) {
	length := len(req)
	res = make([]string, 0, length)
	for k := range req {
		res = append(res, k)
	}
	return
}

func StrStrToKeys(req map[string]string) (res []string) {
	length := len(req)
	res = make([]string, 0, length)
	for k := range req {
		res = append(res, k)
	}
	return
}

func StrInt64ToKeys(req map[string]int64) (res []string) {
	length := len(req)
	res = make([]string, 0, length)
	for k := range req {
		res = append(res, k)
	}
	return
}

func StrIntToKeys(req map[string]int) (res []string) {
	length := len(req)
	res = make([]string, 0, length)
	for k := range req {
		res = append(res, k)
	}
	return
}

func IntIntToValues(req map[int]int) (res []int) {
	length := len(req)
	res = make([]int, 0, length)
	for _, v := range req {
		res = append(res, v)
	}
	return
}

func IntStrToValues(req map[int]string) (res []string) {
	length := len(req)
	res = make([]string, 0, length)
	for _, v := range req {
		res = append(res, v)
	}
	return
}

func Int64Int64ToValues(req map[int64]int64) (res []int64) {
	length := len(req)
	res = make([]int64, 0, length)
	for _, v := range req {
		res = append(res, v)
	}
	return
}

func Int64StrToValues(req map[int64]string) (res []string) {
	length := len(req)
	res = make([]string, 0, length)
	for _, v := range req {
		res = append(res, v)
	}
	return
}

func StrStrToValues(req map[string]string) (res []string) {
	length := len(req)
	res = make([]string, 0, length)
	for _, v := range req {
		res = append(res, v)
	}
	return
}

func StrInt64ToValues(req map[string]int64) (res []int64) {
	length := len(req)
	res = make([]int64, 0, length)
	for _, v := range req {
		res = append(res, v)
	}
	return
}

func StrIntToValues(req map[string]int) (res []int) {
	length := len(req)
	res = make([]int, 0, length)
	for _, v := range req {
		res = append(res, v)
	}
	return
}
