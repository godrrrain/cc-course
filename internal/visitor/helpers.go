package visitor

import (
	"fmt"
	"os"

	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

func (v *IRVisitor) WriteToFile(filename string) error {
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

func (v *IRVisitor) getArrayElementPtr(vi *VariableInfo, indices []value.Value) value.Value {
	if vi == nil || vi.LLVMValue == nil {
		v.Errors = append(v.Errors, fmt.Errorf("nil variable info or LLVM value in array access"))
		return nil
	}

	effIdx := make([]value.Value, len(indices))
	for dim, raw := range indices {
		idx := v.ensureI32(raw)
		if dim < len(vi.Bounds) {
			lower := constant.NewInt(types.I32, vi.Bounds[dim].Lower)
			idx = v.currentBlock.NewSub(idx, lower)
		}
		effIdx[dim] = idx
	}

	args := append([]value.Value{constant.NewInt(types.I32, 0)}, effIdx...)
	return v.currentBlock.NewGetElementPtr(vi.Type, vi.LLVMValue, args...)
}
