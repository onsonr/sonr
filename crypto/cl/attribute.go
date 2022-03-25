/*
 * Copyright 2017 XLAB d.o.o.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package cl

import (
	"fmt"
	"math/big"
	"strconv"
)

// AttrCount holds the number of known, committed and
// hidden parameters.
type AttrCount struct {
	Known     int
	Committed int
	Hidden    int
}

func NewAttrCount(known, committed, hidden int) *AttrCount {
	return &AttrCount{
		Known:     known,
		Committed: committed,
		Hidden:    hidden,
	}
}

func (c *AttrCount) String() string {
	return fmt.Sprintf("known: %d\ncommitted: %d\nhidden: %d\n",
		c.Known, c.Committed, c.Hidden)
}

// CredAttr represents an attribute for the CL scheme.
type CredAttr interface {
	GetValue() interface{}
	FromInternalValue(*big.Int) (interface{}, error)
	UpdateValue(interface{}) error
	InternalValue() *big.Int
	SetInternalValue() error
	IsKnown() bool
	HasVal() bool
	GetName() string
	String() string
}

// attr is part of a credential (RawCredential). In the case of digital identity credential,
// attributes could be for example Name, Gender, Date of Birth. In the case of a credential allowing
// access to some internet service (like electronic newspaper), attributes could be
// Type (for example only news related to politics) of the service and Date of Expiration.
type attr struct {
	Name   string
	Known  bool
	valSet bool
	val    *big.Int
}

func newAttr(name string, known bool) *attr {
	return &attr{
		Name:   name,
		Known:  known,
		valSet: false,
	}
}

func (a *attr) IsKnown() bool {
	return a.Known
}

func (a *attr) InternalValue() *big.Int {
	return a.val
}

func (a *attr) HasVal() bool {
	return a.valSet
}

func (a *attr) GetName() string {
	return a.Name
}

func (a *attr) String() string {
	tag := "known"
	if !a.IsKnown() {
		tag = "revealed"
	}
	return fmt.Sprintf("%s (%s)", a.Name, tag)
}

type Int64Attr struct {
	val int64
	*attr
}

func NewEmptyInt64Attr(name string, known bool) *Int64Attr {
	return &Int64Attr{
		attr: newAttr(name, known),
	}
}

func NewInt64Attr(name string, val int64, known bool) (*Int64Attr,
	error) {
	a := &Int64Attr{
		val:  val,
		attr: newAttr(name, known),
	}
	if err := a.SetInternalValue(); err != nil {
		return nil, err
	}

	return a, nil
}

func (a *Int64Attr) SetInternalValue() error {
	a.attr.val = big.NewInt(int64(a.val)) // FIXME
	a.valSet = true
	return nil
}

func (a *Int64Attr) GetValue() interface{} {
	return a.val
}

func (a *Int64Attr) FromInternalValue(val *big.Int) (interface{}, error) {
	return strconv.Atoi(val.String())
}

func (a *Int64Attr) UpdateValue(n interface{}) error {
	switch n.(type) {
	case int:
		a.val = int64(n.(int))
	case int64:
		a.val = n.(int64)
	}
	return a.SetInternalValue()
}

func (a *Int64Attr) String() string {
	return fmt.Sprintf("%s, type = %T", a.attr.String(), a.val)
}

type StrAttr struct {
	val string
	*attr
}

func NewEmptyStrAttr(name string, known bool) *StrAttr {
	return &StrAttr{
		attr: newAttr(name, known),
	}
}

func NewStrAttr(name, val string, known bool) (*StrAttr,
	error) {
	a := &StrAttr{
		val:  val,
		attr: newAttr(name, known),
	}
	if err := a.SetInternalValue(); err != nil {
		return nil, err
	}

	return a, nil
}

func (a *StrAttr) SetInternalValue() error {
	a.attr.val = new(big.Int).SetBytes([]byte(a.val)) // FIXME
	a.valSet = true
	return nil
}

func (a *StrAttr) GetValue() interface{} {
	return a.val
}

func (a *StrAttr) FromInternalValue(val *big.Int) (interface{}, error) {
	return string(val.Bytes()), nil
}

func (a *StrAttr) UpdateValue(s interface{}) error {
	a.val = s.(string)
	return a.SetInternalValue()
}

func (a *StrAttr) String() string {
	return fmt.Sprintf("%s, type = %T", a.attr.String(), a.val)
}

// FIXME make nicer
// Hook to organization?
func ParseAttrs(specs map[string]interface{}) ([]CredAttr, *AttrCount, error) {
	attrs := make([]CredAttr, len(specs))
	var nKnown, nCommitted int

	for name, val := range specs {
		data, ok := val.(map[string]interface{})
		if !ok {
			return nil, nil, fmt.Errorf("invalid configuration")
		}

		t, ok := data["type"]
		if !ok {
			return nil, nil, fmt.Errorf("missing type specifier")
		}

		ind, ok := data["index"] // to make sure attributes are always sent in proper order to the client
		if !ok {
			return nil, nil, fmt.Errorf("missing index specifier")
		}
		index, err := strconv.Atoi(ind.(string))
		if err != nil {
			return nil, nil, fmt.Errorf("index must be string")
		}

		known := true
		k, ok := data["known"]
		if ok {
			res, err := strconv.ParseBool(k.(string))
			if err != nil {
				return nil, nil, fmt.Errorf("known must be true or false")
			}
			known = res
		}

		if known {
			nKnown++
		} else {
			nCommitted++
		}

		switch t {
		case "string":
			a, err := NewStrAttr(name, "", known) // FIXME
			if err != nil {
				return nil, nil, err
			}
			attrs[index] = a
		case "int64":
			a, err := NewInt64Attr(name, 0, known) // FIXME
			if err != nil {
				return nil, nil, err
			}
			attrs[index] = a
		default:
			return nil, nil, fmt.Errorf("unsupported attribute type: %s", t)
		}

	}

	// TODO hidden params
	return attrs, NewAttrCount(nKnown, nCommitted, 0), nil
}
