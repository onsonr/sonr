package marshal

import "encoding/json"

// NormalizeDocument accepts a JSON document and applies (in order) the given normalizers to it.
func NormalizeDocument(document []byte, normalizers ...Normalizer) ([]byte, error) {
	tmp := make(map[string]interface{}, 0)
	if err := json.Unmarshal(document, &tmp); err != nil {
		return nil, err
	}
	for _, normalizer := range normalizers {
		normalizer(tmp)
	}
	return json.Marshal(tmp)
}

type Normalizer func(map[string]interface{})

// KeyAlias returns a Normalizer that converts an aliased key to its original form. E.g. when working with
// LinkedData in JSON form, `@context` is an alias for `context`. This Normalizer would convert the `@context` key
// to `context`.
func KeyAlias(alias string, aliasFor string) Normalizer {
	return func(m map[string]interface{}) {
		for k, v := range m {
			if k == alias {
				m[aliasFor] = v
				delete(m, k)
			}
		}
	}
}

// Plural returns a Normalizer that converts a singular values (string/numeric/bool/object) to an array.
// This makes unmarshalling DID Documents or Verifiable Credentials easier, since those formats allow certain properties
// to be either a singular value or an array of values.
//
// Example input: 												{"message": "Hello, World"}
// Example output (if 'message' is supplied in 'pluralKeys'): 	{"message": ["Hello, World"]}
//
// This function does not support nested keys.
func Plural(key string) Normalizer {
	return func(m map[string]interface{}) {
		if _, isSlice := m[key].([]interface{}); m[key] != nil && !isSlice {
			m[key] = []interface{}{m[key]}
		}
	}
}

// Unplural returns a Normalizer that converts arrays with a single value into a singular value. It is the opposite
// of the Plural normalizer.
func Unplural(key string) Normalizer {
	return func(m map[string]interface{}) {
		if arr, _ := m[key].([]interface{}); len(arr) == 1 {
			m[key] = arr[0]
		}
	}
}

// PluralValueOrMap returns a Normalizer that behaves like Plural but leaves maps as simply a map. In other words,
// it only turns singular values into an array, except maps.
func PluralValueOrMap(key string) Normalizer {
	return func(m map[string]interface{}) {
		value := m[key]
		if value == nil {
			return
		} else if _, isMap := value.(map[string]interface{}); isMap {
			return
		} else if _, isSlice := value.([]interface{}); !isSlice {
			m[key] = []interface{}{m[key]}
		}
	}
}
