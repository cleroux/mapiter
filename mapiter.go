// mapiter provides a linter that reports code that iterates over a map.
package mapiter

import (
	"go/ast"
	"go/types"
	"strings"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "mapiter",
	Doc:  "reports code that iterates over a map",
	Run:  run,
}

var includeTests bool

func init() {
	Analyzer.Flags.BoolVar(&includeTests, "tests", true, "ignore _test.go files")
}

func run(pass *analysis.Pass) (interface{}, error) {

	for _, file := range pass.Files {

		if !includeTests && strings.HasSuffix(pass.Fset.File(file.Pos()).Name(), "_test.go") {
			continue
		}

		ast.Inspect(file, func(n ast.Node) bool {
			if n == nil {
				return false
			}

			// Look for range expression
			if r, ok := n.(*ast.RangeStmt); ok {

				// Check if range value is a map
				if _, ok := pass.TypesInfo.TypeOf(r.X).(*types.Map); ok {
					pass.Reportf(r.TokPos, "iterating over a map")
					return false
				}
			}

			return true
		})
	}

	return nil, nil
}
