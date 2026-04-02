package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"os"
	"path/filepath"
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
	hash, err := hashFile(filePath)

	if err != nil {
		return "", err
	}

	objectPath := filepath.Join(repoRoot, ".myvcs", "objects", hash)

	_, err = os.Stat(objectPath)
	if err == nil {
		return hash, nil
	}

	if !os.IsNotExist(err) {
		return "", err
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	err = os.MkdirAll(filepath.Dir(objectPath), 0755)
	if err != nil {
		return "", err
	}

	err = os.WriteFile(objectPath, data, 0644)
	if err != nil {
		return "", err
	}

	return hash, nil
}

func readBlob(repoRoot string, hash string) ([]byte, error) {
	filePath := filepath.Join(repoRoot, ".myvcs", "objects", hash)

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func hashCommit(commit Commit) (string, error) {
	commitCopy := commit
	commitCopy.ID = ""

	data, err := json.Marshal(commitCopy)
	if err != nil {
		return "", err
	}

	sum := sha256.Sum256(data)
	return hex.EncodeToString(sum[:]), nil
}
