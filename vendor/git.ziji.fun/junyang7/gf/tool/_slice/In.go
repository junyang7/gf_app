package _slice

func In[K Slice](element K, elementList []K) bool {
	for _, v := range elementList {
		if v == element {
			return true
		}
	}
	return false
}
