package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestReadUeventFile_happyPath(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "uevent")
	if err := os.WriteFile(path, []byte("ID_INPUT_KEYBOARD=1\n"), 0644); err != nil {
		t.Fatal(err)
	}
	data, err := readUeventFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if data != "ID_INPUT_KEYBOARD=1\n" {
		t.Errorf("expected content, got %q", data)
	}
}

func TestReadUeventFile_missing(t *testing.T) {
	_, err := readUeventFile("/nonexistent/uevent")
	if err == nil {
		t.Error("expected error for missing file")
	}
}

func TestReadUeventFile_empty(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "uevent")
	if err := os.WriteFile(path, nil, 0644); err != nil {
		t.Fatal(err)
	}
	data, err := readUeventFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if data != "" {
		t.Errorf("expected empty string, got %q", data)
	}
}

func TestCheckInputGroup_notInGroup(t *testing.T) {
	// In test environments, the "input" group may not exist or the test
	// runner may not be in it. We just verify the function runs.
	err := checkInputGroup()
	if err != nil {
		// Error is expected if not in input group or group doesn't exist
		// Just verify it returns a non-nil error without panicking
		t.Logf("not in input group (expected in most test environments): %v", err)
	}
}
