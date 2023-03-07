package controller

// Sign signs the data with the given account.
func (w *DIDControllerImpl) Sign(data []byte) ([]byte, error) {
	return w.primaryAccount.Sign(data)
}

// Verify verifies the signature with the given account.
func (w *DIDControllerImpl) Verify(data, sig []byte) (bool, error) {
	return w.primaryAccount.Verify(data, sig)
}

// SignWithAccount signs the data with the given account.
func (w *DIDControllerImpl) SignWithAccount(data []byte, accountName string) ([]byte, error) {
	acc, err := w.GetAccount(accountName)
	if err != nil {
		return nil, err
	}
	return acc.Sign(data)
}

// VerifyWithAccount verifies the signature with the given account.
func (w *DIDControllerImpl) VerifyWithAccount(data, sig []byte, accountName string) (bool, error) {
	acc, err := w.GetAccount(accountName)
	if err != nil {
		return false, err
	}
	return acc.Verify(data, sig)
}
