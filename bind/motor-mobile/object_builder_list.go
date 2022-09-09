package motor

import (
	"errors"
	"fmt"

	ct "github.com/sonr-io/sonr/third_party/types/common"
	_ "golang.org/x/mobile/bind"
)

func AddListBool(name, fieldName string, value bool) error {
	if instance == nil {
		return ct.ErrMotorWalletNotInitialized
	}

	builder, ok := objectBuilders[name]
	if !ok {
		return fmt.Errorf("no object builder with name '%s'", name)
	}

	if !builder.Has(fieldName) {
		if err := builder.Set(fieldName, []bool{}); err != nil {
			return fmt.Errorf("error creating list: %s", err)
		}
	}

	list, ok := builder.Get(fieldName).([]bool)
	if !ok {
		return fmt.Errorf("field '%s' is not of boolean type", fieldName)
	}
	list = append(list, value)

	return builder.Set(fieldName, list)
}

func AddListInt(name, fieldName string, value int) error {
	if instance == nil {
		return ct.ErrMotorWalletNotInitialized
	}

	builder, ok := objectBuilders[name]
	if !ok {
		return fmt.Errorf("no object builder with name '%s'", name)
	}

	if !builder.Has(fieldName) {
		if err := builder.Set(fieldName, []int{}); err != nil {
			return fmt.Errorf("error creating list: %s", err)
		}
	}

	list, ok := builder.Get(fieldName).([]int)
	if !ok {
		return fmt.Errorf("field '%s' is not of int type", fieldName)
	}
	list = append(list, value)

	return builder.Set(fieldName, list)
}

func AddListFloat(name, fieldName string, value float32) error {
	if instance == nil {
		return ct.ErrMotorWalletNotInitialized
	}

	builder, ok := objectBuilders[name]
	if !ok {
		return fmt.Errorf("no object builder with name '%s'", name)
	}

	if !builder.Has(fieldName) {
		if err := builder.Set(fieldName, []float32{}); err != nil {
			return fmt.Errorf("error creating list: %s", err)
		}
	}

	list, ok := builder.Get(fieldName).([]float32)
	if !ok {
		return fmt.Errorf("field '%s' is not of float type", fieldName)
	}
	list = append(list, value)

	return builder.Set(fieldName, list)
}

func AddListString(name, fieldName, value string) error {
	if instance == nil {
		return ct.ErrMotorWalletNotInitialized
	}

	builder, ok := objectBuilders[name]
	if !ok {
		return fmt.Errorf("no object builder with name '%s'", name)
	}

	if value == "" {
		return errors.New("value cannot be empty")
	}

	if !builder.Has(fieldName) {
		if err := builder.Set(fieldName, []string{}); err != nil {
			return fmt.Errorf("error creating list: %s", err)
		}
	}

	list, ok := builder.Get(fieldName).([]string)
	if !ok {
		return fmt.Errorf("field '%s' is not of string type", fieldName)
	}
	list = append(list, value)

	return builder.Set(fieldName, list)
}

func AddListBytes(name, fieldName string, value []byte) error {
	if instance == nil {
		return ct.ErrMotorWalletNotInitialized
	}

	builder, ok := objectBuilders[name]
	if !ok {
		return fmt.Errorf("no object builder with name '%s'", name)
	}

	if value == nil {
		return errors.New("value cannot be nil")
	}

	if !builder.Has(fieldName) {
		if err := builder.Set(fieldName, [][]byte{}); err != nil {
			return fmt.Errorf("error creating list: %s", err)
		}
	}

	list, ok := builder.Get(fieldName).([][]byte)
	if !ok {
		return fmt.Errorf("field '%s' is not of bytes type", fieldName)
	}
	list = append(list, value)

	return builder.Set(fieldName, list)
}

func RemoveListItem(name, fieldName string, index int) error {
	if instance == nil {
		return ct.ErrMotorWalletNotInitialized
	}

	builder, ok := objectBuilders[name]
	if !ok {
		return fmt.Errorf("no object builder with name '%s'", name)
	}

	list, ok := builder.Get(fieldName).([]interface{})
	if !ok || list == nil {
		return fmt.Errorf("no list field with name '%s'", fieldName)
	}

	if index < 0 || index >= len(list) {
		return fmt.Errorf("index %d of of range %d", index, len(list))
	}

	list = append(list[:index], list[index+1:])
	return builder.Set(fieldName, list)
}
