package loglinter

import (
	"go/ast"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
	"github.com/ADM307/loglinter/pkg/loglinter/rules"
	"strings"
)

const doc = "loglinter: simple linter for log messages compatible with log/slog and go.uber.org/zap"

var Analyzer = &analysis.Analyzer{
	Name:     "loglinter",
	Doc:      doc,
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

var allowedLogFuncs = map[string]bool{
	"Info": true, "Debug": true, "Warn": true, "Error": true,
	"DPanic": true, "Panic": true, "Fatal": true,
	"Infow": true, "Debugw": true, "Warnw": true, "Errorw": true,
	"Dpanicw": true, "Panicw": true, "Fatalw": true,
}

var targetPackages = map[string]bool{
	"log/slog":        true,
	"go.uber.org/zap": true,
}

func run(pass *analysis.Pass) (interface{}, error) {
	insp := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.CallExpr)(nil),
	}

	insp.Preorder(nodeFilter, func(n ast.Node) {
		call := n.(*ast.CallExpr)
		checkCall(pass, call)
	})

	return nil, nil
}

func checkCall(pass *analysis.Pass, call *ast.CallExpr) {
	obj, ok := getFuncObject(pass, call)
	if !ok {
		return
	}

	if !isLogFunc(pass, obj) {
		return
	}

	if len(call.Args) == 0 {
		return
	}

	lit, ok := call.Args[0].(*ast.BasicLit)
	if !ok || lit.Kind != token.STRING {
		return
	}

	msg := strings.Trim(lit.Value, "\"")

	pos := lit.Pos()

	for _, rule := range rules.All() {
		rule(pass, pos, msg)
	}
}

func getFuncObject(pass *analysis.Pass, call *ast.CallExpr) (types.Object, bool) {
	var obj types.Object

	switch fun := call.Fun.(type) {
	case *ast.Ident:
		obj = pass.TypesInfo.Uses[fun]
	case *ast.SelectorExpr:
		obj = pass.TypesInfo.Uses[fun.Sel]
	default:
		return nil, false
	}

	return obj, obj != nil
}

func isLogFunc(pass *analysis.Pass, obj types.Object) bool {
	funcObj, ok := obj.(*types.Func)
	if !ok {
		return false
	}

	pkg := funcObj.Pkg()
	if pkg == nil {
		return false
	}

	if !targetPackages[pkg.Path()] {
		return false
	}

	name := funcObj.Name()
	return allowedLogFuncs[name]
}
