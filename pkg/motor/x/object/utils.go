package object

func arrayContains(arr []string, val interface{}) bool {

	for _, v := range arr {
		if v == val {
			return true
		}
	}

	return false
}
