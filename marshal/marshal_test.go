package marshal

import (
	"os"
	"path/filepath"
	"sync"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestMarshal(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get wd: %s", err)
	}

	testdata := filepath.Join(wd, "testdata")
	analysistest.Run(t, testdata, Analyzer, "marshal")
}

func TestMarshalCustom(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get wd: %s", err)
	}

	testdata := filepath.Join(wd, "testdata")
	marshalerspath = "testdata/src/marshalcustom/marshalers.txt"

	// reset once so it loads the marshalers from the new path
	once = sync.Once{}
	analysistest.Run(t, testdata, Analyzer, "marshalcustom")
}
