package motor

import (
	"errors"
	"fmt"

	"github.com/sonr-io/sonr/pkg/motor/x/object"
	_ "golang.org/x/mobile/bind"
)

/*
#include <stdlib.h>
*/
import "C"

func AddListBool(n, f *C.char, v C.int) error {
	if instance == nil {
		return errWalletNotExists
	}

	builder, fieldName, err := findBuilder(n, f)
	if err != nil {
		return err
	}

	if !builder.Has(fieldName) {
		if err := builder.Set(fieldName, []bool{}); err != nil {
			return fmt.Errorf("error creating list: %s", err)
		}
	}

	var list []bool
	if l, ok := builder.Get(fieldName).([]bool); !ok {
		return fmt.Errorf("field '%s' is not of boolean type", fieldName)
	} else {
		list = l
	}

	var value bool
	if int(v) == 0 {
		value = false
	} else {
		value = true
	}
	list = append(list, value)

	return builder.Set(fieldName, list)
}

func AddListInt(n, f *C.char, v C.int) error {
	if instance == nil {
		return errWalletNotExists
	}

	builder, fieldName, err := findBuilder(n, f)
	if err != nil {
		return err
	}

	if !builder.Has(fieldName) {
		if err := builder.Set(fieldName, []int{}); err != nil {
			return fmt.Errorf("error creating list: %s", err)
		}
	}

	var list []int
	if l, ok := builder.Get(fieldName).([]int); !ok {
		return fmt.Errorf("field '%s' is not of int type", fieldName)
	} else {
		list = l
	}
	list = append(list, int(v))

	return builder.Set(fieldName, list)
}

func AddListFloat(n, f *C.char, v C.float) error {
	if instance == nil {
		return errWalletNotExists
	}

	builder, fieldName, err := findBuilder(n, f)
	if err != nil {
		return err
	}

	if !builder.Has(fieldName) {
		if err := builder.Set(fieldName, []float32{}); err != nil {
			return fmt.Errorf("error creating list: %s", err)
		}
	}

	var list []float32
	if l, ok := builder.Get(fieldName).([]float32); !ok {
		return fmt.Errorf("field '%s' is not of float type", fieldName)
	} else {
		list = l
	}
	list = append(list, float32(v))

	return builder.Set(fieldName, list)
}

func AddListString(n, f, v *C.char) error {
	if instance == nil {
		return errWalletNotExists
	}

	builder, fieldName, err := findBuilder(n, f)
	if err != nil {
		return err
	}

	if v == nil {
		return errors.New("value cannot be nil")
	}

	if !builder.Has(fieldName) {
		if err := builder.Set(fieldName, []string{}); err != nil {
			return fmt.Errorf("error creating list: %s", err)
		}
	}

	var list []string
	if l, ok := builder.Get(fieldName).([]string); !ok {
		return fmt.Errorf("field '%s' is not of string type", fieldName)
	} else {
		list = l
	}

	value := C.GoString(v)
	list = append(list, value)

	return builder.Set(fieldName, list)
}

func AddListBytes(n, f *C.char, v []byte) error {
	if instance == nil {
		return errWalletNotExists
	}

	builder, fieldName, err := findBuilder(n, f)
	if err != nil {
		return err
	}

	if v == nil {
		return errors.New("value cannot be nil")
	}

	if !builder.Has(fieldName) {
		if err := builder.Set(fieldName, [][]byte{}); err != nil {
			return fmt.Errorf("error creating list: %s", err)
		}
	}

	var list [][]byte
	if l, ok := builder.Get(fieldName).([][]byte); !ok {
		return fmt.Errorf("field '%s' is not of bytes type", fieldName)
	} else {
		list = l
	}

	list = append(list, v)

	return builder.Set(fieldName, list)
}

func findBuilder(n, f *C.char) (*object.ObjectBuilder, string, error) {
	if n == nil {
		return nil, "", errors.New("name cannot be nil")
	}
	if f == nil {
		return nil, "", errors.New("field name cannot be nil")
	}

	name := C.GoString(n)
	if builder, ok := objectBuilders[name]; !ok {
		return nil, "", fmt.Errorf("no object builder with name '%s'", name)
	} else {
		return builder, C.GoString(f), nil
	}
}
