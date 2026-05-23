package visitor

import (
	"fmt"

	"github.com/AskaryanKarine/BMSTU-CC/cource/internal/parser"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/value"
)

func (v *IRVisitor) VisitStatements(ctx *parser.StatementsContext) interface{} {
	for _, stmt := range ctx.AllStatement() {
		v.Visit(stmt)
		if v.currentBlock == nil || v.currentBlock.Term != nil {
			break
		}
	}
	return nil
}

func (v *IRVisitor) VisitStatement(ctx *parser.StatementContext) interface{} {
	if ctx.NonLabelledStatement() != nil {
		v.Visit(ctx.NonLabelledStatement())
	}
	return nil
}

func (v *IRVisitor) VisitNonLabelledStatement(ctx *parser.NonLabelledStatementContext) interface{} {
	switch {
	case ctx.Block() != nil:
		v.Visit(ctx.Block())
	case ctx.VariableDeclaration() != nil:
		v.Visit(ctx.VariableDeclaration())
	case ctx.IfStatement() != nil:
		v.Visit(ctx.IfStatement())
	case ctx.ForStatement() != nil:
		v.Visit(ctx.ForStatement())
	case ctx.WhileStatement() != nil:
		v.Visit(ctx.WhileStatement())
	case ctx.DoStatement() != nil:
		v.Visit(ctx.DoStatement())
	case ctx.SwitchStatement() != nil:
		v.Visit(ctx.SwitchStatement())
	case ctx.TryStatement() != nil:
		v.Visit(ctx.TryStatement())
	case ctx.BreakStatement() != nil:
		v.Visit(ctx.BreakStatement())
	case ctx.ContinueStatement() != nil:
		v.Visit(ctx.ContinueStatement())
	case ctx.ReturnStatement() != nil:
		v.Visit(ctx.ReturnStatement())
	case ctx.YieldStatement() != nil:
		v.Visit(ctx.YieldStatement())
	case ctx.YieldEachStatement() != nil:
		v.Visit(ctx.YieldEachStatement())
	case ctx.ExpressionStatement() != nil:
		v.Visit(ctx.ExpressionStatement())
	case ctx.AssertStatement() != nil:
		v.Visit(ctx.AssertStatement())
	case ctx.LocalFunctionDeclaration() != nil:
		v.Visit(ctx.LocalFunctionDeclaration())
	case ctx.RethrowStatement() != nil:
		v.Visit(ctx.RethrowStatement())
	}
	return nil
}

func (v *IRVisitor) VisitBlock(ctx *parser.BlockContext) interface{} {
	v.enterScope()
	defer v.exitScope()

	if ctx.Statements() != nil {
		v.Visit(ctx.Statements())
	}
	return nil
}

func (v *IRVisitor) VisitIfStatement(ctx *parser.IfStatementContext) interface{} {
	if ctx.Expression() == nil {
		return nil
	}

	condVal := v.Visit(ctx.Expression())
	cond, ok := condVal.(value.Value)
	if !ok {
		v.Errors = append(v.Errors, fmt.Errorf("invalid condition in if statement"))
		return nil
	}

	parent := v.currentBlock

	thenBlock := v.currentFunc.NewBlock(v.freshLabel("if.then"))
	mergeBlock := v.currentFunc.NewBlock(v.freshLabel("if.end"))

	var elseBlock *ir.Block
	if ctx.ELSE_() != nil && len(ctx.AllStatement()) > 1 {
		elseBlock = v.currentFunc.NewBlock(v.freshLabel("if.else"))
	} else {
		elseBlock = mergeBlock
	}

	parent.NewCondBr(cond, thenBlock, elseBlock)

	v.pushBlock(thenBlock)
	v.Visit(ctx.Statement(0))
	if v.currentBlock.Term == nil {
		v.currentBlock.NewBr(mergeBlock)
	}
	v.popBlock()

	if ctx.ELSE_() != nil && len(ctx.AllStatement()) > 1 {
		v.pushBlock(elseBlock)
		v.Visit(ctx.Statement(1))
		if v.currentBlock.Term == nil {
			v.currentBlock.NewBr(mergeBlock)
		}
		v.popBlock()
	}

	v.currentBlock = mergeBlock
	return nil
}

func (v *IRVisitor) VisitWhileStatement(ctx *parser.WhileStatementContext) interface{} {
	if ctx.Expression() == nil {
		return nil
	}

	loopHeader := v.currentFunc.NewBlock(v.freshLabel("while.header"))
	loopBody := v.currentFunc.NewBlock(v.freshLabel("while.body"))
	loopExit := v.currentFunc.NewBlock(v.freshLabel("while.exit"))

	v.currentBlock.NewBr(loopHeader)

	v.pushBlock(loopHeader)
	cond := v.Visit(ctx.Expression())
	condVal, ok := cond.(value.Value)
	if !ok {
		v.Errors = append(v.Errors, fmt.Errorf("invalid condition in while statement"))
		return nil
	}
	v.currentBlock.NewCondBr(condVal, loopBody, loopExit)
	v.popBlock()

	v.pushBlock(loopBody)
	v.currentScope.setMeta(loopExitVar, loopExit)
	if ctx.Statement() != nil {
		v.Visit(ctx.Statement())
	}
	if v.currentBlock.Term == nil {
		v.currentBlock.NewBr(loopHeader)
	}
	v.popBlock()

	v.currentBlock = loopExit
	return nil
}

func (v *IRVisitor) VisitDoStatement(ctx *parser.DoStatementContext) interface{} {
	if ctx.Expression() == nil {
		return nil
	}

	loopBody := v.currentFunc.NewBlock(v.freshLabel("do.body"))
	loopCheck := v.currentFunc.NewBlock(v.freshLabel("do.check"))
	loopExit := v.currentFunc.NewBlock(v.freshLabel("do.exit"))

	v.currentBlock.NewBr(loopBody)

	v.pushBlock(loopBody)
	v.currentScope.setMeta(loopExitVar, loopExit)
	if ctx.Statement() != nil {
		v.Visit(ctx.Statement())
	}
	if v.currentBlock.Term == nil {
		v.currentBlock.NewBr(loopCheck)
	}
	v.popBlock()

	v.pushBlock(loopCheck)
	cond := v.Visit(ctx.Expression())
	condVal, ok := cond.(value.Value)
	if !ok {
		v.Errors = append(v.Errors, fmt.Errorf("invalid condition in do statement"))
		return nil
	}
	v.currentBlock.NewCondBr(condVal, loopBody, loopExit)
	v.popBlock()

	v.currentBlock = loopExit
	return nil
}

func (v *IRVisitor) VisitForStatement(ctx *parser.ForStatementContext) interface{} {
	if ctx.ForLoopParts() == nil {
		return nil
	}

	loopHeader := v.currentFunc.NewBlock(v.freshLabel("for.header"))
	loopBody := v.currentFunc.NewBlock(v.freshLabel("for.body"))
	loopStep := v.currentFunc.NewBlock(v.freshLabel("for.step"))
	loopExit := v.currentFunc.NewBlock(v.freshLabel("for.exit"))

	v.currentBlock.NewBr(loopHeader)

	v.pushBlock(loopHeader)
	v.currentBlock.NewBr(loopBody)
	v.popBlock()

	v.pushBlock(loopBody)
	v.currentScope.setMeta(loopExitVar, loopExit)
	if ctx.Statement() != nil {
		v.Visit(ctx.Statement())
	}
	if v.currentBlock.Term == nil {
		v.currentBlock.NewBr(loopStep)
	}
	v.popBlock()

	v.pushBlock(loopStep)
	v.currentBlock.NewBr(loopHeader)
	v.popBlock()

	v.currentBlock = loopExit
	return nil
}

func (v *IRVisitor) VisitSwitchStatement(ctx *parser.SwitchStatementContext) interface{} {
	if ctx.Expression() == nil {
		return nil
	}

	postSwitch := v.currentFunc.NewBlock(v.freshLabel("switch.end"))

	for _, sc := range ctx.AllSwitchCase() {
		v.Visit(sc)
	}
	if len(ctx.AllDefaultCase()) > 0 {
		v.Visit(ctx.DefaultCase(0))
	}

	v.currentBlock.NewBr(postSwitch)
	v.currentBlock = postSwitch
	return nil
}

func (v *IRVisitor) VisitSwitchCase(ctx *parser.SwitchCaseContext) interface{} {
	v.Errors = append(v.Errors, fmt.Errorf("switch cases not yet fully implemented"))
	if ctx.Statements() != nil {
		v.Visit(ctx.Statements())
	}
	return nil
}

func (v *IRVisitor) VisitDefaultCase(ctx *parser.DefaultCaseContext) interface{} {
	if ctx.Statements() != nil {
		v.Visit(ctx.Statements())
	}
	return nil
}

func (v *IRVisitor) VisitTryStatement(ctx *parser.TryStatementContext) interface{} {
	v.Errors = append(v.Errors, fmt.Errorf("try-catch not yet implemented"))
	if ctx.Block() != nil {
		v.Visit(ctx.Block())
	}
	return nil
}

func (v *IRVisitor) VisitBreakStatement(ctx *parser.BreakStatementContext) interface{} {
	if m, ok := v.currentScope.getMeta(loopExitVar); ok {
		if blk, ok2 := m.(*ir.Block); ok2 {
			if v.currentBlock.Term == nil {
				v.currentBlock.NewBr(blk)
			}
			unreachable := v.currentFunc.NewBlock(v.freshLabel("unreachable"))
			v.currentBlock = unreachable
			return nil
		}
	}
	return nil
}

func (v *IRVisitor) VisitContinueStatement(ctx *parser.ContinueStatementContext) interface{} {
	v.Errors = append(v.Errors, fmt.Errorf("continue not yet fully implemented"))
	return nil
}

func (v *IRVisitor) VisitReturnStatement(ctx *parser.ReturnStatementContext) interface{} {
	if v.currentFunc == nil {
		return nil
	}

	if ctx.Expression() != nil {
		retVal := v.Visit(ctx.Expression())
		if val, ok := retVal.(value.Value); ok {
			if v.currentBlock.Term == nil {
				v.currentBlock.NewRet(val)
			}
		}
	} else {
		if v.currentBlock.Term == nil {
			v.currentBlock.NewRet(nil)
		}
	}

	unreachable := v.currentFunc.NewBlock(v.freshLabel("unreachable"))
	v.currentBlock = unreachable
	return nil
}

func (v *IRVisitor) VisitYieldStatement(ctx *parser.YieldStatementContext) interface{} {
	v.Errors = append(v.Errors, fmt.Errorf("yield not yet implemented"))
	return nil
}

func (v *IRVisitor) VisitYieldEachStatement(ctx *parser.YieldEachStatementContext) interface{} {
	v.Errors = append(v.Errors, fmt.Errorf("yield* not yet implemented"))
	return nil
}

func (v *IRVisitor) VisitRethrowStatement(ctx *parser.RethrowStatementContext) interface{} {
	v.Errors = append(v.Errors, fmt.Errorf("rethrow not yet implemented"))
	return nil
}

func (v *IRVisitor) VisitLocalFunctionDeclaration(ctx *parser.LocalFunctionDeclarationContext) interface{} {
	v.Errors = append(v.Errors, fmt.Errorf("local function declarations not yet implemented"))
	return nil
}

func (v *IRVisitor) VisitAssertStatement(ctx *parser.AssertStatementContext) interface{} {
	return nil
}

func (v *IRVisitor) VisitExpressionStatement(ctx *parser.ExpressionStatementContext) interface{} {
	if ctx.Expression() != nil {
		v.Visit(ctx.Expression())
	}
	return nil
}
