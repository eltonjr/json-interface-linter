package main

import (
	"errors"
	"reflect"
	"strconv"
	"testing"
)

func TestReadMarshalers(t *testing.T) {
	expectedMarshalers := []marshaler{
		{"encode/json.Marshal", 0},
		{"encode/json.Encode", 0},
		{"github.com/gin-gonic/gin.JSON", 0},
		{"myencoder.Encode", 1},
		{"f", 0},
		{"x.y", 1},
	}

	marshalers, err := ReadMarshalers("testdata/valid.txt")
	if err != nil {
		t.Fatalf("Failed to read marshalers: %s", err)
	}

	if !reflect.DeepEqual(marshalers, expectedMarshalers) {
		t.Errorf("Expected marshalers %v, got %v", expectedMarshalers, marshalers)
	}
}

func TestEmptyFile(t *testing.T) {
	var expectedMarshalers []marshaler

	marshalers, err := ReadMarshalers("testdata/empty.txt")
	if err != nil {
		t.Fatalf("Failed to read marshalers: %s", err)
	}

	if !reflect.DeepEqual(marshalers, expectedMarshalers) {
		t.Errorf("Expected marshalers %v, got %v", expectedMarshalers, marshalers)
	}
}

func TestInvalidBracket(t *testing.T) {
	_, err := ReadMarshalers("testdata/invalid_bracket.txt")
	if err == nil {
		t.Fatal("Invalid file should fail")
	}

	if err != ErrMissingClosingBracket {
		t.Errorf("Expected error %v, got %v", ErrMissingClosingBracket, err)
	}
}

func TestInvalidInt(t *testing.T) {
	_, err := ReadMarshalers("testdata/invalid_int.txt")
	if err == nil {
		t.Fatal("Invalid file should fail")
	}

	if !errors.Is(err, strconv.ErrSyntax) {
		t.Errorf("Expected error %v, got %v", strconv.ErrSyntax, err)
	}
}

func TestUseDefaultIfEmptyString(t *testing.T) {
	expectedMarshalers := defaultMarshalers

	marshalers, err := ReadMarshalers("")
	if err != nil {
		t.Fatalf("Failed to read marshalers: %s", err)
	}

	if !reflect.DeepEqual(marshalers, expectedMarshalers) {
		t.Errorf("Expected marshalers %v, got %v", expectedMarshalers, marshalers)
	}
}
