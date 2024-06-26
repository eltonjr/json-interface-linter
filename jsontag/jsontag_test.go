package jsontag

import (
	"flag"
	"os"
	"path/filepath"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

var defaultFlags = flag.FlagSet{}

func TestJSONTag(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get wd: %s", err)
	}

	testdata := filepath.Join(wd, "testdata")
	analysistest.Run(t, testdata, Analyzer(defaultFlags), "jsontag")
}
