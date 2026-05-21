package visitor

import (
	"fmt"
	"github.com/AskaryanKarine/BMSTU-CC/cource/internal/parser"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

type loopKind int

const (
	loopInfinite loopKind = iota
	loopWhile
	loopFor
	loopTimes
)

type loopInfo struct {
	kind    loopKind
	indVar  *VariableInfo
	endVal  value.Value
	stepVal value.Value
	condVal value.Value
}

func (v *IRVisitor) VisitLoopSpecifier(ctx *parser.LoopSpecifierContext) interface{} {
	preHeader := v.currentBlock

	if ctx == nil {
		return &loopInfo{kind: loopInfinite, condVal: constant.NewInt(types.I1, 1)}
	}

	switch {
	case ctx.FOR() != nil:
		idName := NormalizeFunctionName(ctx.ID().GetText())
		startIfc := v.Visit(ctx.Expression(0))
		startVal, ok := startIfc.(value.Value)
		if !ok {
			v.Errors = append(v.Errors, fmt.Errorf("FOR start expression not value"))
			return &loopInfo{kind: loopInfinite, condVal: constant.NewInt(types.I1, 0)}
		}
		endIfc := v.Visit(ctx.Expression(1))
		endVal, ok := endIfc.(value.Value)
		if !ok {
			v.Errors = append(v.Errors, fmt.Errorf("FOR end expression not value"))
			return &loopInfo{kind: loopInfinite, condVal: constant.NewInt(types.I1, 0)}
		}
		var step value.Value
		if ctx.STEP() != nil {
			stepIfc := v.Visit(ctx.Expression(2))
			step, ok = stepIfc.(value.Value)
			if !ok {
				v.Errors = append(v.Errors, fmt.Errorf("FOR step expression not value"))
				step = constant.NewInt(types.I32, 1)
			}
		} else {
			step = constant.NewInt(types.I32, 1)
		}
		alloca := preHeader.NewAlloca(startVal.Type())
		preHeader.NewStore(startVal, alloca)
		ind := &VariableInfo{Name: idName, Type: startVal.Type(), LLVMValue: alloca}
		_ = v.currentScope.Set(idName, ind)
		return &loopInfo{kind: loopFor, indVar: ind, endVal: endVal, stepVal: step}

	case ctx.WHILE() != nil:
		return &loopInfo{kind: loopWhile}

	case ctx.TIMES() != nil:
		cntIfc := v.Visit(ctx.Expression(0))
		countVal, ok := cntIfc.(value.Value)
		if !ok {
			v.Errors = append(v.Errors, fmt.Errorf("TIMES expression not value"))
			countVal = constant.NewInt(types.I32, 0)
		}
		alloca := preHeader.NewAlloca(countVal.Type())
		preHeader.NewStore(countVal, alloca)
		ind := &VariableInfo{Name: v.freshLabel("times"), Type: countVal.Type(), LLVMValue: alloca}
		_ = v.currentScope.Set(ind.Name, ind)
		return &loopInfo{kind: loopTimes, indVar: ind}
	}

	return &loopInfo{kind: loopInfinite, condVal: constant.NewInt(types.I1, 1)}
}

func (v *IRVisitor) VisitEndLoopCondition(ctx *parser.EndLoopConditionContext) interface{} {
	if ctx == nil {
		return constant.NewInt(types.I1, 0)
	}
	return v.Visit(ctx.Expression())
}

func (v *IRVisitor) VisitLoopStatement(ctx *parser.LoopStatementContext) interface{} {
	v.enterScope()
	defer v.exitScope()

	condBlk := v.currentFunc.NewBlock(v.freshLabel("loop.cond"))
	bodyBlk := v.currentFunc.NewBlock(v.freshLabel("loop.body"))
	stepBlk := v.currentFunc.NewBlock(v.freshLabel("loop.step"))
	exitBlk := v.currentFunc.NewBlock(v.freshLabel("loop.end"))
	v.currentScope.setMeta(loopExitVar, exitBlk)

	if v.currentBlock.Term == nil {
		v.currentBlock.NewBr(condBlk)
	}

	var (
		info *loopInfo
		ok   bool
	)
	if ls := ctx.LoopSpecifier(); ls != nil {
		info, ok = v.Visit(ls).(*loopInfo)
		if !ok {
			v.Errors = append(v.Errors, fmt.Errorf("loop specifier did not return loopInfo"))
			return nil
		}
	} else {
		info = &loopInfo{kind: loopInfinite}
	}

	v.pushBlock(condBlk)
	var condVal value.Value
	switch info.kind {
	case loopFor:
		cur := condBlk.NewLoad(info.indVar.Type, info.indVar.LLVMValue)
		condVal = condBlk.NewICmp(enum.IPredSLE, cur, info.endVal)
	case loopTimes:
		cur := condBlk.NewLoad(info.indVar.Type, info.indVar.LLVMValue)
		zero := constant.NewInt(types.I32, 0)
		condVal = condBlk.NewICmp(enum.IPredSGT, cur, zero)
	case loopWhile:
		condIfc := v.Visit(ctx.LoopSpecifier().Expression(0))
		tmp, ok := condIfc.(value.Value)
		if !ok {
			v.Errors = append(v.Errors, fmt.Errorf("while condition not value"))
			tmp = constant.NewInt(types.I1, 0)
		}
		condVal = tmp
	default:
		condVal = constant.NewInt(types.I1, 1)
	}
	info.condVal = condVal
	condBlk.NewCondBr(condVal, bodyBlk, exitBlk)
	v.popBlock()

	v.pushBlock(bodyBlk)
	v.Visit(ctx.StatementSequence())

	var toStepBlk bool = true
	if ctx.EndLoopCondition() != nil {
		endIfc := v.Visit(ctx.EndLoopCondition())
		endCond, ok := endIfc.(value.Value)
		if !ok {
			v.Errors = append(v.Errors, fmt.Errorf("end loop condition is not value"))
			endCond = constant.NewInt(types.I1, 0)
		}
		bodyBlk = v.currentBlock
		if bodyBlk.Term == nil {
			bodyBlk.NewCondBr(endCond, exitBlk, stepBlk)
			toStepBlk = false
		}
	}
	if toStepBlk {
		if v.currentBlock.Term == nil {
			v.currentBlock.NewBr(stepBlk)
		}
	}
	v.popBlock()

	v.pushBlock(stepBlk)
	switch info.kind {
	case loopFor:
		cur := stepBlk.NewLoad(info.indVar.Type, info.indVar.LLVMValue)
		next := stepBlk.NewAdd(cur, info.stepVal)
		stepBlk.NewStore(next, info.indVar.LLVMValue)
	case loopTimes:
		cur := stepBlk.NewLoad(info.indVar.Type, info.indVar.LLVMValue)
		one := constant.NewInt(types.I32, 1)
		next := stepBlk.NewSub(cur, one)
		stepBlk.NewStore(next, info.indVar.LLVMValue)
	}
	stepBlk.NewBr(condBlk)
	v.popBlock()

	v.currentBlock = exitBlk
	return nil
}
