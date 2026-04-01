package main

import (
	"os"
	"testing"
)

func TestHashFile(t *testing.T) {
	file, err := os.CreateTemp("", "testfile")

	if err != nil {
		t.Fatalf("could not create file: %v", err)
	}

	defer os.Remove(file.Name())

	_, err = file.Write([]byte("hello world"))
	if err != nil {
		t.Fatalf("could not write to file: %v", err)
	}

	if err := file.Close(); err != nil {
		t.Fatalf("failed to close file: %v", err)
	}

	hash, err := hashFile(file.Name())
	if err != nil {
		t.Fatalf("could not hash file: %v", err)
	}

	expected := "b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9"

	if hash != expected {
		t.Errorf("expected %s, got %s", expected, hash)
	}
}
