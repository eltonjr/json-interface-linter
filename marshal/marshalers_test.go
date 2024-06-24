package marshal

import (
	"reflect"
	"testing"
)

func TestReadMarshalers(t *testing.T) {
	expectedMarshalers := []marshaler{
		{"encode/json.Marshal", 0},
		{"encode/json.Encode", 0},
		{"github.com/gin-gonic/gin.JSON", 0},
		{"myencoder.Encode", 1},
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
