package visitor

import (
	"fmt"
	"github.com/AskaryanKarine/BMSTU-CC/cource/internal/parser"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"strconv"
	"strings"
)

func (v *IRVisitor) VisitLiteral(ctx *parser.LiteralContext) interface{} {
	switch {
	case ctx.INTEGER() != nil:
		text := ctx.INTEGER().GetText()
		val, err := strconv.ParseInt(text, 10, 64)
		if err != nil {
			v.Errors = append(v.Errors, fmt.Errorf("invalid integer literal: %s", text))
			return nil
		}
		return constant.NewInt(types.I32, val)

	case ctx.REAL() != nil:
		text := ctx.REAL().GetText()
		val, err := strconv.ParseFloat(text, 64)
		if err != nil {
			v.Errors = append(v.Errors, fmt.Errorf("invalid real literal: %s", text))
			return nil
		}
		return constant.NewFloat(types.Double, val)

	case ctx.STRING() != nil:
		text := ctx.STRING().GetText()
		unquoted := stripQuotes(text)
		return v.defineGlobalString(unquoted)

	case ctx.CHAR_LITERAL() != nil:
		text := ctx.CHAR_LITERAL().GetText()
		unquoted := stripQuotes(text)
		if len(unquoted) != 1 {
			v.Errors = append(v.Errors, fmt.Errorf("invalid char literal: %s", text))
			return nil
		}
		return constant.NewInt(types.I8, int64(unquoted[0]))

	case ctx.TRUE() != nil:
		return constant.NewInt(types.I1, 1)

	case ctx.FALSE() != nil:
		return constant.NewInt(types.I1, 0)

	case ctx.NEWLINE_CONST() != nil:
		return v.defineGlobalString("\n")

	default:
		v.Errors = append(v.Errors, fmt.Errorf("unknown literal"))
		return nil
	}
}

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

func (v *IRVisitor) VisitQualifiedIdentifier(ctx *parser.QualifiedIdentifierContext) interface{} {
	codeName := ctx.ID().GetText()
	llvmName := NormalizeFunctionName(codeName)

	if vi, ok := v.currentScope.Get(llvmName); ok {
		return vi
	}

	if fn, ok := v.funcs[llvmName]; ok {
		return fn
	}

	v.Errors = append(v.Errors, fmt.Errorf("undefined identifier: %s", codeName))
	return nil
}

func (v *IRVisitor) VisitArgumentList(ctx *parser.ArgumentListContext) interface{} {
	var args []value.Value
	for _, expr := range ctx.AllExpression() {
		valIfc := v.Visit(expr)
		if valIfc == nil {
			continue
		}

		switch val := valIfc.(type) {
		case *VariableInfo:
			if ptrTy, ok := val.LLVMValue.Type().(*types.PointerType); ok {
				val := v.currentBlock.NewLoad(ptrTy.ElemType, val.LLVMValue)
				args = append(args, val)
			} else {
				args = append(args, val.LLVMValue)
			}

		case value.Value:
			if arrTy, ok := val.Type().(*types.ArrayType); ok {
				ptr := v.currentBlock.NewAlloca(arrTy)
				v.currentBlock.NewStore(val, ptr)
				args = append(args, ptr)
			} else {
				args = append(args, val)
			}

		default:
			v.Errors = append(v.Errors, fmt.Errorf("invalid argument expression"))
			return nil
		}
	}
	return args
}
