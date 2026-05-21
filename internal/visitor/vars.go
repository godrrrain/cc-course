package visitor

import (
	"fmt"
	"github.com/AskaryanKarine/BMSTU-CC/cource/internal/parser"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

func (v *IRVisitor) VisitVariableDeclaration(ctx *parser.VariableDeclarationContext) interface{} {
	typVal := v.Visit(ctx.TypeSpecifier())
	typ, ok := typVal.(types.Type)
	if !ok || typ == nil {
		v.Errors = append(v.Errors, fmt.Errorf("invalid or missing type in variable declaration"))
		return nil
	}

	varsVal := v.Visit(ctx.VariableList())
	vars, ok := varsVal.([]*VariableInfo)
	if !ok {
		v.Errors = append(v.Errors, fmt.Errorf("invalid variable list in variable declaration"))
		return nil
	}

	for _, vi := range vars {
		if len(vi.Bounds) > 0 {
			t := typ
			for i := len(vi.Bounds) - 1; i >= 0; i-- {
				dim := vi.Bounds[i]
				length := uint64(dim.Upper - dim.Lower + 1)
				t = types.NewArray(length, t)
			}
			vi.Type = t
		} else {
			vi.Type = typ
		}

		initVal := vi.LLVMValue

		alloca := v.currentBlock.NewAlloca(vi.Type)
		vi.LLVMValue = alloca

		if err := v.currentScope.Set(vi.Name, vi); err != nil {
			v.Errors = append(v.Errors, fmt.Errorf("variable %s already declared: %v", vi.Name, err))
		}
		if initVal != nil {
			v.currentBlock.NewStore(initVal, alloca)
		}
	}

	return nil
}

func (v *IRVisitor) VisitVariableList(ctx *parser.VariableListContext) interface{} {
	var vars []*VariableInfo
	for _, itemCtx := range ctx.AllVariableDeclarationItem() {
		res := v.Visit(itemCtx)
		if res == nil {
			continue
		}
		vi, ok := res.(*VariableInfo)
		if !ok {
			v.Errors = append(v.Errors, fmt.Errorf("expected *VariableInfo from VisitVariableDeclarationItem, got %T", res))
			continue
		}
		vars = append(vars, vi)
	}
	return vars
}

func (v *IRVisitor) VisitArrayBounds(ctx *parser.ArrayBoundsContext) interface{} {
	lowIfc := v.Visit(ctx.Expression(0))
	upIfc := v.Visit(ctx.Expression(1))
	lowC, lok := lowIfc.(value.Value)
	upC, uok := upIfc.(value.Value)
	if !lok || !uok {
		v.Errors = append(v.Errors, fmt.Errorf("array bounds are not constant values"))
		return nil
	}
	low, ok1 := constIntValue(lowC)
	up, ok2 := constIntValue(upC)
	if !ok1 || !ok2 {
		v.Errors = append(v.Errors, fmt.Errorf("array bounds must be integer constants"))
		return nil
	}
	return ArrayBounds{Lower: low, Upper: up}
}

func (v *IRVisitor) VisitVariableDeclarationItem(ctx *parser.VariableDeclarationItemContext) interface{} {
	name := NormalizeFunctionName(ctx.ID().GetText())
	vi := &VariableInfo{Name: name}

	if ctx.LBRACK() != nil {
		for _, ab := range ctx.AllArrayBounds() {
			b, ok := v.Visit(ab).(ArrayBounds)
			if !ok {
				v.Errors = append(v.Errors, fmt.Errorf("invalid array bounds in variable declaration"))
				return nil
			}
			vi.Bounds = append(vi.Bounds, b)
		}
	}

	if ctx.EQ() != nil {
		if initVal, ok := v.Visit(ctx.Expression()).(value.Value); ok {
			vi.LLVMValue = initVal
		}
	}
	return vi
}
