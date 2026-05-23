package visitor

import (
	"fmt"

	"github.com/AskaryanKarine/BMSTU-CC/cource/internal/parser"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

func (v *IRVisitor) VisitVariableDeclarationList(ctx *parser.VariableDeclarationListContext) interface{} {
	for _, item := range ctx.AllVariableDeclarationItem() {
		v.Visit(item)
	}
	return nil
}

func (v *IRVisitor) VisitVariableDeclarationItem(ctx *parser.VariableDeclarationItemContext) interface{} {
	if ctx.IDENTIFIER() == nil {
		return nil
	}

	name := ctx.IDENTIFIER().GetText()
	var typ types.Type = types.I64

	alloca := v.currentBlock.NewAlloca(typ)
	vi := &VariableInfo{
		Name:      name,
		Type:      typ,
		LLVMValue: alloca,
	}

	if err := v.currentScope.Set(name, vi); err != nil {
		v.Errors = append(v.Errors, fmt.Errorf("variable %s already declared: %v", name, err))
	}

	if ctx.Expression() != nil {
		if initVal, ok := v.Visit(ctx.Expression()).(value.Value); ok {
			v.currentBlock.NewStore(initVal, alloca)
		}
	}

	return vi
}
