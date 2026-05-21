package visitor

import (
	"fmt"
	"github.com/AskaryanKarine/BMSTU-CC/cource/internal/parser"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"strings"
)

func (v *IRVisitor) ensurePrintfFunc() *ir.Func {
	if fn, ok := v.funcs["printf"]; ok {
		return fn
	}
	printf := v.Module.NewFunc(NormalizeFunctionName("printf"), types.I32,
		ir.NewParam("format", types.NewPointer(types.I8)))
	printf.Sig.Variadic = true
	v.funcs["printf"] = printf
	return printf
}

func (v *IRVisitor) buildPrintfFormat(args []value.Value) (value.Value, []value.Value) {
	var formatParts []string
	var formatArgs []value.Value

	for _, arg := range args {
		switch arg.Type().String() {
		case "i32":
			formatParts = append(formatParts, "%d")
			formatArgs = append(formatArgs, arg)
		case "double":
			formatParts = append(formatParts, "%f")
			formatArgs = append(formatArgs, arg)
		case "i1":
			ptrYes := v.defineGlobalString("да")
			ptrNo := v.defineGlobalString("нет")
			strPtr := v.currentBlock.NewSelect(arg, ptrYes, ptrNo)

			formatParts = append(formatParts, "%s")
			formatArgs = append(formatArgs, strPtr)
		case "i8*":
			formatParts = append(formatParts, "%s")
			formatArgs = append(formatArgs, arg)
		default:
			v.Errors = append(v.Errors, fmt.Errorf("unsupported type in OUTPUT: %s", arg.Type().String()))
		}
	}

	formatStr := strings.Join(formatParts, " ") + "\x00"
	strConst := constant.NewCharArrayFromString(formatStr)
	name := fmt.Sprintf(".fmt.str.%d", len(v.Module.Globals))
	globalFormat := v.Module.NewGlobalDef(name, strConst)
	globalFormat.Immutable = true

	formatPtr := v.currentBlock.NewGetElementPtr(
		strConst.Typ,
		globalFormat,
		constant.NewInt(types.I32, 0),
		constant.NewInt(types.I32, 0),
	)
	return formatPtr, formatArgs
}

func (v *IRVisitor) VisitIoArgumentList(ctx *parser.IoArgumentListContext) interface{} {
	var vals []value.Value
	for _, argCtx := range ctx.AllIoArgument() {
		res := v.Visit(argCtx)
		if res == nil {
			continue
		}
		if val, ok := res.(value.Value); ok {
			vals = append(vals, val)
		}
	}
	return vals
}

func (v *IRVisitor) VisitIoArgument(ctx *parser.IoArgumentContext) interface{} {
	if ctx.NEWLINE_CONST() != nil {
		return constant.NewCharArrayFromString("\n")
	}

	expr := v.Visit(ctx.Expression(0))
	val, ok := expr.(value.Value)
	if !ok {
		v.Errors = append(v.Errors, fmt.Errorf("invalid IO expression"))
		return nil
	}
	return val
}
