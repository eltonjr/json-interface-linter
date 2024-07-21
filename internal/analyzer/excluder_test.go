package analyzer

import (
	"reflect"
	"testing"
)

func TestReadExclude(t *testing.T) {
	d, err := ReadExcluders("testdata/valid.txt")
	if err != nil {
		t.Fatalf("Failed to read excluders: %s", err)
	}

	expected := []string{
		"errors",
		"myapp.MyInterface",
		"interface{}",
		"any",
	}

	if len(d) != len(expected) {
		t.Fatalf("Expected %d excluders, got %d", len(expected), len(d))
	}

	if !reflect.DeepEqual(expected, d) {
		t.Errorf("Expected marshalers %v, got %v", expected, d)
	}
}

func TestEmptyFile(t *testing.T) {
	d, err := ReadExcluders("testdata/empty.txt")
	if err != nil {
		t.Fatalf("Failed to read excluders: %s", err)
	}

	if len(d) != 0 {
		t.Fatalf("Expected 0 excluders, got %d", len(d))
	}
}
