package _slice

func Group[K Slice, V Slice](mapList []map[K]V, key K) map[V]map[K]V {
	res := map[V]map[K]V{}
	for _, v := range mapList {
		res[v[key]] = v
	}
	return res
}
