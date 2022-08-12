package motor

import (
	"errors"
	"fmt"

	_ "golang.org/x/mobile/bind"
)

/*
#include <stdlib.h>
*/
import "C"

func NewObjectBuilder(n, d *C.char) error {
	if instance == nil {
		return errWalletNotExists
	}
	if n == nil {
		return errors.New("name cannot be nil")
	}
	if d == nil {
		return errors.New("schema did cannot be nil")
	}

	name := C.GoString(n)
	if _, ok := objectBuilders[name]; ok {
		return fmt.Errorf("object builder exists with name '%s'", name)
	}

	schemaDid := C.GoString(d)
	builder, err := instance.NewObjectBuilder(schemaDid)
	if err != nil {
		return err
	}

	objectBuilders[name] = builder
	return nil
}

func SetObjectLabel(n, l *C.char) error {
	if instance == nil {
		return errWalletNotExists
	}
	if n == nil {
		return errors.New("name cannot be nil")
	}
	if l == nil {
		return errors.New("label cannot be nil")
	}

	name := C.GoString(n)
	if _, ok := objectBuilders[name]; !ok {
		return fmt.Errorf("no object builder with name '%s'", name)
	}

	label := C.GoString(l)
	objectBuilders[name].SetLabel(label)
	return nil
}

// TODO: think about how LIST and ANY must be implemented
// note that a method must be included for each data type
// because passing an interface{} doesn't work across Cgo

func SetBool(n, f *C.char, v C.int) error {
	if instance == nil {
		return errWalletNotExists
	}
	if n == nil {
		return errors.New("name cannot be nil")
	}
	if f == nil {
		return errors.New("field name cannot be nil")
	}

	name := C.GoString(n)
	if _, ok := objectBuilders[name]; !ok {
		return fmt.Errorf("no object builder with name '%s'", name)
	}

	fieldName := C.GoString(f)
	var value bool
	if int(v) == 0 {
		value = false
	} else {
		value = true
	}
	return objectBuilders[name].Set(fieldName, value)
}

func SetInt(n, f *C.char, v C.int) error {
	if instance == nil {
		return errWalletNotExists
	}
	if n == nil {
		return errors.New("name cannot be nil")
	}
	if f == nil {
		return errors.New("field name cannot be nil")
	}

	name := C.GoString(n)
	if _, ok := objectBuilders[name]; !ok {
		return fmt.Errorf("no object builder with name '%s'", name)
	}

	fieldName := C.GoString(f)
	value := int(v)
	return objectBuilders[name].Set(fieldName, value)
}

func SetFloat(n, f *C.char, v C.float) error {
	if instance == nil {
		return errWalletNotExists
	}
	if n == nil {
		return errors.New("name cannot be nil")
	}
	if f == nil {
		return errors.New("field name cannot be nil")
	}

	name := C.GoString(n)
	if _, ok := objectBuilders[name]; !ok {
		return fmt.Errorf("no object builder with name '%s'", name)
	}

	fieldName := C.GoString(f)
	value := float32(v)
	return objectBuilders[name].Set(fieldName, value)
}

func SetString(n, f, v *C.char) error {
	if instance == nil {
		return errWalletNotExists
	}
	if n == nil {
		return errors.New("name cannot be nil")
	}
	if f == nil {
		return errors.New("field name cannot be nil")
	}
	if v == nil {
		return errors.New("value cannot be nil")
	}

	name := C.GoString(n)
	if _, ok := objectBuilders[name]; !ok {
		return fmt.Errorf("no object builder with name '%s'", name)
	}

	fieldName := C.GoString(f)
	value := C.GoString(v)
	return objectBuilders[name].Set(fieldName, value)
}

func SetBytes(n, f *C.char, v []byte) error {
	if instance == nil {
		return errWalletNotExists
	}
	if n == nil {
		return errors.New("name cannot be nil")
	}
	if f == nil {
		return errors.New("field name cannot be nil")
	}
	if v == nil {
		return errors.New("value cannot be nil")
	}

	name := C.GoString(n)
	if _, ok := objectBuilders[name]; !ok {
		return fmt.Errorf("no object builder with name '%s'", name)
	}

	fieldName := C.GoString(f)
	return objectBuilders[name].Set(fieldName, v)
}

func SetLink(n, f, v *C.char) error {
	if instance == nil {
		return errWalletNotExists
	}
	if n == nil {
		return errors.New("name cannot be nil")
	}
	if f == nil {
		return errors.New("field name cannot be nil")
	}
	if v == nil {
		return errors.New("value cannot be nil")
	}

	name := C.GoString(n)
	if _, ok := objectBuilders[name]; !ok {
		return fmt.Errorf("no object builder with name '%s'", name)
	}

	fieldName := C.GoString(f)
	value := C.GoString(v)
	return objectBuilders[name].Set(fieldName, value)
}
