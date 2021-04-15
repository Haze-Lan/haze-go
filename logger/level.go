package logger

var levelMap = make(map[string]int, 5)

func init() {
	levelMap = make(map[string]int, 5)
	levelMap["debug"] = 1
	levelMap["info"] = 2
	levelMap["warn"] = 3
	levelMap["error"] = 4
	levelMap["fatal"] = 5
}
