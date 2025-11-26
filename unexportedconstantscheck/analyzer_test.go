package unexportedconstantscheck

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testCases := []struct {
		desc     string
		patterns string
	}{
		{
			desc:     "default",
			patterns: "simple",
		},
		{
			desc:     "complex",
			patterns: "complex",
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			a := NewAnalyzer()

			analysistest.Run(t, analysistest.TestData(), a, test.patterns)
		})
	}
}

func TestAnalyzerWithSuggestedFixes(t *testing.T) {
	testCases := []struct {
		desc     string
		patterns string
	}{
		{
			desc:     "default",
			patterns: "simple-fix",
		},
		{
			desc:     "complex",
			patterns: "complex-fix",
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			a := NewAnalyzer()

			analysistest.RunWithSuggestedFixes(t, analysistest.TestData(), a, test.patterns)
		})
	}
}
