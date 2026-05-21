package visitor

import (
	"fmt"
	"github.com/AskaryanKarine/BMSTU-CC/cource/internal/parser"
	"github.com/antlr4-go/antlr/v4"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

func (v *IRVisitor) VisitLvalue(ctx *parser.LvalueContext) interface{} {
	if ctx.RETURN_VALUE() != nil {
		if v.currentFunc.Sig.RetType == types.Void {
			v.Errors = append(v.Errors, fmt.Errorf("'знач' used in procedure (void) context"))
			return nil
		}
		vi, ok := v.currentScope.Get(returnNameVar)
		if !ok {
			retAlloca := v.currentBlock.NewAlloca(v.currentFunc.Sig.RetType)
			vi = &VariableInfo{
				Name:      returnNameVar,
				Type:      v.currentFunc.Sig.RetType,
				LLVMValue: retAlloca,
			}
			_ = v.currentScope.Set(returnNameVar, vi)
		}
		return vi.LLVMValue
	}

	identVal := v.Visit(ctx.QualifiedIdentifier())
	vi, ok := identVal.(*VariableInfo)
	if !ok || vi == nil {
		v.Errors = append(v.Errors, fmt.Errorf("invalid lvalue: expected variable"))
		return nil
	}

	if ctx.IndexList() == nil {
		return vi.LLVMValue
	}

	indicesVal := v.Visit(ctx.IndexList())
	rawIndices, ok := indicesVal.([]value.Value)
	if !ok {
		v.Errors = append(v.Errors, fmt.Errorf("invalid index list for array access"))
		return nil
	}

	return v.getArrayElementPtr(vi, rawIndices)
}

func (v *IRVisitor) VisitIndexList(ctx *parser.IndexListContext) interface{} {
	if ctx.COLON() != nil {
		v.Errors = append(v.Errors, fmt.Errorf("array slices are not yet supported"))
		return nil
	}

	exprs := ctx.AllExpression()
	if len(exprs) == 2 {
		a := v.Visit(exprs[0])
		b := v.Visit(exprs[1])
		va, ok1 := a.(value.Value)
		vb, ok2 := b.(value.Value)
		if !ok1 || !ok2 {
			v.Errors = append(v.Errors, fmt.Errorf("invalid expressions in index list"))
			return nil
		}
		return []value.Value{va, vb}
	}

	if len(exprs) == 1 {
		e := v.Visit(exprs[0])
		ve, ok := e.(value.Value)
		if !ok {
			v.Errors = append(v.Errors, fmt.Errorf("invalid expression in index list"))
			return nil
		}
		return []value.Value{ve}
	}

	v.Errors = append(v.Errors, fmt.Errorf("unsupported index list format"))
	return nil
}

func (v *IRVisitor) ensureI32(idx value.Value) value.Value {
	if idx.Type().Equal(types.I32) {
		return idx
	}
	intTy, ok := idx.Type().(*types.IntType)
	if !ok {
		v.Errors = append(v.Errors, fmt.Errorf("index expression must evaluate to integer, got %s", idx.Type().String()))
		return constant.NewInt(types.I32, 0)
	}

	switch {
	case intTy.BitSize > 32:
		return v.currentBlock.NewTrunc(idx, types.I32)
	case intTy.BitSize < 32:
		return v.currentBlock.NewZExt(idx, types.I32)
	default:
		return idx
	}
}

func (v *IRVisitor) VisitExpression(ctx *parser.ExpressionContext) interface{} {
	if ctx.LogicalOrExpression() != nil {
		return v.Visit(ctx.LogicalOrExpression())
	}
	v.Errors = append(v.Errors, fmt.Errorf("invalid expression context"))
	return nil
}

func (v *IRVisitor) VisitExpressionList(ctx *parser.ExpressionListContext) interface{} {
	for _, i := range ctx.AllExpression() {
		v.Visit(i)
	}
	return nil
}

func (v *IRVisitor) VisitPrimaryExpression(ctx *parser.PrimaryExpressionContext) interface{} {
	switch {
	case ctx.Literal() != nil:
		return v.Visit(ctx.Literal())

	case ctx.QualifiedIdentifier() != nil:
		idVal := v.Visit(ctx.QualifiedIdentifier())

		if vi, ok := idVal.(*VariableInfo); ok {
			if _, isPtr := vi.LLVMValue.Type().(*types.PointerType); isPtr {
				if _, isArray := vi.Type.(*types.ArrayType); isArray {
					return vi.LLVMValue
				}
				return v.currentBlock.NewLoad(vi.Type, vi.LLVMValue)
			}
			return vi.LLVMValue
		}

		return idVal

	case ctx.RETURN_VALUE() != nil:
		vi, ok := v.currentScope.Get(returnNameVar)
		if !ok {
			v.Errors = append(v.Errors, fmt.Errorf("'знач' accessed but not defined"))
			return nil
		}
		return v.currentBlock.NewLoad(vi.Type, vi.LLVMValue)

	case ctx.Expression() != nil:
		return v.Visit(ctx.Expression())

	case ctx.ArrayLiteral() != nil:
		v.Errors = append(v.Errors, fmt.Errorf("array literals not implemented"))
		return nil

	default:
		v.Errors = append(v.Errors, fmt.Errorf("invalid primary expression"))
		return nil
	}
}

func (v *IRVisitor) VisitPostfixExpression(ctx *parser.PostfixExpressionContext) interface{} {
	base := v.Visit(ctx.PrimaryExpression())
	val, ok := base.(value.Value)
	if !ok {
		v.Errors = append(v.Errors, fmt.Errorf("invalid primary expression in postfix"))
		return nil
	}

	var vi *VariableInfo
	if ctx.PrimaryExpression().QualifiedIdentifier() != nil {
		if qval, ok := v.Visit(ctx.PrimaryExpression().QualifiedIdentifier()).(*VariableInfo); ok {
			vi = qval
		}
	}

	currentVal := val
	var elementType types.Type
	indexed := false

	for _, child := range ctx.GetChildren()[1:] {
		switch node := child.(type) {
		case antlr.TerminalNode:
			continue

		case *parser.IndexListContext:
			indexed = true
			indicesIfc := v.Visit(node)
			indices, ok := indicesIfc.([]value.Value)
			if !ok {
				v.Errors = append(v.Errors, fmt.Errorf("invalid index list"))
				return nil
			}
			if vi != nil {
				elementType = vi.Type
				for i := 0; i < len(indices); i++ {
					if arrTy, ok := elementType.(*types.ArrayType); ok {
						elementType = arrTy.ElemType
					} else {
						break
					}
				}
				currentVal = v.getArrayElementPtr(vi, indices)
			} else {
				args := append([]value.Value{constant.NewInt(types.I32, 0)}, indices...)
				ptrType, ok := currentVal.Type().(*types.PointerType)
				if !ok {
					v.Errors = append(v.Errors, fmt.Errorf("indexed value is not a pointer"))
					return nil
				}
				elementType = ptrType.ElemType
				for i := 0; i < len(indices); i++ {
					if arrTy, ok := elementType.(*types.ArrayType); ok {
						elementType = arrTy.ElemType
					} else {
						break
					}
				}
				currentVal = v.currentBlock.NewGetElementPtr(ptrType.ElemType, currentVal, args...)
			}

		case *parser.ArgumentListContext:
			argsIfc := v.Visit(node)
			args, ok := argsIfc.([]value.Value)
			if !ok {
				v.Errors = append(v.Errors, fmt.Errorf("invalid argument list"))
				return nil
			}
			fn, ok := currentVal.(*ir.Func)
			if !ok {
				v.Errors = append(v.Errors, fmt.Errorf("attempt to call non-function value"))
				return nil
			}
			currentVal = v.currentBlock.NewCall(fn, args...)

		default:
			v.Errors = append(v.Errors, fmt.Errorf("unexpected postfix expression element: %T", node))
			return nil
		}
	}

	if indexed {
		if ptrType, isPtr := currentVal.Type().(*types.PointerType); isPtr {
			et := elementType
			if et == nil {
				et = ptrType.ElemType
			}
			currentVal = v.currentBlock.NewLoad(et, currentVal)
		}
	}

	return currentVal
}

func (v *IRVisitor) VisitUnaryExpression(ctx *parser.UnaryExpressionContext) interface{} {
	if ctx.PLUS() != nil || ctx.MINUS() != nil || ctx.NOT() != nil {
		expr := v.Visit(ctx.UnaryExpression())
		val, ok := expr.(value.Value)
		if !ok {
			v.Errors = append(v.Errors, fmt.Errorf("invalid operand in unary expression"))
			return nil
		}

		switch {
		case ctx.PLUS() != nil:
			return val

		case ctx.MINUS() != nil:
			switch t := val.Type().(type) {
			case *types.FloatType:
				zero := constant.NewFloat(t, 0)
				return v.currentBlock.NewFSub(zero, val)

			case *types.IntType:
				zero := constant.NewInt(t, 0)
				return v.currentBlock.NewSub(zero, val)

			default:
				v.Errors = append(v.Errors, fmt.Errorf("unsupported type for unary minus"))
				return nil
			}

		case ctx.NOT() != nil:
			return v.currentBlock.NewXor(val, constant.NewInt(types.I1, 1))

		default:
			v.Errors = append(v.Errors, fmt.Errorf("unsupported unary operator"))
			return nil
		}
	}
	return v.Visit(ctx.PostfixExpression())
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

func (v *IRVisitor) VisitPowerExpression(ctx *parser.PowerExpressionContext) interface{} {
	lhsRaw := v.Visit(ctx.UnaryExpression())
	lhs, ok := lhsRaw.(value.Value)
	if !ok {
		v.Errors = append(v.Errors, fmt.Errorf("invalid base operand in power expression"))
		return nil
	}

	if ctx.POWER() == nil {
		return lhs
	}
	rhsRaw := v.Visit(ctx.PowerExpression())
	rhs, ok := rhsRaw.(value.Value)
	if !ok {
		v.Errors = append(v.Errors, fmt.Errorf("invalid exponent operand in power expression"))
		return nil
	}

	isFloat := func(t types.Type) bool {
		_, ok := t.(*types.FloatType)
		return ok
	}
	isFloatBase := isFloat(lhs.Type())
	isFloatExp := isFloat(rhs.Type())
	useFloat := isFloatBase || isFloatExp

	toDouble := func(val value.Value) value.Value {
		if isFloat(val.Type()) && val.Type().Equal(types.Float) {
			return v.currentBlock.NewFPExt(val, types.Double)
		}
		if _, ok := val.Type().(*types.IntType); ok {
			return v.currentBlock.NewSIToFP(val, types.Double)
		}
		return val
	}

	baseD := toDouble(lhs)
	expD := toDouble(rhs)

	resD := v.currentBlock.NewCall(v.ensurePowFunc(), baseD, expD)

	if !useFloat {
		return v.currentBlock.NewFPToSI(resD, types.I32)
	}
	return resD
}

func (v *IRVisitor) VisitMultiplicativeExpression(ctx *parser.MultiplicativeExpressionContext) interface{} {
	lhsRaw := v.Visit(ctx.PowerExpression(0))
	lhs, ok := lhsRaw.(value.Value)
	if !ok {
		v.Errors = append(v.Errors, fmt.Errorf("invalid operand in multiplicative expression"))
		return nil
	}
	for i := 1; i < ctx.GetChildCount(); i += 2 {
		rhs := v.Visit(ctx.PowerExpression((i + 1) / 2))
		rhsVal, ok := rhs.(value.Value)
		if !ok {
			v.Errors = append(v.Errors, fmt.Errorf("invalid right-hand operand in multiplicative expression"))
			return nil
		}

		lhs, rhsVal = v.castToMatch(lhs, rhsVal)

		node := ctx.GetChild(i)
		token, ok := node.(antlr.TerminalNode)
		if !ok {
			v.Errors = append(v.Errors, fmt.Errorf("unexpected non-terminal in multiplicative expression"))
			return nil
		}
		switch token.GetText() {
		case "*":
			switch lhs.Type().(type) {
			case *types.FloatType:
				lhs = v.currentBlock.NewFMul(lhs, rhsVal)
			default:
				lhs = v.currentBlock.NewMul(lhs, rhsVal)
			}
		case "/":
			switch lhs.Type().(type) {
			case *types.FloatType:
				lhs = v.currentBlock.NewFDiv(lhs, rhsVal)
			default:
				lhs = v.currentBlock.NewSDiv(lhs, rhsVal)
			}
		case "div":
			lhs = v.currentBlock.NewSDiv(lhs, rhsVal)
		case "mod":
			lhs = v.currentBlock.NewSRem(lhs, rhsVal)
		default:
			v.Errors = append(v.Errors, fmt.Errorf("unsupported multiplicative operator: %s", token.GetText()))
			return nil
		}
	}
	return lhs
}

func (v *IRVisitor) VisitAdditiveExpression(ctx *parser.AdditiveExpressionContext) interface{} {
	lhsRaw := v.Visit(ctx.MultiplicativeExpression(0))
	lhs, ok := lhsRaw.(value.Value)
	if !ok {
		v.Errors = append(v.Errors, fmt.Errorf("invalid operand in additive expression"))
		return nil
	}
	for i := 1; i < ctx.GetChildCount(); i += 2 {
		rhs := v.Visit(ctx.MultiplicativeExpression((i + 1) / 2))
		rhsVal, ok := rhs.(value.Value)
		if !ok {
			v.Errors = append(v.Errors, fmt.Errorf("invalid right-hand operand in additive expression"))
			return nil
		}

		lhs, rhsVal = v.castToMatch(lhs, rhsVal)

		node := ctx.GetChild(i)
		token, ok := node.(antlr.TerminalNode)
		if !ok {
			v.Errors = append(v.Errors, fmt.Errorf("unexpected non-terminal in additive expression"))
			return nil
		}
		switch token.GetText() {
		case "+":
			switch lhs.Type().(type) {
			case *types.FloatType:
				lhs = v.currentBlock.NewFAdd(lhs, rhsVal)
			default:
				lhs = v.currentBlock.NewAdd(lhs, rhsVal)
			}
		case "-":
			switch lhs.Type().(type) {
			case *types.FloatType:
				lhs = v.currentBlock.NewFSub(lhs, rhsVal)
			default:
				lhs = v.currentBlock.NewSub(lhs, rhsVal)
			}
		default:
			v.Errors = append(v.Errors, fmt.Errorf("unsupported additive operator: %s", token.GetText()))
			return nil
		}
	}
	return lhs
}

func (v *IRVisitor) VisitRelationalExpression(ctx *parser.RelationalExpressionContext) interface{} {
	lhsRaw := v.Visit(ctx.AdditiveExpression(0))
	lhs, ok := lhsRaw.(value.Value)
	if !ok {
		v.Errors = append(v.Errors, fmt.Errorf("invalid operand in relational expression"))
		return nil
	}
	for i := 1; i < ctx.GetChildCount(); i += 2 {
		rhs := v.Visit(ctx.AdditiveExpression((i + 1) / 2))
		rhsVal, ok := rhs.(value.Value)
		if !ok {
			v.Errors = append(v.Errors, fmt.Errorf("invalid right-hand operand in relational expression"))
			return nil
		}

		lhs, rhsVal = v.castToMatch(lhs, rhsVal)

		node := ctx.GetChild(i)
		token, ok := node.(antlr.TerminalNode)
		if !ok {
			v.Errors = append(v.Errors, fmt.Errorf("unexpected non-terminal in relational expression"))
			return nil
		}
		switch token.GetText() {
		case "<":
			switch lhs.Type().(type) {
			case *types.FloatType:
				lhs = v.currentBlock.NewFCmp(enum.FPredOLT, lhs, rhsVal)
			default:
				lhs = v.currentBlock.NewICmp(enum.IPredSLT, lhs, rhsVal)
			}
		case "<=":
			switch lhs.Type().(type) {
			case *types.FloatType:
				lhs = v.currentBlock.NewFCmp(enum.FPredOLE, lhs, rhsVal)
			default:
				lhs = v.currentBlock.NewICmp(enum.IPredSLE, lhs, rhsVal)
			}
		case ">":
			switch lhs.Type().(type) {
			case *types.FloatType:
				lhs = v.currentBlock.NewFCmp(enum.FPredOGT, lhs, rhsVal)
			default:
				lhs = v.currentBlock.NewICmp(enum.IPredSGT, lhs, rhsVal)
			}
		case ">=":
			switch lhs.Type().(type) {
			case *types.FloatType:
				lhs = v.currentBlock.NewFCmp(enum.FPredOGE, lhs, rhsVal)
			default:
				lhs = v.currentBlock.NewICmp(enum.IPredSGE, lhs, rhsVal)
			}
		default:
			v.Errors = append(v.Errors, fmt.Errorf("unsupported relational operator: %s", token.GetText()))
			return nil
		}
	}
	return lhs
}

func (v *IRVisitor) VisitEqualityExpression(ctx *parser.EqualityExpressionContext) interface{} {
	lhsRaw := v.Visit(ctx.RelationalExpression(0))
	lhs, ok := lhsRaw.(value.Value)
	if !ok {
		v.Errors = append(v.Errors, fmt.Errorf("invalid operand in equality expression"))
		return nil
	}

	for i := 1; i < ctx.GetChildCount(); i += 2 {
		rhsRaw := v.Visit(ctx.RelationalExpression((i + 1) / 2))
		rhs, ok := rhsRaw.(value.Value)
		if !ok {
			v.Errors = append(v.Errors, fmt.Errorf("invalid right-hand operand in equality expression"))
			return nil
		}

		lhs, rhs = v.castToMatch(lhs, rhs)

		node := ctx.GetChild(i)
		token, ok := node.(antlr.TerminalNode)
		if !ok {
			v.Errors = append(v.Errors, fmt.Errorf("unexpected non-terminal in equality expression"))
			return nil
		}

		switch token.GetText() {
		case "=":
			switch lhs.Type().(type) {
			case *types.FloatType:
				lhs = v.currentBlock.NewFCmp(enum.FPredOEQ, lhs, rhs)
			default:
				lhs = v.currentBlock.NewICmp(enum.IPredEQ, lhs, rhs)
			}
		case "<>":
			switch lhs.Type().(type) {
			case *types.FloatType:
				lhs = v.currentBlock.NewFCmp(enum.FPredONE, lhs, rhs)
			default:
				lhs = v.currentBlock.NewICmp(enum.IPredNE, lhs, rhs)
			}
		default:
			v.Errors = append(v.Errors, fmt.Errorf("unsupported equality operator: %s", token.GetText()))
			return nil
		}
	}

	return lhs
}

func (v *IRVisitor) VisitLogicalAndExpression(ctx *parser.LogicalAndExpressionContext) interface{} {
	lhsRaw := v.Visit(ctx.EqualityExpression(0))
	lhs, ok := lhsRaw.(value.Value)
	if !ok {
		v.Errors = append(v.Errors, fmt.Errorf("invalid operand for AND"))
		return nil
	}
	for i := 1; i < len(ctx.AllEqualityExpression()); i++ {
		rhs := v.Visit(ctx.EqualityExpression(i))
		rhsVal, ok := rhs.(value.Value)
		if !ok {
			v.Errors = append(v.Errors, fmt.Errorf("invalid right-hand operand for AND"))
			return nil
		}

		lhs, rhsVal = v.castToMatch(lhs, rhsVal)

		lhs = v.currentBlock.NewAnd(lhs, rhsVal)
	}
	return lhs
}

func (v *IRVisitor) VisitLogicalOrExpression(ctx *parser.LogicalOrExpressionContext) interface{} {
	lhsRaw := v.Visit(ctx.LogicalAndExpression(0))
	lhs, ok := lhsRaw.(value.Value)
	if !ok {
		v.Errors = append(v.Errors, fmt.Errorf("invalid operand for OR"))
		return nil
	}
	for i := 1; i < len(ctx.AllLogicalAndExpression()); i++ {
		rhs := v.Visit(ctx.LogicalAndExpression(i))
		rhsVal, ok := rhs.(value.Value)
		if !ok {
			v.Errors = append(v.Errors, fmt.Errorf("invalid right-hand operand for OR"))
			return nil
		}

		lhs, rhsVal = v.castToMatch(lhs, rhsVal)

		lhs = v.currentBlock.NewOr(lhs, rhsVal)
	}
	return lhs
}
