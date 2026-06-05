package visitor

import (
	"fmt"
	"strings"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"

	"cc-course/internal/ast"
)

func (v *ASTIRVisitor) visitExpr(e ast.Expr) value.Value {
	if e == nil {
		return nil
	}
	switch ee := e.(type) {
	case *ast.BinaryExpr:
		return v.visitBinary(ee)
	case *ast.UnaryExpr:
		return v.visitUnary(ee)
	case *ast.PostfixExpr:
		return v.visitPostfix(ee)
	case *ast.AssignExpr:
		return v.visitAssign(ee)
	case *ast.CallExpr:
		return v.visitCall(ee)
	case *ast.Identifier:
		return v.visitIdent(ee)
	case *ast.IntLiteral:
		return constant.NewInt(types.I64, ee.Value)
	case *ast.FloatLiteral:
		return constant.NewFloat(types.Double, ee.Value)
	case *ast.BoolLiteral:
		if ee.Value {
			return constant.NewInt(types.I1, 1)
		}
		return constant.NewInt(types.I1, 0)
	case *ast.StringLiteral:
		return v.defineGlobalString(ee.Value)
	case *ast.NullLiteral:
		return constant.NewNull(types.NewPointer(types.I8))
	case *ast.ListLiteral:
		return v.visitListLit(ee)
	case *ast.IndexExpr:
		return v.visitIndex(ee)
	case *ast.TernaryExpr:
		return v.visitTernary(ee)
	case *ast.InterpolatedStringExpr:
		return v.visitInterpolatedString(ee)
	}
	return nil
}

func (v *ASTIRVisitor) visitBinary(be *ast.BinaryExpr) value.Value {
	lhs := v.visitExpr(be.Left)
	rhs := v.visitExpr(be.Right)
	if lhs == nil || rhs == nil {
		return nil
	}

	switch be.Op {
	case "+", "-", "*", "/", "%":
		lhs, rhs = v.castToMatch(lhs, rhs)
	}

	isFloat := false
	if _, ok := lhs.Type().(*types.FloatType); ok {
		isFloat = true
	}

	switch be.Op {
	case "+":
		if isFloat {
			return v.currentBlock.NewFAdd(lhs, rhs)
		}
		return v.currentBlock.NewAdd(lhs, rhs)
	case "-":
		if isFloat {
			return v.currentBlock.NewFSub(lhs, rhs)
		}
		return v.currentBlock.NewSub(lhs, rhs)
	case "*":
		if isFloat {
			return v.currentBlock.NewFMul(lhs, rhs)
		}
		return v.currentBlock.NewMul(lhs, rhs)
	case "/":
		if isFloat {
			return v.currentBlock.NewFDiv(lhs, rhs)
		}
		return v.currentBlock.NewSDiv(lhs, rhs)
	case "~/":
		return v.currentBlock.NewSDiv(lhs, rhs)
	case "%":
		return v.currentBlock.NewSRem(lhs, rhs)
	case "==":
		if isFloat {
			return v.currentBlock.NewFCmp(enum.FPredOEQ, lhs, rhs)
		}
		return v.currentBlock.NewICmp(enum.IPredEQ, lhs, rhs)
	case "!=":
		if isFloat {
			return v.currentBlock.NewFCmp(enum.FPredONE, lhs, rhs)
		}
		return v.currentBlock.NewICmp(enum.IPredNE, lhs, rhs)
	case "<":
		if isFloat {
			return v.currentBlock.NewFCmp(enum.FPredOLT, lhs, rhs)
		}
		return v.currentBlock.NewICmp(enum.IPredSLT, lhs, rhs)
	case "<=":
		if isFloat {
			return v.currentBlock.NewFCmp(enum.FPredOLE, lhs, rhs)
		}
		return v.currentBlock.NewICmp(enum.IPredSLE, lhs, rhs)
	case ">":
		if isFloat {
			return v.currentBlock.NewFCmp(enum.FPredOGT, lhs, rhs)
		}
		return v.currentBlock.NewICmp(enum.IPredSGT, lhs, rhs)
	case ">=":
		if isFloat {
			return v.currentBlock.NewFCmp(enum.FPredOGE, lhs, rhs)
		}
		return v.currentBlock.NewICmp(enum.IPredSGE, lhs, rhs)
	case "&&":
		return v.currentBlock.NewAnd(lhs, rhs)
	case "||":
		return v.currentBlock.NewOr(lhs, rhs)
	case "&":
		return v.currentBlock.NewAnd(lhs, rhs)
	case "|":
		return v.currentBlock.NewOr(lhs, rhs)
	case "^":
		return v.currentBlock.NewXor(lhs, rhs)
	case "<<":
		return v.currentBlock.NewShl(lhs, rhs)
	case ">>":
		return v.currentBlock.NewAShr(lhs, rhs)
	}
	return nil
}

func (v *ASTIRVisitor) visitUnary(ue *ast.UnaryExpr) value.Value {
	operand := v.visitExpr(ue.Operand)
	if operand == nil {
		return nil
	}

	switch ue.Op {
	case "-":
		if ft, ok := operand.Type().(*types.FloatType); ok {
			zero := constant.NewFloat(ft, 0)
			return v.currentBlock.NewFSub(zero, operand)
		}
		if it, ok := operand.Type().(*types.IntType); ok {
			zero := constant.NewInt(it, 0)
			return v.currentBlock.NewSub(zero, operand)
		}
	case "!":
		if _, ok := operand.Type().(*types.IntType); ok {
			return v.currentBlock.NewXor(operand, constant.NewInt(types.I1, 1))
		}
	case "~":
		if it, ok := operand.Type().(*types.IntType); ok {
			return v.currentBlock.NewXor(operand, constant.NewInt(it, -1))
		}
	}
	return nil
}

func (v *ASTIRVisitor) visitPostfix(pe *ast.PostfixExpr) value.Value {
	switch target := pe.Target.(type) {
	case *ast.Identifier:
		name := target.Name
		if vi, ok := v.currentScope.Get(name); ok {
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
				if pe.Op == "++" {
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
		v.Errors = append(v.Errors, fmt.Errorf("undefined variable: %s", name))
		return nil
	}
	v.Errors = append(v.Errors, fmt.Errorf("postfix expression on non-identifier not supported"))
	return nil
}

func (v *ASTIRVisitor) visitAssign(ae *ast.AssignExpr) value.Value {
	rhs := v.visitExpr(ae.Value)
	if rhs == nil {
		return nil
	}

	switch target := ae.Target.(type) {
	case *ast.Identifier:
		if vi, ok := v.currentScope.Get(target.Name); ok {
			v.currentBlock.NewStore(rhs, vi.LLVMValue)
		} else {
			v.Errors = append(v.Errors, fmt.Errorf("undefined variable: %s", target.Name))
		}
	case *ast.IndexExpr:
		listVal := v.visitExpr(target.Target)
		if listVal == nil {
			v.Errors = append(v.Errors, fmt.Errorf("invalid list in indexed assignment"))
			return rhs
		}
		dataPtr := v.currentBlock.NewExtractValue(listVal, 1)
		idxVal := v.visitExpr(target.Index)
		if idxVal == nil {
			v.Errors = append(v.Errors, fmt.Errorf("invalid index expression"))
			return rhs
		}
		elemPtr := v.currentBlock.NewGetElementPtr(types.I64, dataPtr, idxVal)
		v.currentBlock.NewStore(rhs, elemPtr)
	}
	return rhs
}

func (v *ASTIRVisitor) visitCall(ce *ast.CallExpr) value.Value {
	var fn *ir.Func

	switch fun := ce.Function.(type) {
	case *ast.Identifier:
		name := fun.Name
		var ok bool
		fn, ok = v.funcs[name]
		if !ok {
			v.Errors = append(v.Errors, fmt.Errorf("undefined function: %s", name))
			return nil
		}
	default:
		v.Errors = append(v.Errors, fmt.Errorf("function calls on expressions not yet implemented"))
		return nil
	}

	if fn.Name() == "printf" {
		return v.buildPrintCall(fn, ce.Args)
	}

	var args []value.Value
	for _, a := range ce.Args {
		argVal := v.visitExpr(a)
		if argVal != nil {
			args = append(args, argVal)
		}
	}

	return v.currentBlock.NewCall(fn, args...)
}

func (v *ASTIRVisitor) buildPrintCall(fn *ir.Func, args []ast.Expr) value.Value {
	if len(args) == 0 {
		return nil
	}

	first := args[0]

	if is, ok := first.(*ast.InterpolatedStringExpr); ok {
		return v.handleInterpolatedPrintAST(fn, is)
	}

	argVal := v.visitExpr(first)
	if argVal == nil {
		return nil
	}

	var fmtStr string
	switch argVal.Type().(type) {
	case *types.IntType:
		fmtStr = "%lld\n"
	case *types.FloatType:
		fmtStr = "%f\n"
	default:
		fmtStr = "%s\n"
	}
	formatStr := v.defineGlobalString(fmtStr)
	return v.currentBlock.NewCall(fn, formatStr, argVal)
}

func (v *ASTIRVisitor) handleInterpolatedPrintAST(fn *ir.Func, is *ast.InterpolatedStringExpr) value.Value {
	var formatBuf strings.Builder
	var formatArgs []value.Value

	for _, p := range is.Parts {
		if p.IsExpr {
			if id, ok := p.Expr.(*ast.Identifier); ok {
				if vi, ok2 := v.currentScope.Get(id.Name); ok2 {
					val := v.currentBlock.NewLoad(vi.Type, vi.LLVMValue)
					formatArgs = append(formatArgs, val)
					if _, ok := vi.Type.(*types.PointerType); ok {
						formatBuf.WriteString("%s")
					} else if _, ok := vi.Type.(*types.FloatType); ok {
						formatBuf.WriteString("%f")
					} else {
						formatBuf.WriteString("%lld")
					}
				} else {
					v.Errors = append(v.Errors, fmt.Errorf("undefined variable: %s", id.Name))
					formatBuf.WriteString(id.Name)
				}
			}
		} else {
			formatBuf.WriteString(p.Text)
		}
	}
	formatBuf.WriteString("\n")

	formatStr := v.defineGlobalString(formatBuf.String())
	return v.currentBlock.NewCall(fn, append([]value.Value{formatStr}, formatArgs...)...)
}

func (v *ASTIRVisitor) visitIdent(id *ast.Identifier) value.Value {
	if vi, ok := v.currentScope.Get(id.Name); ok {
		if ptrTy, ok := vi.LLVMValue.Type().(*types.PointerType); ok {
			return v.currentBlock.NewLoad(ptrTy.ElemType, vi.LLVMValue)
		}
		return vi.LLVMValue
	}
	if fn, ok := v.funcs[id.Name]; ok {
		return fn
	}
	v.Errors = append(v.Errors, fmt.Errorf("undefined identifier: %s", id.Name))
	return nil
}

func (v *ASTIRVisitor) visitListLit(ll *ast.ListLiteral) value.Value {
	listTy := types.NewStruct(types.I64, types.NewPointer(types.I64))
	zero := constant.NewInt(types.I64, 0)

	if len(ll.Elements) == 0 {
		alloc := v.currentBlock.NewAlloca(listTy)
		nullPtr := constant.NewNull(types.NewPointer(types.I64))
		val := constant.NewZeroInitializer(listTy)
		withLen := constant.NewInsertValue(val, zero, 0)
		withPtr := constant.NewInsertValue(withLen, nullPtr, 1)
		v.currentBlock.NewStore(withPtr, alloc)
		return v.currentBlock.NewLoad(listTy, alloc)
	}

	var elemVals []value.Value
	for _, el := range ll.Elements {
		ev := v.visitExpr(el)
		if ev != nil {
			elemVals = append(elemVals, ev)
		}
	}

	count := int64(len(elemVals))
	arrTy := types.NewArray(uint64(count), types.I64)
	arrAlloc := v.currentBlock.NewAlloca(arrTy)

	for i, ev := range elemVals {
		idx := constant.NewInt(types.I64, int64(i))
		ptr := v.currentBlock.NewGetElementPtr(arrTy, arrAlloc, zero, idx)
		v.currentBlock.NewStore(ev, ptr)
	}

	firstPtr := v.currentBlock.NewGetElementPtr(arrTy, arrAlloc, zero, zero)

	listAlloc := v.currentBlock.NewAlloca(listTy)
	inserted := v.currentBlock.NewInsertValue(constant.NewZeroInitializer(listTy), constant.NewInt(types.I64, count), 0)
	inserted = v.currentBlock.NewInsertValue(inserted, firstPtr, 1)
	v.currentBlock.NewStore(inserted, listAlloc)

	return v.currentBlock.NewLoad(listTy, listAlloc)
}

func (v *ASTIRVisitor) visitIndex(ix *ast.IndexExpr) value.Value {
	sv := v.visitExpr(ix.Target)
	if sv == nil {
		v.Errors = append(v.Errors, fmt.Errorf("invalid list value"))
		return nil
	}
	idxVal := v.visitExpr(ix.Index)
	if idxVal == nil {
		v.Errors = append(v.Errors, fmt.Errorf("invalid list index"))
		return nil
	}
	dataPtr := v.currentBlock.NewExtractValue(sv, 1)
	elemPtr := v.currentBlock.NewGetElementPtr(types.I64, dataPtr, idxVal)
	return v.currentBlock.NewLoad(types.I64, elemPtr)
}

func (v *ASTIRVisitor) visitTernary(t *ast.TernaryExpr) value.Value {
	condVal := v.visitExpr(t.Condition)
	if condVal == nil {
		v.Errors = append(v.Errors, fmt.Errorf("invalid condition in ternary expression"))
		return nil
	}

	thenBlock := v.currentFunc.NewBlock(v.freshLabel("tern.then"))
	elseBlock := v.currentFunc.NewBlock(v.freshLabel("tern.else"))
	mergeBlock := v.currentFunc.NewBlock(v.freshLabel("tern.end"))

	v.currentBlock.NewCondBr(condVal, thenBlock, elseBlock)

	v.pushBlock(thenBlock)
	thenVal := v.visitExpr(t.Then)
	if v.currentBlock.Term == nil {
		v.currentBlock.NewBr(mergeBlock)
	}
	v.popBlock()

	v.pushBlock(elseBlock)
	elseVal := v.visitExpr(t.Else)
	if v.currentBlock.Term == nil {
		v.currentBlock.NewBr(mergeBlock)
	}
	v.popBlock()

	v.currentBlock = mergeBlock
	phi := v.currentBlock.NewPhi()
	phi.Incs = append(phi.Incs, ir.NewIncoming(thenVal, thenBlock))
	phi.Incs = append(phi.Incs, ir.NewIncoming(elseVal, elseBlock))
	return phi
}

func (v *ASTIRVisitor) visitInterpolatedString(is *ast.InterpolatedStringExpr) value.Value {
	var parts []string
	for _, p := range is.Parts {
		if p.IsExpr {
			if id, ok := p.Expr.(*ast.Identifier); ok {
				parts = append(parts, "$"+id.Name)
			}
		} else {
			parts = append(parts, p.Text)
		}
	}
	combined := strings.Join(parts, "")
	return v.defineGlobalString(combined)
}
