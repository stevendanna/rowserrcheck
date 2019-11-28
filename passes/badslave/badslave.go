package badslave

import (
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/buildssa"
	"golang.org/x/tools/go/ssa"
)

func NewAnalyzer(sqlPkgs ...string) *analysis.Analyzer {
	return &analysis.Analyzer{
		Name: "badslave",
		Doc:  Doc,
		Run:  NewRun(),
		Requires: []*analysis.Analyzer{
			buildssa.Analyzer,
		},
	}
}

const (
	Doc = ""
)

type runner struct {
	pass *analysis.Pass
}

func NewRun() func(pass *analysis.Pass) (interface{}, error) {
	return func(pass *analysis.Pass) (interface{}, error) {
		r := &runner{}
		r.run(pass)
		return nil, nil
	}
}

// run executes an analysis for the pass. The receiver is passed
// by value because this func is called in parallel for different passes.
func (r runner) run(pass *analysis.Pass) (interface{}, error) {
	pssa := pass.ResultOf[buildssa.Analyzer].(*buildssa.SSA)
	for _, f := range pssa.SrcFuncs {
		for _, b := range f.Blocks {
			for _, ins := range b.Instrs {
				if call, ok := ins.(*ssa.Call); ok {
					args := call.Call.Args
					for i := range args {
						if i == 0 {
							continue
						}
						if args[i] == args[i-1] {
							pass.Reportf(ins.Pos(), "xx")
						}
						println(args[i].Name(), args[i-1].Name())
					}
				}
			}
		}
		for _, param := range f.Params {
			println(param.Type().String())
		}
	}
	return nil, nil
}

func doBlock(block *ssa.BasicBlock) {

}
