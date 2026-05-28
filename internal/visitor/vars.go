package visitor

import (
	"fmt"

	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"

	"cc-course/internal/parser"
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

	var initVal value.Value
	if ctx.Expression() != nil {
		if val, ok := v.Visit(ctx.Expression()).(value.Value); ok {
			initVal = val
		}
	}

	var typ types.Type = types.I64
	if initVal != nil {
		typ = initVal.Type()
	}

	alloca := v.currentBlock.NewAlloca(typ)
	vi := &VariableInfo{
		Name:      name,
		Type:      typ,
		LLVMValue: alloca,
	}

	if err := v.currentScope.Set(name, vi); err != nil {
		v.Errors = append(v.Errors, fmt.Errorf("variable %s already declared: %v", name, err))
	}

	if initVal != nil {
		v.currentBlock.NewStore(initVal, alloca)
	}

	return vi
}
