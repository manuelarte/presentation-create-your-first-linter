package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/manuelarte/presentation-create-your-first-linter/unexportedconstantscheck"
)

func main() {
	singlechecker.Main(unexportedconstantscheck.NewAnalyzer())
}
