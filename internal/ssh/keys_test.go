package ssh

import (
	"os"
	"path/filepath"
	"testing"
)

/**
 * Tests host key generation and persistence.
 */
func TestLoadOrGenerateHostKey(t *testing.T) {
	// Use temp directory
	tempDir := t.TempDir()
	keyPath := filepath.Join(tempDir, "test_host_key")

	// First call: should generate new key
	signer1, err := LoadOrGenerateHostKey(keyPath)
	if err != nil {
		t.Fatalf("Failed to generate key: %v", err)
	}
	if signer1 == nil {
		t.Fatal("Signer should not be nil")
	}

	// Verify file was created
	if _, err := os.Stat(keyPath); os.IsNotExist(err) {
		t.Error("Key file was not created")
	}

	// Second call: should load existing key
	signer2, err := LoadOrGenerateHostKey(keyPath)
	if err != nil {
		t.Fatalf("Failed to load key: %v", err)
	}

	// Both signers should produce the same public key
	if string(signer1.PublicKey().Marshal()) != string(signer2.PublicKey().Marshal()) {
		t.Error("Loaded key doesn't match generated key")
	}
}
