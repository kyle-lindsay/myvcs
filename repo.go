package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
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

func readCommit(repoRoot string, id string) (Commit, error) {
	filePath := filepath.Join(repoRoot, ".myvcs", "commits", id+".json")

	data, err := os.ReadFile(filePath)
	if err != nil {
		return Commit{}, err
	}

	var contents Commit

	err = json.Unmarshal(data, &contents)
	if err != nil {
		return Commit{}, err
	}

	return contents, nil
}

func logCommits(repoRoot string) error {
	head, err := readHEAD(repoRoot)
	if err != nil {
		return err
	}

	if head == "" {
		fmt.Println("No commits yet")
		return nil
	}

	current := head

	for current != "" {
		commit, err := readCommit(repoRoot, current)
		if err != nil {
			return err
		}

		fmt.Println("commit", commit.ID)
		fmt.Println("Date:", commit.Timestamp)
		fmt.Println("Message:", commit.Message)
		fmt.Println()

		current = commit.Parent
	}

	return nil
}

func checkoutCommit(repoRoot string, id string) error {
	commit, err := readCommit(repoRoot, id)
	if err != nil {
		return err
	}

	if err = clearWorkingTree(repoRoot); err != nil {
		return err
	}

	if err = restoreWorkingTree(repoRoot, commit.Files); err != nil {
		return err
	}

	headFilePath := filepath.Join(repoRoot, ".myvcs", "HEAD")
	if err = os.WriteFile(headFilePath, []byte(id), 0644); err != nil {
		return err
	}

	return nil
}

func clearWorkingTree(repoRoot string) error {
	err := filepath.WalkDir(repoRoot, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() && d.Name() == ".myvcs" {
			return filepath.SkipDir
		}

		if d.IsDir() {
			return nil
		}

		if err := os.Remove(path); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func restoreWorkingTree(repoRoot string, files []FileEntry) error {
	for _, entry := range files {
		data, err := readBlob(repoRoot, entry.Hash)
		if err != nil {
			return err
		}

		fullPath := filepath.Join(repoRoot, entry.Path)

		if err = os.MkdirAll(filepath.Dir(fullPath), 0755); err != nil {
			return err
		}

		if err = os.WriteFile(fullPath, data, 0644); err != nil {
			return err
		}
	}

	return nil
}
