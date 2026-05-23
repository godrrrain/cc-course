package visitor

import (
	"fmt"
	"strconv"

	"github.com/AskaryanKarine/BMSTU-CC/cource/internal/parser"
	"github.com/antlr4-go/antlr/v4"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

func (v *IRVisitor) VisitExpression(ctx *parser.ExpressionContext) interface{} {
	switch {
	case ctx.AssignableExpression() != nil && ctx.AssignmentOperator() != nil:
		if ae, ok := ctx.AssignableExpression().(*parser.AssignableExpressionContext); ok {
			return v.VisitAssignableExpression(ae)
		}
	case ctx.ConditionalExpression() != nil:
		return v.Visit(ctx.ConditionalExpression())
	case ctx.ThrowExpression() != nil:
		return v.Visit(ctx.ThrowExpression())
	}
	return nil
}

func (v *IRVisitor) VisitConditionalExpression(ctx *parser.ConditionalExpressionContext) interface{} {
	if ctx.IfNullExpression() == nil {
		return nil
	}
	val := v.Visit(ctx.IfNullExpression())
	if ctx.QUES() != nil && len(ctx.AllExpressionWithoutCascade()) == 2 {
		condVal, ok := val.(value.Value)
		if !ok {
			v.Errors = append(v.Errors, fmt.Errorf("invalid condition in ternary expression"))
			return nil
		}
		thenBlock := v.currentFunc.NewBlock("tern.then")
		elseBlock := v.currentFunc.NewBlock("tern.else")
		mergeBlock := v.currentFunc.NewBlock("tern.end")

		v.currentBlock.NewCondBr(condVal, thenBlock, elseBlock)

		v.pushBlock(thenBlock)
		thenVal := v.Visit(ctx.ExpressionWithoutCascade(0))
		thenV, _ := thenVal.(value.Value)
		if v.currentBlock.Term == nil {
			v.currentBlock.NewBr(mergeBlock)
		}
		v.popBlock()

		v.pushBlock(elseBlock)
		elseVal := v.Visit(ctx.ExpressionWithoutCascade(1))
		elseV, _ := elseVal.(value.Value)
		if v.currentBlock.Term == nil {
			v.currentBlock.NewBr(mergeBlock)
		}
		v.popBlock()

		v.currentBlock = mergeBlock
		phi := v.currentBlock.NewPhi()
		phi.Incs = append(phi.Incs, ir.NewIncoming(thenV, thenBlock))
		phi.Incs = append(phi.Incs, ir.NewIncoming(elseV, elseBlock))
		return phi
	}
	return val
}

func (v *IRVisitor) VisitIfNullExpression(ctx *parser.IfNullExpressionContext) interface{} {
	val := v.Visit(ctx.LogicalOrExpression(0))
	for i := 1; i < len(ctx.AllLogicalOrExpression()); i++ {
		rhs := v.Visit(ctx.LogicalOrExpression(i))
		_ = rhs
	}
	return val
}

func (v *IRVisitor) VisitLogicalOrExpression(ctx *parser.LogicalOrExpressionContext) interface{} {
	val := v.Visit(ctx.LogicalAndExpression(0))
	for i := 1; i < len(ctx.AllLogicalAndExpression()); i++ {
		rhs := v.Visit(ctx.LogicalAndExpression(i))
		lhsV, ok1 := val.(value.Value)
		rhsV, ok2 := rhs.(value.Value)
		if ok1 && ok2 {
			val = v.currentBlock.NewOr(lhsV, rhsV)
		}
	}
	return val
}

func (v *IRVisitor) VisitLogicalAndExpression(ctx *parser.LogicalAndExpressionContext) interface{} {
	val := v.Visit(ctx.BitwiseOrExpression(0))
	for i := 1; i < len(ctx.AllBitwiseOrExpression()); i++ {
		rhs := v.Visit(ctx.BitwiseOrExpression(i))
		lhsV, ok1 := val.(value.Value)
		rhsV, ok2 := rhs.(value.Value)
		if ok1 && ok2 {
			val = v.currentBlock.NewAnd(lhsV, rhsV)
		}
	}
	return val
}

func (v *IRVisitor) VisitBitwiseOrExpression(ctx *parser.BitwiseOrExpressionContext) interface{} {
	val := v.Visit(ctx.BitwiseXorExpression(0))
	for i := 1; i < len(ctx.AllBitwiseXorExpression()); i++ {
		rhs := v.Visit(ctx.BitwiseXorExpression(i))
		lhsV, ok1 := val.(value.Value)
		rhsV, ok2 := rhs.(value.Value)
		if ok1 && ok2 {
			val = v.currentBlock.NewOr(lhsV, rhsV)
		}
	}
	return val
}

func (v *IRVisitor) VisitBitwiseXorExpression(ctx *parser.BitwiseXorExpressionContext) interface{} {
	val := v.Visit(ctx.BitwiseAndExpression(0))
	for i := 1; i < len(ctx.AllBitwiseAndExpression()); i++ {
		rhs := v.Visit(ctx.BitwiseAndExpression(i))
		lhsV, ok1 := val.(value.Value)
		rhsV, ok2 := rhs.(value.Value)
		if ok1 && ok2 {
			val = v.currentBlock.NewXor(lhsV, rhsV)
		}
	}
	return val
}

func (v *IRVisitor) VisitBitwiseAndExpression(ctx *parser.BitwiseAndExpressionContext) interface{} {
	val := v.Visit(ctx.ShiftExpression(0))
	for i := 1; i < len(ctx.AllShiftExpression()); i++ {
		rhs := v.Visit(ctx.ShiftExpression(i))
		lhsV, ok1 := val.(value.Value)
		rhsV, ok2 := rhs.(value.Value)
		if ok1 && ok2 {
			val = v.currentBlock.NewAnd(lhsV, rhsV)
		}
	}
	return val
}

func (v *IRVisitor) VisitShiftExpression(ctx *parser.ShiftExpressionContext) interface{} {
	val := v.Visit(ctx.AdditiveExpression(0))
	for i := 1; i < ctx.GetChildCount(); i += 2 {
		rhs := v.Visit(ctx.AdditiveExpression((i + 1) / 2))
		lhsV, ok1 := val.(value.Value)
		rhsV, ok2 := rhs.(value.Value)
		if !ok1 || !ok2 {
			continue
		}
		node := ctx.GetChild(i)
		if token, ok := node.(antlr.TerminalNode); ok {
			switch token.GetText() {
			case "<<":
				val = v.currentBlock.NewShl(lhsV, rhsV)
			case ">>":
				val = v.currentBlock.NewAShr(lhsV, rhsV)
			}
		}
	}
	return val
}

func (v *IRVisitor) VisitAdditiveExpression(ctx *parser.AdditiveExpressionContext) interface{} {
	val := v.Visit(ctx.MultiplicativeExpression(0))
	for i := 1; i < ctx.GetChildCount(); i += 2 {
		rhs := v.Visit(ctx.MultiplicativeExpression((i + 1) / 2))
		lhsV, ok1 := val.(value.Value)
		rhsV, ok2 := rhs.(value.Value)
		if !ok1 || !ok2 {
			continue
		}
		lhsV, rhsV = v.castToMatch(lhsV, rhsV)
		node := ctx.GetChild(i)
		if token, ok := node.(antlr.TerminalNode); ok {
			switch token.GetText() {
			case "+":
				switch lhsV.Type().(type) {
				case *types.FloatType:
					val = v.currentBlock.NewFAdd(lhsV, rhsV)
				default:
					val = v.currentBlock.NewAdd(lhsV, rhsV)
				}
			case "-":
				switch lhsV.Type().(type) {
				case *types.FloatType:
					val = v.currentBlock.NewFSub(lhsV, rhsV)
				default:
					val = v.currentBlock.NewSub(lhsV, rhsV)
				}
			}
		}
	}
	return val
}

func (v *IRVisitor) VisitMultiplicativeExpression(ctx *parser.MultiplicativeExpressionContext) interface{} {
	val := v.Visit(ctx.UnaryExpression(0))
	for i := 1; i < ctx.GetChildCount(); i += 2 {
		rhs := v.Visit(ctx.UnaryExpression((i + 1) / 2))
		lhsV, ok1 := val.(value.Value)
		rhsV, ok2 := rhs.(value.Value)
		if !ok1 || !ok2 {
			continue
		}
		lhsV, rhsV = v.castToMatch(lhsV, rhsV)
		node := ctx.GetChild(i)
		if token, ok := node.(antlr.TerminalNode); ok {
			switch token.GetText() {
			case "*":
				switch lhsV.Type().(type) {
				case *types.FloatType:
					val = v.currentBlock.NewFMul(lhsV, rhsV)
				default:
					val = v.currentBlock.NewMul(lhsV, rhsV)
				}
			case "/":
				switch lhsV.Type().(type) {
				case *types.FloatType:
					val = v.currentBlock.NewFDiv(lhsV, rhsV)
				default:
					val = v.currentBlock.NewSDiv(lhsV, rhsV)
				}
			case "~/":
				val = v.currentBlock.NewSDiv(lhsV, rhsV)
			case "%":
				val = v.currentBlock.NewSRem(lhsV, rhsV)
			}
		}
	}
	return val
}

func (v *IRVisitor) VisitUnaryExpression(ctx *parser.UnaryExpressionContext) interface{} {
	switch {
	case ctx.PLUS() != nil:
		return v.Visit(ctx.UnaryExpression())
	case ctx.MINUS() != nil && ctx.UnaryExpression() != nil:
		val := v.Visit(ctx.UnaryExpression())
		if vv, ok := val.(value.Value); ok {
			switch vv.Type().(type) {
			case *types.FloatType:
				zero := constant.NewFloat(vv.Type().(*types.FloatType), 0)
				return v.currentBlock.NewFSub(zero, vv)
			case *types.IntType:
				zero := constant.NewInt(vv.Type().(*types.IntType), 0)
				return v.currentBlock.NewSub(zero, vv)
			}
		}
		return nil
	case ctx.NOT() != nil:
		val := v.Visit(ctx.UnaryExpression())
		if vv, ok := val.(value.Value); ok {
			if _, ok := vv.Type().(*types.IntType); ok {
				return v.currentBlock.NewXor(vv, constant.NewInt(types.I1, 1))
			}
		}
		return nil
	case ctx.TILDE() != nil:
		val := v.Visit(ctx.UnaryExpression())
		if vv, ok := val.(value.Value); ok {
			return v.currentBlock.NewXor(vv, constant.NewInt(vv.Type().(*types.IntType), -1))
		}
		return nil
	case ctx.AwaitExpression() != nil:
		return v.Visit(ctx.AwaitExpression())
	case ctx.PostfixExpression() != nil:
		return v.Visit(ctx.PostfixExpression())
	}
	return nil
}

func (v *IRVisitor) VisitPostfixExpression(ctx *parser.PostfixExpressionContext) interface{} {
	if ctx.AssignableExpression() != nil && (ctx.PLUS_PLUS() != nil || ctx.MINUS_MINUS() != nil) {
		val := v.Visit(ctx.AssignableExpression())
		if vi, ok := val.(*VariableInfo); ok {
			loaded := v.currentBlock.NewLoad(vi.Type, vi.LLVMValue)
			var delta value.Value
			switch vi.Type.(type) {
			case *types.FloatType:
				delta = constant.NewFloat(vi.Type.(*types.FloatType), 1)
			case *types.IntType:
				delta = constant.NewInt(vi.Type.(*types.IntType), 1)
			}
			if delta != nil {
				var newVal value.Value
				if ctx.PLUS_PLUS() != nil {
					if _, ok := vi.Type.(*types.FloatType); ok {
						newVal = v.currentBlock.NewFAdd(loaded, delta)
					} else {
						newVal = v.currentBlock.NewAdd(loaded, delta)
					}
				} else {
					if _, ok := vi.Type.(*types.FloatType); ok {
						newVal = v.currentBlock.NewFSub(loaded, delta)
					} else {
						newVal = v.currentBlock.NewSub(loaded, delta)
					}
				}
				v.currentBlock.NewStore(newVal, vi.LLVMValue)
			}
			return loaded
		}
	}
	if ctx.Primary() != nil {
		currentVal := v.Visit(ctx.Primary())
		for _, sel := range ctx.AllSelector() {
			if ap := sel.ArgumentPart(); ap != nil {
				fn, ok := currentVal.(*ir.Func)
				if !ok {
					v.Errors = append(v.Errors, fmt.Errorf("attempt to call non-function value"))
					return nil
				}
				argVal := v.Visit(ap)
				args, _ := argVal.([]value.Value)
				if fn.Name() == "printf" && len(args) > 0 {
					formatStr := v.defineGlobalString("%s\n")
					callArgs := []value.Value{formatStr}
					callArgs = append(callArgs, args...)
					currentVal = v.currentBlock.NewCall(fn, callArgs...)
				} else {
					currentVal = v.currentBlock.NewCall(fn, args...)
				}
			} else if sel.LBRACKET() != nil && sel.Expression() != nil {
				v.Visit(sel.Expression())
			}
		}
		return currentVal
	}
	return nil
}

func (v *IRVisitor) VisitPrimary(ctx *parser.PrimaryContext) interface{} {
	switch {
	case ctx.ThisExpression() != nil:
		return v.Visit(ctx.ThisExpression())
	case ctx.IDENTIFIER() != nil:
		name := ctx.IDENTIFIER().GetText()
		if vi, ok := v.currentScope.Get(name); ok {
			if ptrTy, ok := vi.LLVMValue.Type().(*types.PointerType); ok {
				return v.currentBlock.NewLoad(ptrTy.ElemType, vi.LLVMValue)
			}
			return vi.LLVMValue
		}
		if fn, ok := v.funcs[name]; ok {
			return fn
		}
		v.Errors = append(v.Errors, fmt.Errorf("undefined identifier: %s", name))
		return nil
	case ctx.Literal() != nil:
		return v.Visit(ctx.Literal())
	case ctx.NewExpr() != nil:
		return v.Visit(ctx.NewExpr())
	case ctx.ConstExpression() != nil:
		return v.Visit(ctx.ConstExpression())
	case ctx.FunctionExpression() != nil:
		return v.Visit(ctx.FunctionExpression())
	case ctx.Expression() != nil:
		return v.Visit(ctx.Expression())
	}
	return nil
}

func (v *IRVisitor) VisitLiteral(ctx *parser.LiteralContext) interface{} {
	switch {
	case ctx.NullLiteral() != nil:
		return v.Visit(ctx.NullLiteral())
	case ctx.BooleanLiteral() != nil:
		return v.Visit(ctx.BooleanLiteral())
	case ctx.NumericLiteral() != nil:
		return v.Visit(ctx.NumericLiteral())
	case ctx.StringLiteral() != nil:
		return v.Visit(ctx.StringLiteral())
	case ctx.ListLiteral() != nil:
		return v.Visit(ctx.ListLiteral())
	case ctx.SetOrMapLiteral() != nil:
		return v.Visit(ctx.SetOrMapLiteral())
	}
	return nil
}

func (v *IRVisitor) VisitNullLiteral(ctx *parser.NullLiteralContext) interface{} {
	return constant.NewNull(types.NewPointer(types.I8))
}

func (v *IRVisitor) VisitBooleanLiteral(ctx *parser.BooleanLiteralContext) interface{} {
	if ctx.TRUE_() != nil {
		return constant.NewInt(types.I1, 1)
	}
	return constant.NewInt(types.I1, 0)
}

func (v *IRVisitor) VisitNumericLiteral(ctx *parser.NumericLiteralContext) interface{} {
	if ctx.NUMBER() != nil {
		text := ctx.NUMBER().GetText()
		if val, err := strconv.ParseFloat(text, 64); err == nil {
			if val == float64(int64(val)) {
				return constant.NewInt(types.I64, int64(val))
			}
			return constant.NewFloat(types.Double, val)
		}
	}
	if ctx.HEX_NUMBER() != nil {
		text := ctx.HEX_NUMBER().GetText()
		if val, err := strconv.ParseInt(text, 0, 64); err == nil {
			return constant.NewInt(types.I64, val)
		}
	}
	return constant.NewInt(types.I64, 0)
}

func (v *IRVisitor) VisitStringLiteral(ctx *parser.StringLiteralContext) interface{} {
	text := ctx.GetText()
	unquoted := stripQuotes(text)
	return v.defineGlobalString(unquoted)
}

func (v *IRVisitor) VisitListLiteral(ctx *parser.ListLiteralContext) interface{} {
	v.Errors = append(v.Errors, fmt.Errorf("list literals not yet implemented"))
	return nil
}

func (v *IRVisitor) VisitSetOrMapLiteral(ctx *parser.SetOrMapLiteralContext) interface{} {
	v.Errors = append(v.Errors, fmt.Errorf("map/set literals not yet implemented"))
	return nil
}

func (v *IRVisitor) VisitAssignableExpression(ctx *parser.AssignableExpressionContext) interface{} {
	if ctx.IDENTIFIER() != nil {
		name := ctx.IDENTIFIER().GetText()
		if vi, ok := v.currentScope.Get(name); ok {
			return vi
		}
		v.Errors = append(v.Errors, fmt.Errorf("undefined variable: %s", name))
		return nil
	}
	if ctx.Primary() != nil {
		return v.Visit(ctx.Primary())
	}
	return nil
}

func (v *IRVisitor) VisitSelector(ctx *parser.SelectorContext) interface{} {
	if ctx.DOT() != nil && ctx.IDENTIFIER() != nil {
		_ = ctx.IDENTIFIER().GetText()
	}
	if ctx.LBRACKET() != nil && ctx.Expression() != nil {
		return v.Visit(ctx.Expression())
	}
	return nil
}

func (v *IRVisitor) VisitThrowExpression(ctx *parser.ThrowExpressionContext) interface{} {
	v.Errors = append(v.Errors, fmt.Errorf("throw expressions not yet implemented"))
	return nil
}

func (v *IRVisitor) VisitAwaitExpression(ctx *parser.AwaitExpressionContext) interface{} {
	v.Errors = append(v.Errors, fmt.Errorf("await expressions not yet implemented"))
	return v.Visit(ctx.UnaryExpression())
}

func (v *IRVisitor) VisitThisExpression(ctx *parser.ThisExpressionContext) interface{} {
	v.Errors = append(v.Errors, fmt.Errorf("this expressions not yet implemented"))
	return nil
}

func (v *IRVisitor) VisitNewExpr(ctx *parser.NewExprContext) interface{} {
	v.Errors = append(v.Errors, fmt.Errorf("new expressions not yet implemented"))
	return nil
}

func (v *IRVisitor) VisitConstExpression(ctx *parser.ConstExpressionContext) interface{} {
	v.Errors = append(v.Errors, fmt.Errorf("const expressions not yet implemented"))
	return nil
}

func (v *IRVisitor) VisitFunctionExpression(ctx *parser.FunctionExpressionContext) interface{} {
	v.Errors = append(v.Errors, fmt.Errorf("function expressions not yet implemented"))
	return nil
}

func (v *IRVisitor) VisitExpressionList(ctx *parser.ExpressionListContext) interface{} {
	var vals []value.Value
	for _, expr := range ctx.AllExpression() {
		if val, ok := v.Visit(expr).(value.Value); ok {
			vals = append(vals, val)
		}
	}
	return vals
}

func (v *IRVisitor) ensureI32(idx value.Value) value.Value {
	if idx.Type().Equal(types.I32) {
		return idx
	}
	if intTy, ok := idx.Type().(*types.IntType); ok {
		if intTy.BitSize > 32 {
			return v.currentBlock.NewTrunc(idx, types.I32)
		}
		if intTy.BitSize < 32 {
			return v.currentBlock.NewZExt(idx, types.I32)
		}
	}
	return idx
}

func (v *IRVisitor) ensurePowFunc() *ir.Func {
	if fn, ok := v.funcs["pow"]; ok {
		return fn
	}
	pow := v.Module.NewFunc("pow", types.Double,
		ir.NewParam("base", types.Double),
		ir.NewParam("exp", types.Double))
	v.funcs["pow"] = pow
	return pow
}
