package functional

func ContainsString(strings []string, candidate string) bool {
	for _, s := range strings {
		if s == candidate {
			return true
		}
	}
	return false
}
