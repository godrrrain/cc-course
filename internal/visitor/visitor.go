package visitor

import (
	"fmt"

	"github.com/AskaryanKarine/BMSTU-CC/cource/internal/parser"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"

	"github.com/antlr4-go/antlr/v4"
)

type IRVisitor struct {
	*parser.BaseDart2ParserVisitor

	Errors []error
	Module *ir.Module

	currentFunc  *ir.Func
	currentScope *Scope
	currentBlock *ir.Block
	blockStack   []*ir.Block
	funcs        map[string]*ir.Func
}

func NewIRVisitor() *IRVisitor {
	module := ir.NewModule()
	module.TargetTriple = "x86_64-pc-windows-msvc19.39.33523"

	v := &IRVisitor{
		BaseDart2ParserVisitor: &parser.BaseDart2ParserVisitor{},

		Errors: make([]error, 0),
		Module: module,

		currentFunc:  nil,
		currentScope: nil,
		funcs:        make(map[string]*ir.Func),
	}

	v.registerBuiltins()

	return v
}

func (v *IRVisitor) registerBuiltins() {
	printf := v.Module.NewFunc("printf", types.I32,
		ir.NewParam("format", types.NewPointer(types.I8)))
	printf.Sig.Variadic = true
	v.funcs["printf"] = printf
	v.funcs["print"] = printf

	system := v.Module.NewFunc("system", types.I32,
		ir.NewParam("command", types.NewPointer(types.I8)))
	v.funcs["system"] = system
}

func (v *IRVisitor) GetModule() *ir.Module {
	return v.Module
}

func (v *IRVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *IRVisitor) VisitCompilationUnit(ctx *parser.CompilationUnitContext) interface{} {
	// First pass: declare all function prototypes (for recursive calls)
	for _, item := range ctx.AllTopLevelItem() {
		if item.TopLevelDeclaration() != nil &&
			item.TopLevelDeclaration().FunctionDeclaration() != nil {
			v.declareFunctionPrototype(item.TopLevelDeclaration().FunctionDeclaration())
		}
	}

	// Second pass: visit all top-level items
	v.enterScope()
	defer v.exitScope()

	for _, item := range ctx.AllTopLevelItem() {
		v.Visit(item)
	}

	return nil
}

func (v *IRVisitor) VisitTopLevelItem(ctx *parser.TopLevelItemContext) interface{} {
	switch {
	case ctx.TopLevelDeclaration() != nil:
		v.Visit(ctx.TopLevelDeclaration())
	case ctx.ImportOrExportStatement() != nil:
		v.Visit(ctx.ImportOrExportStatement())
	case ctx.LibraryStatement() != nil:
		v.Visit(ctx.LibraryStatement())
	case ctx.PartStatement() != nil:
		v.Visit(ctx.PartStatement())
	}
	return nil
}

func (v *IRVisitor) VisitTopLevelDeclaration(ctx *parser.TopLevelDeclarationContext) interface{} {
	switch {
	case ctx.FunctionDeclaration() != nil:
		v.Visit(ctx.FunctionDeclaration())
	case ctx.VariableDeclaration() != nil:
		v.Visit(ctx.VariableDeclaration())
	case ctx.ClassDeclaration() != nil:
		v.Visit(ctx.ClassDeclaration())
	case ctx.EnumDeclaration() != nil:
		v.Visit(ctx.EnumDeclaration())
	case ctx.MixinDeclaration() != nil:
		v.Visit(ctx.MixinDeclaration())
	case ctx.ExtensionDeclaration() != nil:
		v.Visit(ctx.ExtensionDeclaration())
	case ctx.TypeAliasDeclaration() != nil:
		v.Visit(ctx.TypeAliasDeclaration())
	case ctx.GetterDeclaration() != nil:
		v.Visit(ctx.GetterDeclaration())
	case ctx.SetterDeclaration() != nil:
		v.Visit(ctx.SetterDeclaration())
	}
	return nil
}

func (v *IRVisitor) VisitImportOrExportStatement(ctx *parser.ImportOrExportStatementContext) interface{} {
	return nil
}

func (v *IRVisitor) VisitLibraryStatement(ctx *parser.LibraryStatementContext) interface{} {
	return nil
}

func (v *IRVisitor) VisitPartStatement(ctx *parser.PartStatementContext) interface{} {
	return nil
}

func (v *IRVisitor) VisitClassDeclaration(ctx *parser.ClassDeclarationContext) interface{} {
	v.Errors = append(v.Errors, fmt.Errorf("class declarations not yet implemented"))
	return nil
}

func (v *IRVisitor) VisitEnumDeclaration(ctx *parser.EnumDeclarationContext) interface{} {
	v.Errors = append(v.Errors, fmt.Errorf("enum declarations not yet implemented"))
	return nil
}

func (v *IRVisitor) VisitMixinDeclaration(ctx *parser.MixinDeclarationContext) interface{} {
	v.Errors = append(v.Errors, fmt.Errorf("mixin declarations not yet implemented"))
	return nil
}

func (v *IRVisitor) VisitExtensionDeclaration(ctx *parser.ExtensionDeclarationContext) interface{} {
	v.Errors = append(v.Errors, fmt.Errorf("extension declarations not yet implemented"))
	return nil
}

func (v *IRVisitor) VisitTypeAliasDeclaration(ctx *parser.TypeAliasDeclarationContext) interface{} {
	return nil
}

func (v *IRVisitor) VisitGetterDeclaration(ctx *parser.GetterDeclarationContext) interface{} {
	v.Errors = append(v.Errors, fmt.Errorf("getter declarations not yet implemented"))
	return nil
}

func (v *IRVisitor) VisitSetterDeclaration(ctx *parser.SetterDeclarationContext) interface{} {
	v.Errors = append(v.Errors, fmt.Errorf("setter declarations not yet implemented"))
	return nil
}

func (v *IRVisitor) VisitFunctionDeclaration(ctx *parser.FunctionDeclarationContext) interface{} {
	if ctx.IDENTIFIER() == nil {
		return nil
	}

	funcName := ctx.IDENTIFIER().GetText()

	var retType types.Type = types.Void
	if ctx.ReturnType() != nil {
		if t := v.Visit(ctx.ReturnType()); t != nil {
			if tt, ok := t.(types.Type); ok {
				retType = tt
			}
		}
	}

	if funcName == "main" {
		retType = types.I32
	}

	var irParams []*ir.Param
	if fpl := ctx.FormalParameterList(); fpl != nil {
		if params := v.Visit(fpl); params != nil {
			if paramInfos, ok := params.([]*VariableInfo); ok {
				for _, p := range paramInfos {
					irParams = append(irParams, ir.NewParam(p.Name, p.Type))
				}
			}
		}
	}

	fn, exists := v.funcs[funcName]
	if !exists {
		fn = v.Module.NewFunc(funcName, retType, irParams...)
		v.funcs[funcName] = fn
	} else if len(fn.Blocks) > 0 {
		return nil
	}

	if ctx.FunctionBody() != nil {
		entry := fn.NewBlock("entry")
		v.currentFunc = fn
		v.currentBlock = entry
		v.pushBlock(entry)

		v.enterScope()
		defer v.exitScope()

		if !types.Equal(fn.Sig.RetType, types.Void) {
			retAlloca := entry.NewAlloca(fn.Sig.RetType)
			_ = v.currentScope.Set(returnNameVar, &VariableInfo{
				Name:      returnNameVar,
				Type:      fn.Sig.RetType,
				LLVMValue: retAlloca,
			})
		}

		if fpl := ctx.FormalParameterList(); fpl != nil {
			if params := v.Visit(fpl); params != nil {
				if paramInfos, ok := params.([]*VariableInfo); ok {
					for i, pInfo := range paramInfos {
						llvmParam := fn.Params[i]
						alloca := entry.NewAlloca(pInfo.Type)
						entry.NewStore(llvmParam, alloca)
						pInfo.LLVMValue = alloca
						_ = v.currentScope.Set(pInfo.Name, pInfo)
					}
				}
			}
		}

		if funcName == "main" {
			if systemFn, ok := v.funcs["system"]; ok {
				cmd := v.defineGlobalString("chcp 65001 > nul")
				v.currentBlock.NewCall(systemFn, cmd)
			}
		}

		v.Visit(ctx.FunctionBody())

		if v.currentBlock.Term == nil {
			if !types.Equal(fn.Sig.RetType, types.Void) {
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

	return fn
}

func (v *IRVisitor) VisitFunctionBody(ctx *parser.FunctionBodyContext) interface{} {
	switch {
	case ctx.LBRACE() != nil && ctx.Statements() != nil:
		v.Visit(ctx.Statements())
	case ctx.ARROW() != nil && ctx.Expression() != nil:
		retVal := v.Visit(ctx.Expression())
		if val, ok := retVal.(value.Value); ok {
			if v.currentBlock.Term == nil {
				v.currentBlock.NewRet(val)
			}
		}
	}
	return nil
}

func (v *IRVisitor) VisitFormalParameterList(ctx *parser.FormalParameterListContext) interface{} {
	var params []*VariableInfo
	if nfps := ctx.NormalFormalParameters(); nfps != nil {
		for _, p := range nfps.AllNormalFormalParameter() {
			if param := v.Visit(p); param != nil {
				if infos, ok := param.([]*VariableInfo); ok {
					params = append(params, infos...)
				}
			}
		}
	}
	return params
}

func (v *IRVisitor) VisitNormalFormalParameter(ctx *parser.NormalFormalParameterContext) interface{} {
	if ctx.NormalFormalParameterNoMetadata() != nil {
		return v.Visit(ctx.NormalFormalParameterNoMetadata())
	}
	return nil
}

func (v *IRVisitor) VisitNormalFormalParameterNoMetadata(ctx *parser.NormalFormalParameterNoMetadataContext) interface{} {
	switch {
	case ctx.SimpleFormalParameter() != nil:
		return v.Visit(ctx.SimpleFormalParameter())
	case ctx.FunctionFormalParameter() != nil:
		return v.Visit(ctx.FunctionFormalParameter())
	case ctx.FieldFormalParameter() != nil:
		return v.Visit(ctx.FieldFormalParameter())
	}
	return nil
}

func (v *IRVisitor) VisitSimpleFormalParameter(ctx *parser.SimpleFormalParameterContext) interface{} {
	if ctx.DeclaredIdentifier() != nil {
		return v.Visit(ctx.DeclaredIdentifier())
	}
	if ctx.IDENTIFIER() != nil {
		name := ctx.IDENTIFIER().GetText()
		return []*VariableInfo{{Name: name, Type: types.I64}}
	}
	return nil
}

func (v *IRVisitor) VisitDeclaredIdentifier(ctx *parser.DeclaredIdentifierContext) interface{} {
	var typ types.Type = types.I64
	if ctx.FinalConstVarOrType() != nil {
		if t := v.Visit(ctx.FinalConstVarOrType()); t != nil {
			if tt, ok := t.(types.Type); ok {
				typ = tt
			}
		}
	}
	if ctx.IDENTIFIER() != nil {
		name := ctx.IDENTIFIER().GetText()
		return []*VariableInfo{{Name: name, Type: typ}}
	}
	return nil
}

func (v *IRVisitor) VisitFinalConstVarOrType(ctx *parser.FinalConstVarOrTypeContext) interface{} {
	if ctx.Type_() != nil {
		return v.Visit(ctx.Type_())
	}
	return types.I64
}

func (v *IRVisitor) VisitReturnType(ctx *parser.ReturnTypeContext) interface{} {
	if ctx.Type_() != nil {
		return v.Visit(ctx.Type_())
	}
	return types.Void
}

func (v *IRVisitor) VisitVariableDeclaration(ctx *parser.VariableDeclarationContext) interface{} {
	if ctx.VariableDeclarationList() != nil {
		return v.Visit(ctx.VariableDeclarationList())
	}
	return nil
}

func (v *IRVisitor) declareFunctionPrototype(ctx parser.IFunctionDeclarationContext) {
	if ctx.IDENTIFIER() == nil {
		return
	}

	funcName := ctx.IDENTIFIER().GetText()

	if _, exists := v.funcs[funcName]; exists {
		return
	}

	var retType types.Type = types.Void
	if ctx.ReturnType() != nil {
		if t := v.Visit(ctx.ReturnType()); t != nil {
			if tt, ok := t.(types.Type); ok {
				retType = tt
			}
		}
	}

	if funcName == "main" {
		retType = types.I32
	}

	var irParams []*ir.Param
	if fpl := ctx.FormalParameterList(); fpl != nil {
		if params := v.Visit(fpl); params != nil {
			if paramInfos, ok := params.([]*VariableInfo); ok {
				for _, p := range paramInfos {
					irParams = append(irParams, ir.NewParam(p.Name, p.Type))
				}
			}
		}
	}

	fn := v.Module.NewFunc(funcName, retType, irParams...)
	v.funcs[funcName] = fn
}

func (v *IRVisitor) newBlock(label string) *ir.Block {
	if v.currentFunc == nil {
		return nil
	}
	return v.currentFunc.NewBlock(v.freshLabel(label))
}

func (v *IRVisitor) freshLabel(prefix string) string {
	if v.currentFunc == nil {
		return prefix
	}
	return fmt.Sprintf("%s.%d", prefix, len(v.currentFunc.Blocks))
}

func (v *IRVisitor) pushBlock(block *ir.Block) {
	v.blockStack = append(v.blockStack, v.currentBlock)
	v.currentBlock = block
}

func (v *IRVisitor) popBlock() {
	if len(v.blockStack) == 0 {
		return
	}
	v.currentBlock = v.blockStack[len(v.blockStack)-1]
	v.blockStack = v.blockStack[:len(v.blockStack)-1]
}
