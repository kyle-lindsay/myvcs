package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestBuildSnapshot(t *testing.T) {
	tempDir, err := os.MkdirTemp("", "snapshot-test")

	if err != nil {
		t.Fatalf("failed to create temoporary directory: %v", err)
	}

	defer os.RemoveAll(tempDir)

	filePath := filepath.Join(tempDir, "a.txt")
	err = os.WriteFile(filePath, []byte("hello"), 0644)

	if err != nil {
		t.Fatalf("failed to create file %v: %v", filePath, err)
	}

	subDir := filepath.Join(tempDir, "sub")
	err = os.MkdirAll(subDir, 0755)

	if err != nil {
		t.Fatalf("failed to create directory %v: %v", subDir, err)
	}

	filePath = filepath.Join(tempDir, "sub", "b.txt")
	err = os.WriteFile(filePath, []byte("goodbye"), 0644)

	if err != nil {
		t.Fatalf("failed to create file %v: %v", filePath, err)
	}

	err = os.MkdirAll(filepath.Join(tempDir, ".myvcs"), 0755)

	if err != nil {
		t.Fatalf("failed to create directory .myvcs: %v", err)
	}

	err = os.WriteFile(filepath.Join(tempDir, ".myvcs", "HEAD"), []byte(""), 0644)

	if err != nil {
		t.Fatalf("failed to create file .myvcs/HEAD: %v", err)
	}

	entries, err := buildSnapshot(tempDir)
	if err != nil {
		t.Fatalf("buildSnapshot failed: %v", err)
	}

	paths := map[string]bool{}

	for _, entry := range entries {
		paths[entry.Path] = true
	}

	if !paths["a.txt"] {
		t.Errorf("expected a.txt to be present")
	}

	if !paths["sub/b.txt"] {
		t.Errorf("expected b.txt to be present")
	}

	if paths[".myvcs/HEAD"] {
		t.Errorf("did not expect .myvcs/HEAD to be present")
	}

}
