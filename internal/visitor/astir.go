package visitor

import (
	"fmt"
	"os"
	"strings"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"

	"cc-course/internal/ast"
)

type ASTIRVisitor struct {
	Errors []error
	Module *ir.Module

	currentFunc  *ir.Func
	currentScope *Scope
	currentBlock *ir.Block
	blockStack   []*ir.Block
	funcs        map[string]*ir.Func
	runtime      map[string]value.Value
}

func NewASTIRVisitor() *ASTIRVisitor {
	module := ir.NewModule()
	module.TargetTriple = "x86_64-pc-windows-msvc19.39.33523"

	v := &ASTIRVisitor{
		Errors: make([]error, 0),
		Module: module,

		currentFunc:  nil,
		currentScope: nil,
		funcs:        make(map[string]*ir.Func),
		runtime:      make(map[string]value.Value),
	}

	v.registerBuiltins()
	v.declareRuntimeFuncs()
	v.funcs["readString"] = v.defineReadString()
	v.funcs["readInt"] = v.defineReadInt()

	return v
}

func (v *ASTIRVisitor) registerBuiltins() {
	printf := v.Module.NewFunc("printf", types.I32,
		ir.NewParam("format", types.NewPointer(types.I8)))
	printf.Sig.Variadic = true
	v.funcs["printf"] = printf
	v.funcs["print"] = printf

	system := v.Module.NewFunc("system", types.I32,
		ir.NewParam("command", types.NewPointer(types.I8)))
	v.funcs["system"] = system
}

func (v *ASTIRVisitor) declareRuntimeFuncs() {
	iob := v.Module.NewFunc("__acrt_iob_func", types.NewPointer(types.I8),
		ir.NewParam("fd", types.I32))
	v.runtime["__acrt_iob_func"] = iob

	fgets := v.Module.NewFunc("fgets", types.NewPointer(types.I8),
		ir.NewParam("buf", types.NewPointer(types.I8)),
		ir.NewParam("n", types.I32),
		ir.NewParam("stream", types.NewPointer(types.I8)))
	v.runtime["fgets"] = fgets

	malloc := v.Module.NewFunc("malloc", types.NewPointer(types.I8),
		ir.NewParam("size", types.I64))
	v.runtime["malloc"] = malloc

	strlen := v.Module.NewFunc("strlen", types.I64,
		ir.NewParam("s", types.NewPointer(types.I8)))
	v.runtime["strlen"] = strlen

	atoll := v.Module.NewFunc("atoll", types.I64,
		ir.NewParam("s", types.NewPointer(types.I8)))
	v.runtime["atoll"] = atoll
}

func (v *ASTIRVisitor) defineReadString() *ir.Func {
	fn := v.Module.NewFunc("readString", types.NewPointer(types.I8))

	entry := fn.NewBlock("")

	malloc := v.runtime["malloc"].(*ir.Func)
	buf := entry.NewCall(malloc, constant.NewInt(types.I64, 256))

	iob := v.runtime["__acrt_iob_func"].(*ir.Func)
	stdin := entry.NewCall(iob, constant.NewInt(types.I32, 0))

	fgets := v.runtime["fgets"].(*ir.Func)
	entry.NewCall(fgets, buf, constant.NewInt(types.I32, 256), stdin)

	strlen := v.runtime["strlen"].(*ir.Func)
	lenVal := entry.NewCall(strlen, buf)

	checkNl := fn.NewBlock("")
	done := fn.NewBlock("")

	cmp := entry.NewICmp(enum.IPredSGT, lenVal, constant.NewInt(types.I64, 0))
	entry.NewCondBr(cmp, checkNl, done)

	lastIdx := checkNl.NewSub(lenVal, constant.NewInt(types.I64, 1))
	lastPtr := checkNl.NewGetElementPtr(types.I8, buf, lastIdx)
	lastChar := checkNl.NewLoad(types.I8, lastPtr)
	isNl := checkNl.NewICmp(enum.IPredEQ, lastChar, constant.NewInt(types.I8, 10))
	strip := fn.NewBlock("")
	checkNl.NewCondBr(isNl, strip, done)

	strip.NewStore(constant.NewInt(types.I8, 0), lastPtr)
	strip.NewBr(done)

	done.NewRet(buf)

	return fn
}

func (v *ASTIRVisitor) defineReadInt() *ir.Func {
	fn := v.Module.NewFunc("readInt", types.I64)

	entry := fn.NewBlock("")
	buf := entry.NewAlloca(types.NewArray(64, types.I8))
	zero := constant.NewInt(types.I64, 0)
	ptr := entry.NewGetElementPtr(types.NewArray(64, types.I8), buf, zero, zero)

	iob := v.runtime["__acrt_iob_func"].(*ir.Func)
	stdin := entry.NewCall(iob, constant.NewInt(types.I32, 0))

	fgets := v.runtime["fgets"].(*ir.Func)
	bufSize := constant.NewInt(types.I32, 64)
	entry.NewCall(fgets, ptr, bufSize, stdin)

	atoll := v.runtime["atoll"].(*ir.Func)
	val := entry.NewCall(atoll, ptr)

	entry.NewRet(val)

	return fn
}

// ----- Type conversion -----

func (v *ASTIRVisitor) astTypeToLLVM(t *ast.Type) types.Type {
	if t == nil {
		return types.I64
	}
	switch {
	case t.Name == "int":
		return types.I64
	case t.Name == "double":
		return types.Double
	case t.Name == "bool":
		return types.I1
	case t.Name == "String":
		return types.NewPointer(types.I8)
	case t.Name == "void":
		return types.Void
	case strings.HasPrefix(t.Name, "List"):
		return types.NewStruct(types.I64, types.NewPointer(types.I64))
	default:
		return types.I64
	}
}

// ----- Helpers -----

func (v *ASTIRVisitor) defineGlobalString(s string) value.Value {
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

func (v *ASTIRVisitor) enterScope() {
	v.currentScope = NewScope(v.currentScope)
}

func (v *ASTIRVisitor) exitScope() {
	if v.currentScope == nil {
		v.Errors = append(v.Errors, fmt.Errorf("attempt to exit scope when no scope is active"))
		return
	}
	v.currentScope = v.currentScope.parent
}

func (v *ASTIRVisitor) newBlock(label string) *ir.Block {
	if v.currentFunc == nil {
		return nil
	}
	return v.currentFunc.NewBlock(v.freshLabel(label))
}

func (v *ASTIRVisitor) freshLabel(prefix string) string {
	if v.currentFunc == nil {
		return prefix
	}
	return fmt.Sprintf("%s.%d", prefix, len(v.currentFunc.Blocks))
}

func (v *ASTIRVisitor) pushBlock(block *ir.Block) {
	v.blockStack = append(v.blockStack, v.currentBlock)
	v.currentBlock = block
}

func (v *ASTIRVisitor) popBlock() {
	if len(v.blockStack) == 0 {
		return
	}
	v.currentBlock = v.blockStack[len(v.blockStack)-1]
	v.blockStack = v.blockStack[:len(v.blockStack)-1]
}

func (v *ASTIRVisitor) castToMatch(lhs, rhs value.Value) (value.Value, value.Value) {
	lt, rt := lhs.Type(), rhs.Type()

	if lt.Equal(rt) {
		return lhs, rhs
	}

	if _, ok := lt.(*types.IntType); ok {
		if _, isFloat := rt.(*types.FloatType); isFloat {
			lhs = v.currentBlock.NewSIToFP(lhs, rt)
			return lhs, rhs
		}
	}
	if _, ok := rt.(*types.IntType); ok {
		if _, isFloat := lt.(*types.FloatType); isFloat {
			rhs = v.currentBlock.NewSIToFP(rhs, lt)
			return lhs, rhs
		}
	}

	return lhs, rhs
}

func (v *ASTIRVisitor) ensureI32(idx value.Value) value.Value {
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

// ----- Top-level visit -----

func (v *ASTIRVisitor) VisitProgram(prog *ast.Program) {
	v.enterScope()
	defer v.exitScope()

	for _, d := range prog.Declarations {
		if fd, ok := d.(*ast.FuncDecl); ok {
			v.declareFuncPrototype(fd)
		}
	}

	for _, d := range prog.Declarations {
		switch dd := d.(type) {
		case *ast.FuncDecl:
			v.visitFuncDecl(dd)
		case *ast.VarDecl:
			v.visitVarDecl(dd)
		}
	}
}

func (v *ASTIRVisitor) GetModule() *ir.Module {
	return v.Module
}

func (v *ASTIRVisitor) WriteToFile(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer f.Close()

	if _, err := v.Module.WriteTo(f); err != nil {
		return fmt.Errorf("failed to write LLVM IR to file: %w", err)
	}
	return nil
}

// ----- Function declarations -----

func (v *ASTIRVisitor) declareFuncPrototype(fd *ast.FuncDecl) {
	if _, exists := v.funcs[fd.Name]; exists {
		return
	}

	retType := v.astTypeToLLVM(fd.ReturnType)
	if fd.Name == "main" {
		retType = types.I32
	}

	var irParams []*ir.Param
	for _, p := range fd.Params {
		irParams = append(irParams, ir.NewParam(p.Name, v.astTypeToLLVM(p.Type)))
	}

	fn := v.Module.NewFunc(fd.Name, retType, irParams...)
	v.funcs[fd.Name] = fn
}

func (v *ASTIRVisitor) visitFuncDecl(fd *ast.FuncDecl) {
	fn, exists := v.funcs[fd.Name]
	if !exists {
		retType := v.astTypeToLLVM(fd.ReturnType)
		if fd.Name == "main" {
			retType = types.I32
		}
		var irParams []*ir.Param
		for _, p := range fd.Params {
			irParams = append(irParams, ir.NewParam(p.Name, v.astTypeToLLVM(p.Type)))
		}
		fn = v.Module.NewFunc(fd.Name, retType, irParams...)
		v.funcs[fd.Name] = fn
	} else if len(fn.Blocks) > 0 {
		return
	}

	entry := fn.NewBlock("entry")
	v.currentFunc = fn
	v.currentBlock = entry
	v.pushBlock(entry)

	v.enterScope()
	defer v.exitScope()

	if !fn.Sig.RetType.Equal(types.Void) {
		retAlloca := entry.NewAlloca(fn.Sig.RetType)
		var zeroVal value.Value
		switch t := fn.Sig.RetType.(type) {
		case *types.IntType:
			zeroVal = constant.NewInt(t, 0)
		case *types.FloatType:
			zeroVal = constant.NewFloat(t, 0)
		default:
			zeroVal = constant.NewZeroInitializer(t)
		}
		entry.NewStore(zeroVal, retAlloca)
		v.currentScope.Set(returnNameVar, &VariableInfo{
			Name:      returnNameVar,
			Type:      fn.Sig.RetType,
			LLVMValue: retAlloca,
		})
	}

	for i, p := range fd.Params {
		llvmParam := fn.Params[i]
		alloca := entry.NewAlloca(v.astTypeToLLVM(p.Type))
		entry.NewStore(llvmParam, alloca)
		v.currentScope.Set(p.Name, &VariableInfo{
			Name:      p.Name,
			Type:      v.astTypeToLLVM(p.Type),
			LLVMValue: alloca,
		})
	}

	if fd.Name == "main" {
		if systemFn, ok := v.funcs["system"]; ok {
			cmd := v.defineGlobalString("chcp 65001 > nul")
			v.currentBlock.NewCall(systemFn, cmd)
		}
	}

	if fd.Body != nil {
		v.visitBlockStmt(fd.Body)
	}

	if fd.BodyExpr != nil {
		retVal := v.visitExpr(fd.BodyExpr)
		if v.currentBlock.Term == nil {
			v.currentBlock.NewRet(retVal)
		}
	}

	if v.currentBlock.Term == nil {
		if !fn.Sig.RetType.Equal(types.Void) {
			if retVar, ok := v.currentScope.Get(returnNameVar); ok {
				retVal := v.currentBlock.NewLoad(retVar.Type, retVar.LLVMValue)
				v.currentBlock.NewRet(retVal)
			} else {
				zero := constant.NewInt(fn.Sig.RetType.(*types.IntType), 0)
				v.currentBlock.NewRet(zero)
			}
		} else {
			v.currentBlock.NewRet(nil)
		}
	}

	v.popBlock()
	v.currentFunc = nil
}

func (v *ASTIRVisitor) visitVarDecl(vd *ast.VarDecl) {
	var initVal value.Value
	if vd.Init != nil {
		initVal = v.visitExpr(vd.Init)
	}

	typ := v.astTypeToLLVM(vd.Type)
	if initVal != nil {
		typ = initVal.Type()
	}

	alloca := v.currentBlock.NewAlloca(typ)
	vi := &VariableInfo{
		Name:      vd.Name,
		Type:      typ,
		LLVMValue: alloca,
	}

	if err := v.currentScope.Set(vd.Name, vi); err != nil {
		v.Errors = append(v.Errors, fmt.Errorf("variable %s already declared: %v", vd.Name, err))
	}

	if initVal != nil {
		v.currentBlock.NewStore(initVal, alloca)
	}
}
