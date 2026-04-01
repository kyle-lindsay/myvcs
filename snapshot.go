package main

import (
	"io/fs"
	"path/filepath"
)

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
