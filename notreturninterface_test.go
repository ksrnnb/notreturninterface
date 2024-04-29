package notreturninterface_test

import (
	"testing"

	"github.com/gostaticanalysis/testutil"
	"github.com/ksrnnb/notreturninterface"
	"golang.org/x/tools/go/analysis/analysistest"
)

func init() {
	notreturninterface.Analyzer.Flags.Set("ignore", "any")
}

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	analysistest.Run(t, testdata, notreturninterface.Analyzer, "a")
}
