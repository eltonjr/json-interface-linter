package marshal

import (
	"flag"
	"os"
	"path/filepath"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

var defaultFlags = flag.FlagSet{}

func TestMarshal(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get wd: %s", err)
	}

	testdata := filepath.Join(wd, "testdata")
	ma, _ := Analyzer(defaultFlags)
	analysistest.Run(t, testdata, ma, "marshal")
}

func TestMarshalCustom(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get wd: %s", err)
	}

	testdata := filepath.Join(wd, "testdata")
	fs := flag.NewFlagSet("marshal", flag.ExitOnError)
	fs.String("marshalers", filepath.Join(wd, "testdata/src/marshalcustom/marshalers.txt"), "path to marshalers file")
	ma, err := Analyzer(*fs)
	if err != nil {
		t.Fatalf("Failed to create analyzer: %s", err)
	}

	analysistest.Run(t, testdata, ma, "marshalcustom")
}
