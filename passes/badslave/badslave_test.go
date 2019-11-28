package badslave_test

import (
	"testing"

	"github.com/jingyugao/rowserrcheck/passes/badslave"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, badslave.NewAnalyzer(), "a")
}
