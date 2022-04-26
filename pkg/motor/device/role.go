package device

import "fmt"

// Role is the type of the node (Client, Highway)
type Role int

const (
	// StubMode_LIB is the Node utilized by Mobile and Web Clients
	Role_UNSPECIFIED Role = iota

	// StubMode_CLI is the Node utilized by CLI Clients
	Role_TEST

	// Role_MOTOR is for a Motor Node
	Role_MOTOR

	// Role_HIGHWAY is for a Highway Node
	Role_HIGHWAY
)

// Motor returns true if the node has a client.
func (m Role) IsMotor() bool {
	return m == Role_MOTOR
}

// Highway returns true if the node has a highway stub.
func (m Role) IsHighway() bool {
	return m == Role_HIGHWAY
}

// Prefix returns golog prefix for the node.
func (m Role) Prefix() string {
	var name string
	switch m {
	case Role_HIGHWAY:
		name = "highway"
	case Role_MOTOR:
		name = "motor"
	case Role_TEST:
		name = "test"
	default:
		name = "unknown"
	}
	return fmt.Sprintf("[SONR.%s] ", name)
}
