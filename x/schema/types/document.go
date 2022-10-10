package types

import (
	"fmt"
)

func (d *DocumentValue) GetValue() interface{} {
	switch d.Kind {
	case Kind_BOOL:
		if d.BoolValue != nil {
			return d.BoolValue.Value
		}
	case Kind_BYTES:
		if d.BytesValue != nil {
			return d.BytesValue.Value
		}
	case Kind_INT:
		if d.IntValue != nil {
			return int64(d.IntValue.Value)
		}
	case Kind_FLOAT:
		if d.FloatValue != nil {
			return d.FloatValue.Value
		}
	case Kind_STRING:

		if d.StringValue != nil {
			return d.StringValue.Value
		}
	case Kind_LINK:
		if d.LinkValue != nil {
			return d.LinkValue.Value
		}
	case Kind_LIST:
		if d.ListValue != nil {
			listVal, err := resolveArrayValues(d.ListValue.Value)
			if err != nil {
				return nil
			}
			return listVal
		}
	default:
		return nil
	}
	return nil
}

func resolveArrayValues(vals []*DocumentValue) ([]interface{}, error) {
	arr := make([]interface{}, 0)
	for _, val := range vals {
		if val.BoolValue != nil {
			arr = append(arr, val.BoolValue.Value)
		} else if val.StringValue != nil {
			arr = append(arr, val.StringValue.Value)
		} else if val.IntValue != nil {
			arr = append(arr, val.IntValue.Value)
		} else if val.FloatValue != nil {
			arr = append(arr, val.FloatValue.Value)
		} else if val.BytesValue != nil {
			arr = append(arr, val.BytesValue.Value)
		} else if val.ListValue != nil {
			v, err := resolveArrayValues(val.ListValue.Value)
			if err != nil {
				return nil, err
			}
			arr = append(arr, v)
		} else if val.LinkValue != nil {
			arr = append(arr, val.LinkValue.Value)
		} else {
			return nil, fmt.Errorf("unknown list value: %s", val)
		}
	}

	return arr, nil
}
