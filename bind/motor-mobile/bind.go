package motor

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	mtr "github.com/sonr-io/sonr/pkg/motor"
	mt "github.com/sonr-io/sonr/pkg/motor/types"
	"github.com/sonr-io/sonr/pkg/motor/x/object"
	_ "golang.org/x/mobile/bind"
)

/*
#include <stdlib.h>
*/
import "C"

var (
	errWalletExists    = errors.New("mpc wallet already exists")
	errWalletNotExists = errors.New("mpc wallet does not exist")
)

var (
	instance       mtr.MotorNode
	objectBuilders map[string]*object.ObjectBuilder
)

func Init(buf []byte) ([]byte, error) {
	// Unmarshal the request
	var req mt.InitializeRequest
	if err := json.Unmarshal(buf, &req); err != nil {
		return nil, err
	}

	// Check if public key provided
	if req.DeviceKeyprintPub == nil {
		// Create Motor instance
		instance = mtr.EmptyMotor(req.DeviceId)

		// init objectBuilders
		objectBuilders = make(map[string]*object.ObjectBuilder)

		// Return Initialization Response
		resp := mt.InitializeResponse{
			Success: true,
		}
		return json.Marshal(resp)
	}
	return nil, errors.New("loading existing account not implemented")
}

func CreateAccount(buf []byte) ([]byte, error) {
	if instance == nil {
		return nil, errWalletNotExists
	}
	// decode request
	var request mt.CreateAccountRequest
	if err := json.Unmarshal(buf, &request); err != nil {
		return nil, fmt.Errorf("unmarshal request: %s", err)
	}

	if res, err := instance.CreateAccount(request); err == nil {
		return json.Marshal(res)
	} else {
		return nil, err
	}
}

func Login(buf []byte) ([]byte, error) {
	if instance == nil {
		return nil, errWalletNotExists
	}

	// decode request
	var request mt.LoginRequest
	if err := json.Unmarshal(buf, &request); err != nil {
		return nil, fmt.Errorf("error unmarshalling request: %s", err)
	}

	if res, err := instance.Login(request); err == nil {
		return json.Marshal(res)
	} else {
		return nil, err
	}
}

func CreateSchema(buf []byte) ([]byte, error) {
	if instance == nil {
		return nil, errWalletNotExists
	}

	var request mt.CreateSchemaRequest
	if err := json.Unmarshal(buf, &request); err != nil {
		return nil, fmt.Errorf("unmarshal request: %s", err)
	}

	if res, err := instance.CreateSchema(request); err == nil {
		return json.Marshal(res)
	} else {
		return nil, err
	}
}

func QueryWhatIs(buf []byte) ([]byte, error) {
	if instance == nil {
		return nil, errWalletNotExists
	}

	var request mt.QueryWhatIsRequest
	if err := json.Unmarshal(buf, &request); err != nil {
		return nil, fmt.Errorf("unmarshal request: %s", err)
	}

	if res, err := instance.QueryWhatIs(context.Background(), request); err == nil {
		return json.Marshal(res)
	} else {
		return nil, err
	}
}

// Address returns the address of the wallet.
func Address() string {
	if instance == nil {
		return ""
	}
	wallet := instance.GetWallet()
	if wallet == nil {
		return ""
	}
	addr, err := wallet.Address()
	if err != nil {
		return ""
	}
	return addr
}

// Balance returns the balance of the wallet.
func Balance() int {
	return int(instance.GetBalance())
}

// func Connect() error {
// 	if instance == nil {
// 		return errWalletNotExists
// 	}
// 	h, err := host.NewDefaultHost(context.Background(), config.DefaultConfig(config.Role_MOTOR))
// 	if err != nil {
// 		return err
// 	}
// 	instance.host = h
// 	return nil
// }

// DidDoc returns the DID document as JSON
func DidDoc() string {
	if instance == nil {
		return ""
	}
	doc := instance.GetDIDDocument()
	if doc == nil {
		return ""
	}
	buf, err := doc.MarshalJSON()
	if err != nil {
		return ""
	}
	return string(buf)
}

/**
 * OBJECT BUILDER
 */

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

func RemoveObjectField(n, f *C.char) error {
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
	fieldName := C.GoString(f)
	if builder, ok := objectBuilders[name]; !ok {
		return fmt.Errorf("no object builder with name '%s'", name)
	} else {
		builder.Remove(fieldName)
	}
	return nil
}

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

func RemoveListItem(n, f *C.char, i C.int) error {
	if instance == nil {
		return errWalletNotExists
	}

	builder, fieldName, err := findBuilder(n, f)
	if err != nil {
		return err
	}

	var list []interface{}
	var ok bool
	if list, ok = builder.Get(fieldName).([]interface{}); !ok || list == nil {
		return fmt.Errorf("no list field with name '%s'", fieldName)
	}

	index := int(i)
	if index < 0 || index >= len(list) {
		return fmt.Errorf("index %d of of range %d", index, len(list))
	}

	list = append(list[:index], list[index+1:])
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
