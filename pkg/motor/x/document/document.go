package document

type Document map[string]interface{}

func (o Document) GetFieldNames() []string {
	keys := make([]string, 0)
	for k := range o {
		keys = append(keys, k)
	}
	return keys
}

func (o Document) GetValues() []interface{} {
	values := make([]interface{}, 0)
	for _, v := range o {
		values = append(values, v)
	}
	return values
}
