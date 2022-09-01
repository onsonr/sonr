package motor

import (
	"encoding/json"
	"fmt"

	"github.com/sonr-io/sonr/pkg/motor/x/object"
	_ "golang.org/x/mobile/bind"
)

var (
	objectBuilders map[string]*object.ObjectBuilder
)

func NewObjectBuilder(name, schemaDid string) error {
	if instance == nil {
		return errWalletNotExists
	}

	if _, ok := objectBuilders[name]; ok {
		return fmt.Errorf("object builder exists with name '%s'", name)
	}

	builder, err := instance.NewObjectBuilder(schemaDid)
	if err != nil {
		return err
	}

	objectBuilders[name] = builder
	return nil
}

func SetObjectLabel(name, label string) error {
	if instance == nil {
		return errWalletNotExists
	}

	if _, ok := objectBuilders[name]; !ok {
		return fmt.Errorf("no object builder with name '%s'", name)
	}

	objectBuilders[name].SetLabel(label)
	return nil
}

func SetBool(name, fieldName string, v int) error {
	if instance == nil {
		return errWalletNotExists
	}

	if _, ok := objectBuilders[name]; !ok {
		return fmt.Errorf("no object builder with name '%s'", name)
	}

	var value bool
	if int(v) == 0 {
		value = false
	} else {
		value = true
	}
	return objectBuilders[name].Set(fieldName, value)
}

func SetInt(name, fieldName string, v int) error {
	if instance == nil {
		return errWalletNotExists
	}

	if _, ok := objectBuilders[name]; !ok {
		return fmt.Errorf("no object builder with name '%s'", name)
	}

	value := int(v)
	return objectBuilders[name].Set(fieldName, value)
}

func SetFloat(name, fieldName string, v float32) error {
	if instance == nil {
		return errWalletNotExists
	}

	if _, ok := objectBuilders[name]; !ok {
		return fmt.Errorf("no object builder with name '%s'", name)
	}

	value := float32(v)
	return objectBuilders[name].Set(fieldName, value)
}

func SetString(name, fieldName, value string) error {
	if instance == nil {
		return errWalletNotExists
	}

	if _, ok := objectBuilders[name]; !ok {
		return fmt.Errorf("no object builder with name '%s'", name)
	}

	return objectBuilders[name].Set(fieldName, value)
}

func SetBytes(name, fieldName string, v []byte) error {
	if instance == nil {
		return errWalletNotExists
	}

	if _, ok := objectBuilders[name]; !ok {
		return fmt.Errorf("no object builder with name '%s'", name)
	}

	return objectBuilders[name].Set(fieldName, v)
}

func SetLink(name, fieldName, value string) error {
	if instance == nil {
		return errWalletNotExists
	}

	if _, ok := objectBuilders[name]; !ok {
		return fmt.Errorf("no object builder with name '%s'", name)
	}

	return objectBuilders[name].Set(fieldName, value)
}

func RemoveObjectField(name, fieldName string) error {
	if instance == nil {
		return errWalletNotExists
	}

	if builder, ok := objectBuilders[name]; !ok {
		return fmt.Errorf("no object builder with name '%s'", name)
	} else {
		builder.Remove(fieldName)
	}
	return nil
}

func BuildObject(name string) ([]byte, error) {
	if instance == nil {
		return nil, errWalletNotExists
	}

	builder, ok := objectBuilders[name]
	if !ok {
		return nil, fmt.Errorf("no object builder with name '%s'", name)
	}

	res, err := builder.Build()
	if err != nil {
		return nil, err
	}

	// Using JSON marshalling here for arbitrary object types
	return json.Marshal(res)
}

func UploadObject(name string) ([]byte, error) {
	if instance == nil {
		return nil, errWalletNotExists
	}

	builder, ok := objectBuilders[name]
	if !ok {
		return nil, fmt.Errorf("no object builder with name '%s'", name)
	}

	res, err := builder.Upload()
	if err != nil {
		return nil, err
	}

	return res.Marshal()
}

func GetObject(cid string) ([]byte, error) {
	if instance == nil {
		return nil, errWalletNotExists
	}

	res, err := instance.QueryObject(cid)
	if err != nil {
		return nil, err
	}

	// Using JSON marshalling here for arbitrary object types
	return json.Marshal(res)
}
