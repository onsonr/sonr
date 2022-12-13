package types

import "strings"

func (kt KeyType) FormatString() string {
	str := kt.String()
	ptrs := strings.Split(str, "_")
	result := ""
	for i, ptr := range ptrs {
		if i > 0 {
			result += strings.Title(strings.ToLower(ptr))
		}
	}
	return result
}

func (kt ProofType) FormatString() string {
	str := kt.String()
	ptrs := strings.Split(str, "_")
	result := ""
	for i, ptr := range ptrs {
		if i > 0 {
			result += strings.Title(strings.ToLower(ptr))
		}
	}
	return result
}

func (kt ServiceType) FormatString() string {
	str := kt.String()
	ptrs := strings.Split(str, "_")
	result := ""
	for i, ptr := range ptrs {
		if i > 0 {
			result += strings.Title(strings.ToLower(ptr))
		}
	}
	return result
}
