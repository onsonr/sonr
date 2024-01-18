package abi

import (
	"fmt"
	"math/big"
	"reflect"
	"strings"

	"golang.org/x/crypto/sha3"

	util "github.com/sonrhq/sonr/pkg/hexutil"
)

type ABI struct {
	Constructor *Method
	Methods     map[string]*Method
	Events      map[string]*Event
}

func (a *ABI) Pack(name string, params ...interface{}) ([]byte, error) {
	if name == "" {
		return nil, fmt.Errorf("constructors are not supported yet")
	}
	if a.Methods[name] == nil {
		return nil, fmt.Errorf("method not found")
	}
	m := a.Methods[name]
	var value []byte
	value = append(value, m.SigId()...)
	value = append(value, m.Inputs.Pack(params...)...)
	return value, nil
}

func (a *ABI) PackParams(name string, params ...interface{}) ([]byte, error) {
	if name == "" {
		return nil, fmt.Errorf("constructors are not supported yet")
	}
	if a.Methods[name] == nil {
		return nil, fmt.Errorf("method not found")
	}
	m := a.Methods[name]
	var value []byte
	value = append(value, m.Inputs.Pack(params...)...)
	return value, nil
}

type Method struct {
	Name    string
	Const   bool
	Inputs  Arguments
	Outputs Arguments
}

func (m *Method) SigId() []byte {
	// function foo(uint32 a, int b)    =    "foo(uint32,int256)"
	types := make([]string, len(m.Inputs))
	for i, v := range m.Inputs {
		types[i] = v.Type
	}
	functionStr := fmt.Sprintf("%v(%v)", m.Name, strings.Join(types, ","))
	keccak256 := sha3.NewLegacyKeccak256()
	keccak256.Write([]byte(functionStr))
	return keccak256.Sum(nil)[:4]
}

type Event struct {
	Name      string
	Anonymous bool
	Inputs    Arguments
}

type Arguments []Argument

type Argument struct {
	Name    string
	Type    string
	Indexed bool // indexed is only used by events
}

func (arg Arguments) Pack(params ...interface{}) []byte {
	if len(arg) != len(params) {
		fmt.Errorf("inconsistent number of parameters")
	}

	var value []byte
	i := 0
	for _, v := range arg {
		p := params[i]
		i++
		switch v.Type {
		case "uint256", "uint128", "uint64", "uint32", "uint", "int256", "int128", "int64", "int32", "int":
			va := reflect.ValueOf(p).Interface().(*big.Int) // new(big.Int).SetUint64(reflect.ValueOf(p).Uint()) //params[k].(big.Int)
			value = append(value, PaddedBigBytes(U256(va), 32)...)
			// value = append()
		case "string":
			// packBytesSlice([]byte(reflectValue.String()), reflectValue.Len())
			va := new(big.Int).SetBytes([]byte(reflect.ValueOf(p).String())) //
			value = append(value, PaddedBigBytes(U256(va), 32)...)
		case "address":
			addr := reflect.ValueOf(p).String()
			// addrByte, _ := hex.DecodeString(addr)

			va := new(big.Int).SetBytes(util.RemoveZeroHex(addr))
			value = append(value, PaddedBigBytes(U256(va), 32)...)
		}
	}
	return value
}
