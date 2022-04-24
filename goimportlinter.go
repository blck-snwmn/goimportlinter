package goimportlinter

import (
	"fmt"
	"go/ast"

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
	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.File:
			fmt.Println("file is ", n.Name.Name)
			fmt.Println("file path is ", pass.Fset.File(n.Pos()))
			for _, im := range n.Imports {
				fmt.Println("path", im.Path.Value)
			}
			pass.Reportf(n.Pos(), fmt.Sprintf("%+v", n.Name))
		}
	})

	return nil, nil
}
