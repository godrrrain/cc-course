package visitor

import (
	"fmt"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

type ParamDirection int

const (
	ParamIn ParamDirection = iota
	ParamOut
	ParamInOut
)

type VariableInfo struct {
	Name      string
	LLVMValue value.Value
	Type      types.Type
	Direction ParamDirection
	Bounds    []ArrayBounds
}

type ArrayBounds struct {
	Lower int64
	Upper int64
}

type Scope struct {
	parent *Scope
	vars   map[string]*VariableInfo
	meta   map[string]interface{}
}

func NewScope(parent *Scope) *Scope {
	return &Scope{
		parent: parent,
		vars:   make(map[string]*VariableInfo),
		meta:   make(map[string]interface{}),
	}
}

func (s *Scope) Get(name string) (*VariableInfo, bool) {
	v, ok := s.vars[name]
	if !ok && s.parent != nil {
		return s.parent.Get(name)
	}
	return v, ok
}

func (s *Scope) Set(name string, val *VariableInfo) error {
	if _, exists := s.vars[name]; exists {
		return fmt.Errorf("variable %s already declared in this scope", name)
	}
	s.vars[name] = val
	return nil
}

func (s *Scope) setMeta(key string, v interface{}) {
	if s.meta == nil {
		s.meta = make(map[string]interface{})
	}
	s.meta[key] = v
}

func (s *Scope) getMeta(key string) (interface{}, bool) {
	for cur := s; cur != nil; cur = cur.parent {
		if cur.meta != nil {
			if v, ok := cur.meta[key]; ok {
				return v, true
			}
		}
	}
	return nil, false
}

func (v *IRVisitor) enterScope() {
	v.currentScope = NewScope(v.currentScope)
}

func (v *IRVisitor) exitScope() {
	if v.currentScope == nil {
		v.Errors = append(v.Errors, fmt.Errorf("attempt to exit scope when no scope is active"))
		return
	}
	v.currentScope = v.currentScope.parent
}

type FuncWithParams struct {
	Func   *ir.Func
	Params []*VariableInfo
}

const (
	returnNameVar = "znach"
	loopExitVar   = "__loop_exit"
)
