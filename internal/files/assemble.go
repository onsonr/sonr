package files

import (
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"

	"github.com/onsonr/sonr/internal/db"
)

var (
	kServiceWorkerFileName = "sw.js"
	kVaultFileName         = "vault.wasm"
	kIndexFileName         = "index.html"
)

func Assemble(dir string) error {
	err := os.MkdirAll(dir, 0o755)
	if err != nil {
		return err
	}

	// Write the vault file
	if err := writeVaultWASM(filepath.Join(dir, kVaultFileName)); err != nil {
		return err
	}

	// Write the service worker file
	if err := writeServiceWorkerJS(filepath.Join(dir, kServiceWorkerFileName)); err != nil {
		return err
	}

	// Write the index file
	if err := writeIndexHTML(filepath.Join(dir, kIndexFileName)); err != nil {
		return err
	}

	// Initialize the database
	if err := initializeDatabase(dir); err != nil {
		return err
	}

	return nil
}

func initializeDatabase(dir string) error {
	db, err := db.Open(db.New(db.WithDir(dir)))
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}
	defer db.Close()

	// You can add some initial data here if needed
	// For example:
	// err = db.AddChain("Ethereum", "1")
	// if err != nil {
	// 	return fmt.Errorf("failed to add initial chain: %w", err)
	// }

	return nil
}
