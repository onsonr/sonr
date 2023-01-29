package types

import (
	"encoding/json"
	"errors"
	fmt "fmt"
)

func resolveVerificationRelationships(relationships []*VerificationRelationship, methods []*VerificationMethod) error {
	for i, relationship := range relationships {
		if relationship.Reference != "" {
			continue
		}
		if resolved := resolveVerificationRelationship(relationship.Reference, methods); resolved == nil {
			return fmt.Errorf("unable to resolve %s: %s", verificationMethodKey, relationship.Reference)
		} else {
			relationships[i] = resolved
			relationships[i].Reference = relationship.Reference
		}
	}
	return nil
}

func resolveVerificationRelationship(reference string, methods []*VerificationMethod) *VerificationRelationship {
	for _, method := range methods {
		if method.Id == reference {
			return &VerificationRelationship{VerificationMethod: method}
		}
	}
	return nil
}

func (v VerificationRelationship) MarshalJSON() ([]byte, error) {
	if v.Reference != "" {
		return json.Marshal(*v.VerificationMethod)
	} else {
		return json.Marshal(v.Reference)
	}
}

func (v *VerificationRelationship) UnmarshalJSON(b []byte) error {
	// try to figure out if the item is an object of a string
	type Alias VerificationRelationship
	switch b[0] {
	case '{':
		tmp := Alias{VerificationMethod: &VerificationMethod{}}
		err := json.Unmarshal(b, &tmp)
		if err != nil {
			return fmt.Errorf("could not parse verificationRelation method: %w", err)
		}
		*v = (VerificationRelationship)(tmp)
	case '"':
		err := json.Unmarshal(b, &v.Reference)
		if err != nil {
			return fmt.Errorf("could not parse verificationRelation key relation DId:%w", err)
		}
	default:
		return errors.New("verificationRelation is invalid")
	}
	return nil
}
