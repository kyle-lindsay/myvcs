package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
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

func createCommit(repoRoot string, message string) (string, error) {
	entries, err := buildSnapshot(repoRoot)
	if err != nil {
		return "", err
	}

	for i, entry := range entries {
		filePath := filepath.Join(repoRoot, entry.Path)

		hash, err := storeBlob(repoRoot, filePath)
		if err != nil {
			return "", err
		}

		entries[i].Hash = hash
	}

	parent, err := readHEAD(repoRoot)
	if err != nil {
		return "", err
	}

	commit := Commit{
		Parent:    parent,
		Message:   message,
		Timestamp: time.Now(),
		Files:     entries,
	}

	id, err := hashCommit(commit)
	if err != nil {
		return "", err
	}
	commit.ID = id

	commitPath := filepath.Join(repoRoot, ".myvcs", "commits", id+".json")

	data, err := json.MarshalIndent(commit, "", "  ")
	if err != nil {
		return "", err
	}

	err = os.MkdirAll(filepath.Dir(commitPath), 0755)
	if err != nil {
		return "", err
	}

	err = os.WriteFile(commitPath, data, 0644)
	if err != nil {
		return "", err
	}

	headPath := filepath.Join(repoRoot, ".myvcs", "HEAD")
	err = os.WriteFile(headPath, []byte(id), 0644)
	if err != nil {
		return "", err
	}

	return id, nil
}

func readHEAD(repoRoot string) (string, error) {
	headPath := filepath.Join(repoRoot, ".myvcs", "HEAD")

	data, err := os.ReadFile(headPath)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(data)), nil
}
