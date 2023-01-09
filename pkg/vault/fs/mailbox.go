package fs

import "errors"

// A method that returns a list of messages.
func (c *Config) ListMessages() ([][]byte, error) {
	return nil, errors.New("Method not implemented")
}

// Sending a message to a user.
func (c *Config) SendMessage(to []byte, message []byte) error {
	return errors.New("Method not implemented")
}
