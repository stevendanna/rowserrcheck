package rowserr_test

import (
	"testing"

	"github.com/jingyugao/rowserr/passes/rowserr"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, rowserr.Analyzer, "a")
}
