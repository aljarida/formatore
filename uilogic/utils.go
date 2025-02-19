package uilogic

func has(arr []string, s string) bool {
	for _, t := range arr {
		if s == t {
			return true
		}
	}
	return false
}

func isQuit(s string) bool {
	return has(quitTokens, s)
}

func isDone(s string) bool {
	return has(doneTokens, s)
}
