package schemas

import (
	"fmt"
	"reflect"

	"github.com/ipld/go-ipld-prime/datamodel"
	"github.com/ipld/go-ipld-prime/node/basicnode"
	st "github.com/sonr-io/sonr/x/schema/types"
)

func (as *SchemaImpl) BuildNodesFromDefinition(label, schemaDid string, object map[string]interface{}) error {
	if as.fields == nil {
		return errSchemaFieldsInvalid
	}

	if err := as.VerifyDocument(object); err != nil {
		return fmt.Errorf("%s: %s", errSchemaFieldsInvalid, err)
	}

	// Create IPLD Noded
	np := basicnode.Prototype.Any
	nb := np.NewBuilder()     // Create a builder.
	ma, err := nb.BeginMap(3) // label, schema DID, and document
	if err != nil {
		return err
	}

	labelAs, err := ma.AssembleEntry(st.IPLD_LABEL)
	if err != nil {
		return err
	}
	if err := labelAs.AssignString(label); err != nil {
		return err
	}

	schemaDidAs, err := ma.AssembleEntry(st.IPLD_SCHEMA_DID)
	if err != nil {
		return err
	}
	if err := schemaDidAs.AssignString(schemaDid); err != nil {
		return err
	}

	docAs, err := ma.AssembleEntry(st.IPLD_DOCUMENT)
	if err != nil {
		return err
	}
	valueAs, err := docAs.BeginMap(int64(len(as.fields)))
	if err != nil {
		return err
	}

	for _, t := range as.fields {
		k := t.Name
		valueAs.AssembleKey().AssignString(k)
		if t.GetKind() != st.Kind_LINK {
			err = as.AssignValueToNode(t.FieldKind, valueAs, object[k])
			if err != nil {
				return fmt.Errorf("assign value to node: %s", err)
			}
		} else if t.GetKind() == st.Kind_LINK {
			err := as.BuildSchemaFromLink(t.FieldKind.LinkDid, valueAs, object[t.Name].(map[string]interface{}))
			if err != nil {
				return fmt.Errorf("build schema from link: %s", err)
			}
		}
	}

	buildErr := valueAs.Finish()
	if buildErr != nil {
		return buildErr
	}
	buildErr = ma.Finish()
	if buildErr != nil {
		return buildErr
	}

	as.nodes = nb.Build()

	return nil
}

func (as *SchemaImpl) AssignValueToNode(kind *st.SchemaFieldKind, ma datamodel.MapAssembler, value interface{}) error {
	switch kind.GetKind() {
	case st.Kind_STRING:
		val := value.(string)
		ma.AssembleValue().AssignString(val)
	case st.Kind_INT:
		switch v := value.(type) {
		case int:
			val := int64(v)
			ma.AssembleValue().AssignInt(val)
		case int32:
			val := int64(v)
			ma.AssembleValue().AssignInt(val)
		case int64:
			val := int64(v)
			ma.AssembleValue().AssignInt(val)
		}
	case st.Kind_FLOAT:
		switch v := value.(type) {
		case float64:
			ma.AssembleValue().AssignFloat(v)
		case float32:
			ma.AssembleValue().AssignFloat(float64(v))
		}
	case st.Kind_BOOL:
		val := value.(bool)
		ma.AssembleValue().AssignBool(val)
	case st.Kind_BYTES:
		val := value.([]byte)
		ma.AssembleValue().AssignBytes(val)
	case st.Kind_LIST:
		val := make([]interface{}, 0)
		s := reflect.ValueOf(value)
		for i := 0; i < s.Len(); i++ {
			val = append(val, s.Index(i).Interface())
		}
		n, err := as.BuildNodeFromList(val, kind.ListKind)
		if err != nil {
			return fmt.Errorf("build node from list: %s", err)
		}
		ma.AssembleValue().AssignNode(n)
	default:
		return errSchemaFieldsInvalid
	}

	return nil
}

func (as *SchemaImpl) BuildSchemaFromLink(key string, ma datamodel.MapAssembler, value map[string]interface{}) error {
	if as.subWhatIs[key] == nil {
		return errNodeNotFound
	}

	sd, ok := as.subWhatIs[key]
	if !ok {
		return fmt.Errorf("sub WhatIs '%s' not found", key)
	}

	err := as.VerifySubObject(sd.Schema.Fields, value)

	if err != nil {
		return err
	}

	// Create IPLD Node
	np := basicnode.Prototype.Any
	nb := np.NewBuilder() // Create a builder.
	lma, err := nb.BeginMap(int64(len(value)))

	if err != nil {
		return err
	}

	for _, f := range sd.Schema.Fields {
		lma.AssembleKey().AssignString(f.Name)
		if f.GetKind() != st.Kind_LINK {
			err := as.AssignValueToNode(f.FieldKind, lma, value[f.Name])
			if err != nil {
				return err
			}
		} else if f.GetKind() == st.Kind_LINK {
			err = as.BuildSchemaFromLink(f.FieldKind.LinkDid, lma, value[f.Name].(map[string]interface{}))
			if err != nil {
				return err
			}
		}

	}

	lma.Finish()
	n := nb.Build()
	ma.AssembleValue().AssignNode(n)

	return nil
}

func (as *SchemaImpl) BuildSchemaFromLinkForList(key string, ma datamodel.ListAssembler, value map[string]interface{}) error {
	if as.subWhatIs[key] == nil {
		return errNodeNotFound
	}

	sd := as.subWhatIs[key]

	err := as.VerifySubObject(sd.Schema.Fields, value)

	if err != nil {
		return fmt.Errorf("verify subobject: %s", err)
	}

	// Create IPLD Node
	np := basicnode.Prototype.Any
	nb := np.NewBuilder() // Create a builder.
	lma, err := nb.BeginMap(int64(len(value)))

	if err != nil {
		return err
	}

	for _, f := range sd.Schema.Fields {
		lma.AssembleKey().AssignString(f.Name)
		if f.GetKind() != st.Kind_LINK {
			err := as.AssignValueToNode(f.FieldKind, lma, value[f.Name])
			if err != nil {
				return err
			}
		} else if f.GetKind() == st.Kind_LINK {
			err = as.BuildSchemaFromLink(f.FieldKind.LinkDid, lma, value[f.Name].(map[string]interface{}))
			if err != nil {
				return err
			}
		}

	}

	lma.Finish()
	n := nb.Build()
	ma.AssembleValue().AssignNode(n)

	return nil
}

func (as *SchemaImpl) BuildNodeFromList(lst []interface{}, kind *st.SchemaFieldKind) (datamodel.Node, error) {
	// Create IPLD Node
	np := basicnode.Prototype.Any
	nb := np.NewBuilder() // Create a builder.
	la, err := nb.BeginList(int64(len(lst)))
	if err != nil {
		return nil, err
	}

	if len(lst) < 1 {
		return nb.Build(), nil
	}

	if err := as.VerifyList(lst, kind); err != nil {
		return nil, fmt.Errorf("verify list: %s", err)
	}

	for _, val := range lst {
		switch lst[0].(type) {
		case int:
			lstItem := interface{}(val).(int)
			la.AssembleValue().AssignInt(int64(lstItem))
		case int32:
			lstItem := interface{}(val).(int32)
			la.AssembleValue().AssignInt(int64(lstItem))
		case int64:
			lstItem := interface{}(val).(int64)
			la.AssembleValue().AssignInt(int64(lstItem))
		case float64:
			lstItem := interface{}(val).(float64)
			la.AssembleValue().AssignFloat(lstItem)
		case float32:
			lstItem := interface{}(val).(float32)
			la.AssembleValue().AssignFloat(float64(lstItem))
		case bool:
			lstItem := interface{}(val).(bool)
			la.AssembleValue().AssignBool(lstItem)
		case string:
			lstItem := interface{}(val).(string)
			la.AssembleValue().AssignString(lstItem)
		case []byte:
			lstItem := interface{}(val).([]byte)
			la.AssembleValue().AssignBytes(lstItem)
		case map[string]interface{}:
			if kind.Kind == st.Kind_LINK {
				if err := as.BuildSchemaFromLinkForList(kind.LinkDid, la, val.(map[string]interface{})); err != nil {
					return nil, fmt.Errorf("build schema from link for list: %s", err)
				}
			}
		/*
			The below cases are for handling lists of up to 3 dimensions.
			Within each cases arrays are normalized to match a type of []interface{}
			each generic []interface{} array is then handed back to this function to further resolve types.
			depth is cut off at 3 dimensions due to having to implement explicit type cases here
		*/
		case []map[string]interface{}, []string, []int, []int32, []int64, []float32, []float64:
			value := make([]interface{}, 0)
			s := reflect.ValueOf(val)
			for i := 0; i < s.Len(); i++ {
				value = append(value, s.Index(i).Interface())
			}
			n, err := as.BuildNodeFromList(value, kind.ListKind)
			if err != nil {
				return nil, fmt.Errorf("build node from list: %s", err)
			}
			la.AssembleValue().AssignNode(n)
		default:
			value := make([]interface{}, 0)
			s := reflect.ValueOf(val)
			for i := 0; i < s.Len(); i++ {
				value = append(value, s.Index(i).Interface())
			}
			n, err := as.BuildNodeFromList(value, kind.ListKind)
			if err != nil {
				return nil, fmt.Errorf("build node from list: %s", err)
			}
			la.AssembleValue().AssignNode(n)
		}
	}
	err = la.Finish()

	if err != nil {
		return nil, err
	}

	return nb.Build(), nil
}
