package hexutil

import "strconv"

func ConvertToUint64(v string) uint64 {
	i, _ := strconv.ParseUint(v, 10, 64)
	return i
}
