package util

// IndexOf returns the index of the first instance of a value in a slice
func IndexOf(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}

	return -1
}

// Contains returns true if the string is in the slice
func Contains(vs []string, t string) bool {
	return IndexOf(vs, t) >= 0
}

// Filter returns a new slice containing all strings from the slice that satisfy the predicate
func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// Complement returns a new slice containing all strings from the slice that do not satisfy the predicate
func Complement(vs []string, ts []string) []string {
	return Filter(vs, func(s string) bool {
		return !Contains(ts, s)
	})
}
