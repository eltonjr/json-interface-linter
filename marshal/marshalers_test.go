package marshal

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
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

func BenchmarkParseLine(b *testing.B) {
	buf, err := os.ReadFile("testdata/valid.txt")
	if err != nil {
		b.Fatalf("Failed to read marshalers: %s", err)
		return
	}

	b.ResetTimer()
	lastline := ""
	for i := 0; i < b.N; i++ {
		scanner := bufio.NewScanner(bytes.NewReader(buf))

		for scanner.Scan() {
			line := scanner.Bytes()
			m, _ := parseLine(line)
			lastline = m.functionPath
		}
	}
	fmt.Println(lastline)
}

func TestParseLine(t *testing.T) {
	buf, err := os.ReadFile("testdata/valid.txt")
	if err != nil {
		t.Fatalf("Failed to read marshalers: %s", err)
		return
	}

	scanner := bufio.NewScanner(bytes.NewReader(buf))

	for scanner.Scan() {
		line := scanner.Bytes()
		parseLine(line)
	}
}
