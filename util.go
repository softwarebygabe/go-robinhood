package robinhood

func containsString(l []string, s string) bool {
	for _, si := range l {
		if s == si {
			return true
		}
	}
	return false
}
