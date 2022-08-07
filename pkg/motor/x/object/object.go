package object

type Object map[string]interface{}

func (o Object) GetFieldNames() []string {
	keys := make([]string, 0)
	for k := range o {
		keys = append(keys, k)
	}
	return keys
}

func (o Object) GetValues() []interface{} {
	values := make([]interface{}, 0)
	for _, v := range o {
		values = append(values, v)
	}
	return values
}
