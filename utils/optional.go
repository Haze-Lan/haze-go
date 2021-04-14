package utils

import "strings"

// ValueOfString 判断为空则使用默认值
func ValueOfString(data interface{}, defaultValue string) string {
	if data == nil {
		return defaultValue
	}
	if len(strings.TrimSpace(data.(string))) == 0 {
		return defaultValue
	}
	return data.(string)
}
func ValueOfInt(data interface{}, defaultValue uint64) uint64 {
	if data == nil {
		return defaultValue
	}
	return data.(uint64)
}
