package _slice

import "git.ziji.fun/junyang7/gf/tool/_as"

func Implode[K Slice](elementList []K, separator string) string {
	res := ""
	for i, j := 0, len(elementList); i < j; i++ {
		res += _as.String(elementList[i])
		if i < j-1 {
			res += separator
		}
	}
	return res
}
