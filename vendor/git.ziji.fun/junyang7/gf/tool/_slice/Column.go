package _slice

func Column[K Slice, V Slice](mapList []map[K]V, key K) []V {
	var res []V
	for _, v := range mapList {
		res = append(res, v[key])
	}
	return res
}
