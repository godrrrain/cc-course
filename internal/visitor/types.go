package visitor

import (
	"fmt"
	"github.com/AskaryanKarine/BMSTU-CC/cource/internal/parser"
	"github.com/llir/llvm/ir/types"
)

func (v *IRVisitor) VisitTypeSpecifier(ctx *parser.TypeSpecifierContext) interface{} {
	if ctx.ArrayType() != nil {
		typ := v.Visit(ctx.ArrayType())
		if t, ok := typ.(types.Type); ok {
			return t
		}
		v.Errors = append(v.Errors, fmt.Errorf("invalid array type: %v", ctx.GetText()))
		return types.I32
	}

	baseVal := v.Visit(ctx.BasicType())
	base, ok := baseVal.(types.Type)
	if !ok {
		v.Errors = append(v.Errors, fmt.Errorf("invalid basic type: %v", ctx.GetText()))
		base = types.I32
	}
	if ctx.TABLE_SUFFIX() != nil {
		return types.NewPointer(base)
	}
	return base
}

func (v *IRVisitor) VisitBasicType(ctx *parser.BasicTypeContext) interface{} {
	if ctx.INTEGER_TYPE() != nil {
		return types.I32
	} else if ctx.REAL_TYPE() != nil {
		return types.Double
	} else if ctx.BOOLEAN_TYPE() != nil {
		return types.I1
	} else if ctx.CHAR_TYPE() != nil {
		return types.I8
	} else if ctx.STRING_TYPE() != nil {
		return types.NewPointer(types.I8)
	}
	v.Errors = append(v.Errors, fmt.Errorf("unknown basic type: %s", ctx.GetText()))
	return types.I32
}

func (v *IRVisitor) VisitArrayType(ctx *parser.ArrayTypeContext) interface{} {
	var elementType types.Type

	if ctx.INTEGER_ARRAY_TYPE() != nil {
		elementType = types.I32
	} else if ctx.REAL_ARRAY_TYPE() != nil {
		elementType = types.Double
	} else if ctx.BOOLEAN_ARRAY_TYPE() != nil {
		elementType = types.I1
	} else if ctx.CHAR_ARRAY_TYPE() != nil {
		elementType = types.I8
	} else if ctx.STRING_ARRAY_TYPE() != nil {
		elementType = types.NewPointer(types.I8)
	} else {
		v.Errors = append(v.Errors, fmt.Errorf("unknown array type %s", ctx.GetText()))
		return types.Void
	}

	return elementType
}
