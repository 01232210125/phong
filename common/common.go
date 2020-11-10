package common

func CheckIsExist(lst []string, s string) bool {
	for _, v := range lst {
		if s == v {
			return true
		}
	}
	return false
}
