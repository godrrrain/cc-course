package visitor

import (
	"fmt"
	"github.com/AskaryanKarine/BMSTU-CC/cource/internal/parser"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/value"
)

func (v *IRVisitor) VisitStatementSequence(ctx *parser.StatementSequenceContext) interface{} {
	for _, stmt := range ctx.AllStatement() {
		v.Visit(stmt)
	}
	return nil
}

func (v *IRVisitor) VisitStatement(ctx *parser.StatementContext) interface{} {
	switch {
	case ctx.VariableDeclaration() != nil:
		v.Visit(ctx.VariableDeclaration())
	case ctx.AssignmentStatement() != nil:
		v.Visit(ctx.AssignmentStatement())
	case ctx.IoStatement() != nil:
		v.Visit(ctx.IoStatement())
	case ctx.IfStatement() != nil:
		v.Visit(ctx.IfStatement())
	case ctx.SwitchStatement() != nil:
		v.Visit(ctx.SwitchStatement())
	case ctx.LoopStatement() != nil:
		v.Visit(ctx.LoopStatement())
	case ctx.ExitStatement() != nil:
		v.Visit(ctx.ExitStatement())
	case ctx.StopStatement() != nil:
		v.Visit(ctx.StopStatement())
	default:
		return nil
	}

	return nil
}

func (v *IRVisitor) VisitAssignmentStatement(ctx *parser.AssignmentStatementContext) interface{} {
	if ctx.Lvalue() != nil && ctx.ASSIGN() != nil {
		lhs := v.Visit(ctx.Lvalue())
		rhs := v.Visit(ctx.Expression())

		lhsPtr, ok1 := lhs.(value.Value)
		rhsVal, ok2 := rhs.(value.Value)
		if !ok1 || !ok2 {
			v.Errors = append(v.Errors, fmt.Errorf("invalid assignment operands"))
			return nil
		}

		v.currentBlock.NewStore(rhsVal, lhsPtr)
		return nil
	}

	exprIfc := v.Visit(ctx.Expression())
	if exprIfc == nil {
		return nil
	}

	if fn, ok := exprIfc.(*ir.Func); ok && len(fn.Params) == 0 {
		v.currentBlock.NewCall(fn)
		return nil
	}
	return nil
}

func (v *IRVisitor) VisitIoStatement(ctx *parser.IoStatementContext) interface{} {
	printf := v.ensurePrintfFunc()

	argsVal := v.Visit(ctx.IoArgumentList())
	args, ok := argsVal.([]value.Value)
	if !ok {
		v.Errors = append(v.Errors, fmt.Errorf("invalid IO argument list"))
		return nil
	}

	if ctx.OUTPUT() == nil {
		v.Errors = append(v.Errors, fmt.Errorf("only OUTPUT is currently supported in IO statements"))
		return nil
	}

	formatPtr, formatArgs := v.buildPrintfFormat(args)
	callArgs := append([]value.Value{formatPtr}, formatArgs...)
	v.currentBlock.NewCall(printf, callArgs...)

	return nil
}

func (v *IRVisitor) VisitIfStatement(ctx *parser.IfStatementContext) interface{} {
	condVal := v.Visit(ctx.Expression())
	cond, ok := condVal.(value.Value)
	if !ok {
		v.Errors = append(v.Errors, fmt.Errorf("invalid condition in if statement"))
		return nil
	}

	parent := v.currentBlock

	thenBlock := v.newBlock("if.then")
	mergeBlock := v.newBlock("if.end")

	var elseBlock *ir.Block
	if ctx.ELSE() != nil {
		elseBlock = v.newBlock("if.else")
	} else {
		elseBlock = mergeBlock
	}

	parent.NewCondBr(cond, thenBlock, elseBlock)

	v.pushBlock(thenBlock)
	v.Visit(ctx.StatementSequence(0))
	if v.currentBlock.Term == nil {
		v.currentBlock.NewBr(mergeBlock)
	}
	v.popBlock()

	if ctx.ELSE() != nil {
		v.pushBlock(elseBlock)
		v.Visit(ctx.StatementSequence(1))
		if v.currentBlock.Term == nil {
			v.currentBlock.NewBr(mergeBlock)
		}
		v.popBlock()
	}

	v.currentBlock = mergeBlock

	return nil
}

func (v *IRVisitor) VisitSwitchStatement(ctx *parser.SwitchStatementContext) interface{} {
	postSwitch := v.currentFunc.NewBlock(v.freshLabel("switch_end"))

	nextCondBlock := v.currentBlock

	for idx, c := range ctx.AllCaseBlock() {
		condBlock := v.currentFunc.NewBlock(v.freshLabel(fmt.Sprintf("case%d_cond", idx)))
		bodyBlock := v.currentFunc.NewBlock(v.freshLabel(fmt.Sprintf("case%d_body", idx)))

		if nextCondBlock.Term == nil {
			nextCondBlock.NewBr(condBlock)
		}

		v.pushBlock(condBlock)
		condVal := v.Visit(c.Expression()).(value.Value)
		nextCondBlock = v.currentFunc.NewBlock(v.freshLabel(fmt.Sprintf("case%d_next", idx)))
		v.currentBlock.NewCondBr(condVal, bodyBlock, nextCondBlock)
		v.popBlock()

		v.pushBlock(bodyBlock)
		v.Visit(c.StatementSequence())
		if v.currentBlock.Term == nil {
			v.currentBlock.NewBr(postSwitch)
		}
		v.popBlock()
	}

	if ctx.ELSE() != nil {
		elseBlock := v.currentFunc.NewBlock(v.freshLabel("else_body"))
		if nextCondBlock.Term == nil {
			nextCondBlock.NewBr(elseBlock)
		}
		v.pushBlock(elseBlock)
		v.Visit(ctx.StatementSequence())
		if v.currentBlock.Term == nil {
			v.currentBlock.NewBr(postSwitch)
		}
		v.popBlock()
	} else if nextCondBlock.Term == nil {
		nextCondBlock.NewBr(postSwitch)
	}

	v.currentBlock = postSwitch
	return nil
}
