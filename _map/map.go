package _map

func Keys[K comparable, V any](m map[K]V) (res []K) {
	res = make([]K, 0, len(m))
	for k := range m {
		res = append(res, k)
	}
	return
}

func Values[K comparable, V any](m map[K]V) (res []V) {
	res = make([]V, 0, len(m))
	for _, v := range m {
		res = append(res, v)
	}
	return
}
