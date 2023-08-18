package utils

func IncludeString(arr []string, obj string) bool {
	for _, item := range arr {
		if item == obj {
			return true
		}
	}
	return false
}
