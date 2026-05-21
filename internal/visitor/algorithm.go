package visitor

import (
	"fmt"
	"github.com/AskaryanKarine/BMSTU-CC/cource/internal/parser"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
)

func (v *IRVisitor) declareAlgorithmPrototype(ctx *parser.AlgorithmHeaderContext) {
	if ctx == nil {
		return
	}

	codeFuncName := ctx.AlgorithmNameTokens().GetText()
	llvmFuncName := NormalizeFunctionName(codeFuncName)

	// Первая встреченная функция это main на уровне LLVM.
	if v.mainOriginalName == "" {
		v.mainOriginalName = llvmFuncName
		llvmFuncName = "main"
	}

	if _, exists := v.funcs[llvmFuncName]; exists {
		return
	}

	var retType types.Type = types.Void

	if llvmFuncName == "main" {
		retType = types.I32
	}

	if ts := ctx.TypeSpecifier(); ts != nil {
		if t, ok := v.Visit(ts).(types.Type); ok {
			retType = t
		}
	}

	var irParams []*ir.Param
	if paramList := ctx.ParameterList(); paramList != nil {
		if res := v.Visit(paramList); res != nil {
			if paramInfos, ok := res.([]*VariableInfo); ok {
				for _, p := range paramInfos {
					pt := p.Type
					if _, isArr := pt.(*types.ArrayType); isArr {
						if _, alreadyPtr := pt.(*types.PointerType); !alreadyPtr {
							pt = types.NewPointer(pt)
						}
					}

					if p.Direction == ParamOut || p.Direction == ParamInOut {
						if _, isPtr := pt.(*types.PointerType); !isPtr {
							pt = types.NewPointer(pt)
						}
					}

					irParams = append(irParams, ir.NewParam(p.Name, pt))
				}
			}
		}
	}

	f := v.Module.NewFunc(llvmFuncName, retType, irParams...)
	v.funcs[llvmFuncName] = f
}

func (v *IRVisitor) VisitAlgorithmDefinition(ctx *parser.AlgorithmDefinitionContext) interface{} {
	v.enterScope()
	defer v.exitScope()

	fw := v.Visit(ctx.AlgorithmHeader())
	funcWithParams, ok := fw.(*FuncWithParams)
	if !ok || funcWithParams == nil || funcWithParams.Func == nil {
		return nil
	}
	fn := funcWithParams.Func
	v.currentFunc = fn

	for _, d := range ctx.AllVariableDeclaration() {
		v.Visit(d)
	}
	v.Visit(ctx.AlgorithmBody())

	// добавить возврат, если его нет
	if v.currentBlock.Term == nil {
		if !types.Equal(fn.Sig.RetType, types.Void) {
			if retVar, ok := v.currentScope.Get(returnNameVar); ok {
				retVal := v.currentBlock.NewLoad(retVar.Type, retVar.LLVMValue)
				v.currentBlock.NewRet(retVal)
			} else {
				v.currentBlock.NewRet(constant.NewInt(fn.Sig.RetType.(*types.IntType), 0))
			}
		} else {
			v.currentBlock.NewRet(nil) // void
		}
	}
	v.currentFunc = nil

	return fn
}

func (v *IRVisitor) VisitAlgorithmHeader(ctx *parser.AlgorithmHeaderContext) interface{} {
	name := ctx.AlgorithmNameTokens()
	if name == nil {
		v.Errors = append(v.Errors, fmt.Errorf("algorithm header name is nil"))
		return nil
	}
	userName := name.GetText()
	llvmName := NormalizeFunctionName(userName)

	// Первая встреченная функция это main на уровне LLVM.
	if v.mainOriginalName == llvmName {
		llvmName = "main"
	}

	var paramInfos []*VariableInfo
	if plist := ctx.ParameterList(); plist != nil {
		if raw := v.Visit(plist); raw != nil {
			if infos, ok := raw.([]*VariableInfo); ok {
				paramInfos = infos
			}
		}
	}

	var irParams []*ir.Param
	for _, p := range paramInfos {
		pt := p.Type
		if p.Direction == ParamOut || p.Direction == ParamInOut {
			if _, ok := pt.(*types.PointerType); !ok {
				pt = types.NewPointer(pt)
			}
		}
		irParams = append(irParams, ir.NewParam(p.Name, pt))
	}

	fn, exists := v.funcs[llvmName]
	if !exists {
		var retT types.Type
		retT = types.Void
		if ts := ctx.TypeSpecifier(); ts != nil {
			if t, ok := v.Visit(ts).(types.Type); ok {
				retT = t
			}
		}
		fn = v.Module.NewFunc(llvmName, retT, irParams...)
		v.funcs[llvmName] = fn
	} else if len(fn.Params) == 0 {
		fn.Params = append(fn.Params, irParams...)
	}

	if len(fn.Blocks) == 0 {
		entry := fn.NewBlock("entry")
		v.currentBlock = entry
		v.pushBlock(entry)

		if !types.Equal(fn.Sig.RetType, types.Void) {
			retAlloca := entry.NewAlloca(fn.Sig.RetType)
			_ = v.currentScope.Set(returnNameVar, &VariableInfo{
				Name:      returnNameVar,
				Type:      fn.Sig.RetType,
				LLVMValue: retAlloca,
			})
		}

		for i, pInfo := range paramInfos {
			llvmParam := fn.Params[i]

			if ptrTy, isPtr := pInfo.Type.(*types.PointerType); isPtr {
				pInfo.LLVMValue = llvmParam

				if arrTy, ok := ptrTy.ElemType.(*types.ArrayType); ok {
					pInfo.Type = arrTy
				}
			} else {
				alloca := entry.NewAlloca(pInfo.Type)
				entry.NewStore(llvmParam, alloca)
				pInfo.LLVMValue = alloca
			}

			_ = v.currentScope.Set(pInfo.Name, pInfo)
		}
	} else {
		v.currentBlock = fn.Blocks[0]
		v.pushBlock(v.currentBlock)
	}

	return &FuncWithParams{Func: fn, Params: paramInfos}
}

func (v *IRVisitor) VisitParameterList(ctx *parser.ParameterListContext) interface{} {
	var allParams []*VariableInfo
	for _, decl := range ctx.AllParameterDeclaration() {
		p := v.Visit(decl)
		if p == nil {
			continue
		}

		declParams, ok := p.([]*VariableInfo)
		if !ok {
			v.Errors = append(v.Errors, fmt.Errorf("expected []*VariableInfo from VisitParameterDeclaration, got %T", p))
			continue
		}
		allParams = append(allParams, declParams...)
	}

	return allParams
}

func (v *IRVisitor) VisitParameterDeclaration(ctx *parser.ParameterDeclarationContext) interface{} {
	var direction ParamDirection = ParamIn
	if ctx.IN_PARAM() != nil {
		direction = ParamIn
	}
	if ctx.OUT_PARAM() != nil {
		direction = ParamOut
	}
	if ctx.INOUT_PARAM() != nil {
		direction = ParamInOut
	}

	typVal := v.Visit(ctx.TypeSpecifier())
	typ, ok := typVal.(types.Type)
	if !ok || typ == nil {
		v.Errors = append(v.Errors, fmt.Errorf("invalid or missing type specifier in parameter declaration"))
		return nil
	}

	varListVal := v.Visit(ctx.VariableList())
	varInfos, ok := varListVal.([]*VariableInfo)
	if !ok {
		v.Errors = append(v.Errors, fmt.Errorf("invalid variable list in parameter declaration"))
		return nil
	}

	for _, vi := range varInfos {
		vi.Direction = direction

		if len(vi.Bounds) > 0 {
			t := typ
			for i := len(vi.Bounds) - 1; i >= 0; i-- {
				dim := vi.Bounds[i]
				length := uint64(dim.Upper - dim.Lower + 1)
				t = types.NewArray(length, t)
			}
			vi.Type = types.NewPointer(t)
		} else {
			vi.Type = typ
			if direction == ParamOut || direction == ParamInOut {
				vi.Type = types.NewPointer(vi.Type)
			}
		}
	}

	return varInfos
}

func (v *IRVisitor) VisitAlgorithmBody(ctx *parser.AlgorithmBodyContext) interface{} {
	return v.Visit(ctx.StatementSequence())
}
