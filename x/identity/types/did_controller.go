// Utility functions for DID Controller - https://w3c.github.io/did-core/#controller
// I.e. Verification Material for IPFS Node which stores MPC Configurations
package types

import "errors"

func (d *DidDocument) Alias() string {
	return d.AlsoKnownAs[0]
}

func (d *DidDocument) ControllerCount() int {
	return len(d.Controller)
}

func (d *DidDocument) FindController(did string) (string, error) {
	for _, c := range d.Controller {
		if c == did {
			return c, nil
		}
	}
	return "", errors.New("did not found")
}

func (d *DidDocument) FindContext() []string {
	return d.Context
}

// AddController adds a DID as a controller
func (d *DidDocument) AddController(id string) {
	if d.Controller == nil {
		d.Controller = make([]string, 0)
	}
	d.Controller = append(d.Controller, id)
}

// AddAlias adds a string alias to the document for a .snr domain name into the AlsoKnownAs field
// in the document.
func (d *DidDocument) AddAlias(alias string) {
	if d.AlsoKnownAs == nil {
		d.AlsoKnownAs = make([]string, 0)
	}
	d.AlsoKnownAs = append(d.AlsoKnownAs, alias)
}

// IsController returns whether the given DID is a controller of the DID document.
func (d *DidDocument) IsController(controller string) bool {
	if controller == "" {
		return false
	}
	for _, curr := range d.Controller {
		if curr == controller {
			return true
		}
	}

	return false
}

// ControllersAsString returns all DID controllers as a string array
func (d *DidDocument) ControllersAsString() []string {
	var controllers []string
	for _, controller := range d.Controller {
		controllers = append(controllers, controller)
	}
	return controllers
}
