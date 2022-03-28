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
)

// RawCred represents a credential to be used by application that
// executes the scheme to prove possesion of an anonymous credential.
type RawCred struct {
	attrs                map[int]CredAttr
	attrIndices          map[string]int // positions of attributes amongst all attributes
	attrCount            *AttrCount
	attrKnownIndices     map[string]int // positions of known attributes amongst known attributes
	attrCommittedIndices map[string]int // positions of commited attributes amongst committed attributes
}

func NewRawCred(c *AttrCount) *RawCred {
	return &RawCred{
		attrs:                make(map[int]CredAttr),
		attrIndices:          make(map[string]int),
		attrCount:            c,
		attrKnownIndices:     make(map[string]int),
		attrCommittedIndices: make(map[string]int),
	}
}

// missingAttrs checks whether any of the attributes
// associated with this raw credential was left unset by the client.
func (c *RawCred) missingAttrs() error {
	for _, a := range c.attrs {
		if !a.HasVal() {
			fmt.Println(a.GetName(), " missing")
			return fmt.Errorf(a.GetName())
		}
		fmt.Println(a.GetName(), " ok")
	}
	return nil
}

func (c *RawCred) GetAttr(name string) (CredAttr, error) {
	i, ok := c.attrIndices[name]
	if !ok {
		return nil, fmt.Errorf("no attribute %s in this credential", name)
	}
	return c.attrs[i], nil
}

func (c *RawCred) AddEmptyStrAttr(name string, known bool) error {
	if err := c.validateAttr(name, known); err != nil {
		return err
	}
	i := len(c.attrs)
	empty := NewEmptyStrAttr(name, known)
	c.insertAttr(i, empty)

	return nil
}

func (c *RawCred) AddStrAttr(name, val string, known bool) error {
	if err := c.AddEmptyStrAttr(name, known); err != nil {
		return err
	}

	a, _ := c.GetAttr(name)
	return a.UpdateValue(val)
}

func (c *RawCred) AddInt64Attr(name string, val int64, known bool) error {
	if err := c.AddEmptyInt64Attr(name, known); err != nil {
		return err
	}

	a, _ := c.GetAttr(name)
	return a.UpdateValue(val)
}

func (c *RawCred) AddEmptyInt64Attr(name string, known bool) error {
	if err := c.validateAttr(name, known); err != nil {
		return err
	}
	i := len(c.attrs)
	empty := NewEmptyInt64Attr(name, known)
	c.insertAttr(i, empty)
	return nil
}

// GetKnownVals returns *big.Int values of known attributes.
// The returned elements are ordered by attribute's index.
func (c *RawCred) GetKnownVals() []*big.Int {
	var values []*big.Int
	for i := 0; i < len(c.attrs); i++ { // avoid range to have attributes in proper order
		attr := c.attrs[i]
		if attr.IsKnown() {
			values = append(values, attr.InternalValue())
		}
	}

	return values
}

// GetCommittedVals returns *big.Int values of committed attributes.
// The returned elements are ordered by attribute's index.
func (c *RawCred) GetCommittedVals() []*big.Int {
	var values []*big.Int
	for i := 0; i < len(c.attrs); i++ { // avoid range to have attributes in
		// proper order
		attr := c.attrs[i]
		if !attr.IsKnown() {
			values = append(values, attr.InternalValue())
		}
	}

	return values
}

func (c *RawCred) GetAttrs() map[int]CredAttr {
	return c.attrs
}

func (c *RawCred) GetAttrInternalIndex(attrName string) (int, error) {
	a, err := c.GetAttr(attrName)
	if err != nil {
		return -1, err
	}
	if a.IsKnown() {
		return c.attrKnownIndices[attrName], nil
	} else {
		return c.attrCommittedIndices[attrName], nil
	}
}

func (c *RawCred) insertAttr(i int, a CredAttr) {
	c.attrIndices[a.GetName()] = i
	c.attrs[i] = a
	if a.IsKnown() {
		c.attrKnownIndices[a.GetName()] = len(c.attrKnownIndices)
	} else {
		c.attrCommittedIndices[a.GetName()] = len(c.attrCommittedIndices)
	}
}

func (c *RawCred) validateAttr(name string, known bool) error {
	if known && len(c.GetKnownVals()) >= c.attrCount.Known {
		return fmt.Errorf("known attributes exhausted")
	}

	if !known && len(c.GetCommittedVals()) >= c.attrCount.Committed {
		return fmt.Errorf("committed attributes exhausted")
	}

	if name == "" {
		return fmt.Errorf("attribute's name cannot be empty")
	}

	if c.hasAttr(name) {
		return fmt.Errorf("duplicate attribute, ignoring")
	}

	return nil
}

func (c *RawCred) hasAttr(name string) bool {
	for _, a := range c.attrs {
		if name == a.GetName() {
			return true
		}
	}

	return false
}
