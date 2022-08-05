package types

import (
	"fmt"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sonr-io/sonr/pkg/crypto/did"
)

// ContainsAlias checks if the alias is in the list of aliases of the whois
func (w *WhoIs) ContainsAlias(target string) bool {
	for _, a := range w.Alias {
		if a.GetName() == target {
			return true
		}
	}
	return false
}

// ContainsController checks if the controller is in the list of controllers of the whois
func (w *WhoIs) ContainsController(target string) bool {
	// Checks if the controller is in the list of controllers
	for _, c := range w.Controllers {
		if c == target {
			return true
		}
	}
	return false
}

// OwnerAccAddress returns the owner address of the whois
func (w *WhoIs) OwnerAccAddress() (addr sdk.AccAddress, err error) {
	return sdk.AccAddressFromBech32(w.GetOwner())
}

// UpdateDidBuffer copies the DID document into the WhoIs object if the DID document has the same DID owner
func (w *WhoIs) UpdateDidBuffer(buf []byte) (WhoIs, error) {
	// Trim snr account prefix
	rawCreator := strings.TrimPrefix(w.GetOwner(), "snr")
	docI, err := did.NewDocument(fmt.Sprintf("did:snr:%s", rawCreator))
	if err != nil {
		return *w, err
	}
	err = docI.CopyFromBytes(buf)
	if err != nil {
		return *w, err
	}
	for _, a := range docI.GetAlsoKnownAs() {
		if !w.ContainsAlias(a) {
			w.AddAlsoKnownAs(a, false)
		}
	}

	snrDoc, err := NewDIDDocumentFromPkg(docI)
	if err != nil {
		return *w, err
	}

	w.Controllers = docI.ControllersAsString()
	w.Timestamp = time.Now().Unix()
	w.IsActive = true
	w.DidDocument = snrDoc
	return *w, nil
}

// GetAlsoKnownAs returns the list of aliases of the whois as string array
func (w *WhoIs) GetAlsoKnownAs() []string {
	var aliases []string
	for _, a := range w.Alias {
		aliases = append(aliases, a.GetName())
	}
	return aliases
}

// AddAlsoKnownAs adds an alias to the list of aliases of the whois and the underlying DID document
func (w *WhoIs) AddAlsoKnownAs(as string, updateDoc bool) (WhoIs, error) {
	// Update the WhoIs
	aliases := w.GetAlias()
	if !w.ContainsAlias(as) {
		aliases = append(aliases, &Alias{
			Name:      as,
			IsForSale: false,
			Amount:    -1,
		})
	}
	w.Alias = aliases

	if updateDoc {
		// Update the DID document
		doc := w.GetDidDocument()

		// Add the alias to the DID document
		doc.AlsoKnownAs = w.GetAlsoKnownAs()
		w.DidDocument = doc
	}
	return *w, nil
}

// RemoveAlsoKnownAs removes an alias from the list of aliases of the whois and the underlying DID document
func (w *WhoIs) RemoveAlsoKnownAs(as string, updateDoc bool) (WhoIs, error) {
	// Update the WhoIs
	aliases := w.GetAlias()
	i, _, err := w.FindAliasByName(as)
	if err != nil {
		return *w, err
	}
	aliases = append(aliases[:i], aliases[i+1:]...)
	w.Alias = aliases

	if updateDoc {
		// Update the DID document
		doc := w.GetDidDocument()

		// Remove the alias from the DID document
		doc.AlsoKnownAs = w.GetAlsoKnownAs()
		w.DidDocument = doc
	}
	return *w, nil
}

// FindAliasByName returns the alias and index with the given name or error if not found
func (w *WhoIs) FindAliasByName(name string) (int, *Alias, error) {
	for i, a := range w.Alias {
		if a.GetName() == name {
			return i, a, nil
		}
	}
	return -1, nil, fmt.Errorf("alias %s not found", name)
}

// UpdateAlias updates the alias properties for amount and IsForSale
func (w *WhoIs) UpdateAlias(name string, amount int, isForSale bool) (WhoIs, error) {
	// Update the WhoIs
	i, a, err := w.FindAliasByName(name)
	if err != nil {
		return *w, err
	}
	a.Amount = int32(amount)
	a.IsForSale = isForSale
	w.Alias[i] = a
	return *w, nil
}

// GetDidDocumentBuffer returns the DID document as bytes from JSON
func (w *WhoIs) GetDidDocumentBuffer() ([]byte, error) {
	if w.DidDocument == nil {
		return nil, fmt.Errorf("no DID document found")
	}
	return w.DidDocument.MarshalJSON()
}
