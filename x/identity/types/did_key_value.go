package types

func MapToKeyValueList(m map[string]string) []*KeyValuePair {
	kvs := make([]*KeyValuePair, 0)
	for k, v := range m {
		kvs = append(kvs, &KeyValuePair{Key: k, Value: v})
	}
	return kvs
}

func KeyValueListToMap(kvs []*KeyValuePair) map[string]string {
	m := make(map[string]string)
	for _, kv := range kvs {
		m[kv.Key] = kv.Value
	}
	return m
}
