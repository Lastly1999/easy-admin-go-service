package helps

// IsContainInArray 验证字符串是否存在于数组中
func IsContainInArray(item string, items []string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}
