package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func initialise() error {
	_, err := os.Stat(".myvcs")
	if err == nil {
		return fmt.Errorf("repository already exists in this directory")
	}
	if !os.IsNotExist(err) {
		return err
	}

	if err := os.MkdirAll(".myvcs", 0755); err != nil {
		return err
	}

	if err := os.MkdirAll(".myvcs/objects", 0755); err != nil {
		return err
	}

	if err := os.MkdirAll(".myvcs/commits", 0755); err != nil {
		return err
	}

	if err := os.WriteFile(".myvcs/HEAD", []byte(""), 0644); err != nil {
		return err
	}

	return nil
}

func createCommit(repoRoot string, message string) error {
	entries, err := buildSnapshot(repoRoot)
	if err != nil {
		return err
	}

	for i, entry := range entries {
		filePath := filepath.Join(repoRoot, entry.Path)

		hash, err := storeBlob(repoRoot, filePath)
		if err != nil {
			return err
		}

		entries[i].Hash = hash

	}

	return nil
}

func readHEAD(repoRoot string) (string, error) {

}
