package visitor

import (
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
)

func (v *IRVisitor) declareRuntimeFuncs() {
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

func (v *IRVisitor) defineReadString() *ir.Func {
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

func (v *IRVisitor) defineReadInt() *ir.Func {
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
