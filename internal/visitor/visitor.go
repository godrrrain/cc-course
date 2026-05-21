package visitor

import (
	"github.com/AskaryanKarine/BMSTU-CC/cource/internal/parser"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"runtime"

	"github.com/antlr4-go/antlr/v4"
)

type IRVisitor struct {
	*parser.BaseKumirParserVisitor

	Errors []error
	Module *ir.Module

	currentFunc      *ir.Func
	currentScope     *Scope
	currentBlock     *ir.Block
	blockStack       []*ir.Block
	funcs            map[string]*ir.Func
	mainOriginalName string
}

func NewIRVisitor() *IRVisitor {
	module := ir.NewModule()
	module.TargetTriple = runtimeLLVMTriple()
	return &IRVisitor{
		BaseKumirParserVisitor: &parser.BaseKumirParserVisitor{},

		Errors: make([]error, 0),
		Module: module,

		currentFunc:  nil,
		currentScope: nil,
		funcs:        make(map[string]*ir.Func),
	}
}

func runtimeLLVMTriple() string {
	arch := map[string]string{"arm64": "arm64", "amd64": "x86_64"}[runtime.GOARCH]
	os := map[string]string{"darwin": "apple-macosx14.0.0", "linux": "pc-linux-gnu"}[runtime.GOOS]
	return arch + "-" + os
}

func (v *IRVisitor) GetModule() *ir.Module {
	return v.Module
}

func (v *IRVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *IRVisitor) VisitProgram(ctx *parser.ProgramContext) interface{} {
	v.enterScope()
	defer v.exitScope()

	for _, item := range ctx.AllProgramItem() {
		v.Visit(item)
	}

	for _, item := range ctx.AllAlgorithmDefinition() {
		v.declareAlgorithmPrototype(item.AlgorithmHeader().(*parser.AlgorithmHeaderContext))
	}

	for _, item := range ctx.AllAlgorithmDefinition() {
		v.Visit(item)
	}

	return nil
}

func (v *IRVisitor) VisitExitStatement(ctx *parser.ExitStatementContext) interface{} {
	if m, ok := v.currentScope.getMeta(loopExitVar); ok {
		if blk, ok2 := m.(*ir.Block); ok2 {
			if v.currentBlock.Term == nil {
				v.currentBlock.NewBr(blk)
			}
			unreachable := v.currentFunc.NewBlock(v.freshLabel("unreachable"))
			v.currentBlock = unreachable
			return nil
		}
	}

	if v.currentFunc != nil && v.currentBlock.Term == nil {
		if !types.Equal(v.currentFunc.Sig.RetType, types.Void) {
			zero := constant.NewInt(v.currentFunc.Sig.RetType.(*types.IntType), 0)
			v.currentBlock.NewRet(zero)
		} else {
			v.currentBlock.NewRet(nil)
		}
	}
	return nil
}
