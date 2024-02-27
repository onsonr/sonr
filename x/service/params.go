package service

// DefaultParams returns default module parameters.
func DefaultParams() Params {
	return Params{
		// Set default values here.
	}
}

// Validate does the sanity check on the params.
func (p Params) Validate() error {
	// Sanity check goes here.
	return nil
}

// DefaultRecords returns default records.
func DefaultRecords() []Record {
	return []Record{
		{
			Name:        "Sonr Localhost",
			Origin:      "localhost",
			Description: "Sonr Localhost Chat App",
			Authority:   "#",
		},
	}
}
