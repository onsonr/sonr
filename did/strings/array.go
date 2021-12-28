package strings

func IndexOf(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}

	return -1
}

func Contains(vs []string, t string) bool {
	return IndexOf(vs, t) >= 0
}

func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func Complement(vs []string, ts []string) []string {
	return Filter(vs, func(s string) bool {
		return !Contains(ts, s)
	})
}
