package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func main() {
	argsLength := len(os.Args)

	if argsLength == 1 {
		fmt.Println("No parameters were provided")
		return
	}

	command := os.Args[1]

	switch command {
	case "init":
		if argsLength == 2 {
			if err := initialise(); err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("init requires no parameters")
			return
		}
	default:
		fmt.Println("Unknown command " + command)
		return
	}
}

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

func buildSnapshot(root string) ([]FileEntry, error) {
	entries := []FileEntry{}

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() && d.Name() == ".myvcs" {
			return filepath.SkipDir
		}

		if d.IsDir() {
			return nil
		}

		relPath, err := filepath.Rel(root, path)

		if err != nil {
			return err
		}

		fileHash, err := hashFile(path)

		if err != nil {
			return err
		}

		entry := FileEntry{
			Path: relPath,
			Hash: fileHash,
		}

		entries = append(entries, entry)

		return nil
	})

	if err != nil {
		return nil, err
	}

	return entries, nil
}

func hashFile(path string) (string, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	sum := sha256.Sum256(file)
	hash := hex.EncodeToString(sum[:])

	return hash, nil
}
