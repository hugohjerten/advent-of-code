package utils

func ContainsInt(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func ContainsOnly(ss string, r rune) bool {
	for _, e := range ss {
		if e != r {
			return false
		}
	}
	return true
}
