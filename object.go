package main

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
)

func hashFile(path string) (string, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	sum := sha256.Sum256(file)
	hash := hex.EncodeToString(sum[:])

	return hash, nil
}

func storeBlob(repoRoot string, filePath string) (string, error) {
	return "", nil
}
