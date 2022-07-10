package _slice

func Unique[K Slice](elementList []K) []K {
	var res []K
	filter := map[K]bool{}
	for _, element := range elementList {
		if _, ok := filter[element]; !ok {
			res = append(res, element)
			filter[element] = true
		}
	}
	return res
}
