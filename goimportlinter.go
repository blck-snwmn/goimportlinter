package goimportlinter

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "goimportlinter is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "goimportlinter",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		// (*ast.ImportSpec)(nil),
		(*ast.File)(nil),
	}

	targetDirsToAllows := map[string][]string{
		"domain":  {"domain"},
		"usecase": {"usecase", "domain"},
		"handler": {"handler", "usecase"},
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.File:
			filePath := pass.Fset.File(n.Pos())

			var allows []string
			for t, as := range targetDirsToAllows {
				if strings.Contains(filePath.Name(), t) {
					allows = as
				}
			}
			if len(allows) == 0 {
				return // no check if file is not target
			}
			for _, im := range n.Imports {
				importName := im.Path.Value
				if !hasOwnModulePrefix(importName) {
					continue
				}
				var ok bool
				for _, allow := range allows {
					if strings.Contains(importName, allow) {
						ok = true
						continue
					}
				}
				if ok {
					return
				}
				pass.Reportf(im.Pos(), "this file can't import %s", im.Path.Value)
			}
		}
	})

	return nil, nil
}
func hasOwnModulePrefix(s string) bool {
	// const prefix = "blck-snwmn"
	// return strings.HasPrefix(s, prefix)
	return true
}
