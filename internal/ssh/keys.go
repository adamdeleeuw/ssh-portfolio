package ssh

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/pem"
	"os"
	"path/filepath"

	"github.com/gliderlabs/ssh"
	gossh "golang.org/x/crypto/ssh"
)

/**
 * Loads existing Ed25519 host key or generates a new one if not found.
 *
 * @param path - Filesystem path to store/load the key
 * @return SSH signer for the host key
 * @return error if key operations fail
 * @effects Creates key file on disk if it doesn't exist
 */
func LoadOrGenerateHostKey(path string) (ssh.Signer, error) {
	// Try to load existing key
	if _, err := os.Stat(path); err == nil {
		keyData, err := os.ReadFile(path)
		if err != nil {
			return nil, err
		}

		signer, err := gossh.ParsePrivateKey(keyData)
		if err != nil {
			return nil, err
		}

		return signer, nil
	}

	// Generate new Ed25519 key
	_, privKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, err
	}

	// Convert to SSH format
	signer, err := gossh.NewSignerFromKey(privKey)
	if err != nil {
		return nil, err
	}

	// Ensure directory exists
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, err
	}

	// Save private key in OpenSSH format
	pemBlock, err := gossh.MarshalPrivateKey(privKey, "")
	if err != nil {
		return nil, err
	}

	keyFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return nil, err
	}
	defer keyFile.Close()

	if err := pem.Encode(keyFile, pemBlock); err != nil {
		return nil, err
	}

	return signer, nil
}
