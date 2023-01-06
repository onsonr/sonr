package protocol

// A Protocol is a protocol that can be used to authenticate identities.
//
// The first thing to notice is that the type is an interface. This means that any type that implements
// the methods defined in the interface is also a Protocol.
//
// The second thing to notice is that the type is defined in the Go package. This means that any type
// that implements the methods defined in the interface is also a Protocol.
//
// The third thing to notice is that the type is defined in the Go package. This means that any type
// that implements the methods defined in the interface is also a Protocol.
// @property {string} Name - The name of the protocol.
// @property {string} Version - The version of the protocol.
// @property {error} Init - This is called when the protocol is initialized.
// @property GetChallenge - This method returns a challenge for the given identity.
// @property {error} Register - This is the function that will be called when a user registers.
// @property {error} Authenticate - Authenticates the given identity.
// @property {error} Close - Closes the protocol.
type Protocol interface {
	// Name returns the name of the protocol.
	Name() string

	// Version returns the version of the protocol.
	Version() string

	// Init initializes the protocol.
	Init() error

	// GetChallenge returns a challenge for the given identity.
	GetChallenge(identity string) (string, error)

	// Register registers the given identity.
	Register(identity string, challenge string, response string) error

	// Authenticate authenticates the given identity.
	Authenticate(identity string, challenge string, response string) error

	// Close closes the protocol.
	Close() error
}

