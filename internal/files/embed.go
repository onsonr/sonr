package files

import (
	"context"
	_ "embed"
	"os"
)

//go:embed vault.wasm
var vaultWasmData []byte

func writeServiceWorkerJS(path string) error {
	// Create the service worker file
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the service worker file to the specified path
	err = VaultServiceWorker(kVaultFileName).Render(context.Background(), file)
	if err != nil {
		return err
	}
	return nil
}

func writeVaultWASM(path string) error {
	// Create the vault file
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the embedded vault file to the specified path
	err = os.WriteFile(file.Name(), vaultWasmData, 0o644)
	if err != nil {
		return err
	}
	return nil
}

func writeIndexHTML(path string) error {
	// create the index file
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// write the index file to the specified path
	err = IndexHTML().Render(context.Background(), file)
	if err != nil {
		return err
	}
	return nil
}
