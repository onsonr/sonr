package prompt

type Validator func(string) bool

func AllowAll(s string) bool {
	return true
}

func NonEmpty(s string) bool {
	return s != ""
}

func And(fns ...Validator) Validator {
	return func(s string) bool {
		valid := true
		for _, fn := range fns {
			valid = valid && fn(s)
		}
		return valid
	}
}
