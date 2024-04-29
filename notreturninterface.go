package notreturninterface

import (
	"go/ast"
	"go/types"
	"strings"

	"github.com/gostaticanalysis/analysisutil"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "notreturninterface is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "notreturninterface",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

var ignoreInterfaces string

func init() {
	Analyzer.Flags.StringVar(&ignoreInterfaces, "ignore", "", "comma-separated list of interfaces to ignore")
}

func run(pass *analysis.Pass) (any, error) {
	ignoreInterfacesSet := make(map[string]struct{})
	for _, ignoreInterface := range strings.Split(strings.TrimSpace(ignoreInterfaces), ",") {
		if ignoreInterface != "" {
			ignoreInterfacesSet[ignoreInterface] = struct{}{}
		}
	}

	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
	}
	pass.Report = analysisutil.ReportWithoutIgnore(pass)
	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.FuncDecl:
			if n.Type.Results == nil || n.Type.Results.List == nil {
				return
			}

			for _, field := range n.Type.Results.List {
				typeExpr := pass.TypesInfo.TypeOf(field.Type)

				if typeExpr == nil {
					continue
				}

				if _, ok := typeExpr.Underlying().(*types.Interface); !ok {
					continue
				}

				if typeExpr.String() == "error" {
					continue
				}

				if _, ok := ignoreInterfacesSet[typeExpr.String()]; ok {
					continue
				}

				pass.Reportf(n.Pos(), "function %s must not return interface %s, but struct", n.Name.Name, typeExpr)
			}
		default:
		}
	})

	return nil, nil
}
