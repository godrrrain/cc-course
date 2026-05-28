package visitor

import (
	"fmt"
	"strings"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"

	"cc-course/internal/parser"
)

func (v *IRVisitor) defineGlobalString(s string) value.Value {
	if !strings.HasSuffix(s, "\x00") {
		s += "\x00"
	}
	strConst := constant.NewCharArrayFromString(s)
	name := fmt.Sprintf(".str.literal.%d", len(v.Module.Globals))
	globalStr := v.Module.NewGlobalDef(name, strConst)
	globalStr.Immutable = true
	ptr := v.currentBlock.NewGetElementPtr(
		strConst.Typ,
		globalStr,
		constant.NewInt(types.I32, 0),
		constant.NewInt(types.I32, 0),
	)
	return ptr
}

func (v *IRVisitor) VisitArgumentList(ctx *parser.ArgumentListContext) interface{} {
	var args []value.Value
	if ctx.ExpressionList() != nil {
		if exprArgs := v.Visit(ctx.ExpressionList()); exprArgs != nil {
			if vals, ok := exprArgs.([]value.Value); ok {
				args = append(args, vals...)
			}
		}
	}
	return args
}

func (v *IRVisitor) VisitNamedArgument(ctx *parser.NamedArgumentContext) interface{} {
	if ctx.Expression() != nil {
		return v.Visit(ctx.Expression())
	}
	return nil
}

func (v *IRVisitor) VisitElements(ctx *parser.ElementsContext) interface{} {
	var vals []value.Value
	for _, elem := range ctx.AllElement() {
		if val := v.Visit(elem); val != nil {
			if v, ok := val.(value.Value); ok {
				vals = append(vals, v)
			}
		}
	}
	return vals
}

func (v *IRVisitor) VisitElement(ctx *parser.ElementContext) interface{} {
	switch {
	case ctx.ExpressionElement() != nil:
		return v.Visit(ctx.ExpressionElement())
	case ctx.MapElement() != nil:
		return v.Visit(ctx.MapElement())
	case ctx.SpreadElement() != nil:
		return v.Visit(ctx.SpreadElement())
	case ctx.IfElement() != nil:
		return v.Visit(ctx.IfElement())
	case ctx.ForElement() != nil:
		return v.Visit(ctx.ForElement())
	}
	return nil
}

func (v *IRVisitor) VisitExpressionElement(ctx *parser.ExpressionElementContext) interface{} {
	if ctx.Expression() != nil {
		return v.Visit(ctx.Expression())
	}
	return nil
}

func (v *IRVisitor) VisitMapElement(ctx *parser.MapElementContext) interface{} {
	return nil
}

func (v *IRVisitor) VisitSpreadElement(ctx *parser.SpreadElementContext) interface{} {
	if ctx.Expression() != nil {
		return v.Visit(ctx.Expression())
	}
	return nil
}

func (v *IRVisitor) VisitIfElement(ctx *parser.IfElementContext) interface{} {
	return nil
}

func (v *IRVisitor) VisitForElement(ctx *parser.ForElementContext) interface{} {
	return nil
}

func (v *IRVisitor) VisitSymbolLiteral(ctx *parser.SymbolLiteralContext) interface{} {
	v.Errors = append(v.Errors, fmt.Errorf("symbol literals not yet implemented"))
	return nil
}

func stripQuotes(s string) string {
	if len(s) >= 2 {
		if (s[0] == '"' && s[len(s)-1] == '"') || (s[0] == '\'' && s[len(s)-1] == '\'') {
			return s[1 : len(s)-1]
		}
	}
	return s
}

func (v *IRVisitor) castToMatch(lhs, rhs value.Value) (value.Value, value.Value) {
	lt, rt := lhs.Type(), rhs.Type()

	if lt.Equal(rt) {
		return lhs, rhs
	}

	if _, ok := lt.(*types.IntType); ok {
		if _, isFloat := rt.(*types.FloatType); isFloat {
			lhs = v.currentBlock.NewSIToFP(lhs, rt)
			return lhs, rhs
		}
	}
	if _, ok := rt.(*types.IntType); ok {
		if _, isFloat := lt.(*types.FloatType); isFloat {
			rhs = v.currentBlock.NewSIToFP(rhs, lt)
			return lhs, rhs
		}
	}

	return lhs, rhs
}

func constIntValue(val value.Value) (int64, bool) {
	switch v := val.(type) {
	case *constant.Int:
		return v.X.Int64(), true
	case *ir.InstSub:
		if lhs, ok := v.X.(*constant.Int); ok && lhs.X.Sign() == 0 {
			if rhs, ok := v.Y.(*constant.Int); ok {
				return -rhs.X.Int64(), true
			}
		}
	}
	return 0, false
}

func (v *IRVisitor) VisitAssignmentOperator(ctx *parser.AssignmentOperatorContext) interface{} {
	return nil
}

func (v *IRVisitor) VisitAssignableSelectorPart(ctx *parser.AssignableSelectorPartContext) interface{} {
	return nil
}

func (v *IRVisitor) VisitAssignableSelector(ctx *parser.AssignableSelectorContext) interface{} {
	return nil
}

func (v *IRVisitor) VisitUnconditionalAssignableSelector(ctx *parser.UnconditionalAssignableSelectorContext) interface{} {
	return nil
}

func (v *IRVisitor) VisitCascade(ctx *parser.CascadeContext) interface{} {
	return nil
}

func (v *IRVisitor) VisitCascadeSection(ctx *parser.CascadeSectionContext) interface{} {
	return nil
}

func (v *IRVisitor) VisitCascadeSelector(ctx *parser.CascadeSelectorContext) interface{} {
	return nil
}

func (v *IRVisitor) VisitCascadeSectionTail(ctx *parser.CascadeSectionTailContext) interface{} {
	return nil
}

func (v *IRVisitor) VisitCascadeAssignment(ctx *parser.CascadeAssignmentContext) interface{} {
	return nil
}

func (v *IRVisitor) VisitExpressionWithoutCascade(ctx *parser.ExpressionWithoutCascadeContext) interface{} {
	return nil
}

func (v *IRVisitor) VisitThrowExpressionWithoutCascade(ctx *parser.ThrowExpressionWithoutCascadeContext) interface{} {
	return nil
}

func (v *IRVisitor) VisitFunctionExpressionBody(ctx *parser.FunctionExpressionBodyContext) interface{} {
	return nil
}

func (v *IRVisitor) VisitConstructorDesignation(ctx *parser.ConstructorDesignationContext) interface{} {
	return nil
}

func (v *IRVisitor) VisitTypeNotFunction(ctx *parser.TypeNotFunctionContext) interface{} {
	if ctx.VOID_() != nil {
		return types.Void
	}
	if ctx.TypeNotVoidNotFunction() != nil {
		return v.Visit(ctx.TypeNotVoidNotFunction())
	}
	return types.I64
}

func (v *IRVisitor) VisitTypeIdentifier(ctx *parser.TypeIdentifierContext) interface{} {
	return nil
}

func (v *IRVisitor) VisitFunctionTypeTails(ctx *parser.FunctionTypeTailsContext) interface{} {
	return nil
}

func (v *IRVisitor) VisitFunctionTypeTail(ctx *parser.FunctionTypeTailContext) interface{} {
	return nil
}

func (v *IRVisitor) VisitParameterTypeList(ctx *parser.ParameterTypeListContext) interface{} {
	return nil
}

func (v *IRVisitor) VisitNormalParameterTypes(ctx *parser.NormalParameterTypesContext) interface{} {
	return nil
}

func (v *IRVisitor) VisitOptionalParameterTypes(ctx *parser.OptionalParameterTypesContext) interface{} {
	return nil
}

func (v *IRVisitor) VisitOptionalPositionalParameterTypes(ctx *parser.OptionalPositionalParameterTypesContext) interface{} {
	return nil
}

func (v *IRVisitor) VisitNamedParameterTypes(ctx *parser.NamedParameterTypesContext) interface{} {
	return nil
}

func (v *IRVisitor) VisitNamedParameterType(ctx *parser.NamedParameterTypeContext) interface{} {
	return nil
}

func (v *IRVisitor) VisitTypedIdentifier(ctx *parser.TypedIdentifierContext) interface{} {
	return nil
}

func (v *IRVisitor) VisitOptionalOrNamedFormalParameters(ctx *parser.OptionalOrNamedFormalParametersContext) interface{} {
	return nil
}

func (v *IRVisitor) VisitOptionalPositionalFormalParameters(ctx *parser.OptionalPositionalFormalParametersContext) interface{} {
	return nil
}

func (v *IRVisitor) VisitNamedFormalParameters(ctx *parser.NamedFormalParametersContext) interface{} {
	return nil
}

func (v *IRVisitor) VisitFunctionFormalParameter(ctx *parser.FunctionFormalParameterContext) interface{} {
	return nil
}

func (v *IRVisitor) VisitFieldFormalParameter(ctx *parser.FieldFormalParameterContext) interface{} {
	return nil
}

func (v *IRVisitor) VisitFormalParameterPart(ctx *parser.FormalParameterPartContext) interface{} {
	return nil
}

func (v *IRVisitor) VisitDefaultFormalParameter(ctx *parser.DefaultFormalParameterContext) interface{} {
	return nil
}

func (v *IRVisitor) VisitDefaultNamedParameter(ctx *parser.DefaultNamedParameterContext) interface{} {
	return nil
}

func (v *IRVisitor) VisitArgumentPart(ctx *parser.ArgumentPartContext) interface{} {
	if ctx.Arguments() != nil {
		return v.Visit(ctx.Arguments())
	}
	return nil
}

func (v *IRVisitor) VisitArguments(ctx *parser.ArgumentsContext) interface{} {
	if ctx.ArgumentList() != nil {
		return v.Visit(ctx.ArgumentList())
	}
	return nil
}

func (v *IRVisitor) VisitOperator(ctx *parser.OperatorContext) interface{} {
	return nil
}

func (v *IRVisitor) VisitMetadata(ctx *parser.MetadataContext) interface{} {
	return nil
}

func (v *IRVisitor) VisitMetadatum(ctx *parser.MetadatumContext) interface{} {
	return nil
}

func (v *IRVisitor) VisitDottedIdentifier(ctx *parser.DottedIdentifierContext) interface{} {
	return nil
}

func (v *IRVisitor) VisitIdentifierList(ctx *parser.IdentifierListContext) interface{} {
	return nil
}
