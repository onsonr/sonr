package types

import (
	"fmt"
	"strings"
	"time"

	"github.com/sonr-io/sonr/pkg/did"
)

// UnmarshalDidDocument unmarshals the whois document into a DID document
func (w *WhoIs) UnmarshalDidDocument() (*did.Document, error) {
	doc := did.Document{}
	err := doc.UnmarshalJSON(w.DidDocument)
	if err != nil {
		return nil, err
	}
	return &doc, nil
}

// CopyFromDidDocument copies the DID document into the WhoIs object if the DID document has the same DID owner
func (w *WhoIs) CopyFromDidDocument(doc *did.Document) error {
	if w.Owner != strings.TrimLeft(doc.ID.ID, "did:snr:") {
		return fmt.Errorf("owner mismatch: %s != %s", w.Owner, doc.ID.ID)
	}

	w.Alias = doc.AlsoKnownAs
	w.Controllers = doc.ControllersAsString()
	w.Timestamp = time.Now().Unix()
	w.IsActive = true
	return nil
}
