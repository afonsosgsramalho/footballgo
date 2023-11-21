package utils

func IndexOf(arr []string, val string) int {
	for pos, v := range arr {
		if v == val {
			return pos
		}
	}
	return -1
}
