package uilogic

func has(arr []string, s string) bool {
	for _, t := range arr {
		if s == t {
			return true
		}
	}
	return false
}


func isCancel(s string) bool {
	return has(cancelTokens, s)
}

func isQuit(s string) bool {
	return has(quitTokens, s)
}
