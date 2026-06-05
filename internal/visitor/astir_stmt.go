package visitor

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"

	"cc-course/internal/ast"
)

func (v *ASTIRVisitor) visitStmt(s ast.Stmt) {
	if s == nil {
		return
	}
	switch ss := s.(type) {
	case *ast.BlockStmt:
		v.visitBlockStmt(ss)
	case *ast.IfStmt:
		v.visitIfStmt(ss)
	case *ast.WhileStmt:
		v.visitWhileStmt(ss)
	case *ast.DoWhileStmt:
		v.visitDoWhileStmt(ss)
	case *ast.ForStmt:
		v.visitForStmt(ss)
	case *ast.ForInStmt:
		v.visitForInStmt(ss)
	case *ast.ExprStmt:
		v.visitExprStmt(ss)
	case *ast.ReturnStmt:
		v.visitReturnStmt(ss)
	case *ast.BreakStmt:
		v.visitBreakStmt()
	case *ast.ContinueStmt:
		v.visitContinueStmt()
	case *ast.AssertStmt:
		v.visitAssertStmt(ss)
	case *ast.VarDecl:
		v.visitVarDecl(ss)
	}
}

func (v *ASTIRVisitor) visitBlockStmt(bs *ast.BlockStmt) {
	v.enterScope()
	defer v.exitScope()

	for _, s := range bs.Statements {
		v.visitStmt(s)
		if v.currentBlock == nil || v.currentBlock.Term != nil {
			break
		}
	}
}

func (v *ASTIRVisitor) visitIfStmt(is *ast.IfStmt) {
	if is.Condition == nil {
		return
	}

	condVal := v.visitExpr(is.Condition)
	if condVal == nil {
		v.Errors = append(v.Errors, errInvalidCondition)
		return
	}

	thenBlock := v.currentFunc.NewBlock(v.freshLabel("if.then"))
	mergeBlock := v.currentFunc.NewBlock(v.freshLabel("if.end"))

	var elseBlock *ir.Block
	if is.Else != nil {
		elseBlock = v.currentFunc.NewBlock(v.freshLabel("if.else"))
	} else {
		elseBlock = mergeBlock
	}

	v.currentBlock.NewCondBr(condVal, thenBlock, elseBlock)

	v.pushBlock(thenBlock)
	v.visitStmt(is.Then)
	if v.currentBlock.Term == nil {
		v.currentBlock.NewBr(mergeBlock)
	}
	v.popBlock()

	if is.Else != nil {
		v.pushBlock(elseBlock)
		v.visitStmt(is.Else)
		if v.currentBlock.Term == nil {
			v.currentBlock.NewBr(mergeBlock)
		}
		v.popBlock()
	}

	v.currentBlock = mergeBlock
}

func (v *ASTIRVisitor) visitWhileStmt(ws *ast.WhileStmt) {
	if ws.Condition == nil {
		return
	}

	loopHeader := v.currentFunc.NewBlock(v.freshLabel("while.header"))
	loopBody := v.currentFunc.NewBlock(v.freshLabel("while.body"))
	loopExit := v.currentFunc.NewBlock(v.freshLabel("while.exit"))

	v.currentBlock.NewBr(loopHeader)

	v.pushBlock(loopHeader)
	condVal := v.visitExpr(ws.Condition)
	if condVal == nil {
		v.Errors = append(v.Errors, errInvalidCondition)
		return
	}
	v.currentBlock.NewCondBr(condVal, loopBody, loopExit)
	v.popBlock()

	v.pushBlock(loopBody)
	v.currentScope.setMeta(loopExitVar, loopExit)
	if ws.Body != nil {
		v.visitStmt(ws.Body)
	}
	if v.currentBlock.Term == nil {
		v.currentBlock.NewBr(loopHeader)
	}
	v.popBlock()

	v.currentBlock = loopExit
}

func (v *ASTIRVisitor) visitDoWhileStmt(ds *ast.DoWhileStmt) {
	if ds.Condition == nil {
		return
	}

	loopBody := v.currentFunc.NewBlock(v.freshLabel("do.body"))
	loopCheck := v.currentFunc.NewBlock(v.freshLabel("do.check"))
	loopExit := v.currentFunc.NewBlock(v.freshLabel("do.exit"))

	v.currentBlock.NewBr(loopBody)

	v.pushBlock(loopBody)
	v.currentScope.setMeta(loopExitVar, loopExit)
	if ds.Body != nil {
		v.visitStmt(ds.Body)
	}
	if v.currentBlock.Term == nil {
		v.currentBlock.NewBr(loopCheck)
	}
	v.popBlock()

	v.pushBlock(loopCheck)
	condVal := v.visitExpr(ds.Condition)
	if condVal == nil {
		v.Errors = append(v.Errors, errInvalidCondition)
		return
	}
	v.currentBlock.NewCondBr(condVal, loopBody, loopExit)
	v.popBlock()

	v.currentBlock = loopExit
}

func (v *ASTIRVisitor) visitForStmt(fs *ast.ForStmt) {
	v.enterScope()
	defer v.exitScope()

	// Init
	if fs.Init != nil {
		v.visitStmt(fs.Init)
	}

	loopCond := v.currentFunc.NewBlock(v.freshLabel("for.cond"))
	loopBody := v.currentFunc.NewBlock(v.freshLabel("for.body"))
	loopStep := v.currentFunc.NewBlock(v.freshLabel("for.step"))
	loopExit := v.currentFunc.NewBlock(v.freshLabel("for.exit"))

	v.currentBlock.NewBr(loopCond)

	// Condition
	v.pushBlock(loopCond)
	if fs.Condition != nil {
		condVal := v.visitExpr(fs.Condition)
		if condVal == nil {
			v.Errors = append(v.Errors, errInvalidCondition)
			return
		}
		v.currentBlock.NewCondBr(condVal, loopBody, loopExit)
	} else {
		v.currentBlock.NewBr(loopBody)
	}
	v.popBlock()

	// Body
	v.pushBlock(loopBody)
	v.currentScope.setMeta(loopExitVar, loopExit)
	if fs.Body != nil {
		v.visitStmt(fs.Body)
	}
	if v.currentBlock.Term == nil {
		v.currentBlock.NewBr(loopStep)
	}
	v.popBlock()

	// Step
	v.pushBlock(loopStep)
	if fs.Update != nil {
		v.visitExpr(fs.Update)
	}
	if v.currentBlock.Term == nil {
		v.currentBlock.NewBr(loopCond)
	}
	v.popBlock()

	v.currentBlock = loopExit
}

func (v *ASTIRVisitor) visitForInStmt(fis *ast.ForInStmt) {
	v.enterScope()
	defer v.exitScope()

	loopTyp := v.astTypeToLLVM(fis.VarType)

	iterVal := v.visitExpr(fis.Iterable)
	if iterVal == nil {
		v.Errors = append(v.Errors, errInvalidIterable)
		return
	}

	lenVal := v.currentBlock.NewExtractValue(iterVal, 0)
	dataPtrVal := v.currentBlock.NewExtractValue(iterVal, 1)

	varAlloca := v.currentBlock.NewAlloca(loopTyp)
	v.currentScope.Set(fis.VarName, &VariableInfo{
		Name:      fis.VarName,
		Type:      loopTyp,
		LLVMValue: varAlloca,
	})

	idxAlloca := v.currentBlock.NewAlloca(types.I64)
	zero := constant.NewInt(types.I64, 0)
	v.currentBlock.NewStore(zero, idxAlloca)

	loopCond := v.currentFunc.NewBlock(v.freshLabel("forin.cond"))
	loopBody := v.currentFunc.NewBlock(v.freshLabel("forin.body"))
	loopStep := v.currentFunc.NewBlock(v.freshLabel("forin.step"))
	loopExit := v.currentFunc.NewBlock(v.freshLabel("forin.exit"))

	v.currentBlock.NewBr(loopCond)

	// Condition
	v.pushBlock(loopCond)
	idxLoad := v.currentBlock.NewLoad(types.I64, idxAlloca)
	cond := v.currentBlock.NewICmp(enum.IPredSLT, idxLoad, lenVal)
	v.currentBlock.NewCondBr(cond, loopBody, loopExit)
	v.popBlock()

	// Body
	v.pushBlock(loopBody)
	v.currentScope.setMeta(loopExitVar, loopExit)

	curIdx := v.currentBlock.NewLoad(types.I64, idxAlloca)
	elemPtr := v.currentBlock.NewGetElementPtr(types.I64, dataPtrVal, curIdx)
	elemVal := v.currentBlock.NewLoad(types.I64, elemPtr)
	v.currentBlock.NewStore(elemVal, varAlloca)

	if fis.Body != nil {
		v.visitStmt(fis.Body)
	}
	if v.currentBlock.Term == nil {
		v.currentBlock.NewBr(loopStep)
	}
	v.popBlock()

	// Step
	v.pushBlock(loopStep)
	stepIdx := v.currentBlock.NewLoad(types.I64, idxAlloca)
	nextIdx := v.currentBlock.NewAdd(constant.NewInt(types.I64, 1), stepIdx)
	v.currentBlock.NewStore(nextIdx, idxAlloca)
	if v.currentBlock.Term == nil {
		v.currentBlock.NewBr(loopCond)
	}
	v.popBlock()

	v.currentBlock = loopExit
}

func (v *ASTIRVisitor) visitExprStmt(es *ast.ExprStmt) {
	if es.Expression != nil {
		v.visitExpr(es.Expression)
	}
}

func (v *ASTIRVisitor) visitReturnStmt(rs *ast.ReturnStmt) {
	if v.currentFunc == nil {
		return
	}

	if rs.Value != nil {
		retVal := v.visitExpr(rs.Value)
		if retVal != nil && v.currentBlock.Term == nil {
			v.currentBlock.NewRet(retVal)
		}
	} else {
		if v.currentBlock.Term == nil {
			v.currentBlock.NewRet(nil)
		}
	}

	unreachable := v.currentFunc.NewBlock(v.freshLabel("unreachable"))
	v.currentBlock = unreachable
}

func (v *ASTIRVisitor) visitBreakStmt() {
	if m, ok := v.currentScope.getMeta(loopExitVar); ok {
		if blk, ok2 := m.(*ir.Block); ok2 {
			if v.currentBlock.Term == nil {
				v.currentBlock.NewBr(blk)
			}
			unreachable := v.currentFunc.NewBlock(v.freshLabel("unreachable"))
			v.currentBlock = unreachable
		}
	}
}

func (v *ASTIRVisitor) visitContinueStmt() {
	// Placeholder
}

func (v *ASTIRVisitor) visitAssertStmt(as *ast.AssertStmt) {
	// No-op for now
}

var (
	errInvalidCondition = &irError{"invalid condition"}
	errInvalidIterable  = &irError{"invalid iterable in for-in loop"}
)

type irError struct {
	msg string
}

func (e *irError) Error() string { return e.msg }
