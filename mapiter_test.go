package mapiter

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestMapIter(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), Analyzer, "basic")
}

func TestMapIterDefinedMap(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), Analyzer, "definedmap")
}

func TestMapIterTests(t *testing.T) {
	analyzer := Analyzer
	analyzer.Flags.Set("tests", "false")
	analysistest.Run(t, analysistest.TestData(), analyzer, "testsflag")
}
