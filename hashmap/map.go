package hashmap

type Map struct {
}

func NewMap() *Map {
	return &Map{}
}

func (*Map) IntBoolToKeys(req map[int]bool) (res []int) {
	length := len(req)
	res = make([]int, 0, length)
	for k := range req {
		res = append(res, k)
	}
	return
}

func (*Map) IntIntToKeys(req map[int]int) (res []int) {
	length := len(req)
	res = make([]int, 0, length)
	for k := range req {
		res = append(res, k)
	}
	return
}

func (*Map) IntStrToKeys(req map[int]string) (res []int) {
	length := len(req)
	res = make([]int, 0, length)
	for k := range req {
		res = append(res, k)
	}
	return
}

func (*Map) Int64BoolToKeys(req map[int64]bool) (res []int64) {
	length := len(req)
	res = make([]int64, 0, length)
	for k := range req {
		res = append(res, k)
	}
	return
}

func (*Map) Int64Int64ToKeys(req map[int64]int64) (res []int64) {
	length := len(req)
	res = make([]int64, 0, length)
	for k := range req {
		res = append(res, k)
	}
	return
}

func (*Map) Int64StrToKeys(req map[int64]string) (res []int64) {
	length := len(req)
	res = make([]int64, 0, length)
	for k := range req {
		res = append(res, k)
	}
	return
}

func (*Map) StrBoolToKeys(req map[string]bool) (res []string) {
	length := len(req)
	res = make([]string, 0, length)
	for k := range req {
		res = append(res, k)
	}
	return
}

func (*Map) StrStrToKeys(req map[string]string) (res []string) {
	length := len(req)
	res = make([]string, 0, length)
	for k := range req {
		res = append(res, k)
	}
	return
}

func (*Map) StrInt64ToKeys(req map[string]int64) (res []string) {
	length := len(req)
	res = make([]string, 0, length)
	for k := range req {
		res = append(res, k)
	}
	return
}

func (*Map) StrIntToKeys(req map[string]int) (res []string) {
	length := len(req)
	res = make([]string, 0, length)
	for k := range req {
		res = append(res, k)
	}
	return
}

func (*Map) IntIntToValues(req map[int]int) (res []int) {
	length := len(req)
	res = make([]int, 0, length)
	for _, v := range req {
		res = append(res, v)
	}
	return
}

func (*Map) IntStrToValues(req map[int]string) (res []string) {
	length := len(req)
	res = make([]string, 0, length)
	for _, v := range req {
		res = append(res, v)
	}
	return
}

func (*Map) Int64Int64ToValues(req map[int64]int64) (res []int64) {
	length := len(req)
	res = make([]int64, 0, length)
	for _, v := range req {
		res = append(res, v)
	}
	return
}

func (*Map) Int64StrToValues(req map[int64]string) (res []string) {
	length := len(req)
	res = make([]string, 0, length)
	for _, v := range req {
		res = append(res, v)
	}
	return
}

func (*Map) StrStrToValues(req map[string]string) (res []string) {
	length := len(req)
	res = make([]string, 0, length)
	for _, v := range req {
		res = append(res, v)
	}
	return
}

func (*Map) StrInt64ToValues(req map[string]int64) (res []int64) {
	length := len(req)
	res = make([]int64, 0, length)
	for _, v := range req {
		res = append(res, v)
	}
	return
}

func (*Map) StrIntToValues(req map[string]int) (res []int) {
	length := len(req)
	res = make([]int, 0, length)
	for _, v := range req {
		res = append(res, v)
	}
	return
}
