package did

import (
	"time"
)

type constError string

func (err constError) Error() string {
	return string(err)
}

const (
	// InvalidDIDErr indicates: "The DID supplied to the DID resolution function does not conform to valid syntax. (See § 3.1 DID Syntax.)"
	InvalidDIDErr = constError("supplied DID is invalid")
	// NotFoundErr indicates: "The DID resolver was unable to find the DID document resulting from this resolution request."
	NotFoundErr = constError("supplied DID wasn't found")
	// DeactivatedErr indicates: The DID supplied to the DID resolution function has been deactivated. (See § 7.2.4 Deactivate .)
	DeactivatedErr = constError("supplied DID is deactivated")
)

// Resolver defines the interface for DID resolution as specified by the DID Core specification (https://www.w3.org/TR/did-core/#did-resolution).
type Resolver interface {
	// Resolve tries to resolve the given input DID to its DID Document and Metadata. In addition to errors specific
	// to this resolver it can return InvalidDIDErr, NotFoundErr and DeactivatedErr as specified by the DID Core specification.
	// If no error occurs the DID Document and Medata are returned.
	Resolve(inputDID string) (*Document, *DocumentMetadata, error)
}

// DocumentMedata represents DID Document Metadata as specified by the DID Core specification (https://www.w3.org/TR/did-core/#did-document-metadata-properties).
type DocumentMetadata struct {
	Created    *time.Time
	Updated    *time.Time
	Properties map[string]interface{}
}
