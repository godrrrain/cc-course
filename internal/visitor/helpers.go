package visitor

import (
	"fmt"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
	"os"
	"strings"
	"unicode"

	"github.com/essentialkaos/translit"
)

func containsNonASCII(s string) bool {
	for _, r := range s {
		if r > unicode.MaxASCII {
			return true
		}
	}
	return false
}

func NormalizeFunctionName(name string) string {
	transliterated := translit.EncodeToICAO(name)
	if containsNonASCII(name) {
		transliterated += "_rus"
	}

	var sb strings.Builder
	for i, r := range transliterated {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || r == '_' {
			sb.WriteRune(r)
		} else if i == 0 && unicode.IsDigit(r) {
			sb.WriteRune('_')
			sb.WriteRune(r)
		} else {
			sb.WriteRune('_')
		}
	}
	return sb.String()
}

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

func stripQuotes(s string) string {
	if len(s) >= 2 {
		if (s[0] == '"' && s[len(s)-1] == '"') || (s[0] == '\'' && s[len(s)-1] == '\'') {
			return s[1 : len(s)-1]
		}
	}
	return s
}

func (v *IRVisitor) newBlock(base string) *ir.Block {
	name := fmt.Sprintf("%s.%d", base, len(v.currentFunc.Blocks))
	block := v.currentFunc.NewBlock(name)
	v.currentBlock = block
	return block
}

func (v *IRVisitor) pushBlock(b *ir.Block) {
	v.blockStack = append(v.blockStack, v.currentBlock)
	v.currentBlock = b
}

func (v *IRVisitor) popBlock() {
	n := len(v.blockStack)
	if n == 0 {
		v.currentBlock = nil
		return
	}
	v.currentBlock = v.blockStack[n-1]
	v.blockStack = v.blockStack[:n-1]
}

func (v *IRVisitor) castToMatch(lhs, rhs value.Value) (value.Value, value.Value) {
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

func (v *IRVisitor) freshLabel(prefix string) string {
	return fmt.Sprintf("%s.%d", prefix, len(v.currentFunc.Blocks))
}

func constIntValue(val value.Value) (int64, bool) {
	switch v := val.(type) {
	case *constant.Int:
		return v.X.Int64(), true
	case *ir.InstSub:
		if lhs, ok := v.X.(*constant.Int); ok && lhs.X.Sign() == 0 {
			if rhs, ok := v.Y.(*constant.Int); ok {
				return -rhs.X.Int64(), true
			}
		}
	}
	return 0, false
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
