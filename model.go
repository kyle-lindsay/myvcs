package main

import "time"

type FileEntry struct {
	Path string `json:"path"`
	Hash string `json:"hash"`
}

type Commit struct {
	ID        string      `json:"id"`
	Parent    string      `json:"parent"`
	Message   string      `json:"message"`
	Timestamp time.Time   `json:"timestamp"`
	Files     []FileEntry `json:"files"`
}
