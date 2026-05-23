package visitor

import (
	"github.com/AskaryanKarine/BMSTU-CC/cource/internal/parser"
	"github.com/llir/llvm/ir/types"
)

func (v *IRVisitor) VisitType(ctx *parser.TypeContext) interface{} {
	if ctx.FunctionType() != nil {
		return v.Visit(ctx.FunctionType())
	}
	if ctx.TypeNotFunction() != nil {
		return v.Visit(ctx.TypeNotFunction())
	}
	return types.I64
}

func (v *IRVisitor) VisitTypeNotVoid(ctx *parser.TypeNotVoidContext) interface{} {
	if ctx.FunctionType() != nil {
		return v.Visit(ctx.FunctionType())
	}
	if ctx.TypeNotVoidNotFunction() != nil {
		return v.Visit(ctx.TypeNotVoidNotFunction())
	}
	return types.I64
}

func (v *IRVisitor) VisitTypeNotVoidNotFunction(ctx *parser.TypeNotVoidNotFunctionContext) interface{} {
	if ctx.TypeName() != nil {
		typeName := ctx.TypeName().GetText()
		switch typeName {
		case "int":
			return types.I64
		case "double":
			return types.Double
		case "bool":
			return types.I1
		case "String":
			return types.NewPointer(types.I8)
		case "num":
			return types.Double
		case "dynamic", "var":
			return types.I64
		}
	}
	if ctx.FUNCTION_() != nil {
		return types.NewPointer(types.I8)
	}
	return types.I64
}

func (v *IRVisitor) VisitTypeName(ctx *parser.TypeNameContext) interface{} {
	return ctx.GetText()
}

func (v *IRVisitor) VisitTypeArguments(ctx *parser.TypeArgumentsContext) interface{} {
	if ctx.TypeList() != nil {
		return v.Visit(ctx.TypeList())
	}
	return nil
}

func (v *IRVisitor) VisitTypeList(ctx *parser.TypeListContext) interface{} {
	var typeList []types.Type
	for _, t := range ctx.AllType_() {
		if typ := v.Visit(t); typ != nil {
			if tt, ok := typ.(types.Type); ok {
				typeList = append(typeList, tt)
			}
		}
	}
	return typeList
}

func (v *IRVisitor) VisitFunctionType(ctx *parser.FunctionTypeContext) interface{} {
	return types.NewPointer(types.I8)
}

func (v *IRVisitor) VisitTypeParameters(ctx *parser.TypeParametersContext) interface{} {
	return nil
}

func (v *IRVisitor) VisitTypeParameterList(ctx *parser.TypeParameterListContext) interface{} {
	return nil
}

func (v *IRVisitor) VisitTypeParameter(ctx *parser.TypeParameterContext) interface{} {
	return nil
}
