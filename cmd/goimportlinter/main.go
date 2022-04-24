package main

import (
	"github/blck-snwmn/goimportlinter"

	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(goimportlinter.Analyzer) }
