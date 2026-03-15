package loglinter_test

import (
	"golang.org/x/tools/go/analysis/analysistest"
	"github.com/ADM307/loglinter/pkg/loglinter"
	"testing"
)

func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, loglinter.Analyzer, "basic")
}
