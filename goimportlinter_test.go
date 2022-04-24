package goimportlinter_test

import (
	"testing"

	"github/blck-snwmn/goimportlinter"

	"github.com/gostaticanalysis/testutil"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	analysistest.Run(t, testdata, goimportlinter.Analyzer,
		"a/domain",
		"a/handler",
	)
}
